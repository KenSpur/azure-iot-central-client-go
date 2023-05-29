package iotcentral

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserFunctions(t *testing.T) {
	token := os.Getenv("IOTCENTRAL_API_TOKEN")
	host := os.Getenv("IOTCENTRAL_HOST")
	email := strings.ToLower(os.Getenv("IOTCENTRAL_USER_EMAIL"))

	t.Logf("Host: %s", host)

	client, err := NewClient(&host, &token)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	userReq := UserRequest{
		Email: email,
		Roles: []RoleAssignment{
			{
				Role: "ca310b8d-2f4a-44e0-a36e-957c202cd8d4",
			},
		},
	}

	// CreateUser
	userResponse, err := client.CreateUser(userReq)
	assert.NoError(t, err, "Failed to create user one")

	var roleIds []string
	for _, role := range userResponse.Roles {
		roleIds = append(roleIds, role.Role)
	}

	assert.Equal(t, email, userResponse.Email, "User email does not match")
	assert.Contains(t, roleIds, userReq.Roles[0].Role, "User roles does not contain role one")

	// GetUsers
	userResponse, err = client.GetUser(userResponse.ID)
	assert.NoError(t, err, "Failed to get users")

	var gottenRoleIds []string
	for _, role := range userResponse.Roles {
		gottenRoleIds = append(gottenRoleIds, role.Role)
	}

	assert.Equal(t, email, userResponse.Email, "User email does not match")
	assert.Contains(t, gottenRoleIds, userReq.Roles[0].Role, "User roles does not contain role one")

	// UpdateUser

	orgReqID := "usertestorg"
	organizationRequest := OrganizationRequest{
		DisplayName: "User Test Org",
	}

	orgResponse, _ := client.CreateOrganization(orgReqID, organizationRequest)

	userReq.Roles = append(userReq.Roles, RoleAssignment{
		Role:         "c495eb57-eb18-489e-9802-62c474e5645c",
		Organization: orgResponse.ID,
	})

	userResponse, err = client.UpdateUser(userResponse.ID, userReq)
	assert.NoError(t, err, "Failed to update user one")

	var updatedRoleIds []string
	for _, role := range userResponse.Roles {
		updatedRoleIds = append(updatedRoleIds, role.Role)
	}

	assert.Equal(t, email, userResponse.Email, "User email does not match")
	assert.Contains(t, updatedRoleIds, userReq.Roles[0].Role, "User roles does not contain role one")
	assert.Contains(t, updatedRoleIds, userReq.Roles[1].Role, "User roles does not contain role two")

	_ = client.DeleteOrganization(orgResponse.ID)

	// DeleteUser
	err = client.DeleteUser(userResponse.ID)
	assert.NoError(t, err, "Failed to delete user one")
}
