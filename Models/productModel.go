package models

import "time"

type Product struct {
	Id               int       `json: "id" gorm:"type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT"`
	Sku              string    `json: "sku"`
	Name             string    `json: "name"`
	Stock            int       `json: "stock"`
	Price            int       `json: "price"`
	Image            string    `json: "image"`
	TotalFinalPrice  int       `json: "total_final_price"`
	TotalNormalPrice int       `json: "total_Normal_price"`
	CategoryId       int       `json: "categoryId"`
	DiscountId       int       `json: "discountId"`
	CreatedAT        time.Time `json: "created_at"`
	UpdatedAt        time.Time `json: "updated_at"`
}
