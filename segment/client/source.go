package client

import (
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) ListSources(pageToken string) (Sources, error) {
	request := Request{
		endpoint: "sources",
		result:   new(Sources),
	}
	request.params = &map[string]string{"page_token": pageToken}

	return *request.result.(*Sources), request.Do(client)
}

func (client *Client) ListAllSources() ([]Source, error) {
	token := ""
	var results []Source

	for {
		sources, err := client.ListSources(token)
		if err != nil {
			return nil, err
		}
		if len(sources.Sources) == 0 {
			break
		}
		results = append(results, sources.Sources...)
		token = sources.NextPageToken
	}
	return results, nil
}
func (client *Client) GetSource(name string) (Source, error) {
	request := Request{
		endpoint: "sources/" + name,
		result:   new(Source),
	}
	return *request.result.(*Source), request.Do(client)
}
func (client *Client) CreateSource(name string, catalog string, displayName string, isDev bool) (Source, error) {
	if !strings.HasPrefix(name, "workspaces") {
		name = fmt.Sprintf("workspaces/%s/sources/%s", client.workspace, name)
	}
	var labels map[string]string
	if isDev {
		labels = map[string]string{"environment": "dev"}
	} else {
		labels = map[string]string{"environment": "prod"}
	}
	request := Request{
		method: http.MethodPost,
		body: struct {
			Source Source `json:"source,omitempty"`
		}{
			Source{
				Name:        name,
				CatalogName: catalog,
				DisplayName: displayName,
				Labels:      labels,
			},
		},
		endpoint: "sources",
		result:   new(Source),
	}
	return *request.result.(*Source), request.Do(client)
}
func (client *Client) DeleteSource(name string) error {
	if !strings.HasPrefix(name, "workspaces") {
		name = fmt.Sprintf("workspaces/%s/sources/%s", client.workspace, name)
	}

	request := Request{
		endpoint: name,
		method:   http.MethodDelete,
	}
	return request.Do(client)
}
