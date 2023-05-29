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
	assert.NoError(t, err, "Failed to get role one")

	assert.Equal(t, role.ID, roleID, "Role ID does not match")
	assert.Equal(t, role.DisplayName, "Administrator", "Role DisplayName does not match (Administrator)")
}
