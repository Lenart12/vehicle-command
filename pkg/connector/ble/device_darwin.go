package ble

import (
	"github.com/teslamotors/vehicle-command/internal/log"
	"tinygo.org/x/bluetooth"
)

func IsAdapterError(err error) bool {
	// TODO: Add check for Darwin
	return false
}

func AdapterErrorHelpMessage(err error) string {
	return err.Error()
}

func newAdapter(id string) *bluetooth.Adapter {
	if id != "" {
		// TODO: Add support for Darwin
		log.Warning("BLE adapter ID is not supported on Darwin")
	}

	return bluetooth.DefaultAdapter
}

func (c *Connection) Close() {
	if err := c.device.Disconnect(); err != nil {
		log.Warning("ble: failed to disconnect: %s", err)
	}
}
