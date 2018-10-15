package apimedic

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
)

//Mode is the mode that apimedic is running in (Sandbox or Live).
type Mode int

const (
	//Sandbox uses apimedic's Sandbox
	Sandbox Mode = 0
	//Live uses apimedic's Live mode
	Live Mode = 1
)

type service struct {
	authUrl   string
	healthUrl string
}

var services = map[Mode]service{
	Sandbox: service{
		authUrl:   "https://sandbox-authservice.priaid.ch",
		healthUrl: "https://sandbox-healthservice.priaid.ch",
	},
	Live: service{
		authUrl:   "https://authservice.priaid.ch",
		healthUrl: "https://healthservice.priaid.ch",
	},
}

//Client is a client for accessing apimedic's apis
type Client struct {
	Mode       Mode
	httpClient *http.Client
}

//NewClient returns a new client
func NewClient(m Mode, hc *http.Client) *Client {
	c := hc
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{
		Mode:       m,
		httpClient: c,
	}
}

func (c *Client) LogIn(apikey, secretkey string) (string, error) {
	return "", nil
}

func (c *Client) getAuthorizationHeader(apikey, secretkey string) string {
	return fmt.Sprintf("Bearer %s:%s", apikey, c.computeHash(secretkey))
}

func (c *Client) computeHash(s string) string {
	uri := fmt.Sprintf("%s/login", services[c.Mode].authUrl)
	b := []byte(s)
	h := hmac.New(md5.New, b)
	db := h.Sum([]byte(uri))
	return base64.StdEncoding.EncodeToString(db)
}
