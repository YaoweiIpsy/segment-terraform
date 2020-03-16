package client

import (
	"net/http"
)

func (client *Client) ListDestinations(srcName string) (Destinations, error) {
	request := Request{
		client:   client,
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations",
		result:   new(Destinations),
	}
	return *request.result.(*Destinations), request.Do()
}

func (client *Client) GetDestination(srcName string, destName string) (Destination, error) {
	request := Request{
		client:   client,
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + destName,
		result:   new(Destination),
	}
	return *request.result.(*Destination), request.Do()
}

func (client *Client) CreateDestination(srcName, destName string, enabled bool, configs []DestinationConfig) (Destination, error) {
	path := "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations"
	request := Request{
		client:   client,
		endpoint: path,
		method:   http.MethodPost,
		body: struct {
			Destination Destination `json:"destination,omitempty"`
		}{
			Destination{
				Name:           path + "/" + destName,
				Enabled:        enabled,
				ConnectionMode: "UNSPECIFIED",
				Configs:        configs,
			},
		},
		result: new(Destination),
	}
	return *request.result.(*Destination), request.Do()
}

func (client *Client) DeleteDestination(srcName, destName string) error {
	request := Request{
		client:   client,
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + destName,
		method:   http.MethodDelete,
	}
	return request.Do()
}
