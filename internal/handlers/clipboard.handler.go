package handlers

import (
	"fmt"

	"github.com/MarcelArt/salin-tempel/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.design/x/clipboard"
)

// Create creates a new clipboard
// @Summary Create a new clipboard
// @Description Create a new clipboard
// @Tags Clipboard
// @Accept json
// @Produce json
// @Param Clipboard body models.CopyString true "Clipboard data"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Router /clipboard [post]
func Copy(c *fiber.Ctx) error {
	var copyString models.CopyString
	if err := c.BodyParser(&copyString); err != nil {
		err = fmt.Errorf("invalid json: %w", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	clipboard.Write(clipboard.FmtText, []byte(copyString.Text))

	return c.SendStatus(fiber.StatusCreated)
}
