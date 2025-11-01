package handlers

import (
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/models"
	"github.com/Shamim-cse-pust/REST-API-Project-with-GO/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Global validator instance
var validate = validator.New()

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handles POST /api/v1/users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	// Parse JSON body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JSON format",
			"message": "Please check your request body format",
		})
	}

	// Validate request using struct tags
	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"message": err.Error(),
		})
	}

	// Create user using service
	user, err := h.userService.CreateUser(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
	}

	// Return success response (without password)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user.ToResponse(),
	})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	// Get user ID from URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid number",
		})
	}

	// Get user using service
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "User not found",
			"message": err.Error(),
		})
	}

	// Return success response (without password)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
		"data":    user.ToResponse(),
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	// Get user ID from URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid number",
		})
	}

	var req models.UpdateUserRequest

	// Parse JSON body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid JSON format",
			"message": "Please check your request body format",
		})
	}

	// Validate request using struct tags
	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"message": err.Error(),
		})
	}

	// Update user using service
	user, err := h.userService.UpdateUser(id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update user",
			"message": err.Error(),
		})
	}

	// Return success response (without password)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    user.ToResponse(),
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	// Get user ID from URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid user ID",
			"message": "User ID must be a valid number",
		})
	}

	// Delete user using service
	err = h.userService.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "User not found",
			"message": err.Error(),
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
