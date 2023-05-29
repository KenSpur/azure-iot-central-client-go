package iotcentral

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

/////////////////////////
// GetUser
/////////////////////////

// GetUser - Returns a specifc user
func (c *Client) GetUser(userID string) (*UserResponse, error) {
	user, err := c.getUser(userID)
	if err != nil {
		return nil, err
	}

	if user.Type != "email" {
		return nil, fmt.Errorf("user is not a user")
	}

	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Roles: user.Roles,
	}, nil
}

// GetADGroupUser - Returns a specifc AD Group user
func (c *Client) GetADGroupUser(userID string) (*ADGroupUserResponse, error) {
	user, err := c.getUser(userID)
	if err != nil {
		return nil, err
	}

	if user.Type != "adGroup" {
		return nil, fmt.Errorf("user is not an AD Group")
	}

	return &ADGroupUserResponse{
		ID:       user.ID,
		ObjectID: user.ObjectID,
		Roles:    user.Roles,
		TenantID: user.TenantID,
	}, nil
}

// GetServicePrincipalUser - Returns a specifc Service Principal user
func (c *Client) GetServicePrincipalUser(userID string) (*ServicePrincipalUserResponse, error) {
	user, err := c.getUser(userID)
	if err != nil {
		return nil, err
	}

	if user.Type != "servicePrincipal" {
		return nil, fmt.Errorf("user is not a Service Principal")
	}

	return &ServicePrincipalUserResponse{
		ID:       user.ID,
		ObjectID: user.ObjectID,
		Roles:    user.Roles,
		TenantID: user.TenantID,
	}, nil
}

func (c *Client) getUser(userID string) (*userResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/users/%s?api-version=2022-10-31-preview", c.HostURL, userID), nil)
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

	user := userResponse{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/////////////////////////
// CreateUser
/////////////////////////

// CreateUser - Returns a specifc user
func (c *Client) CreateUser(userReq UserRequest) (*UserResponse, error) {
	userID := uuid.New().String()

	user := userRequest{
		Email: userReq.Email,
		Roles: userReq.Roles,
		Type:  "email",
	}

	userResponse, err := c.createUser(userID, user)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:    userResponse.ID,
		Email: userResponse.Email,
		Roles: userResponse.Roles,
	}, nil
}

// CreateADGroupUser - Returns a specifc AD Group user
func (c *Client) CreateADGroupUser(userReq ADGroupUserRequest) (*ADGroupUserResponse, error) {
	userID := uuid.New().String()

	user := userRequest{
		ObjectID: userReq.ObjectID,
		Roles:    userReq.Roles,
		TenantID: userReq.TenantID,
		Type:     "adGroup",
	}

	userResponse, err := c.createUser(userID, user)
	if err != nil {
		return nil, err
	}

	return &ADGroupUserResponse{
		ID:       userResponse.ID,
		ObjectID: userResponse.ObjectID,
		Roles:    userResponse.Roles,
		TenantID: userResponse.TenantID,
	}, nil
}

// CreateServicePrincipalUser - Returns a specifc Service Principal user
func (c *Client) CreateServicePrincipalUser(userReq ServicePrincipalUserRequest) (*ServicePrincipalUserResponse, error) {
	userID := uuid.New().String()

	user := userRequest{
		ObjectID: userReq.ObjectID,
		Roles:    userReq.Roles,
		TenantID: userReq.TenantID,
		Type:     "servicePrincipal",
	}

	userResponse, err := c.createUser(userID, user)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipalUserResponse{
		ID:       userResponse.ID,
		ObjectID: userResponse.ObjectID,
		Roles:    userResponse.Roles,
		TenantID: userResponse.TenantID,
	}, nil
}

func (c *Client) createUser(userID string, user userRequest) (*userResponse, error) {
	d, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/users/%s?api-version=2022-10-31-preview", c.HostURL, userID), strings.NewReader(string(d)))
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

	userResponse := userResponse{}
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

/////////////////////////
// UpdateUser
/////////////////////////

// UpdateUser - Returns a specifc user
func (c *Client) UpdateUser(userID string, userReq UserRequest) (*UserResponse, error) {
	user := userRequest{
		Email: userReq.Email,
		Roles: userReq.Roles,
		Type:  "email",
	}

	userResponse, err := c.updateUser(userID, user)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:    userResponse.ID,
		Email: userResponse.Email,
		Roles: userResponse.Roles,
	}, nil
}

// UpdateADGroupUser - Returns a specifc AD Group user
func (c *Client) UpdateADGroupUser(userID string, userReq ADGroupUserRequest) (*ADGroupUserResponse, error) {
	user := userRequest{
		ObjectID: userReq.ObjectID,
		Roles:    userReq.Roles,
		TenantID: userReq.TenantID,
		Type:     "adGroup",
	}

	userResponse, err := c.updateUser(userID, user)
	if err != nil {
		return nil, err
	}

	return &ADGroupUserResponse{
		ID:       userResponse.ID,
		ObjectID: userResponse.ObjectID,
		Roles:    userResponse.Roles,
		TenantID: userResponse.TenantID,
	}, nil
}

// UpdateServicePrincipalUser - Returns a specifc Service Principal user
func (c *Client) UpdateServicePrincipalUser(userID string, userReq ServicePrincipalUserRequest) (*ServicePrincipalUserResponse, error) {
	user := userRequest{
		ObjectID: userReq.ObjectID,
		Roles:    userReq.Roles,
		TenantID: userReq.TenantID,
		Type:     "servicePrincipal",
	}

	userResponse, err := c.updateUser(userID, user)
	if err != nil {
		return nil, err
	}

	return &ServicePrincipalUserResponse{
		ID:       userResponse.ID,
		ObjectID: userResponse.ObjectID,
		Roles:    userResponse.Roles,
		TenantID: userResponse.TenantID,
	}, nil
}

func (c *Client) updateUser(userID string, user userRequest) (*userResponse, error) {
	d, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/users/%s?api-version=2022-10-31-preview", c.HostURL, userID), strings.NewReader(string(d)))
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

	userResponse := userResponse{}
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

/////////////////////////
// DeleteUser
/////////////////////////

// DeleteUser - Deletes a user
func (c *Client) DeleteUser(userID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/users/%s?api-version=2022-10-31-preview", c.HostURL, userID), nil)
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
