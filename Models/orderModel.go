package models

import "time"

type Order struct {
	Id             int       `json: "id" gorm:"type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey"`
	CashierID      int       `json: "cashierId"`
	PaymentTypesId int       `json: "paymentTypesId"`
	TotalPrice     int       `json: "totalPrice"`
	TotalPaid      int       `json: "totalPaid"`
	TotalReturn    int       `json: "totalReturn"`
	ReceiptId      string    `json: "receiptId"`
	IsDownload     int       `json: "isDownload"`
	ProductId      string    `json: "productId"`
	Quantities     string    `json: "quantities"`
	CreatedAT      time.Time `json: "created_at"`
	UpdatedAt      time.Time `json: "updated_at"`
}
