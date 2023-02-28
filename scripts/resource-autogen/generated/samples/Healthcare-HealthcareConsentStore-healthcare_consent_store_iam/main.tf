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
resource "google_healthcare_dataset" "dataset" {
  location = "us-central1"
  name     = "my-dataset"
}

resource "google_healthcare_consent_store" "my-consent" {
  dataset = google_healthcare_dataset.dataset.id
  name    = "my-consent-store"
}

resource "google_service_account" "test-account" {
  account_id   = "my-account"
  display_name = "Test Service Account"
}

resource "google_healthcare_consent_store_iam_member" "test-iam" {
  dataset          = google_healthcare_dataset.dataset.id
  consent_store_id = google_healthcare_consent_store.my-consent.name
  role             = "roles/editor"
  member           = "serviceAccount:${google_service_account.test-account.email}"
}
```
