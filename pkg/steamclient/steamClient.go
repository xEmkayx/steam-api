package steamclient

import (
	"errors"
	"log/slog"
	"net/http"
)

/*
This Client is used to send the requests to steam
*/
type Client struct {
	Key        string       // Access-/API-Key for the Steam API
	HttpClient *http.Client // An Http-Client to send requests with. Customizable
}

// Create a Client, without Key
func NewClientWithoutKey(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{HttpClient: httpClient}
}

/*
Create New Client with the provided API-key and a custom http.Client
*/
func New(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{Key: key, HttpClient: httpClient}
}

func (c Client) IsKeySet() bool {
	return c.Key != ""
}

// General method to send a GET request
func (c Client) getRequest(urlStr string) (resp *http.Response, err error) {
	if c.HttpClient == nil {
		return nil, errors.New("the HttpClient should is not defined")
	}
	slog.Debug("Sending GET-Request to " + urlStr)
	return c.HttpClient.Get(urlStr)
}
