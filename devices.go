package iotcentral

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetDevices - Returns a list of devices
func (c *Client) GetDevices() ([]DeviceResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/devices?api-version=2022-10-31-preview", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", statusCode, body)
	}

	deviceCollection := deviceCollectionResponse{}
	err = json.Unmarshal(body, &deviceCollection)
	if err != nil {
		return nil, err
	}

	devices := deviceCollection.Value

	return devices, nil
}

// GetDevice - Returns a specifc device
func (c *Client) GetDevice(deviceID string) (*DeviceResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/devices/%s?api-version=2022-10-31-preview", c.HostURL, deviceID), nil)
	if err != nil {
		return nil, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", statusCode, body)
	}

	device := DeviceResponse{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

// CreateDevice - Create new device
func (c *Client) CreateDevice(deviceID string, device DeviceRequest) (*DeviceResponse, error) {
	d, err := json.Marshal(device)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/devices/%s?api-version=2022-10-31-preview", c.HostURL, deviceID), strings.NewReader(string(d)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", statusCode, body)
	}

	deviceResponse := DeviceResponse{}
	err = json.Unmarshal(body, &deviceResponse)
	if err != nil {
		return nil, err
	}

	return &deviceResponse, nil
}

// UpdateDevice - Updates a device
func (c *Client) UpdateDevice(deviceID string, device DeviceRequest) (*DeviceResponse, error) {
	d, err := json.Marshal(device)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/devices/%s?api-version=2022-10-31-preview", c.HostURL, deviceID), strings.NewReader(string(d)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", statusCode, body)
	}

	deviceResponse := DeviceResponse{}
	err = json.Unmarshal(body, &deviceResponse)
	if err != nil {
		return nil, err
	}

	return &deviceResponse, nil
}

// DeleteDevice - Deletes a device
func (c *Client) DeleteDevice(deviceID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/devices/%s?api-version=2022-10-31-preview", c.HostURL, deviceID), nil)
	if err != nil {
		return err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if statusCode != http.StatusNoContent {
		return fmt.Errorf("status: %d, body: %s", statusCode, body)
	}

	return nil
}
