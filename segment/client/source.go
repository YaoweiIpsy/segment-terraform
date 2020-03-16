package client

import (
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) ListSources(pageToken string) (Sources, error) {
	request := Request{
		client:   client,
		endpoint: "workspaces/" + client.workspace + "/sources",
		result:   new(Sources),
	}
	if len(pageToken) > 5 {
		request.params = &map[string]string{"page_token": pageToken}
	}
	return *request.result.(*Sources), request.Do()
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
		client:   client,
		endpoint: "workspaces/" + client.workspace + "/sources/" + name,
		result:   new(Source),
	}
	return *request.result.(*Source), request.Do()
}
func (client *Client) CreateSource(name string, catalog string) (Source, error) {
	if !strings.HasPrefix(name, "workspaces") {
		name = fmt.Sprintf("workspaces/%s/sources/%s", client.workspace, name)
	}
	request := Request{
		client: client,
		method: http.MethodPost,
		body: struct {
			Source Source `json:"source,omitempty"`
		}{
			Source{
				Name:        name,
				CatalogName: catalog,
			},
		},
		result: new(Source),
	}

	return *request.result.(*Source), request.Do()

}
func (client *Client) DeleteSource(name string) error {
	if !strings.HasPrefix(name, "workspaces") {
		name = fmt.Sprintf("workspaces/%s/sources/%s", client.workspace, name)
	}

	request := Request{
		client:   client,
		endpoint: name,
		method:   http.MethodDelete,
	}
	return request.Do()
}
