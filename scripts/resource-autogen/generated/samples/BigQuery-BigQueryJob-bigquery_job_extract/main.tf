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
resource "google_bigquery_table" "source-one" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.source-one.dataset_id
  table_id   = "job_extract_table"

  schema = <<EOF
[
  {
    "name": "name",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "post_abbr",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "date",
    "type": "DATE",
    "mode": "NULLABLE"
  }
]
EOF
}

resource "google_bigquery_dataset" "source-one" {
  dataset_id    = "job_extract_dataset"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
}

resource "google_storage_bucket" "dest" {
  name          = "job_extract_bucket"
  location      = "US"
  force_destroy = true
}

resource "google_bigquery_job" "job" {
  job_id     = "job_extract"

  extract {
    destination_uris = ["${google_storage_bucket.dest.url}/extract"]

    source_table {
      project_id = google_bigquery_table.source-one.project
      dataset_id = google_bigquery_table.source-one.dataset_id
      table_id   = google_bigquery_table.source-one.table_id
    }

    destination_format = "NEWLINE_DELIMITED_JSON"
    compression = "GZIP"
  }
}
```
