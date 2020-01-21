package teamcity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/paranoidd/teamcity-sdk-go/teamcity"

	"fmt"
)

func dataSourceAgentPool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAgentPoolRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAgentPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)

	pool, err := client.GetAgentPoolByName(d.Get("name").(string))
	if err != nil {
		return err
	}

	if pool == nil || len(pool.Name) == 0 {
		return fmt.Errorf("no matching Agent Pool found")
	}

	id := fmt.Sprintf("%d", pool.ID)
	d.SetId(id)
	d.Set("name", pool.Name)

	return nil
}
