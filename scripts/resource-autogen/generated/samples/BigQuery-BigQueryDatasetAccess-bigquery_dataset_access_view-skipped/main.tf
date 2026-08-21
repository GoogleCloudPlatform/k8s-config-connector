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
resource "google_bigquery_dataset_access" "access" {
  dataset_id    = google_bigquery_dataset.private.dataset_id
  view {
    project_id = google_bigquery_table.public.project
    dataset_id = google_bigquery_dataset.public.dataset_id
    table_id   = google_bigquery_table.public.table_id
  }
}

resource "google_bigquery_dataset" "private" {
  dataset_id = "example_dataset"
}

resource "google_bigquery_dataset" "public" {
  dataset_id = "example_dataset2"
}

resource "google_bigquery_table" "public" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.public.dataset_id
  table_id   = "example_table"

  view {
    query          = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]"
    use_legacy_sql = false
  }
}
```
