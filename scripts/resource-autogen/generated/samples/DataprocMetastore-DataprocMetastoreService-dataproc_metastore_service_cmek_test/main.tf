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
data "google_project" "project" {}

data "google_storage_project_service_account" "gcs_account" {}


resource "google_dataproc_metastore_service" "default" {
  service_id = "example-service"
  location   = "us-central1"

  encryption_config {
    kms_key = google_kms_crypto_key.crypto_key.id
  }

  hive_metastore_config {
    version = "3.1.2"
  }

  depends_on = [google_kms_crypto_key_iam_binding.crypto_key_binding]
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "example-key"
  key_ring = google_kms_key_ring.key_ring.id

  purpose  = "ENCRYPT_DECRYPT"
}

resource "google_kms_key_ring" "key_ring" {
  name     = "example-keyring"
  location = "us-central1"
}

resource "google_kms_crypto_key_iam_binding" "crypto_key_binding" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-metastore.iam.gserviceaccount.com",
    "serviceAccount:${data.google_storage_project_service_account.gcs_account.email_address}"
  ]
}
```
