package handlers

import (
	"net/http"

	"EMSproject/database"
	"EMSproject/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func CreateSite(c *echo.Context) error {
	var site models.Site

	if err := c.Bind(&site); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	site.ID = uuid.New().String()

	query := `
	INSERT INTO site (id, name, address, type, created_at, updated_at)
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

func GetSites(c *echo.Context) error {
	rows, err := database.DB.Query("SELECT id, name, address, type, created_at, updated_at FROM site")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	defer rows.Close()

	sites := make([]models.Site, 0)

	for rows.Next() {
		var site models.Site
		if err := rows.Scan(
			&site.ID,
			&site.Name,
			&site.Address,
			&site.Type,
			&site.CreatedAt,
			&site.UpdatedAt,
		); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan site data"})
		}
		sites = append(sites, site)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error iterating rows"})
	}

	return c.JSON(http.StatusOK, sites)
}

func GetSiteByID(c *echo.Context) error {
	site_id := c.Param("site_id")

	var site models.Site

	err := database.DB.QueryRow("SELECT id, name, address, type, created_at, updated_at FROM site WHERE id = $1", site_id).Scan(
		&site.ID,
		&site.Name,
		&site.Address,
		&site.Type,
		&site.CreatedAt,
		&site.UpdatedAt,
	)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, site)
}

func UpdateSite(c *echo.Context) error {
	site_id := c.Param("site_id")

	var site models.Site

	if err := c.Bind(&site); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	query := `
	UPDATE site
	SET name = $1, address = $2, type = $3, updated_at = NOW()
	WHERE id = $4
	`
	_, err := database.DB.Exec(
		query,
		site.Name,
		site.Address,
		site.Type,
		site_id,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, site)
}

func DeleteSite(c *echo.Context) error {
	site_id := c.Param("site_id")

	_, err := database.DB.Exec("DELETE FROM site WHERE id = $1", site_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Site deleted successfully"})
}
