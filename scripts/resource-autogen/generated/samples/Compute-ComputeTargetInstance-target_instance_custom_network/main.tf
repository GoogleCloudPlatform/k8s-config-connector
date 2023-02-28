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
resource "google_compute_target_instance" "custom_network" {
  provider = google-beta
  name     = "custom-network"
  instance = google_compute_instance.target-vm.id
  network  = data.google_compute_network.target-vm.self_link
}

data "google_compute_network" "target-vm" {
  provider = google-beta
  name = "default"
}

data "google_compute_image" "vmimage" {
  provider = google-beta
  family  = "debian-10"
  project = "debian-cloud"
}

resource "google_compute_instance" "target-vm" {
  provider = google-beta
  name         = "custom-network-target-vm"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.vmimage.self_link
    }
  }

  network_interface {
    network = "default"
  }
}
```
