package models

type SimpleResponse struct {
	Message string `json:"message"`
	Date    string `json:"date, omitempty"`
}

type PaymentResponse struct {
	Message string `json:"message"`
	Payments []Payment `json:"payments"`
}


type Payment struct {
	Author    string	`json:"author"`
	ProductName string	`json:"productName"`
	Product   Product	`gorm:"foreignkey:ProductName;association_foreignkey:Name"`
	Sum		  string	`json:"sum"`
}

func (Payment) TableName() string {
	return "payments"
}

type Product struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (Product) TableName() string {
	return "products"
}