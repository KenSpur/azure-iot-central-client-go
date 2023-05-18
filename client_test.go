package iotcentral

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	token := os.Getenv("TOKEN")
	host := os.Getenv("HOST")

	t.Logf("Host: %s", host)

	client, err := NewClient(&host, &token)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	devices, err := client.GetDevices()
	if err != nil {
		t.Fatalf("Failed to get devices: %v", err)
	}

	t.Logf("Found %d devices.", len(devices))

	for index, device := range devices {
		t.Logf("Device %d: %s", index, device.DisplayName)
	}
}
