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
resource "google_data_fusion_instance" "cmek" {
  name   = "my-instance"
  region = "us-central1"
  type   = "BASIC"

  crypto_key_config {
    key_reference = google_kms_crypto_key.crypto_key.id
  }

  depends_on = [google_kms_crypto_key_iam_binding.crypto_key_binding]
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "my-instance"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  name     = "my-instance"
  location = "us-central1"
}

resource "google_kms_crypto_key_iam_binding" "crypto_key_binding" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datafusion.iam.gserviceaccount.com"
  ]
}

data "google_project" "project" {}
```
