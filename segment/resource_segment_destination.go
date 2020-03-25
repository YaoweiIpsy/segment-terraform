package segment

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	SegmentApi "segment-terraform/segment/client"
	"strings"
)

func resourceSegmentDestination() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connection_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UNSPECIFIED",
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"configs": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"list": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
			},
		},
		Create: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			dstName := r.Get("destination_name").(string)
			connectionMode := r.Get("connection_mode").(string)
			enabled := r.Get("enabled").(bool)
			configs := r.Get("configs").(*schema.Set)

			var configsList []SegmentApi.DestinationConfig

			if configs != nil {
				for _, config := range configs.List() {
					c := SegmentApi.DestinationConfig{
						Name:  config.(map[string]interface{})["name"].(string),
						Value: config.(map[string]interface{})["value"],
						Type:  config.(map[string]interface{})["type"].(string),
					}
					if c.Type == "list" {
						c.Value = config.(map[string]interface{})["list"]
					}
					configsList = append(configsList, c)
				}
			}
			destination, err := client.CreateDestination(srcName, dstName, connectionMode, enabled, configsList...)
			if err == nil {
				r.SetId(destination.Name)
				_ = r.Set("source_name", srcName)
				_, err = client.GetDestination(srcName, dstName)
			}
			return err
		},
		Read: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			name := r.Get("destination_name").(string)
			destination, err := client.GetDestination(srcName, name)
			if err == nil {
				_ = r.Set("enabled", destination.Enabled)
				_ = r.Set("configs", destination.Configs)
				_ = r.Set("connection_mode", destination.ConnectionMode)
			}
			return err
		},
		Update: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			dstName := r.Get("destination_name").(string)
			enabled := r.Get("enabled").(bool)
			configs := r.Get("configs").(*schema.Set)

			var configsList []SegmentApi.DestinationConfig

			if configs != nil {
				for _, config := range configs.List() {
					c := SegmentApi.DestinationConfig{
						Name:  config.(map[string]interface{})["name"].(string),
						Value: config.(map[string]interface{})["value"],
						Type:  config.(map[string]interface{})["type"].(string),
					}
					if c.Type == "list" {
						c.Value = config.(map[string]interface{})["list"]
					}
					configsList = append(configsList, c)
				}
			}
			destination, err := client.UpdateDestination(srcName, dstName, enabled, configsList...)
			if err == nil {
				r.SetId(destination.Name)
				_ = r.Set("source_name", srcName)
				_, err = client.GetDestination(srcName, dstName)
			}
			return err
		},
		Delete: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			name := r.Get("destination_name").(string)
			return client.DeleteDestination(srcName, name)
		},
		Importer: &schema.ResourceImporter{
			State: func(r *schema.ResourceData, meta interface{}) (data []*schema.ResourceData, err error) {
				client := meta.(*SegmentApi.Client)
				s := strings.SplitN(r.Id(), "/", 2)
				destination, err := client.GetDestination(s[0], s[1])
				if err != nil {
					return nil, fmt.Errorf("invalid source: %q; err: %v", r.Id(), err)
				}

				_ = r.Set("source_name", s[0])
				_ = r.Set("destination_name", s[1])
				_ = r.Set("enabled", destination.Enabled)
				_ = r.Set("configs", destination.Configs)
				_ = r.Set("connection_mode", destination.ConnectionMode)

				return []*schema.ResourceData{r}, nil
			},
		},
	}
}
