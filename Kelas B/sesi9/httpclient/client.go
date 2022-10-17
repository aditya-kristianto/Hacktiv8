package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	client  *http.Client
	baseUrl string
}

func NewClient(timeout time.Duration, baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
		client: &http.Client{
			Timeout: timeout * time.Second,
		},
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	data, err := c.req(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) Post(url string) ([]byte, error) {
	var payload = map[string]interface{}{
		"userId": 1,
		"title":  "MNC B",
		"body":   "Ini dari MNC B",
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	data, err := c.req(url, http.MethodPost, bytePayload)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) Put(url string) {
	c.req(url, http.MethodPut, []byte{})
}

func (c *Client) Delete(url string) {
	c.req(url, http.MethodDelete, nil)
}

func (c *Client) req(url string, method string, data []byte) ([]byte, error) {

	req, err := http.NewRequest(method, c.baseUrl+url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	var dataInterface interface{}
	err = json.NewDecoder(resp.Body).Decode(&dataInterface)
	if err != nil {
		return nil, err
	}

	dataByte, err := json.Marshal(dataInterface)
	if err != nil {
		return nil, err
	}

	return dataByte, nil
}
