package iotcentral

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetDevices - Returns a list of devices
func (c *Client) GetDevices() ([]Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/devices?api-version=2022-10-31-preview", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	deviceCollection := deviceCollection{}
	err = json.Unmarshal(body, &deviceCollection)
	if err != nil {
		return nil, err
	}

	devices := deviceCollection.Value

	return devices, nil
}
