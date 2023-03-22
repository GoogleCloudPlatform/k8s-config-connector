package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccRuntimeconfigConfigDatasource_basic(t *testing.T) {
	t.Parallel()

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRuntimeconfigDatasourceConfig(RandString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceStateMatchesResourceState("data.google_runtimeconfig_config.default", "google_runtimeconfig_config.default"),
				),
			},
		},
	})
}

func testAccRuntimeconfigDatasourceConfig(suffix string) string {
	return fmt.Sprintf(`
resource "google_runtimeconfig_config" "default" {
	name        = "runtime-%s"
	description = "runtime-%s"
}

data "google_runtimeconfig_config" "default" {
  name = google_runtimeconfig_config.default.name
}
`, suffix, suffix)
}
