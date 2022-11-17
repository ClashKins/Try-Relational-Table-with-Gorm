package controllers

import (
	"LATIHAN1/database"
	"LATIHAN1/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetAllProducts(c *gin.Context) {
	var db = database.GetDB()

	var marts []models.Mart
	err := db.Find(&marts).Error

	if err != nil {
		fmt.Println("Error getting data from marts :", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": marts})
}

func GetOneProducts(c *gin.Context) {
	var db = database.GetDB()
	
	var oneProduct models.Mart
	
	err := db.First(&oneProduct, "Id = ?", c.Param("id")).Error
	if err !=  nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data one": oneProduct})
}

func CreateMarts(c *gin.Context) {
	var db = database.GetDB()
	var input models.Mart
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
	}
	martinput := models.Mart{Bulan: input.Bulan, Makanan: input.Makanan, Minuman: input.Minuman, Harga: input.Harga}
	db.Create(martinput)
	c.JSON(http.StatusOK, gin.H{"data" : martinput})
}

func UpdateMarts(c *gin.Context){
	var db = database.GetDB()
	var mart models.Mart
	err := db.First(&mart, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input models.Mart
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&mart).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": mart})
}

func DeleteMart(c *gin.Context) {
	var db = database.GetDB()
	var deleteMart models.Mart
	err := db.First(&deleteMart, "Id = ?", c.Param("id")).Error;
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&deleteMart)
	c.JSON(http.StatusOK, gin.H{"data": true})
}