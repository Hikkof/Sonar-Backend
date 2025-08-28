package main

import (
	"app/controllers"
	"app/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	storage.InitializeDataBase()

	const userId = "/user/:id"

	e := echo.New()
	//products
	e.GET("/product/:id", controllers.GetProduct)
	e.GET("/products", controllers.GetProducts)
	e.POST("/product", controllers.CreateProduct)
	e.PUT("/product", controllers.UpdateProduct)
	e.DELETE("/delete_product/:id", controllers.DeleteProduct)
	//user
	e.GET(userId, controllers.GetUser)
	e.GET("/users", controllers.GetUsers)
	e.POST("/user", controllers.CreateUser)
	e.PUT(userId, controllers.UpdateUser)
	e.DELETE(userId, controllers.DeleteUser)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
