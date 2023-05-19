package iotcentral

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetOrganizations - Returns a list of organizations
func (c *Client) GetOrganizations() ([]OrganizationResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/organizations?api-version=2022-10-31-preview", c.HostURL), nil)
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

	organizationCollection := organizationCollectionResponse{}
	err = json.Unmarshal(body, &organizationCollection)
	if err != nil {
		return nil, err
	}

	organizations := organizationCollection.Value

	return organizations, nil
}

// GetOrganization - Returns a specifc organization
func (c *Client) GetOrganization(organizationID string) (*OrganizationResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/organizations/%s?api-version=2022-10-31-preview", c.HostURL, organizationID), nil)
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

	organization := OrganizationResponse{}
	err = json.Unmarshal(body, &organization)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// CreateOrganization - Create new organization
func (c *Client) CreateOrganization(organization OrganizationRequest) (*OrganizationResponse, error) {
	o, err := json.Marshal(organization)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/organizations?api-version=2022-10-31-preview", c.HostURL), strings.NewReader(string(o)))
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

	organizationResponse := OrganizationResponse{}
	err = json.Unmarshal(body, &organizationResponse)
	if err != nil {
		return nil, err
	}

	return &organizationResponse, nil
}

// UpdateOrganization - Updates an organization
func (c *Client) UpdateOrganization(organizationID string, organization OrganizationRequest) (*OrganizationResponse, error) {
	d, err := json.Marshal(organization)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/organizations/%s?api-version=2022-10-31-preview", c.HostURL, organizationID), strings.NewReader(string(d)))
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

	organizationResponse := OrganizationResponse{}
	err = json.Unmarshal(body, &organizationResponse)
	if err != nil {
		return nil, err
	}

	return &organizationResponse, nil
}

// DeleteOrganization - Deletes an organization
func (c *Client) DeleteOrganization(organizationID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/organizations/%s?api-version=2022-10-31-preview", c.HostURL, organizationID), nil)
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
