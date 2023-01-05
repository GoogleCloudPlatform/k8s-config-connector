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
locals {
  project = "my-project-name" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-bq-geojson-sample"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "object" {
  name   = "geojson-data.jsonl"
  bucket = google_storage_bucket.bucket.name
  content = <<EOF
{"type":"Feature","properties":{"continent":"Europe","region":"Scandinavia"},"geometry":{"type":"Polygon","coordinates":[[[-30.94,53.33],[33.05,53.33],[33.05,71.86],[-30.94,71.86],[-30.94,53.33]]]}}
{"type":"Feature","properties":{"continent":"Africa","region":"West Africa"},"geometry":{"type":"Polygon","coordinates":[[[-23.91,0],[11.95,0],[11.95,18.98],[-23.91,18.98],[-23.91,0]]]}}
EOF
}

resource "google_bigquery_table" "foo" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.bar.dataset_id
  table_id   = "job_load_table"
}

resource "google_bigquery_dataset" "bar" {
  dataset_id                  = "job_load_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_job" "job" {
  job_id     = "job_load"

  labels = {
    "my_job" = "load"
  }

  load {
    source_uris = [
      "gs://${google_storage_bucket_object.object.bucket}/${google_storage_bucket_object.object.name}"
    ]

    destination_table {
      project_id = google_bigquery_table.foo.project
      dataset_id = google_bigquery_table.foo.dataset_id
      table_id   = google_bigquery_table.foo.table_id
    }

    write_disposition = "WRITE_TRUNCATE"
    autodetect = true
    source_format = "NEWLINE_DELIMITED_JSON"
    json_extension = "GEOJSON"
  }

  depends_on = ["google_storage_bucket_object.object"]
}
```
