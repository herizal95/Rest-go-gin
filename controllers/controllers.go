package controllers

import (
	"net/http"
	"pustaka-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateBio struct {
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Jurusan   string `json:"jurusan"`
	Pekerjaan string `json:"pekerjaan"`
}
type UpdateBio struct {
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Jurusan   string `json:"jurusan"`
	Pekerjaan string `json:"pekerjaan"`
}
type Responce struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []models.Biodata
}

// Get All Biodata
func BiodataView(c *gin.Context) {

	var data []models.Biodata // []models : mengambil semua isi pada models Biodata
	db := c.MustGet("db").(*gorm.DB)

	db.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Create new Biodata
func BiodataAdd(c *gin.Context) {

	var input CreateBio
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := models.Biodata{
		Nama:      input.Nama,
		Alamat:    input.Alamat,
		Jurusan:   input.Jurusan,
		Pekerjaan: input.Pekerjaan,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Get Biodata/:id
func BiodataGetId(c *gin.Context) {

	var data []models.Biodata

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}

// Delete Biodata/:id
func BiodataDelete(c *gin.Context) {

	var data models.Biodata

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak ditemukan!"})
		return
	}
	db.Delete(&data)

	c.JSON(http.StatusOK, gin.H{"data": true})

}

// Update Biodata/:id
func BiodataUpdate(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var data models.Biodata

	if errr := db.Where("id = ?", c.Param("id")).First(&data).Error; errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Data tidak ditemukan!"})
		return
	}

	//validasi input
	var input UpdateBio
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	var inputUpdate models.Biodata
	inputUpdate.Nama = input.Nama
	inputUpdate.Alamat = input.Alamat
	inputUpdate.Jurusan = input.Jurusan
	inputUpdate.Pekerjaan = input.Pekerjaan

	db.Model(&data).Updates(inputUpdate)

	c.JSON(http.StatusOK, gin.H{"doto": data})

}
