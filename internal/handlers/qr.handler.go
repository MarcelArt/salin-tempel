package handlers

import (
	"fmt"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// Get Device QR Code
// @Summary Get Device QR Code
// @Description Get Device QR Code
// @Tags QRCode
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /qr [get]
func QRCode(c *fiber.Ctx) error {
	ip := getLocalIP()

	url := fmt.Sprintf("%s://%s:3000", c.Protocol(), ip)
	log.Println(url)

	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		err = fmt.Errorf("failed to generate QR code: %w", err)
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "image/png")
	return c.Status(fiber.StatusOK).Send(png)
}
