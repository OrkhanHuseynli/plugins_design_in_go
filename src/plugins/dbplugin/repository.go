package dbplugin

import (
	"github.com/jinzhu/gorm"
	"plugins_design_in_go/src/models"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) AddPayment(payment *models.Payment)  {
	r.db.Create(payment)
}

func (r *Repository) GetPaymentsByAuthor(author string)  *[]models.Payment {
	var payments []models.Payment
	r.db.Where("author = ?", author).Find(&payments)
	return  &payments
}