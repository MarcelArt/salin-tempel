//go:build linux
// +build linux

package device

import (
	"fmt"
	"os/user"

	"github.com/zcalusic/sysinfo"
)

func GetDeviceInfo() (Device, error) {
	current, err := user.Current()
	if err != nil {
		err = fmt.Errorf("error getting user info: %w", err)
		return Device{}, err
	}

	var si sysinfo.SysInfo
	si.GetSysInfo()

	device := Device{
		User:    current.Name,
		OS:      si.OS.Name,
		Product: si.Product.Name,
		Type:    "unix",
	}
	return device, nil
}
