package belajarHandler

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/internal/model"
)

func GetUsers(c *fiber.Ctx) error {
	var users []model.Pasien
	// Find all users in database
	result := database.DB.Find(&users)
	// Check for errors during query execution
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return list of users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data User Berhasil Ditampilkan!",
		"data":    users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	// Parse request body
	var user model.Pasien
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	// Insert new user into database
	result := database.DB.Create(&user)
	// Check for errors during insertion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Berhasil Ditambahkan!",
		"data":    user,
	})
}
func GetUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")
	// Find user by id_user in database
	var user model.Pasien
	result := database.DB.First(&user, id)
	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan!",
		})
	}
	// Check for errors during query
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return user
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    user,
	})
}
func UpdateUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")
	// Find user by id_user in database
	var user model.Pasien
	result := database.DB.First(&user, id)
	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan",
		})
	}
	// Parse request body
	var newUser model.Pasien
	if err := c.BodyParser(&newUser); err != nil {
		return err
	}
	// Update user in database
	result = database.DB.Model(&user).Updates(newUser)
	// Check for errors during update
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Berhasil Diperbarui!",
		"data":    user,
	})
}
func DeleteUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")
	// Find user by id_user in database
	var user model.Pasien
	result := database.DB.First(&user, id)
	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan",
		})
	}
	// Delete user from database
	result = database.DB.Delete(&user)
	// Check for errors during deletion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Berhasil Dihapus!",
		"data":    user,
	})
}
