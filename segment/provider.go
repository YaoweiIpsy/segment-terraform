package segment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	segmentApi "segment-terraform/segment/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Description: "The Access Token used to connect to Segment",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEGMENT_ACCESS_TOKEN", nil),
			},
			"workspace": {
				Type:        schema.TypeString,
				Description: "The Segment workspace to manage",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEGMENT_WORKSPACE", nil),
			},
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return segmentApi.NewClient(d.Get("access_token").(string), d.Get("workspace").(string)), nil
		},
		ResourcesMap: map[string]*schema.Resource{
			"segment_source":             resourceSegmentSource(),
			"segment_destination":        resourceSegmentDestination(),
			"segment_destination_filter": resourceSegmentDestinationFilter(),
		},
	}
}
