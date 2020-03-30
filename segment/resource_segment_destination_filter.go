package segment

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	SegmentApi "segment-terraform/segment/client"
	"strings"
)

func resourceSegmentDestinationFilter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition": {
				Type:     schema.TypeString,
				Required: true,
			},
			"actions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
		Create: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			dstName := r.Get("destination_name").(string)
			condition := r.Get("condition").(string)
			actions := r.Get("actions").([]interface{})
			title := r.Get("title").(string)
			description := r.Get("description").(string)
			enabled := r.Get("enabled").(bool)

			var actionStrs []string
			for _, action := range actions {
				actionStrs = append(actionStrs, action.(string))
			}
			filter, err := client.CreateDestinationFilter(srcName, dstName, condition, title, description, enabled, actionStrs...)
			if err == nil {
				r.SetId(filter.Name[strings.LastIndex(filter.Name, "/")+1:])
				_ = r.Set("source_name", srcName)
				_, err = client.GetDestinationFilter(srcName, dstName, r.Id())
			}
			return err
		},
		Read: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			dstName := r.Get("destination_name").(string)
			filterId := r.Id()[strings.LastIndex(r.Id(), "/")+1:]
			filter, err := client.GetDestinationFilter(srcName, dstName, filterId)
			if err == nil {
				_ = r.Set("enabled", filter.Enabled)
				_ = r.Set("title", filter.Title)
				_ = r.Set("condition", filter.If)
				var actionStrs []string
				for _, act := range filter.Actions {
					actStr, err := json.Marshal(act)
					if err != nil {
						return err
					}
					actionStrs = append(actionStrs, string(actStr))
				}
				_ = r.Set("actions", actionStrs)
			}
			return err
		},
		Update: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			dstName := r.Get("destination_name").(string)
			condition := r.Get("condition").(string)
			actions := r.Get("actions").([]interface{})
			title := r.Get("title").(string)
			description := r.Get("description").(string)
			enabled := r.Get("enabled").(bool)
			filterId := r.Id()[strings.LastIndex(r.Id(), "/")+1:]
			var actionStrs []string
			for _, action := range actions {
				actionStrs = append(actionStrs, action.(string))
			}

			filter, err := client.UpdateDestinationFilter(srcName, dstName, filterId, condition, title, description, enabled, actionStrs...)
			if err == nil {
				r.SetId(filter.Name[strings.LastIndex(filter.Name, "/")+1:])
				_ = r.Set("source_name", srcName)
				_, err = client.GetDestinationFilter(srcName, dstName, filterId)
			}
			return err
		},
		Delete: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			name := r.Get("destination_name").(string)
			return client.DeleteDestinationFilter(srcName, name, r.Id())
		},
		Importer: &schema.ResourceImporter{
			State: func(r *schema.ResourceData, meta interface{}) (data []*schema.ResourceData, err error) {
				client := meta.(*SegmentApi.Client)
				s := strings.SplitN(r.Id(), "/", 2)
				filter, err := client.GetDestinationFilter(s[0], s[1], s[2])
				if err != nil {
					return nil, err
				}
				_ = r.Set("enabled", filter.Enabled)
				_ = r.Set("title", filter.Title)
				_ = r.Set("condition", filter.If)
				var actionStrs []string
				for _, act := range filter.Actions {
					actStr, err := json.Marshal(act)
					if err != nil {
						return nil, err
					}
					actionStrs = append(actionStrs, string(actStr))
				}
				_ = r.Set("actions", actionStrs)

				return []*schema.ResourceData{r}, nil
			},
		},
	}
}
