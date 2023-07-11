// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccVmwareenginePrivateCloud_vmwareEnginePrivateCloudUpdate(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"region":          "southamerica-west1", // using region with low node utilization.
		"random_suffix":   acctest.RandString(t, 10),
		"organization":    envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
	}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareenginePrivateCloudDestroyProducer(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testPrivateCloudUpdateConfig(context, "description1", 3),
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
			{
				Config: testPrivateCloudUpdateConfig(context, "description2", 4), // Expand PC
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
			{
				Config: testPrivateCloudUpdateConfig(context, "description2", 3), // Shrink PC
			},
			{
				ResourceName:            "google_vmwareengine_private_cloud.vmw-engine-pc",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name"},
			},
		},
	})
}

func testPrivateCloudUpdateConfig(context map[string]interface{}, description string, nodeCount int) string {
	context["node_count"] = nodeCount
	context["description"] = description

	return acctest.Nprintf(`
resource "google_vmwareengine_network" "default-nw" {
   provider      	   = google-beta
   project           = google_project_service.acceptance.project
   name              = "%{region}-default"
   location          = "%{region}"
   type              = "LEGACY"
}

resource "google_vmwareengine_private_cloud" "vmw-engine-pc" {
  location = "%{region}-a"
  name = "tf-test-sample-pc%{random_suffix}"
  project = google_project_service.acceptance.project
  provider = google-beta
  description = "%{description}"
  network_config {
    management_cidr = "192.168.30.0/24"
    vmware_engine_network = google_vmwareengine_network.default-nw.id
  }
  management_cluster {
    cluster_id = "tf-test-sample-mgmt-cluster-custom-core-count%{random_suffix}"
    node_type_configs {
      node_type_id = "standard-72"
      node_count = "%{node_count}"
      custom_core_count = 32
    }
  }
}

# there can be only 1 Legacy network per region for a given project, so creating new project to isolate tests.
resource "google_project" "acceptance" {
  name            = "tf-test-%{random_suffix}"
  provider        = google-beta
  project_id      = "tf-test-%{random_suffix}"
  org_id          = "%{organization}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "acceptance" {
  project  = google_project.acceptance.project_id
  provider = google-beta
  service  = "vmwareengine.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.acceptance]

  create_duration = "60s"
}
`, context)
}

func testAccCheckVmwareenginePrivateCloudDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vmwareengine_private_cloud" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}
			config := acctest.GoogleProviderConfig(t)
			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VmwareengineBasePath}}projects/{{project}}/locations/{{location}}/privateClouds/{{name}}")
			if err != nil {
				return err
			}
			billingProject := ""
			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}
			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("VmwareenginePrivateCloud still exists at %s", url)
			}
		}
		return nil
	}
}
