package main

import (
	"EMSproject/database"
	"EMSproject/handlers"

	"github.com/labstack/echo/v5"
)

func main() {
	database.Connect()

	e := echo.New()

	e.POST("/sites", handlers.CreateSite)
	e.GET("/sites", handlers.GetSites)
	e.GET("/sites/:site_id", handlers.GetSiteByID)
	e.PUT("/sites/:site_id", handlers.UpdateSite)
	e.DELETE("/sites/:site_id", handlers.DeleteSite)

	e.POST("/sites/:site_id/meters", handlers.CreateMeter)
	e.GET("/sites/:site_id/meters", handlers.GetMetersBySiteID)
	e.GET("/meters/:meter_id", handlers.GetMeterByID)
	e.PUT("/meters/:meter_id", handlers.UpdateMeter)
	e.DELETE("/meters/:meter_id", handlers.DeleteMeter)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error(err.Error())
	}
}
