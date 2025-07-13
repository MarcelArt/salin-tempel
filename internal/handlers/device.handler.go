package handlers

import (
	"os/user"

	"github.com/MarcelArt/salin-tempel/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/zcalusic/sysinfo"
)

// Get Device
// @Summary Get Device
// @Description Get Device
// @Tags Device
// @Accept json
// @Produce json
// @Success 200 {object} models.Device
// @Router / [get]
func GetDevice(c *fiber.Ctx) error {
	current, err := user.Current()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var si sysinfo.SysInfo
	si.GetSysInfo()

	device := models.Device{
		User:    current.Name,
		OS:      si.OS.Name,
		Product: si.Product.Name,
	}

	return c.Status(fiber.StatusOK).JSON(device)
}
