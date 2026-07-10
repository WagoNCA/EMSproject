package handlers

import (
	"database/sql"
	"net/http"

	"EMSproject/database"
	"EMSproject/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func CreateMeter(c *echo.Context) error {
	var meter models.Meter

	if err := c.Bind(&meter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	meter.ID = uuid.New().String()

	query := `
	INSERT INTO meter (id, site_id, unit, type, created_at, updated_at)
	VALUES ($1, $2, $3, $4, NOW(), NOW())
	`

	_, err := database.DB.Exec(
		query,
		meter.ID,
		meter.SiteID,
		meter.Unit,
		meter.Type,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create meter"})
	}

	return c.JSON(http.StatusCreated, meter)
}

func GetMetersBySiteID(c *echo.Context) error {
	site_id := c.Param("site_id")

	rows, err := database.DB.Query("SELECT id, name, type, site_id, created_at, updated_at FROM meter WHERE site_id = $1", site_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve meters"})
	}

	defer rows.Close()

	var meters []models.Meter

	for rows.Next() {
		var meter models.Meter
		if err := rows.Scan(&meter.ID, &meter.Unit, &meter.Type, &meter.SiteID, &meter.CreatedAt, &meter.UpdatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan meter data"})
		}
		meters = append(meters, meter)
	}

	return c.JSON(http.StatusOK, meters)
}

func GetMeterByID(c *echo.Context) error {
	meter_id := c.Param("id")

	row := database.DB.QueryRow("SELECT id, site_id, unit, type, created_at, updated_at FROM Meter WHERE id = $1", meter_id)

	var meter models.Meter
	if err := row.Scan(&meter.ID, &meter.SiteID, &meter.Unit, &meter.Type, &meter.CreatedAt, &meter.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Meter not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve meter"})
	}

	return c.JSON(http.StatusOK, meter)
}

func UpdateMeter(c *echo.Context) error {
	meter_id := c.Param("id")

	var meter models.Meter
	if err := c.Bind(&meter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	query := `
	UPDATE Meter
	SET unit = $1, type = $2, site_id = $3, updated_at = NOW()
	WHERE id = $4
	`

	result, err := database.DB.Exec(
		query,
		meter.Unit,
		meter.Type,
		meter.SiteID,
		meter_id,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update meter"})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rows affected"})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Meter not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Meter updated successfully"})
}

func DeleteMeter(c *echo.Context) error {
	meter_id := c.Param("id")

	result, err := database.DB.Exec("DELETE FROM Meter WHERE id = $1", meter_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete meter"})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve rows affected"})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Meter not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Meter deleted successfully"})
}
