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
data "google_client_config" "current" {}

resource "google_kms_key_ring" "apigee_keyring" {
  name     = "apigee-keyring"
  location = "us-central1"
}

resource "google_kms_crypto_key" "apigee_key" {
  name            = "apigee-key"
  key_ring        = google_kms_key_ring.apigee_keyring.id

  lifecycle {
    prevent_destroy = true
  }
}

resource "google_project_service_identity" "apigee_sa" {
  provider = google-beta
  project  = google_project.project.project_id
  service  = google_project_service.apigee.service
}

resource "google_kms_crypto_key_iam_binding" "apigee_sa_keyuser" {
  crypto_key_id = google_kms_crypto_key.apigee_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:${google_project_service_identity.apigee_sa.email}",
  ]
}

resource "google_apigee_organization" "org" {
  analytics_region                     = "us-central1"
  display_name                         = "apigee-org"
  description                          = "Terraform-provisioned Apigee Org without VPC Peering."
  project_id                           = data.google_client_config.current.project
  disable_vpc_peering                  = true
  runtime_database_encryption_key_name = google_kms_crypto_key.apigee_key.id

  depends_on = [
    google_kms_crypto_key_iam_binding.apigee_sa_keyuser,
  ]
}
```
