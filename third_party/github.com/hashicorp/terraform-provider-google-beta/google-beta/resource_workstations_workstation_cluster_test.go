package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccorkstationsWorkstationCluster_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckWorkstationsWorkstationClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationCluster_workstationClusterBasicExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag"},
			},
			{
				Config: testAccorkstationsWorkstationCluster_update(context),
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

func testAccorkstationsWorkstationCluster_update(context map[string]interface{}) string {
	return Nprintf(`
resource "google_workstations_workstation_cluster" "default" {
  provider   		      	 = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location   		        = "us-central1"

  labels = {
	  update = "true"
  }
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}
`, context)
}
