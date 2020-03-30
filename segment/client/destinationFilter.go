package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) ListDestinationFilters(srcName string, dstName string) (DestinationFilters, error) {
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + dstName + "/filters",
		result:   new(DestinationFilters),
	}
	return *request.result.(*DestinationFilters), request.Do(client)
}

func (client *Client) GetDestinationFilter(srcName string, dstName string, filterId string) (DestinationFilter, error) {
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + dstName + "/filters/" + filterId,
		result:   new(DestinationFilter),
	}
	return *request.result.(*DestinationFilter), request.Do(client)
}

func (client *Client) CreateDestinationFilter(srcName string, dstName string, ifs string, title string, description string, enabled bool, actions ...string) (DestinationFilter, error) {
	var filter DestinationFilter
	var filterActions []DestinationFilterAction
	actionsStr := fmt.Sprintf("[%s]", strings.Join(actions, ","))
	err := json.Unmarshal([]byte(actionsStr), &filterActions)
	if err != nil {
		return filter, err
	}
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + dstName + "/filters/",
		method:   http.MethodPost,
		body: struct {
			Filter DestinationFilter `json:"filter,omitempty"`
		}{
			DestinationFilter{
				If:          ifs,
				Actions:     filterActions,
				Title:       title,
				Description: description,
				Enabled:     enabled,
			},
		},
		result: &filter,
	}
	return filter, request.Do(client)

}
func (client *Client) UpdateDestinationFilter(srcName string, dstName string, filterId string, ifs string, title string, description string, enabled bool, actions ...string) (DestinationFilter, error) {
	var filter DestinationFilter
	var filterActions []DestinationFilterAction
	actionsStr := fmt.Sprintf("[%s]", strings.Join(actions, ","))
	err := json.Unmarshal([]byte(actionsStr), &filterActions)
	if err != nil {
		return filter, err
	}
	if err != nil {
		return filter, err
	}
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + dstName + "/filters/" + filterId,
		method:   http.MethodPatch,
		body: struct {
			Filter     DestinationFilter `json:"filter,omitempty"`
			UpdateMask struct {
				Paths []string `json:"paths,omitempty"`
			} `json:"update_mask,omitempty"`
		}{
			Filter: DestinationFilter{
				If:          ifs,
				Actions:     filterActions,
				Title:       title,
				Description: description,
				Enabled:     enabled,
			},
			UpdateMask: struct {
				Paths []string `json:"paths,omitempty"`
			}{
				Paths: []string{"if", "actions", "title", "description", "enabled"},
			},
		},
		result: &filter,
	}
	return filter, request.Do(client)
}
func (client *Client) DeleteDestinationFilter(srcName string, dstName string, filterId string) error {
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations/" + dstName + "/filters/" + filterId,
		method:   http.MethodDelete,
	}
	return request.Do(client)
}
