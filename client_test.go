package dialonce

import (
	"testing"

	"github.com/segmentio/go-env"
	"github.com/stretchr/testify/assert"
)

var token = env.GetDefault("DIALONCE_ACCESS_TOKEN", "AsvAUbSo37MgnDl4HKBYz66LWuOow3o3")

func testClient() *Client {
	c := Init(token)
	c.SetBaseURL("http://api.dialonce.io/")
	return c
}

func TestClient_SetAccessToken(t *testing.T) {
	c := testClient()
	c.SetAccessToken("testAccessToken")

	assert.Equal(t, c.Config.AccessToken, "testAccessToken")
}

func TestClient_SetBaseURL(t *testing.T) {
	c := testClient()
	c.SetBaseURL("http://custom.base.url/")

	assert.Equal(t, c.Config.BaseURL, "http://custom.base.url/")
}

func TestClient_call(t *testing.T) {
	//client := testClient()

}
