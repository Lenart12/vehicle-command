package ble

import (
	"errors"

	"github.com/go-ble/ble"
)

func newDevice(ble.Addr) (ble.Device, error) {
	return nil, errors.New("not supported on Windows")
}
