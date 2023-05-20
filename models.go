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

// Organization -
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
