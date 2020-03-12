package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"sample_graphql_in_go/src/models"
)

type DbManager struct {
	Database *gorm.DB
}

func NewDatabase() *DbManager {
	db, err := gorm.Open("mysql", "user:password@(db:3306)/db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	migrateSchemas(db)
	return &DbManager{Database: db}
}

func migrateSchemas(db *gorm.DB) {
	db.AutoMigrate(&models.Payment{})
}

func (d *DbManager) AddPayment(payment *models.Payment)  {
	d.Database.Create(payment)
}

func (d *DbManager) GetPaymentsByAuthor(author string)  *[]models.Payment {
	var payments []models.Payment
	d.Database.Where("author = ?", author).Find(&payments)
	return  &payments
}