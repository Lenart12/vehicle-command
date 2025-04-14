package ble

import (
	"github.com/teslamotors/vehicle-command/internal/log"
	"tinygo.org/x/bluetooth"
)

func IsAdapterError(_ error) bool {
	// TODO: Add check for Darwin
	return false
}

func AdapterErrorHelpMessage(err error) string {
	return err.Error()
}

func newAdapter(id *string) (*bluetooth.Adapter, error) {
	if id != nil && *id != "" {
		log.Warning("Darwin does not support specifying a Bluetooth adapter ID")
		return nil, ErrAdapterInvalidID
	}

	return bluetooth.DefaultAdapter, nil
}

var (
	deviceCharacteristicWrite = bluetooth.DeviceCharacteristic.Write
)

func (c *Connection) Close() {
	if err := c.device.Disconnect(); err != nil {
		log.Warning("ble: failed to disconnect: %s", err)
	}
}
