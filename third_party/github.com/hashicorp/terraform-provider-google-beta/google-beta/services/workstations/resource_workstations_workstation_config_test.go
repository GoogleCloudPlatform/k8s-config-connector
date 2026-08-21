// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package workstations_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWorkstationsWorkstationConfig_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_basic(context),
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

func testAccWorkstationsWorkstationConfig_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_displayName(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"display_name":  "Display Name N",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_displayName(context, ""),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_displayName(context, "2"),
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

func testAccWorkstationsWorkstationConfig_displayName(context map[string]interface{}, update string) string {
	context["display_name"] = context["display_name"].(string) + update
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"
  display_name           = "%{display_name} %{random_suffix}"

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_persistentDirectories(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_persistentDirectories(context),
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

func testAccWorkstationsWorkstationConfig_persistentDirectories(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  persistent_directories {
	mount_path = "/home"
  }

  labels = {
	foo = "bar"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigBasicExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_update(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_workstationConfigBasicExample(context),
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

func testAccWorkstationsWorkstationConfig_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location   		         = "us-central1"

  labels = {
	foo = "bar"
  }
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location   		     = "us-central1"

  host {
    gce_instance {
      machine_type                 = "n1-standard-4"
      boot_disk_size_gb            = 35
      disable_public_ip_addresses  = true
      enable_nested_virtualization = true
    }
  }

  labels = {
	foo = "bar"
  }

  lifecycle {
    prevent_destroy = true
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_updateHostDetails(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_updateHostDetailsDefault(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_updateHostDetailsUpdated(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_updateHostDetailsUnsetInstanceConfigs(context),
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

func testAccWorkstationsWorkstationConfig_updateHostDetailsDefault(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  host {
    gce_instance {
      machine_type                = "e2-standard-2"
      boot_disk_size_gb           = 35
      pool_size                   = 0

      disable_public_ip_addresses = false

      shielded_instance_config {
        enable_secure_boot          = false
        enable_vtpm                 = false
        enable_integrity_monitoring = false
      }

      confidential_instance_config {
        enable_confidential_compute = false
      }
    }
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_updateHostDetailsUpdated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  host {
    gce_instance {
      machine_type                = "n2d-standard-2"
      boot_disk_size_gb           = 35
      pool_size                   = 1

      disable_public_ip_addresses = true
      tags = ["foo", "bar"]

      shielded_instance_config {
        enable_secure_boot          = true
        enable_vtpm                 = true
        enable_integrity_monitoring = true
      }

      confidential_instance_config {
        enable_confidential_compute = true
      }
    }
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_updateHostDetailsUnsetInstanceConfigs(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  host {
    gce_instance {
      machine_type                = "n2d-standard-2"
      boot_disk_size_gb           = 35
      pool_size                   = 1

      disable_public_ip_addresses = true
      tags = ["foo", "bar"]

      shielded_instance_config {}
      confidential_instance_config {}
    }
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_updateWorkingDir(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_withCustomWorkingDir(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_unsetWorkingDir(context),
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

func testAccWorkstationsWorkstationConfig_withCustomWorkingDir(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  container {
    image       = "us-central1-docker.pkg.dev/cloud-workstations-images/predefined/code-oss:latest"
    working_dir = "/test"
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_unsetWorkingDir(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  container {
    image       = "us-central1-docker.pkg.dev/cloud-workstations-images/predefined/code-oss:latest"
  }
}
`, context)
}

func TestAccWorkstationsWorkstationConfig_updatePersistentDirectorySourceSnapshot(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationConfig_withSourceDiskSnapshot(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccWorkstationsWorkstationConfig_withUpdatedSourceDiskSnapshot(context),
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

func testAccWorkstationsWorkstationConfig_withSourceDiskSnapshot(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_compute_disk" "test_source_disk" {
  provider = google-beta
  name     = "tf-test-workstation-source-disk%{random_suffix}"
  size     = 10
  type     = "pd-ssd"
  zone     = "us-central1-a"
}

resource "google_compute_snapshot" "test_source_snapshot" {
  provider    = google-beta
  name        = "tf-test-workstation-source-snapshot%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  zone        = "us-central1-a"
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  persistent_directories {
    mount_path = "/home"

    gce_pd {
      source_snapshot = google_compute_snapshot.test_source_snapshot.id
      reclaim_policy  = "DELETE"
    }
  }
}
`, context)
}

func testAccWorkstationsWorkstationConfig_withUpdatedSourceDiskSnapshot(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_compute_disk" "test_source_disk" {
  provider = google-beta
  name     = "tf-test-workstation-source-disk%{random_suffix}"
  size     = 10
  type     = "pd-ssd"
  zone     = "us-central1-a"
}

resource "google_compute_snapshot" "test_source_snapshot" {
  provider    = google-beta
  name        = "tf-test-workstation-source-snapshot%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  zone        = "us-central1-a"
}

resource "google_compute_snapshot" "test_source_snapshot2" {
  provider    = google-beta
  name        = "tf-test-workstation-source-snapshot2%{random_suffix}"
  source_disk = google_compute_disk.test_source_disk.name
  zone        = "us-central1-a"
}

resource "google_workstations_workstation_cluster" "default" {
  provider                   = google-beta
  workstation_cluster_id     = "tf-test-workstation-cluster%{random_suffix}"
  network                    = google_compute_network.default.id
  subnetwork                 = google_compute_subnetwork.default.id
  location                   = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "tf-test-workstation-config%{random_suffix}"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = "us-central1"

  persistent_directories {
    mount_path = "/home"

    gce_pd {
      source_snapshot = google_compute_snapshot.test_source_snapshot2.id
      reclaim_policy  = "RETAIN"
    }
  }
}
`, context)
}
