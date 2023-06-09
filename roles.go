package iotcentral

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetRoles - Returns a list of roles
func (c *Client) GetRoles() ([]RoleResponse, error) {
	url := fmt.Sprintf("%s/api/roles?api-version=2022-10-31-preview", c.HostURL)
	var allRoles []RoleResponse

	for url != "" {
		req, err := http.NewRequest("GET", url, nil)
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

		roleCollection := roleCollectionResponse{}
		err = json.Unmarshal(body, &roleCollection)
		if err != nil {
			return nil, err
		}

		allRoles = append(allRoles, roleCollection.Value...)

		// Update the URL to the next page if it's available, otherwise break the loop
		url = roleCollection.NextLink
	}

	return allRoles, nil
}

// GetRole - Returns a specifc role
func (c *Client) GetRole(roleID string) (*RoleResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/roles/%s?api-version=2022-10-31-preview", c.HostURL, roleID), nil)
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

	role := RoleResponse{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// GetRoleByName - Returns a specifc role
func (c *Client) GetRoleByName(name string) (*RoleResponse, error) {
	cleanName := cleanName(name)
	roles, err := c.GetRoles()
	if err != nil {
		return nil, err
	}

	roleResponse := RoleResponse{}
	for _, role := range roles {
		if role.DisplayName == cleanName {
			roleResponse = role
			break
		}
	}

	if roleResponse.ID == "" {
		return nil, fmt.Errorf("role not found for name: %s", name)
	}

	return &roleResponse, nil
}

func cleanName(name string) string {
	switch name {
	case "App Administrator", "App Builder", "App Operator":
		name = strings.TrimPrefix(name, "App ")
	case "Org Administrator":
		name = "Org Admin"
	}

	return name
}
