// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesMesh_update(t *testing.T) {
	t.Parallel()

	meshName := fmt.Sprintf("tf-test-mesh-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesMeshDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesMesh_basic(meshName),
			},
			{
				ResourceName:      "google_network_services_mesh.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesMesh_update(meshName),
			},
			{
				ResourceName:      "google_network_services_mesh.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesMesh_basic(meshName string) string {
	return fmt.Sprintf(`
resource "google_network_services_mesh" "foobar" {
  name        = "%s"
  description = "my description"
}
`, meshName)
}

func testAccNetworkServicesMesh_update(meshName string) string {
	return fmt.Sprintf(`
resource "google_network_services_mesh" "foobar" {
  name        = "%s"
  description = "update description"
  labels      = {
    foo = "bar"
  } 
}
`, meshName)
}
