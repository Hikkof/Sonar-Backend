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
	const cartId = "/cart/:id"
	const categoryId = "/category/:id"
	const m5Id = "/m5/:id"

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
	//cart
	e.GET(cartId, controllers.GetCart)
	e.POST("/cart", controllers.CreateCart)
	e.PUT(cartId, controllers.UpdateCart)
	e.DELETE(cartId, controllers.DeleteCart)
	//category
	e.GET(categoryId, controllers.GetCategory)
	e.POST("/category", controllers.CreateCategory)
	e.PUT(categoryId, controllers.UpdateCategory)
	e.DELETE(categoryId, controllers.DeleteCategory)
	//m5
	e.GET(m5Id, controllers.GetM5)
	e.POST("/m5", controllers.CreateM5)
	e.PUT(m5Id, controllers.UpdateM5)
	e.DELETE(m5Id, controllers.DeleteM5)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
