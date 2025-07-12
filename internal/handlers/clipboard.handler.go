package handlers

import (
	"fmt"
	"io"

	"github.com/MarcelArt/salin-tempel/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/skratchdot/open-golang/open"
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

// Create creates a new picture clipboard
// @Summary Create a new picture clipboard
// @Description Create a new picture clipboard
// @Tags Clipboard
// @Accept json
// @Produce json
// @Param file formData file true "Picture file"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Router /clipboard/img [post]
func CopyPicture(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		err = fmt.Errorf("invalid file: %w", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	file, err := fileHeader.Open()
	if err != nil {
		err = fmt.Errorf("invalid file: %w", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("invalid file: %w", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	clipboard.Write(clipboard.FmtImage, fileContent)
	return c.SendStatus(fiber.StatusCreated)
}

// Open link in a browser
// @Summary Open link in a browser
// @Description Open link in a browser
// @Tags Clipboard
// @Accept json
// @Produce json
// @Param Clipboard body models.CopyString true "Link"
// @Success 204 {object} string
// @Failure 400 {object} string
// @Router /clipboard/open [post]
func OpenLink(c *fiber.Ctx) error {
	var link models.CopyString
	if err := c.BodyParser(&link); err != nil {
		err = fmt.Errorf("invalid json: %w", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	open.Run(link.Text)

	return c.SendStatus(fiber.StatusNoContent)
}
