/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

resource "google_compute_instance" "foo" {
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
  can_ip_forward      = false
  deletion_protection = false
  enable_display      = false
  labels = {
    cnrm-lease-expiration = "1601998420"
    cnrm-lease-holder-id  = "btpp498colih6qs1pe5g"
  }
  machine_type = "n1-standard-1"
  name         = "computetargetpool-dep4"
  network_interface {
    network            = "https://www.googleapis.com/compute/v1/projects/my-project/global/networks/computetargetpool-dep"
    network_ip         = "10.2.0.5"
    subnetwork         = "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/subnetworks/computetargetpool-dep"
    subnetwork_project = "my-project"
  }
  project = "my-project"
  scheduling {
    automatic_restart   = true
    min_node_cpus       = 0
    on_host_maintenance = "MIGRATE"
    preemptible         = false
  }
  zone = "us-central1-f"
}
