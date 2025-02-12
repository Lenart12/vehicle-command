package ble

import (
	"strings"

	"tinygo.org/x/bluetooth"
)

func IsAdapterError(err error) bool {
	// D-Bus not found
	if strings.Contains(err.Error(), "dbus") && strings.HasSuffix(err.Error(), "no such file or directory") {
		return true
	}
	// D-Bus is running but org.bluez is not found
	if strings.Contains(err.Error(), "The name org.bluez was not provided by any .service files") {
		return true
	}
	return false
}

func AdapterErrorHelpMessage(err error) string {
	return "Failed to initialize BLE adapter: \n\t" + err.Error() + "\n" +
		"Make sure bluez and dbus are installed and running.\n" +
		"If running in a container, make sure the container has access to the host's D-Bus socket. (e.g. -v /var/run/dbus:/var/run/dbus)"
}

func newAdapter(id string) *bluetooth.Adapter {
	if id != "" {
		return bluetooth.NewAdapter(id)
	}

	return bluetooth.DefaultAdapter
}
