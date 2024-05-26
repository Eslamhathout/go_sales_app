package models

import "time"

type Category struct {
	Id        uint      `json: "id" gorm:"type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey"go`
	Name      string    `json: "name"`
	CreatedAT time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}
