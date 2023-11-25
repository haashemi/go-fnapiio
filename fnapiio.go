package fnapiio

import "net/http"

const Host = "https://fortniteapi.io"

// Client holds the api token and http client for the requests.
type Client struct {
	token  string
	client *http.Client
	lang   Lang
}

// New returns a new Client with the passed token.
func New(token string) *Client {
	return &Client{token: token}
}

// NewWithHTTPClient returns a new Client with the passed http-client and token.
func NewWithHTTPClient(token string, client *http.Client) *Client {
	return &Client{token: token, client: client}
}
