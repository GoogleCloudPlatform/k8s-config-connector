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
resource "google_network_management_connectivity_test" "instance-test" {
  name = "conn-test-instances"
  source {
    instance = google_compute_instance.source.id
  }

  destination {
    instance = google_compute_instance.destination.id
  }

  protocol = "TCP"
}

resource "google_compute_instance" "source" {
  name = "source-vm"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.debian_9.id
    }
  }

  network_interface {
    network = google_compute_network.vpc.id
    access_config {
    }
  }
}

resource "google_compute_instance" "destination" {
  name = "dest-vm"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.debian_9.id
    }
  }

  network_interface {
    network = google_compute_network.vpc.id
    access_config {
    }
  }
}

resource "google_compute_network" "vpc" {
  name = "conn-test-net"
}

data "google_compute_image" "debian_9" {
  family  = "debian-11"
  project = "debian-cloud"
}
```
