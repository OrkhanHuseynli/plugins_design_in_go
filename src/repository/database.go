package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"plugins_design_in_go/src/models"
)

type DbManager struct {
	Database *gorm.DB
}

func NewDatabase() *DbManager {
	db, err := gorm.Open("mysql", "user:user@(db:3307)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	migrateSchemas(db)
	return &DbManager{Database: db}
}

func migrateSchemas(db *gorm.DB) {
	db.AutoMigrate(&models.Payment{}, &models.Product{})
}

func (d *DbManager) AddPayment(payment *models.Payment)  {
	d.Database.Create(payment)
}

func (d *DbManager) GetPaymentsByAuthor(author string)  *[]models.Payment {
	var payments []models.Payment
	d.Database.Where("author = ?", author).Find(&payments)
	return  &payments
}