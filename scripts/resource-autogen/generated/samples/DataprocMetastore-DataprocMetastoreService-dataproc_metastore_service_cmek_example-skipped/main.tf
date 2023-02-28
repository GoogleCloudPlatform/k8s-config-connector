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
resource "google_dataproc_metastore_service" "default" {
  service_id = "example-service"
  location   = "us-central1"

  encryption_config {
    kms_key = google_kms_crypto_key.crypto_key.id
  }

  hive_metastore_config {
    version = "3.1.2"
  }
}

resource "google_kms_crypto_key" "crypto_key" {
  provider = google-beta
  name     = "example-key"
  key_ring = google_kms_key_ring.key_ring.id

  purpose  = "ENCRYPT_DECRYPT"
}

resource "google_kms_key_ring" "key_ring" {
  provider = google-beta
  name     = "example-keyring"
  location = "us-central1"
}
```
