package main

import (
	"app/controllers"
	"app/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	storage.InitializeDataBase()

	const user_id = "/user/:id"
	const cart_id ="/cart/:id"
	const category_id = "/category/:id"
	const m5_id = "/m5/:id"

	e := echo.New()
	//products
	e.GET("/product/:id", controllers.GetProduct)
	e.GET("/products", controllers.GetProducts)
	e.POST("/product", controllers.CreateProduct)
	e.PUT("/product", controllers.UpdateProduct)
	e.DELETE("/delete_product/:id", controllers.DeleteProduct)
	//user
	e.GET(user_id, controllers.GetUser)
	e.GET("/users", controllers.GetUsers)
	e.POST("/user", controllers.CreateUser)
	e.PUT(user_id, controllers.UpdateUser)
	e.DELETE(user_id, controllers.DeleteUser)
	//cart
	e.GET(cart_id, controllers.GetCart)
	e.POST("/cart", controllers.CreateCart)
	e.PUT(cart_id, controllers.UpdateCart)
	e.DELETE(cart_id, controllers.DeleteCart)
	//category
	e.GET(category_id, controllers.GetCategory)
	e.POST("/category", controllers.CreateCategory)
	e.PUT(category_id, controllers.UpdateCategory)
	e.DELETE(category_id, controllers.DeleteCategory)
	//m5
	e.GET(m5_id, controllers.GetM5)
	e.POST("/m5", controllers.CreateM5)
	e.PUT(m5_id, controllers.UpdateM5)
	e.DELETE(m5_id, controllers.DeleteM5)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
