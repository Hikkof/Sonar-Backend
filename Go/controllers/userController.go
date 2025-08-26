package controllers

import (
	"app/models"
	"app/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var user models.User
	storage.Database.First(&user, id)

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	var users []models.User
	storage.Database.Preload("ProductID").Find(&users)

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return err
	}
	storage.Database.Create(user)

	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var user models.User

	if err := storage.Database.First(&user, id).Error; err != nil {
		return err
	}
	updateData := new(models.User)

	if err := c.Bind(updateData); err != nil {
		return err
	}
	user.Name = updateData.Name
	user.Password = updateData.Password

	if err := storage.Database.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}
	var user models.User

	if err := storage.Database.First(&user, id).Error; err != nil {
		return err
	}
	storage.Database.Delete(&user)

	return c.NoContent(http.StatusNoContent)
}
