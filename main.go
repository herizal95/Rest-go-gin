package main

import (
	"pustaka-api/models"
	"pustaka-api/routes"
	"pustaka-api/services"
)

func main() {

	db := services.SetupDB()
	db.AutoMigrate(&models.Biodata{}, &models.Barang{})

	routes := routes.SettingRoutes(db)
	routes.Run()
}
