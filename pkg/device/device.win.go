//go:build windows
// +build windows

package device

import (
	"fmt"
	"os/user"
)

func GetDeviceInfo() (Device, error) {
	current, err := user.Current()
	if err != nil {
		err = fmt.Errorf("error getting user info: %w", err)
		return Device{}, err
	}

	device := Device{
		User:    current.Name,
		OS:      "Windows",
		Product: "Unknown",
	}
	return device, nil
}
