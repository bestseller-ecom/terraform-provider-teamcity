package teamcity

import (
	"fmt"
	//"log"
	//"regexp"
	//"strings"

	//"math/rand"
	// "reflect"
	"testing"
	//"time"

	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/bestseller-ecom/teamcity-sdk-go/teamcity"
	"github.com/bestseller-ecom/teamcity-sdk-go/types"
)

var testAccVcsRoot = `
resource "teamcity_vcs_root" "bar" {
  project = "Single"
  name = "bar"

  properties = {
    url    = "https://github.com/bestseller-ecom/teamcity-sdk-go"
    branch = "refs/heads/master"
  }
}`

func TestAccVcsRoot_basic(t *testing.T) {
	var v types.VcsRoot

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVcsRoot,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVcsRootExists("teamcity_vcs_root.bar", &v),
				),
			},
		},
	})
}

func testAccCheckVcsRpptDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*teamcity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "teamcity_vcs_root" {
			continue
		}

		// Try to find the Group
		var err error
		vcs, err := client.GetVcsRoot(rs.Primary.ID)

		if err == nil && vcs == nil {
			continue
		}

		if err == nil {
			return fmt.Errorf("VCS Root still exists")
		}

		return err
	}

	return nil
}

func testAccCheckVcsRootExists(n string, v *types.VcsRoot) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Project ID is set")
		}

		client := testAccProvider.Meta().(*teamcity.Client)

		vcs, err := client.GetVcsRoot(rs.Primary.ID)
		if err != nil {
			return err
		}

		if vcs == nil {
			return fmt.Errorf("VCS Root not found")
		}

		*v = *vcs

		return nil
	}
}
