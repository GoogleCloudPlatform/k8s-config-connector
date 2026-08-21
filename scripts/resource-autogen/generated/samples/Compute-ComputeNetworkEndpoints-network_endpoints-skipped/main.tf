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

```hcl
resource "google_compute_network_endpoints" "default-endpoints" {
  network_endpoint_group = google_compute_network_endpoint_group.neg.name

  network_endpoints {
    instance   = google_compute_instance.endpoint-instance1.name
    port       = google_compute_network_endpoint_group.neg.default_port
    ip_address = google_compute_instance.endpoint-instance1.network_interface[0].network_ip
  }
  network_endpoints {
    instance   = google_compute_instance.endpoint-instance2.name
    port       = google_compute_network_endpoint_group.neg.default_port
    ip_address = google_compute_instance.endpoint-instance2.network_interface[0].network_ip
  }
}

data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_instance" "endpoint-instance1" {
  name         = "endpoint-instance1"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.default.id
    access_config {
    }
  }
}

resource "google_compute_instance" "endpoint-instance2" {
  name         = "endpoint-instance2"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.default.id
    access_config {
    }
  }
}

resource "google_compute_network_endpoint_group" "group" {
  name         = "my-lb-neg"
  network      = google_compute_network.default.id
  subnetwork   = google_compute_subnetwork.default.id
  default_port = "90"
  zone         = "us-central1-a"
}

resource "google_compute_network" "default" {
  name                    = "neg-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "neg-subnetwork"
  ip_cidr_range = "10.0.0.1/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
```
