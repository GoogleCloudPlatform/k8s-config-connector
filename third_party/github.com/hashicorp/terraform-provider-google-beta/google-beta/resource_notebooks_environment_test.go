// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotebooksEnvironment_create(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("%d", RandInt(t))
	name := fmt.Sprintf("tf-env-%s", prefix)

	VcrTest(t, resource.TestCase{
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksEnvironment_create(name),
			},
			{
				ResourceName:      "google_notebooks_environment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNotebooksEnvironment_create(name string) string {
	return fmt.Sprintf(`

resource "google_notebooks_environment" "test" {
  name = "%s"
  location = "us-west1-a"
  container_image {
    repository = "gcr.io/deeplearning-platform-release/base-cpu"
  }  
}
`, name)
}
