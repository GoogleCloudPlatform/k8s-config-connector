// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package tpuv2_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccTpuV2Vm_update(t *testing.T) {
	t.Parallel()

	randomSuffix := acctest.RandString(t, 10)
	context := map[string]interface{}{
		"random_suffix": randomSuffix,
		"description":   "Old description original context",
	}
	updatedContext := map[string]interface{}{
		"random_suffix": randomSuffix,
		"description":   "New description updated context",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckTpuV2VmDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccTpuV2Vm_full(context),
			},
			{
				ResourceName:            "google_tpu_v2_vm.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
			{
				Config: testAccTpuV2Vm_update(updatedContext),
			},
			{
				ResourceName:            "google_tpu_v2_vm.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
		},
	})
}

func testAccTpuV2Vm_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_tpu_v2_runtime_versions" "available" {
  provider = google-beta
}

data "google_tpu_v2_accelerator_types" "available" {
  provider = google-beta
}

resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name = "tf-test-test-tpu%{random_suffix}"
  zone = "us-central1-c"
  description = "%{description}"

  runtime_version  = "tpu-vm-tf-2.13.0"
  accelerator_type = "v2-8"
}
`, context)
}

func testAccTpuV2Vm_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_tpu_v2_runtime_versions" "available" {
  provider = google-beta
}

data "google_tpu_v2_accelerator_types" "available" {
  provider = google-beta
}

resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name = "tf-test-test-tpu%{random_suffix}"
  zone = "us-central1-c"
  description = "%{description}"

  runtime_version  = "tpu-vm-tf-2.13.0"
  accelerator_type = "v2-8"
}
`, context)
}
