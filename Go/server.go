package main

import (
	"app/controllers"
	"app/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	storage.InitializeDataBase()

	e := echo.New()
	//products
	e.GET("/product/:id", controllers.GetProduct)
	e.GET("/products", controllers.GetProducts)
	e.POST("/product", controllers.CreateProduct)
	e.PUT("/product", controllers.UpdateProduct)
	e.DELETE("/delete_product/:id", controllers.DeleteProduct)
	//user
	e.GET("/user/:id", controllers.GetUser)
	e.GET("/users", controllers.GetUsers)
	e.POST("/user", controllers.CreateUser)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)
	//cart
	e.GET("/cart/:id", controllers.GetCart)
	e.POST("/cart", controllers.CreateCart)
	e.PUT("/cart/:id", controllers.UpdateCart)
	e.DELETE("/cart/:id", controllers.DeleteCart)
	//category
	e.GET("/cart/:id", controllers.GetCategory)
	e.POST("/cart", controllers.CreateCategory)
	e.PUT("/cart/:id", controllers.UpdateCategory)
	e.DELETE("/cart/:id", controllers.DeleteCategory)
	//m5
	e.GET("/cart/:id", controllers.GetM5)
	e.POST("/cart", controllers.CreateM5)
	e.PUT("/cart/:id", controllers.UpdateM5)
	e.DELETE("/cart/:id", controllers.DeleteM5)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
