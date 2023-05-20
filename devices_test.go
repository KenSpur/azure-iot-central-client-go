package iotcentral

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceFunctions(t *testing.T) {
	token := os.Getenv("IOTCENTRAL_API_TOKEN")
	host := os.Getenv("IOTCENTRAL_HOST")

	t.Logf("Host: %s", host)

	client, err := NewClient(&host, &token)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	deviceOneID := "example-device-1"
	deviceTwoID := "example-device-2"

	deviceOneRequest := DeviceRequest{
		DisplayName: "Example Device 1",
	}

	deviceTwoRequest := DeviceRequest{
		DisplayName: "Example Device 2",
		Simulated:   true,
	}

	// CreateDevice
	deviceOne, err := client.CreateDevice(deviceOneID, deviceOneRequest)
	assert.NoError(t, err, "Failed to create device one")

	deviceTwo, err := client.CreateDevice(deviceTwoID, deviceTwoRequest)
	assert.NoError(t, err, "Failed to create device two")

	assert.Equal(t, deviceOne.DisplayName, deviceOneRequest.DisplayName, "Device one display name does not match")
	assert.Equal(t, deviceTwo.DisplayName, deviceTwoRequest.DisplayName, "Device two display name does not match")
	assert.Equal(t, deviceTwo.Simulated, deviceTwoRequest.Simulated, "Device two simulated does not match")

	// GetDevices
	devices, err := client.GetDevices()
	assert.NoError(t, err, "Failed to get devices")

	var displayNames []string
	for _, device := range devices {
		displayNames = append(displayNames, device.DisplayName)
	}

	assert.Contains(t, displayNames, deviceOne.DisplayName, "Display names does not contain device one")
	assert.Contains(t, displayNames, deviceTwo.DisplayName, "Display names does not contain device two")

	// GetDevice
	deviceOne, err = client.GetDevice(deviceOneID)
	assert.NoError(t, err, "Failed to get device one")

	deviceTwo, err = client.GetDevice(deviceTwoID)
	assert.NoError(t, err, "Failed to get device two")

	assert.Equal(t, deviceOne.DisplayName, deviceOneRequest.DisplayName, "Device one display name does not match")
	assert.Equal(t, deviceTwo.DisplayName, deviceTwoRequest.DisplayName, "Device two display name does not match")

	// UpdateDevice
	deviceOneRequest.DisplayName = "Example Device 1 Updated"

	deviceOne, err = client.UpdateDevice(deviceOneID, deviceOneRequest)
	assert.NoError(t, err, "Failed to update device one")

	assert.Equal(t, deviceOne.DisplayName, deviceOneRequest.DisplayName, "Device one display name does not match")

	// DeleteDevice
	err = client.DeleteDevice(deviceOneID)
	assert.NoError(t, err, "Failed to delete device one")

	err = client.DeleteDevice(deviceTwoID)
	assert.NoError(t, err, "Failed to delete device two")

	devices, _ = client.GetDevices()

	var newDisplayNames []string
	for _, device := range devices {
		newDisplayNames = append(newDisplayNames, device.DisplayName)
	}

	assert.NotContains(t, newDisplayNames, deviceOne.DisplayName, "Display names contains device one")
	assert.NotContains(t, newDisplayNames, deviceTwo.DisplayName, "Display names contains device two")
}
