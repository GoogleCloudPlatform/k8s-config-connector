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
resource "google_compute_instance" "vm" {
  provider     = google-beta
  name         = "my-vm"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_machine_image" "image" {
  provider        = google-beta
  name            = "my-image"
  source_instance = google_compute_instance.vm.self_link
  machine_image_encryption_key {
    kms_key_name = google_kms_crypto_key.crypto_key.id
  }
}

resource "google_kms_crypto_key" "crypto_key" {
  provider = google-beta
  name     = "key"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  provider = google-beta
  name     = "keyring"
  location = "us"
}
```
