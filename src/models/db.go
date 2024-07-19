package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func OpenDatabaseConnection() {
	var err error
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Douala", host, username, password, databaseName, port)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}
}

func AutoMigrateModels() {
	err := Database.AutoMigrate(&User{}, &Product{}, &Comment{}, &Role{}, &Category{}, &Image{})
	if err != nil {
		println(err)
		return
	}
}

func SeedData() {
	var roles = []Role{{Name: "admin", Description: "Administrator role"}, {Name: "manager", Description: "Manager role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var hastPassword, err = bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		hastPassword = []byte(os.Getenv("ADMIN_PASSWORD"))
	}
	var user = []User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: string(hastPassword), RoleID: 1, Age: ""}}
	Database.Save(&roles)
	Database.Save(&user)
}
