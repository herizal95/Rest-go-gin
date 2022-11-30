package routes

import (
	"pustaka-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SettingRoutes(db *gorm.DB) *gin.Engine {

	routes := gin.Default()
	routes.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	// Crud Biodata
	routes.GET("/biodata", controllers.BiodataView)
	routes.POST("/biodata", controllers.BiodataAdd)
	routes.GET("/biodata/:id", controllers.BiodataGetId)
	routes.DELETE("/biodata/:id", controllers.BiodataDelete)
	routes.PATCH("/biodata/:id", controllers.BiodataUpdate)

	// Crud DataBarang
	routes.GET("/barang", controllers.BarangView)
	routes.POST("/barang", controllers.BarangAdd)
	routes.GET("/barang/:id", controllers.BarangGetID)
	routes.DELETE("/barang/:id", controllers.BarangDelete)
	routes.PATCH("/barang/:id", controllers.BarangUpdate)

	return routes

}
