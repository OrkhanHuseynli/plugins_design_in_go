package dbplugin

import (
	"github.com/jinzhu/gorm"
	"plugins_design_in_go/src/models"
)

type MockRepository struct {
	db         *gorm.DB
	MockObject models.Payment
}

func  (r *MockRepository) GetDB() *gorm.DB {
	return r.db
}

//func  (r *MockRepository) setDB(db  *gorm.DB) {
//	r.db = db
//}

func (r *MockRepository) AddPayment(payment *models.Payment)  {
	r.db.Create(payment)
}

func (r *MockRepository) GetPaymentsByAuthor(author string)  []models.Payment {
	ps :=  []models.Payment{r.MockObject}
	return ps
}

func (r *MockRepository) DeletePaymentsByAuthor(author string)  bool {
	return true
}

func (r *MockRepository) GetPaymentsByAuthorEagerly(author string) []models.Payment {
	ps :=  []models.Payment{r.MockObject}
	return  ps
}