package controllers

import (
	"app/models"
	"app/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//var m5Number = 0

func GetM5(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	var m5 models.M5
	storage.Database.First(&m5, id)
	/*
		if err := storage.Database.First(&m5, id).Error; err != nil {
			return err
		}
	*/
	return c.JSON(http.StatusOK, m5)
}

func CreateM5(c echo.Context) error {
	m5 := new(models.M5)
	if err := c.Bind(m5); err != nil {
		return err
	}
	storage.Database.Create(m5)
	/*
		if err := storage.Database.Create(m5).Error; err != nil {
			return err
		}
	*/
	return c.JSON(http.StatusCreated, m5)
}

func UpdateM5(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	var m5 models.M5
	if err := storage.Database.First(&m5, id).Error; err != nil {
		return err
	}

	updateData := new(models.M5)
	if err := c.Bind(updateData); err != nil {
		return err
	}

	m5.Name = updateData.Name

	if err := storage.Database.Save(&m5).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m5)
}

func DeleteM5(c echo.Context) error {
	//id, _ := strconv.Atoi(c.Param("id"))
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid m5 ID"})
	}
	var m5 models.M5
	if err := storage.Database.First(&m5, id).Error; err != nil {
		return err
	}
	storage.Database.Delete(&m5)
	return c.NoContent(http.StatusNoContent)
}
