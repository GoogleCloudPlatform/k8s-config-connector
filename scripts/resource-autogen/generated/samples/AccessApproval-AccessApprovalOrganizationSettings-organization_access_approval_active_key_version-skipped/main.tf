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
resource "google_project" "my_project" {
  name       = "My Project"
  project_id = "your-project-id"
  org_id     = "123456789"
}

resource "google_kms_key_ring" "key_ring" {
  name     = "key-ring"
  location = "global"
  project  = google_project.my_project.project_id
}

resource "google_kms_crypto_key" "crypto_key" {
  name = "crypto-key"
  key_ring = google_kms_key_ring.key_ring.id
  purpose = "ASYMMETRIC_SIGN"

  version_template {
    algorithm = "EC_SIGN_P384_SHA384"
  }
}

data "google_access_approval_organization_service_account" "service_account" {
  organization_id = "123456789"
}

resource "google_kms_crypto_key_iam_member" "iam" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.signerVerifier"
  member        = "serviceAccount:${data.google_access_approval_organization_service_account.service_account.account_email}"
}

data "google_kms_crypto_key_version" "crypto_key_version" {
  crypto_key = google_kms_crypto_key.crypto_key.id
}

resource "google_organization_access_approval_settings" "organization_access_approval" {
  organization_id     = "123456789"
  active_key_version  = data.google_kms_crypto_key_version.crypto_key_version.name

  enrolled_services {
  	cloud_product = "all"
  }

  depends_on = [google_kms_crypto_key_iam_member.iam]
}
```
