package iotcentral

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	if host == nil || *host == "" {
		return nil, errors.New("host is nil or empty")
	}
	if token == nil || *token == "" {
		return nil, errors.New("token is nil or empty")
	}

	c := Client{
		HostURL:    *host,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      *token,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, int, error) {
	if c == nil {
		return nil, 0, errors.New("client is nil")
	}

	token := c.Token
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return body, res.StatusCode, err
}
