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
resource "google_filestore_instance" "instance" {
  name     = "test-instance"
  location = "us-central1"
  tier     = "ENTERPRISE"

  file_shares {
    capacity_gb = 1024
    name        = "share1"
  }

  networks {
    network = "default"
    modes   = ["MODE_IPV4"]
  }

  kms_key_name = google_kms_crypto_key.filestore_key.id
}

resource "google_kms_key_ring" "filestore_keyring" {
  name     = "filestore-keyring"
  location = "us-central1"
}

resource "google_kms_crypto_key" "filestore_key" {
  name            = "filestore-key"
  key_ring        = google_kms_key_ring.filestore_keyring.id
}
```
