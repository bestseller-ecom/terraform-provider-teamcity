package teamcity

import (
	// "fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/flatmap"
	"github.com/hashicorp/terraform/helper/schema"

	// "errors"
	"log"

	"github.com/bestseller-ecom/teamcity-sdk-go/teamcity"
	"github.com/bestseller-ecom/teamcity-sdk-go/types"
	// "reflect"
)

func resourceVcsRoot() *schema.Resource {
	return &schema.Resource{
		Create: resourceVcsRootCreate,
		Read:   resourceVcsRootRead,
		Update: resourceVcsRootUpdate,
		Delete: resourceVcsRootDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"project": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "_Root",
				ForceNew:     true,
				ValidateFunc: teamcity.ValidateVcsRootID,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcs_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "jetbrains.git",
			},
			"properties": &schema.Schema{
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: optionalSuffixReturnSuppress,
			},
		},
	}
}

/*
   ID         string     `json:"id"`
   Name       string     `json:"name"`
   VcsName    string     `json:"vcsName"`
   Href       string     `json:"href"`
   ProjectID  ProjectId  `json:"project"`
  Properties Properties `json:"properties"`
*/

func resourceVcsRootCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)

	vcsProject := d.Get("project").(string)
	vcsName := d.Get("name").(string)
	vcsProvider := d.Get("vcs_provider").(string)
	vcsProperties := flatmap.Flatten(d.Get("properties").(map[string]interface{}))
	spew.Dump(vcsProperties)

	vcs := types.VcsRoot{
		ProjectID:  types.ProjectId(vcsProject),
		Name:       vcsName,
		VcsName:    vcsProvider,
		Properties: types.Properties(vcsProperties),
	}
	err := client.CreateVcsRoot(&vcs)
	if err != nil {
		return err
	}

	id := vcs.ID
	d.SetId(id)
	d.Set("vcs_provider", vcs.VcsName)
	return nil
}

func resourceVcsRootRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("Reading vcs_root resource %q", d.Id())
	client := meta.(*teamcity.Client)
	vcs, err := client.GetVcsRoot(d.Id())
	if err != nil {
		return err
	}
	if vcs == nil {
		d.SetId("")
		return nil
	}

	d.Set("project", vcs.ProjectID)
	d.Set("name", vcs.Name)
	d.Set("vcs_provider", vcs.VcsName)
	d.Set("properties", vcs.Properties)

	return nil
}

func resourceVcsRootUpdate(d *schema.ResourceData, meta interface{}) error {
	// return errors.New("Update is not supported")
	client := meta.(*teamcity.Client)

	id := d.Id()
	d.Partial(true)

	tfProperties := d.Get("properties").(map[string]interface{})

	vcsProperties := types.Properties(flatmap.Flatten(tfProperties))

	if d.HasChange("properties") {
		if err := client.ReplaceAllVcsRootProperties(id, &vcsProperties); err != nil {
			return err
		}
		d.SetPartial("properties")
	}

	d.Partial(false)
	return nil
}

func resourceVcsRootDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teamcity.Client)
	return client.DeleteVcsRoot(d.Id())
}
