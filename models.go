package iotcentral

/////////////////////////
// Device
/////////////////////////

type deviceCollectionResponse struct {
	Value    []DeviceResponse `json:"value"`
	NextLink string           `json:"nextLink,omitempty"`
}

// DeviceResponse -
type DeviceResponse struct {
	ID            string   `json:"id"`
	Etag          string   `json:"etag"`
	DisplayName   string   `json:"displayName"`
	Template      string   `json:"template,omitempty"`
	Simulated     bool     `json:"simulated"`
	Provisioned   bool     `json:"provisioned"`
	Enabled       bool     `json:"enabled"`
	Organizations []string `json:"organizations,omitempty"`
}

// DeviceRequest -
type DeviceRequest struct {
	Etag          string   `json:"etag,omitempty"`
	DisplayName   string   `json:"displayName"`
	Template      string   `json:"template,omitempty"`
	Simulated     bool     `json:"simulated,omitempty"`
	Enabled       bool     `json:"enabled,omitempty"`
	Organizations []string `json:"organizations,omitempty"`
}

/////////////////////////
// Organization
/////////////////////////

type organizationCollectionResponse struct {
	Value    []OrganizationResponse `json:"value"`
	NextLink string                 `json:"nextLink,omitempty"`
}

// OrganizationResponse -
type OrganizationResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Parent      string `json:"parent,omitempty"`
}

// OrganizationRequest -
type OrganizationRequest struct {
	DisplayName string `json:"displayName,omitempty"`
	Parent      string `json:"parent,omitempty"`
}

/////////////////////////
// Role
/////////////////////////

type roleCollectionResponse struct {
	Value    []RoleResponse `json:"value"`
	NextLink string         `json:"nextLink,omitempty"`
}

// RoleResponse -
type RoleResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

/////////////////////////
// User
/////////////////////////

type userRequest struct {
	Email    string           `json:"email,omitempty"`
	ObjectID string           `json:"objectId,omitempty"`
	Roles    []RoleAssignment `json:"roles"`
	TenantID string           `json:"tenantId,omitempty"`
	Type     string           `json:"type"`
}

type userResponse struct {
	ID       string           `json:"id"`
	Email    string           `json:"email,omitempty"`
	ObjectID string           `json:"objectId,omitempty"`
	Roles    []RoleAssignment `json:"roles"`
	TenantID string           `json:"tenantId,omitempty"`
	Type     string           `json:"type"`
}

// UserRequest -
type UserRequest struct {
	Email string           `json:"email"`
	Roles []RoleAssignment `json:"roles"`
}

// ADGroupUserRequest -
type ADGroupUserRequest struct {
	ObjectID string           `json:"objectId"`
	Roles    []RoleAssignment `json:"roles"`
	TenantID string           `json:"tenantId"`
}

// ServicePrincipalUserRequest -
type ServicePrincipalUserRequest struct {
	ObjectID string           `json:"objectId"`
	Roles    []RoleAssignment `json:"roles"`
	TenantID string           `json:"tenantId"`
}

// UserResponse -
type UserResponse struct {
	ID    string           `json:"id"`
	Email string           `json:"email"`
	Roles []RoleAssignment `json:"roles"`
}

// ADGroupUserResponse -
type ADGroupUserResponse struct {
	ID       string           `json:"id"`
	ObjectID string           `json:"objectId"`
	Roles    []RoleAssignment `json:"roles"`
	TenantID string           `json:"tenantId"`
}

// ServicePrincipalUserResponse -
type ServicePrincipalUserResponse struct {
	ID       string           `json:"id"`
	ObjectID string           `json:"objectId"`
	Roles    []RoleAssignment `json:"roles"`
	TenantID string           `json:"tenantId"`
}

/////////////////////////
// RoleAssignment
/////////////////////////

// RoleAssignment -
type RoleAssignment struct {
	Role         string `json:"role"`
	Organization string `json:"organization,omitempty"`
}
