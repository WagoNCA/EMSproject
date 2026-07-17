package handlers

import (
	"EMSproject/database"
	"EMSproject/models"
	"context"
	"fmt"
	"net/http"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo/v5"
)

func CreateReading(c *echo.Context) error {
	var reading models.MeterReading

	if err := c.Bind(&reading); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	writeAPI := database.InfluxClient.WriteAPIBlocking(
		database.InfluxOrg,
		database.InfluxBucket,
	)

	point := influxdb2.NewPoint(
		"meter_readings",
		map[string]string{
			"meter_id": reading.MeterID,
		},
		map[string]interface{}{
			"value": reading.Value,
		},
		time.Now(),
	)

	err := writeAPI.WritePoint(
		context.Background(),
		point,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, reading)
}

func GetReadings(c *echo.Context) error {
	meterID := c.Param("meter_id")
	from := c.QueryParam("from")
	to := c.QueryParam("to")

	_, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "invalid from date"})
	}

	_, err = time.Parse(time.RFC3339, to)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "invalid to date"})
	}

	if from == "" || to == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "from and to are required"})
	}

	queryAPI := database.InfluxClient.QueryAPI(
		database.InfluxOrg,
	)

	query := fmt.Sprintf(`
from(bucket: "%s")
	|> range(start: time(v: "%s"), stop: time(v: "%s"))
	|> filter(fn: (r) => r._measurement == "meter_readings")
	|> filter(fn: (r) => r.meter_id == "%s")
`,
		database.InfluxBucket,
		from,
		to,
		meterID,
	)

	result, err := queryAPI.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	readings := []map[string]interface{}{}

	for result.Next() {
		readings = append(readings, map[string]interface{}{
			"time":  result.Record().Time(),
			"value": result.Record().Value(),
		})
	}

	if result.Err() != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Err().Error()})
	}

	return c.JSON(http.StatusOK, readings)
}
