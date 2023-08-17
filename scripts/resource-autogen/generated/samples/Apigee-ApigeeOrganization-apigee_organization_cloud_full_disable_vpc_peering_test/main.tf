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
resource "google_project" "project" {
  provider = google-beta

  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}

resource "google_project_service" "apigee" {
  provider = google-beta

  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  provider = google-beta

  project = google_project.project.project_id
  service = "compute.googleapis.com"
}

resource "google_project_service" "kms" {
  provider = google-beta

  project = google_project.project.project_id
  service = "cloudkms.googleapis.com"
}

resource "google_kms_key_ring" "apigee_keyring" {
  provider = google-beta

  name       = "apigee-keyring"
  location   = "us-central1"
  project    = google_project.project.project_id
  depends_on = [google_project_service.kms]
}

resource "google_kms_crypto_key" "apigee_key" {
  provider = google-beta

  name            = "apigee-key"
  key_ring        = google_kms_key_ring.apigee_keyring.id
}

resource "google_project_service_identity" "apigee_sa" {
  provider = google-beta

  project = google_project.project.project_id
  service = google_project_service.apigee.service
}

resource "google_kms_crypto_key_iam_binding" "apigee_sa_keyuser" {
  provider = google-beta

  crypto_key_id = google_kms_crypto_key.apigee_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:${google_project_service_identity.apigee_sa.email}",
  ]
}

resource "google_apigee_organization" "org" {
  provider = google-beta

  display_name                         = "apigee-org"
  description                          = "Terraform-provisioned Apigee Org without VPC Peering."
  analytics_region                     = "us-central1"
  project_id                           = google_project.project.project_id
  disable_vpc_peering                  = true
  billing_type                         = "EVALUATION"
  runtime_database_encryption_key_name = google_kms_crypto_key.apigee_key.id
  properties {
    property {
      name = "features.mart.connect.enabled"
      value = "true"
    }
    property {
      name = "features.hybrid.enabled"
      value = "true"
    }
  }

  depends_on = [
    google_kms_crypto_key_iam_binding.apigee_sa_keyuser,
  ]
}
```
