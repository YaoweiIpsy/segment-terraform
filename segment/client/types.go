package client

import "time"

// Workspace defines the struct for the workspace object
type Workspace struct {
	Name        string    `json:"name,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	ID          string    `json:"id,omitempty"`
	CreateTime  time.Time `json:"create_time,omitempty"`
}

// Sources defines the struct for the sources object
type Sources struct {
	Sources       []Source `json:"sources,omitempty"`
	NextPageToken string   `json:"next_page_token,omitempty"`
}

// Source defines the struct for the source object
type Source struct {
	Name          string        `json:"name,omitempty"`
	CatalogName   string        `json:"catalog_name,omitempty"`
	Parent        string        `json:"parent,omitempty"`
	WriteKeys     []string      `json:"write_keys,omitempty"`
	LibraryConfig LibraryConfig `json:"library_config,omitempty"`
	CreateTime    time.Time     `json:"create_time,omitempty"`
}

// LibraryConfig contains information about a source's library
type LibraryConfig struct {
	MetricsEnabled       bool   `json:"metrics_enabled,omitempty"`
	RetryQueue           bool   `json:"retry_queue,omitempty"`
	CrossDomainIDEnabled bool   `json:"cross_domain_id_enabled,omitempty"`
	APIHost              string `json:"api_host,omitempty"`
}

// Destinations defines the struct for the destination object
type Destinations struct {
	Destinations []Destination `json:"destinations,omitempty"`
}

// Destination defines the struct for the destination object
type Destination struct {
	Name           string              `json:"name,omitempty"`
	Parent         string              `json:"parent,omitempty"`
	DisplayName    string              `json:"display_name,omitempty"`
	Enabled        bool                `json:"enabled,omitempty"`
	ConnectionMode string              `json:"connection_mode,omitempty"`
	Configs        []DestinationConfig `json:"config,omitempty"`
	CreateTime     time.Time           `json:"create_time,omitempty"`
	UpdateTime     time.Time           `json:"update_time,omitempty"`
}

// DestinationConfig contains information about how a Destination is configured
type DestinationConfig struct {
	Name        string      `json:"name,omitempty"`
	DisplayName string      `json:"display_name,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	Type        string      `json:"type,omitempty"`
}
