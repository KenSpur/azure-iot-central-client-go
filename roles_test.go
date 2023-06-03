package iotcentral

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleFunctions(t *testing.T) {
	token := os.Getenv("IOTCENTRAL_API_TOKEN")
	host := os.Getenv("IOTCENTRAL_HOST")

	t.Logf("Host: %s", host)

	client, err := NewClient(&host, &token)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// GetRoles
	roles, err := client.GetRoles()
	assert.NoError(t, err, "Failed to get roles")

	assert.NotEmpty(t, roles, "Roles is empty")

	// GetRole
	roleID := "ca310b8d-2f4a-44e0-a36e-957c202cd8d4"
	role, err := client.GetRole(roleID)
	assert.NoError(t, err, "Failed to get role")

	assert.Equal(t, roleID, role.ID, "Role ID does not match")
	assert.Equal(t, "Administrator", role.DisplayName, "Role DisplayName does not match")

	// GetRoleByName
	roleName := "App Builder"
	role, err = client.GetRoleByName(roleName)

	assert.NoError(t, err, "Failed to get role by name")

	assert.Equal(t, "Builder", role.DisplayName, "Role DisplayName does not match")
	assert.Equal(t, "344138e9-8de4-4497-8c54-5237e96d6aaf", role.ID, "Role ID does not match")

	roleName = "Org Administrator"
	role, err = client.GetRoleByName(roleName)

	assert.NoError(t, err, "Failed to get role by name")

	assert.Equal(t, "Org Admin", role.DisplayName, "Role DisplayName does not match")
	assert.Equal(t, "c495eb57-eb18-489e-9802-62c474e5645c", role.ID, "Role ID does not match")
}
