// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleBasicUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePreUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePostUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePreUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "default" {
  region      = "us-west2"
  name        = "tf-test%{random_suffix}"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.default.name
  region          = "us-west2"
  description     = "basic rule pre update"
  action          = "allow"
  priority        = 100
  preview         = false
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
    }
  }
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRulePostUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "default" {
  region      = "us-west2"
  name        = "tf-test%{random_suffix}"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule" {
  security_policy = google_compute_region_security_policy.default.name
  region          = "us-west2"
  description     = "basic rule post update"
  action          = "deny(403)"
  priority        = 100
  preview         = true
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["172.16.0.0/12"]
    }
  }
}
`, context)
}

func TestAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate(t *testing.T) {
	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionSecurityPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchBasic(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate2(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchBasic(context),
			},
			{
				ResourceName:      "google_compute_region_security_policy_rule.policy_rule_network_match",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policyddosprotection" {
  region      = "us-west2"
  name        = "tf-test-policyddos%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}
  
resource "google_compute_network_edge_security_service" "edge_sec_service" {
  region          = "us-west2"
  name            = "tf-test-edgesec%{random_suffix}"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}
  
resource "google_compute_region_security_policy" "policynetworkmatch" {
  region      = "us-west2"
  name        = "tf-test-polnetmatch%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}
  
resource "google_compute_region_security_policy_rule" "policy_rule_network_match" {
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.policynetworkmatch.name
  priority = 100
  network_match {
    src_ip_ranges    = ["10.10.0.0/16"]
  }
  action = "allow"
  preview = true
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policyddosprotection" {
  region      = "us-west2"
  name        = "tf-test-policyddos%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}
  
resource "google_compute_network_edge_security_service" "edge_sec_service" {
  region          = "us-west2"
  name            = "tf-test-edgesec%{random_suffix}"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}
  
resource "google_compute_region_security_policy" "policynetworkmatch" {
  region      = "us-west2"
  name        = "tf-test-polnetmatch%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}
  
resource "google_compute_region_security_policy_rule" "policy_rule_network_match" {
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.policynetworkmatch.name
  priority = 100
  network_match {
    src_ip_ranges    = ["10.10.0.0/16"]
    src_asns         = [6939]
    src_ports        = [443]
    src_region_codes = ["US"]
    ip_protocols     = ["UDP"]
    dest_ip_ranges   = ["10.0.0.0/8"]
    dest_ports       = [80]
    user_defined_fields {
      name = "SIG1_AT_0"
      values = ["0x8700"]
    }
  }
  action = "allow"
  preview = true
}
`, context)
}

func testAccComputeRegionSecurityPolicyRule_regionSecurityPolicyRuleNetworkMatchUpdate2(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_security_policy" "policyddosprotection" {
  region      = "us-west2"
  name        = "tf-test-policyddos%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}
  
resource "google_compute_network_edge_security_service" "edge_sec_service" {
  region          = "us-west2"
  name            = "tf-test-edgesec%{random_suffix}"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}
  
resource "google_compute_region_security_policy" "policynetworkmatch" {
  region      = "us-west2"
  name        = "tf-test-polnetmatch%{random_suffix}"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}
  
resource "google_compute_region_security_policy_rule" "policy_rule_network_match" {
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.policynetworkmatch.name
  priority = 100
  network_match {
    src_ip_ranges    = ["10.0.0.0/8"]
    src_asns         = [15169]
    src_ports        = [80]
    src_region_codes = ["AU"]
    ip_protocols     = ["TCP"]
    dest_ip_ranges   = ["10.10.0.0/16"]
    dest_ports       = [443]
    user_defined_fields {
      name = "SIG2_AT_8"
      values = ["0x8700","0x8F00"]
    }
  }
  action = "allow"
  preview = true
}
`, context)
}
