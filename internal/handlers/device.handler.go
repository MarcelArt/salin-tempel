package handlers

import (
	"fmt"

	"github.com/MarcelArt/salin-tempel/pkg/device"
	"github.com/gofiber/fiber/v2"
)

// Get Device
// @Summary Get Device
// @Description Get Device
// @Tags Device
// @Accept json
// @Produce json
// @Success 200 {object} device.Device
// @Router / [get]
func GetDevice(c *fiber.Ctx) error {
	d, err := device.GetDeviceInfo()
	if err != nil {
		err = fmt.Errorf("failed getting device info: %w", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(d)
}
