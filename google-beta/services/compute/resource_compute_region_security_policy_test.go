// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeRegionSecurityPolicy_regionSecurityPolicyBasicUpdateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicy_basic(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy.regionSecPolicy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicy_update(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy.regionSecPolicy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicy_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "regionSecPolicy" {
  name        = "tf-test%{random_suffix}"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR_NETWORK"

  ddos_protection_config {
    ddos_protection = "STANDARD"
  }
}
`, context)
}

func testAccComputeRegionSecurityPolicy_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "regionSecPolicy" {
  name        = "tf-test%{random_suffix}"
  description = "basic update region security policy"
  type        = "CLOUD_ARMOR_NETWORK"

  ddos_protection_config {
    ddos_protection = "ADVANCED"
  }
}
`, context)
}

func TestAccComputeRegionSecurityPolicy_regionSecurityPolicyUserDefinedFieldsUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicy_withoutUserDefinedFields(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy.regionSecPolicy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicy_withUserDefinedFields(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy.regionSecPolicy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicy_withUserDefinedFieldsUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy.regionSecPolicy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicy_withoutUserDefinedFields(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy.regionSecPolicy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicy_withoutUserDefinedFields(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "regionSecPolicy" {
  name        = "tf-test%{random_suffix}"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR_NETWORK"
}
`, context)
}

func testAccComputeRegionSecurityPolicy_withUserDefinedFields(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "regionSecPolicy" {
  name        = "tf-test%{random_suffix}"
  description = "basic update region security policy"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
}
`, context)
}

func testAccComputeRegionSecurityPolicy_withUserDefinedFieldsUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "regionSecPolicy" {
  name        = "tf-test%{random_suffix}"
  description = "basic update region security policy"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "UDP"
    offset = 4
    size = 4
    mask = "0xFFFF"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8700"
  }
}
`, context)
}
