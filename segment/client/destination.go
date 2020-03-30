package client

import (
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) ListDestinations(srcName string) (Destinations, error) {
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations",
		result:   new(Destinations),
	}
	return *request.result.(*Destinations), request.Do(client)
}

func (client *Client) GetDestination(srcName string, destName string) (Destination, error) {
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + destName,
		result:   new(Destination),
	}
	return *request.result.(*Destination), request.Do(client)
}

func (client *Client) CreateDestination(srcName, destName string, connectionMode string, enabled bool, configs ...DestinationConfig) (Destination, error) {
	for index, config := range configs {
		if !strings.HasPrefix(config.Name, "workspaces") {
			configs[index].Name = fmt.Sprintf("workspaces/%s/sources/%s/destinations/%s/config/%s", client.workspace, srcName, destName, config.Name)
		}
	}
	request := Request{
		endpoint: "sources/" + srcName + "/destinations",
		method:   http.MethodPost,
		body: struct {
			Destination Destination `json:"destination,omitempty"`
		}{
			Destination{
				Name:           "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + destName,
				Enabled:        enabled,
				ConnectionMode: connectionMode,
				Configs:        configs,
			},
		},
		result: new(Destination),
	}
	return *request.result.(*Destination), request.Do(client)
}
func (client *Client) UpdateDestination(srcName, destName string, enabled bool, configs ...DestinationConfig) (Destination, error) {
	for index, config := range configs {
		if !strings.HasPrefix(config.Name, "workspaces") {
			configs[index].Name = fmt.Sprintf("workspaces/%s/sources/%s/destinations/%s/config/%s", client.workspace, srcName, destName, config.Name)
		}
	}

	request := Request{
		endpoint: "sources/" + srcName + "/destinations/" + destName,
		method:   http.MethodPatch,
		body: struct {
			Enabled    bool                `json:"enabled,omitempty"`
			Configs    []DestinationConfig `json:"config,omitempty"`
			UpdateMask map[string][]string `json:"update_mask,omitempty"`
		}{
			Enabled: enabled,
			Configs: configs,
			UpdateMask: map[string][]string{
				"Updates": {
					"destination.config",
					"destination.enabled",
				},
			},
		},
		result: new(Destination),
	}
	return *request.result.(*Destination), request.Do(client)
}

func (client *Client) DeleteDestination(srcName, destName string) error {
	request := Request{
		endpoint: fmt.Sprintf("sources/%s/destinations/%s", srcName, destName),
		method:   http.MethodDelete,
	}
	return request.Do(client)
}
