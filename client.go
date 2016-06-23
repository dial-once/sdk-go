package dialonce

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"github.com/segmentio/go-env"
)


var baseURL = env.GetDefault("DIALONCE_BASE_URL", "https://api.dial-once.com/")

// Client represents a Dial Once client.
type Client struct {
	Config 	*ClientConfig
	HTTPClient  *http.Client
	IVR 		*IVR
}

// ClientConfig ...
type ClientConfig struct {
	AccessToken string
	BaseURL     string
}

// SetAccessToken ...
func (c *Client) SetAccessToken(accessToken string) {
	c.Config.AccessToken = accessToken
}

// SetBaseURL ...
func (c *Client) SetBaseURL(baseURL string) {
	c.Config.BaseURL = baseURL
}

// NewConfig with the given access token.
func NewConfig(accessToken string) *ClientConfig {
	return &ClientConfig{
		AccessToken: accessToken,
		BaseURL: baseURL,
	}
}

// New client.
func New(config *ClientConfig) *Client {
	client := &Client{
		Config: config,
		HTTPClient: http.DefaultClient,
	}

	client.IVR = &IVR{client}
	return client
}

// Init ...
func Init(accessToken string) *Client {
	config := NewConfig(accessToken)
	return New(config)
}

// call rpc style endpoint.
func (c *Client) call(verb string, path string, in interface{}) (io.ReadCloser, error) {
	url := c.Config.BaseURL + path
	data := []byte("")

	if in != nil {
		body, err := json.Marshal(in)
		data = body

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(verb, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Config.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 400 {
		return res.Body, err
	}

	defer res.Body.Close()

	return nil, errors.New(http.StatusText(res.StatusCode))
}
