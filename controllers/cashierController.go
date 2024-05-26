package controllers

import (
	"os"
	models "sales_api_go/go_sales_app/Models"
	db "sales_api_go/go_sales_app/config"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Login
func Login(c *fiber.Ctx) error {
	cashierId := c.Params("id")
	if cashierId == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier ID is missing",
			})
	}

	var body_data map[string]string
	err := c.BodyParser(&body_data)
	if err != nil {
		return err
	}

	var cashier models.Cashier
	db.DB.Select("*").Where("id = ?", cashierId).First(&cashier)

	if body_data["passcode"] != cashier.Passcode {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Password is not correct!",
			})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(cashier.Id)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": err,
			})
	}
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Token is generated",
			"data":    tokenString,
		})
}

func CreateCashier(c *fiber.Ctx) error {

	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid body",
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Name should be provided",
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Pass code should be provided",
			})
	}

	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAT: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Created successfully",
			"data":    cashier,
		})

}

func CashierList(c *fiber.Ctx) error {
	var cashier []models.Cashier
	// limit, _ := strconv.Atoi(c.Query("limit"))
	// skip, _ := strconv.Atoi(c.Query("skip"))

	var count int64

	db.DB.Select("*").Find(&cashier).Count(&count)
	return c.Status(200).JSON(
		fiber.Map{
			"success":  true,
			"message":  "Returned successfully",
			"data":     cashier,
			"returned": count,
		})
}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("id")

	//No need for this as in case of not passing an id, it gonna route to create cashier and throughs a diff error "method not allowed" because you are calling it with GET
	if cashierId == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier ID is missing",
			})
	}

	var cashier models.Cashier
	db.DB.Select("*").Where("id = ?", cashierId).First(&cashier)
	//Don't forget to unpack it if you wanna change it's fields names
	cashierData := make(map[string]interface{})
	cashierData["cashier_id"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["created_at"] = cashier.CreatedAT
	cashierData["updated_at"] = cashier.UpdatedAt

	//If you tried to pass any cashier id even if it's not exist: it gonna get you a success msg with cashier id =0, so in order for us to handle wrong cashier ids
	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			})
	}
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Returned successfully",
			"data":    cashier,
		})
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("id")
	var updatedCashier models.Cashier
	err := c.BodyParser(&updatedCashier)
	if err != nil {
		return err
	}

	var cashier models.Cashier
	db.DB.Select("*").Where("id = ?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Not found",
			})
	}
	if updatedCashier.Name != "" {
		cashier.Name = updatedCashier.Name
		cashier.UpdatedAt = time.Time{}
	}

	if updatedCashier.Passcode != "" {
		cashier.Passcode = updatedCashier.Passcode
		cashier.UpdatedAt = time.Time{}

	}

	db.DB.Save(&cashier)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Updated successfully",
			"data":    cashier,
		})

}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("id")

	var cashier models.Cashier
	db.DB.Select("*").Where("id = ?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Not found",
			})
	}

	db.DB.Delete(&cashier)
	//another way: db.DB.Select("*").Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Deleted successfully",
		})
}
