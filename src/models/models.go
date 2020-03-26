package models

import "github.com/jinzhu/gorm"

type SimpleResponse struct {
	Message string `json:"message"`
	Date    string `json:"date, omitempty"`
}

type PaymentResponse struct {
	Message string `json:"message"`
	Payments []Payment `json:"payments"`
}


type Payment struct {
	gorm.Model
	Author    string	//`json:"author"`
	Product   Product	`gorm:"foreignkey:ProductName"`
	Sum		  string	//`json:"sum"`
}

func (Payment) TableName() string {
	return "payments"
}

type Product struct {
	gorm.Model
	ProductName string //`json:"name"`
	Type        string //`json:"type"`
}

func (Product) TableName() string {
	return "products"
}