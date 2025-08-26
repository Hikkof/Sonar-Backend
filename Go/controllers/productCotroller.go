package controllers

import (
	"app/models"
	"app/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var product models.Product
	storage.Database.First(&product, id)

	return c.JSON(http.StatusOK, product)
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	storage.Database.Preload("ProductID").Find(&products)

	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		return err
	}

	if err := storage.Database.Create(product).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var product models.Product

	if err := storage.Database.First(&product, id).Error; err != nil {
		return err
	}
	updateData := new(models.Product)

	if err := c.Bind(updateData); err != nil {
		return err
	}
	product.Name = updateData.Name
	product.Price = updateData.Price
	product.Description = updateData.Description
	product.SellerID = updateData.SellerID

	if err := storage.Database.Save(&product).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var product models.Product

	if err := storage.Database.First(&product, id).Error; err != nil {
		return err
	}
	storage.Database.Delete(&product)

	return c.NoContent(http.StatusNoContent)
}
