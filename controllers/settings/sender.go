package settings

import (
	"bulk-sms/database"
	"bulk-sms/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateSender(c *fiber.Ctx) error {
	var model models.SenderID

	if err := c.BodyParser(&model); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	if model.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Email  cannot be empty",
		})
	}

	if model.Description == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Phone  cannot be empty",
		})
	}
	result := database.DB.Create(&model)
	if result.RowsAffected != 1 {

		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "Record Not Saved",
		})

	}
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Saved",
	})

}
func AllSenders(c *fiber.Ctx) error {
	var model []models.SenderID
	database.DB.Find(&model)
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}

func GetSender(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var model models.SenderID

	database.DB.Find(&model, "title = ?", id)
	if model.Title == "" {
		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "No Record Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}
func UpdateSender(c *fiber.Ctx) error {
	var body models.SenderID

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	var model models.SenderID

	Id := c.Params("id")
	// CHECK AVAILABLE CONTACT
	err := database.DB.First(&model, "title = ?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":  0,
		})
	}

	errUpdate := database.DB.Model(&model).Updates(body).Error //save new changes

	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to update record",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Record Updated",
		"Data":    model,
	})
}

func DeleteSender(c *fiber.Ctx) error {
	Id := c.Params("id")
	var model models.SenderID

	//	CHECK AVAILABLE SenderId
	err := database.DB.Debug().First(&model, "title=?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":  0,
		})
	}

	errDelete := database.DB.Debug().Delete(&model).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to update record",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Record Removed",
		"Status":  1,
	})
}
