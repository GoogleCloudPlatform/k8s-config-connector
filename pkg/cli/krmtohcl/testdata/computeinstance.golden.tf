resource "google_compute_instance" "computetargetpool_dep4" {
  boot_disk {
    auto_delete = true
    device_name = "persistent-disk-0"

    initialize_params {
      image = "https://www.googleapis.com/compute/beta/projects/debian-cloud/global/images/debian-9-stretch-v20200910"
      size  = 10
      type  = "pd-standard"
    }

    mode   = "READ_WRITE"
    source = "https://www.googleapis.com/compute/v1/projects/my-project/zones/us-central1-f/disks/computetargetpool-dep4"
  }

  labels = {
    cnrm-lease-expiration = "1603985453"
    cnrm-lease-holder-id  = "btpp498colih6qs1pe5g"
  }

  machine_type            = "n1-standard-1"
  metadata_startup_script = "echo \"$${test} %%{test}\" > /test.txt"
  name                    = "computetargetpool-dep4"

  network_interface {
    network            = "https://www.googleapis.com/compute/v1/projects/my-project/global/networks/computetargetpool-dep"
    network_ip         = "10.2.0.5"
    subnetwork         = "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/subnetworks/computetargetpool-dep"
    subnetwork_project = "my-project"
  }

  project = "my-project"

  scheduling {
    automatic_restart   = true
    on_host_maintenance = "MIGRATE"
  }

  zone = "us-central1-f"
}
# terraform import google_compute_instance.computetargetpool_dep4 projects/my-project/zones/us-central1-f/instances/computetargetpool-dep4
