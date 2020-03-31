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

func (r *Repository) GetPaymentsByAuthor(author string)  []models.Payment {
	var payments []models.Payment
	r.db.Model(models.Payment{Author:author}).Find(&payments)
	return  payments
}

func (r *Repository) DeletePaymentsByAuthor(author string)  bool {
	r.db.Where("author = ?", author).Delete(&models.Payment{})
	return true
}

func (r *Repository) GetPaymentsByAuthorEagerly(author string) []models.Payment {
	var payments []models.Payment
	r.db.Model(models.Payment{Author:author}).Preload("Product").Find(&payments)
	return  payments
}