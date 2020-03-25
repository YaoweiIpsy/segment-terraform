package client

func (client *Client) ListDestinationFilters(srcName string) (Destinations, error) {
	request := Request{
		endpoint: "workspaces/" + client.workspace + "/sources/" + srcName + "/destinations",
		result:   new(Destinations),
	}
	return *request.result.(*Destinations), request.Do(client)
}
