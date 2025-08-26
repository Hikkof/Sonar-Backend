package controllers

import (
	"app/models"
	"app/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var category models.Category
	storage.Database.First(&category, id)

	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	category := new(models.Category)

	if err := c.Bind(category); err != nil {
		return err
	}
	storage.Database.Create(category)

	return c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var category models.Category

	if err := storage.Database.First(&category, id).Error; err != nil {
		return err
	}
	updateData := new(models.Category)

	if err := c.Bind(updateData); err != nil {
		return err
	}
	category.Name = updateData.Name

	if err := storage.Database.Save(&category).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var category models.Category

	if err := storage.Database.First(&category, id).Error; err != nil {
		return err
	}
	storage.Database.Delete(&category)
	return c.NoContent(http.StatusNoContent)
}
