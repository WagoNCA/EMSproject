package database

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	InfluxClient influxdb2.Client
	InfluxOrg    string
	InfluxBucket string
)

func ConnectInflux() {

	url := os.Getenv("INFLUX_URL")
	token := os.Getenv("INFLUX_TOKEN")

	InfluxOrg = os.Getenv("INFLUX_ORG")
	InfluxBucket = os.Getenv("INFLUX_BUCKET")

	InfluxClient = influxdb2.NewClient(
		url,
		token,
	)
}
