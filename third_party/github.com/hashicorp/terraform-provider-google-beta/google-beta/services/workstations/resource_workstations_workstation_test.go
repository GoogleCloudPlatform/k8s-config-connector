// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package workstations_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWorkstationsWorkstation_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstation_basic(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstation_modified(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func testAccWorkstationsWorkstation_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		             = "us-central1"

  labels = {
	  foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"

  labels = {
	  foo = "bar"
  }
}

resource "google_workstations_workstation" "default" {
  workstation_id         = "tf-test-workstation%{random_suffix}"
  workstation_config_id  =  google_workstations_workstation_config.default.workstation_config_id
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"

  labels = {
	  foo = "bar"
  }

  env = {
    name = "bar"
  }
}
`, context)
}

func testAccWorkstationsWorkstation_modified(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		             = "us-central1"

  labels = {
	  foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"

  labels = {
	  foo = "bar"
  }
}

resource "google_workstations_workstation" "default" {
  workstation_id         = "tf-test-workstation%{random_suffix}"
  workstation_config_id  =  google_workstations_workstation_config.default.workstation_config_id
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		         = "us-central1"
  display_name           = "workstation%{random_suffix}"

  labels = {
	  foo = "bar"
  }

  env = {
    name = "test"
  }
}
`, context)
}
