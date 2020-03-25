package segment

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	SegmentApi "segment-terraform/segment/client"
)

func resourceSegmentSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"catalog_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_dev": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"write_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			name := r.Get("source_name").(string)
			catalog := r.Get("catalog_name").(string)
			isDev := r.Get("is_dev").(bool)

			source, err := client.CreateSource(name, catalog, isDev)
			if err == nil {
				r.SetId(source.Name)
				_ = r.Set("catalog_name", source.CatalogName)
				_ = r.Set("write_key", source.WriteKeys[0])
				_, err = client.GetSource(name)
			}
			return err
		},
		Read: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			_, err := client.GetSource(srcName)
			return err
		},
		Delete: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			srcName := r.Get("source_name").(string)
			return client.DeleteSource(srcName)
		},
		Importer: &schema.ResourceImporter{
			State: func(r *schema.ResourceData, meta interface{}) (data []*schema.ResourceData, err error) {
				client := meta.(*SegmentApi.Client)
				s, err := client.GetSource(r.Id())
				if err != nil {
					return nil, fmt.Errorf("invalid source: %q; err: %v", r.Id(), err)
				}
				r.SetId(s.Name)
				_ = r.Set("catalog_name", s.CatalogName)
				return []*schema.ResourceData{r}, nil
			},
		},
	}
}
