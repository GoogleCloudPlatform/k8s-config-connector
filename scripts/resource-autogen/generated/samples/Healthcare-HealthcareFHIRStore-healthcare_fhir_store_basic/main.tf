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
resource "google_healthcare_fhir_store" "default" {
  name    = "example-fhir-store"
  dataset = google_healthcare_dataset.dataset.id
  version = "R4"
  complex_data_type_reference_parsing = "DISABLED"

  enable_update_create           = false
  disable_referential_integrity  = false
  disable_resource_versioning    = false
  enable_history_import          = false
  default_search_handling_strict = false

  notification_config {
    pubsub_topic = google_pubsub_topic.topic.id
  }

  labels = {
    label1 = "labelvalue1"
  }
}

resource "google_pubsub_topic" "topic" {
  name     = "fhir-notifications"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "example-dataset"
  location = "us-central1"
}
```
