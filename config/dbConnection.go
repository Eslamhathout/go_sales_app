package config

import (
	"fmt"
	"os"
	models "sales_api_go/go_sales_app/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connection Error")
	}
	DB = db
	fmt.Println("DB Connected")

	//Call the migrate all
	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Discount{},
		&models.Order{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
	)
}
