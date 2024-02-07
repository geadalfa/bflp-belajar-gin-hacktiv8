package controllers

import (
	"fmt"
	"net/http"

	"belajar-gin/database"
	"belajar-gin/models"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

// CreateCars godoc
// @Summary Post details for a given Id
// @Description Post detials of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Sucess 200 {object} models.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.Car

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create book
	carinput := models.Car{Brand: input.Brand, Model: input.Model, Price: input.Price}
	db.Create(&carinput)

	ctx.JSON(http.StatusOK, gin.H{"data": carinput})
}

// UpdateCars godoc
// @Summary Update car identified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("id")

	db := database.GetDB()
	var carUpdate models.Car

	// Retrieve the car from the database
	if err := db.First(&carUpdate, carID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Bind the JSON data to the carUpdate struct
	if err := ctx.ShouldBindJSON(&carUpdate); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Update the car record
	if err := db.Save(&carUpdate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v has been successfully updated", carID),
	})
}

// GetOneCar godoc
// @Summary: Get details for a given Id
// @Description Get details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetCar(ctx *gin.Context) {
	var db = database.GetDB()

	var carOne models.Car
	err := db.First(&carOne, "Id = ?", ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": carOne})

}

// DeleteCars godoc
// @Summary Delete car identified by the given Id
// @Description Delete the order corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{Id} {delete}
func DeleteCar(c *gin.Context) {
	var db = database.GetDB()

	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetAllCar godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produces json
// @Success 200 {object} models.Car
// @Router /orders [get]
func GetAllCar(ctx *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting datas:", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cars})
}
