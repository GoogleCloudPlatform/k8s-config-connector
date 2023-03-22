package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWorkstationsWorkstationConfig_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
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
	return Nprintf(`
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
		"random_suffix": RandString(t, 10),
		"display_name":  "Display Name N",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
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
	return Nprintf(`
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
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
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
	return Nprintf(`
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
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationConfigDestroyProducer(t),
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
		},
	})
}

func testAccWorkstationsWorkstationConfig_update(context map[string]interface{}) string {
	return Nprintf(`
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
      machine_type                = "e2-standard-4"
      boot_disk_size_gb           = 35
      disable_public_ip_addresses = true
    }
  }

  labels = {
	foo = "bar"
  }
}
`, context)
}
