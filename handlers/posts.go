package handlers

import (
	"blog/database"
	"blog/models"
	"blog/requests"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllPosts(c *fiber.Ctx) error {
	limitParam := c.Params("limit")
	offsetParam := c.Params("offset")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid limit",
		})
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid offset",
		})
	}

	var posts []*models.Post

	if err := database.DB.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving posts",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(posts)
}

var validate = validator.New()

func CreatePosts(c *fiber.Ctx) error {
	var req requests.CreatePostRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		errors := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			var msg string
			switch e.Tag() {
			case "required":
				msg = e.Field() + " wajib diisi"
			case "min":
				msg = e.Field() + " minimal " + e.Param() + " karakter"
			case "oneof":
				msg = e.Field() + " harus salah satu dari: " + e.Param()
			default:
				msg = e.Field() + " tidak valid"
			}
			errors = append(errors, msg)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   errors,
			"message": "failed to validate",
		})
	}

	newPost := models.Post{
		Title:       req.Title,
		Content:     req.Content,
		Category:    req.Category,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		Status:      req.Status,
	}

	database.DB.Create(&newPost)

	return c.Status(fiber.StatusOK).JSON([]string{})
}

func GetPostById(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var post []*models.Post

	result := database.DB.First(&post, id)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Article not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var req requests.CreatePostRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		errors := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			var msg string
			switch e.Tag() {
			case "required":
				msg = e.Field() + " wajib diisi"
			case "min":
				msg = e.Field() + " minimal " + e.Param() + " karakter"
			case "oneof":
				msg = e.Field() + " harus salah satu dari: " + e.Param()
			default:
				msg = e.Field() + " tidak valid"
			}
			errors = append(errors, msg)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   errors,
			"message": "failed to validate",
		})
	}

	result := database.DB.Model(&models.Post{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":        req.Title,
			"content":      req.Content,
			"category":     req.Category,
			"updated_date": time.Now(),
			"status":       req.Status,
		})

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Article not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succes update Article",
	})
}

func DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	result := database.DB.Where("id = ?", id).Delete(&models.Post{})

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Article not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete post successfully",
	})
}
