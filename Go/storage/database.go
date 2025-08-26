package storage

import (
	"app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var Database *gorm.DB

func createProducts() {
	var products = []models.Product{
		{ProductID: 1, Name: "Boots", Price: 109.99, Description: "gr8 description", SellerID: 1},
		{ProductID: 2, Name: "The Thing", Price: 209.99, Description: "a novel", SellerID: 1},
		{ProductID: 3, Name: "Boot", Price: 9.99, Description: "a singular boot", SellerID: 1},
	}
	for i := 0; i < len(products); i++ {
		result := Database.Create(&products[i])
		if result.Error != nil {
			log.Fatal("Error: ", result.Error)
		}
	}
}

func createUsers() {
	var users = []models.User{
		{UserID: 1, Name: "NoNaMe78", Password: "123456"},
	}
	for i := 0; i < len(users); i++ {
		result := Database.Create(&users[i])
		if result.Error != nil {
			log.Fatal("Error: ", result.Error)
		}
	}
}

func InitializeDataBase() {
	var err error
	Database, err = gorm.Open(sqlite.Open("database.Database"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	err = Database.AutoMigrate(&models.Product{}, &models.User{}, &models.Product{}, &models.Category{}, &models.M5{})
	if err != nil {
		log.Fatal("Error: ", err)
	}
	//createProducts()
	//createUsers()
}
