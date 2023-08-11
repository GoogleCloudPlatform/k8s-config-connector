// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNotebooksInstance_create_container(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("%d", acctest.RandInt(t))
	name := fmt.Sprintf("tf-%s", prefix)

	acctest.VcrTest(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_create_container(name),
			},
			{
				ResourceName:            "google_notebooks_instance.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ExpectNonEmptyPlan:      true,
				ImportStateVerifyIgnore: []string{"container_image", "metadata", "vm_image"},
			},
		},
	})
}

func testAccNotebooksInstance_create_container(name string) string {
	return fmt.Sprintf(`

resource "google_notebooks_instance" "test" {
  name = "%s"
  location = "us-west1-a"
  machine_type = "e2-medium"
  metadata = {
    proxy-mode = "service_account"
    terraform  = "true"
  }
  container_image {
    repository = "gcr.io/deeplearning-platform-release/base-cpu"
    tag = "latest"
  }
}
`, name)
}
