package controllers

import (
	"app/models"
	"app/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCart(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var cart models.Cart
	storage.Database.First(&cart, id)

	return c.JSON(http.StatusOK, cart)
}

func CreateCart(c echo.Context) error {
	cart := new(models.Cart)

	if err := c.Bind(cart); err != nil {
		return err
	}
	storage.Database.Create(cart)

	return c.JSON(http.StatusCreated, cart)
}

func AddProductToCart(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var cart models.Cart

	if err := storage.Database.First(&cart, id).Error; err != nil {
		return err
	}
	var product models.Product

	if err := storage.Database.First(&product, id).Error; err != nil {
		return err
	}
	err = storage.Database.Model(&cart).Association("Contents").Append(&product)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cart)
}

func RemoveProductFromCart(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var cart models.Cart

	if err := storage.Database.First(&cart, id).Error; err != nil {
		return err
	}
	var product models.Product

	if err := storage.Database.First(&product, id).Error; err != nil {
		return err
	}
	err = storage.Database.Model(&cart).Association("Contents").Delete(&product)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateCart(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var cart models.Cart

	if err := storage.Database.First(&cart, id).Error; err != nil {
		return err
	}
	updateData := new(models.Cart)

	if err := c.Bind(updateData); err != nil {
		return err
	}
	cart.Contents = updateData.Contents

	if err := storage.Database.Save(&cart).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cart)
}

func DeleteCart(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var cart models.Cart

	if err := storage.Database.First(&cart, id).Error; err != nil {
		return err
	}
	storage.Database.Delete(&cart)

	return c.NoContent(http.StatusNoContent)
}
