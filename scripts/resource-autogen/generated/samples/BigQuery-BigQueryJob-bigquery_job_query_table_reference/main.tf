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
resource "google_bigquery_table" "foo" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.bar.dataset_id
  table_id   = "job_query_table"
}

resource "google_bigquery_dataset" "bar" {
  dataset_id                  = "job_query_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_job" "job" {
  job_id     = "job_query"

  labels = {
    "example-label" ="example-value"
  }

  query {
    query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]"

    destination_table {
      table_id = google_bigquery_table.foo.id
    }

    default_dataset {
      dataset_id = google_bigquery_dataset.bar.id
    }

    allow_large_results = true
    flatten_results = true

    script_options {
      key_result_statement = "LAST"
    }
  }
}
```
