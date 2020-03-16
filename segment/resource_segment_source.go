package segment

import (
	SegmentApi "awesomeProject/segment/client"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
		},
		Create: func(r *schema.ResourceData, meta interface{}) error {
			client := meta.(*SegmentApi.Client)
			name := r.Get("source_name").(string)
			catalog := r.Get("catalog_name").(string)

			source, err := client.CreateSource(name, catalog)
			if err == nil {
				r.SetId(source.Name)
				r.Set("catalog_name", source.CatalogName)
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
				r.Set("catalog_name", s.CatalogName)
				return []*schema.ResourceData{r}, nil
			},
		},
	}
}
