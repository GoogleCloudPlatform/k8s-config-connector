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
resource "google_healthcare_dicom_store" "default" {
  provider = google-beta

  name    = "example-dicom-store"
  dataset = google_healthcare_dataset.dataset.id

  notification_config {
    pubsub_topic = google_pubsub_topic.topic.id
  }

  labels = {
    label1 = "labelvalue1"
  }

  stream_configs {
    bigquery_destination {
      table_uri = "bq://${google_bigquery_dataset.bq_dataset.project}.${google_bigquery_dataset.bq_dataset.dataset_id}.${google_bigquery_table.bq_table.table_id}"
    }
  }  
}

resource "google_pubsub_topic" "topic" {
  provider = google-beta

  name     = "dicom-notifications"
}

resource "google_healthcare_dataset" "dataset" {
  provider = google-beta

  name     = "example-dataset"
  location = "us-central1"
}

resource "google_bigquery_dataset" "bq_dataset" {
  provider = google-beta

  dataset_id    = "dicom_bq_ds"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
  delete_contents_on_destroy = true
}

resource "google_bigquery_table" "bq_table" {
  provider = google-beta

  deletion_protection = false
  dataset_id = google_bigquery_dataset.bq_dataset.dataset_id
  table_id   = "dicom_bq_tb"
}
```
