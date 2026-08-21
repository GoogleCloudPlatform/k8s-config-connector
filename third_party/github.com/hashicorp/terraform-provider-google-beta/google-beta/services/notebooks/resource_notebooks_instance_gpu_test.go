// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package notebooks_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNotebooksInstance_create_gpu(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("%d", acctest.RandInt(t))
	name := fmt.Sprintf("tf-%s", prefix)

	acctest.VcrTest(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_create_gpu(name),
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

func testAccNotebooksInstance_create_gpu(name string) string {
	return fmt.Sprintf(`

resource "google_notebooks_instance" "test" {
  name = "%s"
  location = "us-west1-a"
  machine_type = "n1-standard-1"  // can't be e2 because of accelerator
  metadata = {
    proxy-mode = "service_account"
    terraform  = "true"
  }
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-gpu"
  }
  install_gpu_driver = true
  accelerator_config {
    type         = "NVIDIA_TESLA_T4"
    core_count   = 1
  }
}
`, name)
}
