package controllers

import (
	"net/http"
	"pustaka-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateBarang struct {
	KodeBarang  string `json:"kode_barang"`
	NamaBarang  string `json:"nama_barang"`
	JenisBarang string `json:"jenis_barang"`
	HargaBeli   int    `json:"harga_beli"`
	HargaJual   int    `json:"harga_jual"`
	Kategori    string `json:"kategori"`
}

type UpdateBarang struct {
	KodeBarang  string `json:"kode_barang"`
	NamaBarang  string `json:"nama_barang"`
	JenisBarang string `json:"jenis_barang"`
	HargaBeli   int    `json:"harga_beli"`
	HargaJual   int    `json:"harga_jual"`
	Kategori    string `json:"kategori"`
}

// Get All
func BarangView(con *gin.Context) {
	db := con.MustGet("db").(*gorm.DB)
	var data []models.Barang
	db.Find(&data)

	con.JSON(http.StatusOK, gin.H{
		"responce": 200,
		"status":   1,
		"data":     data,
	})
}

// Create New
func BarangAdd(c *gin.Context) {
	// validasi var input
	var input CreateBarang
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// create barang
	data := models.Barang{
		KodeBarang:  input.KodeBarang,
		NamaBarang:  input.NamaBarang,
		JenisBarang: input.JenisBarang,
		HargaBeli:   input.HargaBeli,
		HargaJual:   input.HargaJual,
		Kategori:    input.Kategori,
	}

	condb := c.MustGet("db").(*gorm.DB)
	condb.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Get :id
func BarangGetID(c *gin.Context) {

	var data []models.Barang

	db := c.MustGet("db").(*gorm.DB)
	if errr := db.Where("id = ?", c.Param("id")).First(&data).Error; errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message error": "Data Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   1,
		"responce": 200,
		"data":     data,
	})

}

// delete :id
func BarangDelete(c *gin.Context) {
	var data models.Barang
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id =?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message error": "Data tidak ada!"})
		return
	}
	db.Delete(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// update :id
func BarangUpdate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var data models.Barang

	if err := db.Where("id=?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "data not found!"})
		return
	}

	var input UpdateBarang
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"msg": err.Error()})
	}

	var inputUpdate models.Barang
	inputUpdate.NamaBarang = input.NamaBarang
	inputUpdate.JenisBarang = input.JenisBarang
	inputUpdate.KodeBarang = input.KodeBarang
	inputUpdate.HargaBeli = input.HargaBeli
	inputUpdate.HargaJual = input.HargaJual
	inputUpdate.Kategori = input.Kategori

	db.Model(&data).Update(inputUpdate)

	c.JSON(http.StatusOK, gin.H{
		"status":   1,
		"responce": 200,
		"data":     data,
	})
}
