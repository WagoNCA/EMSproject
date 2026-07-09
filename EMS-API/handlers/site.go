package handlers

import (
	"net/http"

	"EMSproject/database"
	"EMSproject/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func CreateSite(c echo.Context) error {
	var site models.Site

	if err := c.Bind(&site); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	site.ID = uuid.New().String()

	query := `
	INSERT INTO Site (id, name, address, type, created_at, updated_at)
	VALUES ($1, $2, $3, $4, NOW(), NOW())
	`

	_, err := database.DB.Exec(
		query,
		site.ID,
		site.Name,
		site.Address,
		site.Type,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create site"})
	}

	return c.JSON(http.StatusCreated, site)
}
