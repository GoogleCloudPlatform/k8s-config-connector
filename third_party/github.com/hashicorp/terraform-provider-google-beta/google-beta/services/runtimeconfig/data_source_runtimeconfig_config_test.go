// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package runtimeconfig_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccRuntimeconfigConfigDatasource_basic(t *testing.T) {
	t.Parallel()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRuntimeconfigDatasourceConfig(acctest.RandString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_runtimeconfig_config.default", "google_runtimeconfig_config.default"),
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
