package iotcentral

type deviceCollection struct {
	Value    []Device `json:"value"`
	NextLink string   `json:"nextLink"`
}

// Device -
type Device struct {
	ID            string   `json:"id"`
	Etag          string   `json:"etag"`
	DisplayName   string   `json:"displayName"`
	Template      string   `json:"template"`
	Simulated     bool     `json:"simulated"`
	Provisioned   bool     `json:"provisioned"`
	Enabled       bool     `json:"enabled"`
	Organizations []string `json:"organizations"`
}

// Organization -
// type Organization struct {
// 	ID          string `json:"id"`
// 	DisplayName string `json:"displayName"`
// 	Parent      string `json:"parent"`
// }
