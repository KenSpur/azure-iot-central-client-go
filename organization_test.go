package iotcentral

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationFunctions(t *testing.T) {
	token := os.Getenv("IOTCENTRAL_API_TOKEN")
	host := os.Getenv("IOTCENTRAL_HOST")

	t.Logf("Host: %s", host)

	client, err := NewClient(&host, &token)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	organizationOneID := "example-organization-1"
	organizationTwoID := "example-organization-2"

	organizationOneRequest := OrganizationRequest{
		DisplayName: "Example Organization 1",
	}

	organizationTwoRequest := OrganizationRequest{
		DisplayName: "Example Organization 2",
		Parent:      organizationOneID,
	}

	// CreateOrganization
	organizationOne, err := client.CreateOrganization(organizationOneID, organizationOneRequest)
	assert.NoError(t, err, "Failed to create Organization one")

	organizationTwo, err := client.CreateOrganization(organizationTwoID, organizationTwoRequest)
	assert.NoError(t, err, "Failed to create Organization two")

	assert.Equal(t, organizationOne.DisplayName, organizationOneRequest.DisplayName, "Organization one display name does not match")
	assert.Equal(t, organizationTwo.DisplayName, organizationTwoRequest.DisplayName, "Organization two display name does not match")
	assert.Equal(t, organizationTwo.Parent, organizationTwoRequest.Parent, "Organization two simulated does not match")

	// GetOrganizations
	organizations, err := client.GetOrganizations()
	assert.NoError(t, err, "Failed to get Organizations")

	var displayNames []string
	for _, organization := range organizations {
		displayNames = append(displayNames, organization.DisplayName)
	}

	assert.Contains(t, displayNames, organizationOne.DisplayName, "Display names does not contain Organization one")
	assert.Contains(t, displayNames, organizationTwo.DisplayName, "Display names does not contain Organization two")

	// GetOrganization
	organizationOne, err = client.GetOrganization(organizationOneID)
	assert.NoError(t, err, "Failed to get Organization one")

	organizationTwo, err = client.GetOrganization(organizationTwoID)
	assert.NoError(t, err, "Failed to get Organization two")

	assert.Equal(t, organizationOne.DisplayName, organizationOneRequest.DisplayName, "Organization one display name does not match")
	assert.Equal(t, organizationTwo.DisplayName, organizationTwoRequest.DisplayName, "Organization two display name does not match")

	// UpdateOrganization
	organizationOneRequest.DisplayName = "Example Organization 1 Updated"

	organizationOne, err = client.UpdateOrganization(organizationOneID, organizationOneRequest)
	assert.NoError(t, err, "Failed to update Organization one")

	assert.Equal(t, organizationOne.DisplayName, organizationOneRequest.DisplayName, "Organization one display name does not match")

	// DeleteOrganization
	err = client.DeleteOrganization(organizationTwoID)
	assert.NoError(t, err, "Failed to delete Organization two")

	err = client.DeleteOrganization(organizationOneID)
	assert.NoError(t, err, "Failed to delete Organization one")

	organizations, _ = client.GetOrganizations()

	var newDisplayNames []string
	for _, organization := range organizations {
		newDisplayNames = append(newDisplayNames, organization.DisplayName)
	}

	assert.NotContains(t, newDisplayNames, organizationOne.DisplayName, "Display names contains Organization one")
	assert.NotContains(t, newDisplayNames, organizationTwo.DisplayName, "Display names contains Organization two")
}
