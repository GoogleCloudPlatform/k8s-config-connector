// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"

	compute "google.golang.org/api/compute/v0.beta"
)

func TestAccComputeInstanceFromMachineImage_basic(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_basic(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),

					// Check that fields were set based on the template
					resource.TestCheckResourceAttr(resourceName, "machine_type", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "attached_disk.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.automatic_restart", "false"),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_overrideMetadataDotStartupScript(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_overrideMetadataDotStartupScript(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),
					resource.TestCheckResourceAttr(resourceName, "metadata.startup-script", ""),
				),
			},
		},
	})

}

func TestAccComputeInstanceFromMachineImage_diffProject(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"
	org := envvar.GetTestOrgFromEnv(t)
	billingId := envvar.GetTestBillingAccountFromEnv(t)
	projectID := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_diffProject(projectID, org, billingId, instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),

					// Check that fields were set based on the template
					resource.TestCheckResourceAttr(resourceName, "machine_type", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "attached_disk.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.automatic_restart", "false"),
				),
			},
		},
	})
}

func testAccCheckComputeInstanceFromMachineImageDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_instance_from_machine_image" {
				continue
			}

			_, err := config.NewComputeClient(config.UserAgent).Instances.Get(
				config.Project, rs.Primary.Attributes["zone"], rs.Primary.ID).Do()
			if err == nil {
				return fmt.Errorf("Instance still exists")
			}
		}

		return nil
	}
}

func testAccComputeInstanceFromMachineImage_basic(instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = true
  }

  can_ip_forward = true
}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  // Overrides
  can_ip_forward = false
  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
  }
}
`, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImage_overrideMetadataDotStartupScript(instanceName, generatedInstanceName string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    startup-script = "#!/bin/bash\necho Hello"
  }

}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  // Overrides
  metadata = {
    startup-script = ""
  }
}
`, instanceName, instanceName, generatedInstanceName)
}

func testAccComputeInstanceFromMachineImage_diffProject(projectID, org, billingId, instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_project" "project" {
	provider     = google-beta
	project_id      = "%s"
	name            = "%s"
	org_id          = "%s"
	billing_account = "%s"
}

resource "google_project_service" "service" {
	provider     = google-beta
	project = google_project.project.project_id
	service = "compute.googleapis.com"
	timeouts {
	  create = "30m"
	  update = "40m"
	}
	disable_dependent_services = true
}

resource "google_project_service" "monitoring" {
	provider     = google-beta
	project = google_project.project.project_id
	service = "monitoring.googleapis.com"
	timeouts {
	  create = "30m"
	  update = "40m"
	}
	disable_dependent_services = true

	depends_on = [google_project_service.service]
}

resource "google_compute_instance" "vm" {
  provider     = google-beta
  project = google_project.project.project_id
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = true
  }

  can_ip_forward = true

  depends_on = [google_project_service.monitoring]
}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  project = google_project.project.project_id
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  // Overrides
  can_ip_forward = false
  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
  }
}
`, projectID, projectID, org, billingId, instance, instance, newInstance)
}
