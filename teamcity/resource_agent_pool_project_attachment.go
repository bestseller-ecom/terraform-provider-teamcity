package teamcity

import (
	"github.com/bestseller-ecom/teamcity-sdk-go/teamcity"
	"github.com/bestseller-ecom/teamcity-sdk-go/types"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceAgentPoolProjectAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAgentPoolProjectAttachementCreate,
		Read:   resourceAgentPoolProjectAttachementRead,
		Delete: resourceAgentPoolProjectAttachementDelete,

		Schema: map[string]*schema.Schema{
			"pool": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAgentPoolProjectAttachementCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)

	attachmentPool := d.Get("pool").(int)
	attachmentProject := d.Get("project").(string)

	attachment := types.AgentPoolAttachment{
		ProjectID: attachmentProject,
	}

	create_err := client.CreateAgentPoolProjectAttachment(attachmentPool, &attachment)
	if create_err != nil {
		return create_err
	}

	id := fmt.Sprintf("%d_%s", attachmentPool, attachment.ProjectID)
	d.SetId(id)
	return nil
}

func resourceAgentPoolProjectAttachementRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Reading agent_pool_project_attachment resource %q", d.Id())
	client := meta.(*teamcity.Client)
	pool, err := client.GetAgentPoolById(d.Get("pool").(int))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Read Agent Pools: %q", pool)
	var project types.Project
	project = pool.Projects[d.Get("project").(string)]

	log.Printf("[DEBUG] Read Agent Pool: %q", project)

	if project.ID == "" {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceAgentPoolProjectAttachementDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)
	err := client.DeleteAgentPoolProjectAttachement(d.Get("pool").(int), d.Get("project").(string))
	return err
}
