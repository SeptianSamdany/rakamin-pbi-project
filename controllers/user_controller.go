package controllers

import (
	"net/http"
	"rakamin-project/services"
	"rakamin-project/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

// Register handler
func (uc *UserController) Register(c *fiber.Ctx) error {
	var userInput models.User

	// Parse input
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi input sederhana
	if userInput.Email == "" || userInput.KataSandi == "" || userInput.Nama == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Nama, Email, dan Kata Sandi harus diisi"})
	}

	// Call service
	user, err := uc.UserService.Register(userInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Login handler
func (uc *UserController) Login(c *fiber.Ctx) error {
	var loginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse input
	if err := c.BodyParser(&loginInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi input
	if loginInput.Email == "" || loginInput.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Email dan Password harus diisi"})
	}

	// Call service
	token, err := uc.UserService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}

// Update User handler
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint) // Pastikan user_id valid
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user authentication"})
	}

	var userInput models.User
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Call service
	updatedUser, err := uc.UserService.UpdateUser(userID, userInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(updatedUser)
}

// Delete User handler
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint) // Pastikan user_id valid
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user authentication"})
	}

	// Call service
	err := uc.UserService.DeleteUser(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}