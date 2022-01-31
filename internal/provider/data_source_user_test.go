package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUser,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.mattermost_user.sut", "email", "mattermost@octo.com"),
				),
			},
		},
	})
}

const testAccDataSourceUser = `
data "mattermost_user" "sut" {
  username = "mattermost"
}
`
