package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	url         string
	accessToken string
	workspace   string
}

func NewClient(accessToken string, workspace string) *Client {
	return &Client{
		url:         "https://platform.segmentapis.com/v1beta",
		accessToken: accessToken,
		workspace:   workspace,
	}
}

type Request struct {
	endpoint string
	params   *map[string]string
	body     interface{}
	result   interface{}
	method   string
}

func (request *Request) Do(client *Client) error {
	body := bytes.NewBuffer(nil)
	if request.body != nil {
		enc := json.NewEncoder(body)
		if err := enc.Encode(request.body); err != nil {
			return err
		}
	}
	if !strings.HasPrefix(request.endpoint, "workspaces/") {
		request.endpoint = fmt.Sprintf("workspaces/%s/%s", client.workspace, strings.Trim(request.endpoint, "/"))
	}
	req, err := http.NewRequest(
		request.method,
		fmt.Sprintf("%s/%s", client.url, strings.Trim(request.endpoint, "/")),
		body)
	if err != nil {
		return err
	}
	if request.params != nil {
		q := req.URL.Query()
		for key, value := range *request.params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("Authorization", "Bearer "+client.accessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("%s %s failed: %d - %s", req.Method, req.URL, resp.StatusCode, string(data))
	}
	if request.result == nil {
		return nil
	}
	return json.Unmarshal(data, request.result)
}
