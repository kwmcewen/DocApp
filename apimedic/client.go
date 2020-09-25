package apimedic

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
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
	authURL   string
	healthURL string
}

var services = map[Mode]service{
	Sandbox: service{
		authURL:   "https://sandbox-authservice.priaid.ch",
		healthURL: "https://sandbox-healthservice.priaid.ch",
	},
	Live: service{
		authURL:   "https://authservice.priaid.ch",
		healthURL: "https://healthservice.priaid.ch",
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

//LogIn logs the user into apimedic and returns the token
func (c *Client) LogIn(username, password string) (string, error) {
	var returnVal string = ""
	req, err := http.NewRequest("POST", c.getAuthURL(), nil)
	if err != nil {
		return returnVal, err
	}
	req.Header.Add("Authorization", c.getAuthorizationHeader(username, password))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return returnVal, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		returnVal := string(bodyBytes)
		return returnVal, err
	}

	return returnVal, fmt.Errorf("Response: (%s)", resp.Status)

}

func (c *Client) RequestAndResponseService(resource string, token string) (string, error) {
	req, err := http.NewRequest("GET", c.getHealthURLForRequestAndResponseService(resource)+c.getQueryParameters(token), nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return bodyString, err
	}

	return "", fmt.Errorf("Response: (%s)", resp.Status)
}

func (c *Client) getAuthorizationHeader(username, password string) string {
	return fmt.Sprintf("Bearer %s:%s", username, c.computeHash(password))
}

func (c *Client) getAuthURL() string {
	uri := fmt.Sprintf("%s/login", services[c.Mode].authURL)
	return uri
}

func (c *Client) getHealthURLForRequestAndResponseService(resource string) string {
	return c.getHealthURL(resource + "?") // Need to make this dynamic
}

func (c *Client) getHealthURL(queryParam string) string {
	uri := fmt.Sprintf("%s/"+queryParam, services[c.Mode].healthURL)
	return uri
}

func (c *Client) computeHash(s string) string {
	b := []byte(s)
	h := hmac.New(md5.New, b)
	h.Write([]byte(c.getAuthURL()))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *Client) getQueryParameters(t string) string {
	return "token=" + t + "&language=en-gb&format=json"
}
