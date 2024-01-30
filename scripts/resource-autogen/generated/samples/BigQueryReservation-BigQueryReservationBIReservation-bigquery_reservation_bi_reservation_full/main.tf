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
  table_id   = "table_%{random_suffix}"
}

resource "google_bigquery_table" "foo2" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.bar2.dataset_id
  table_id   = "table2_%{random_suffix}"
}

resource "google_bigquery_dataset" "bar" {
  dataset_id                  = "dataset_%{random_suffix}"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
}

resource "google_bigquery_dataset" "bar2" {
  dataset_id                  = "dataset2_%{random_suffix}"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
}

resource "google_bigquery_bi_reservation" "reservation" {
	location  = "EU"
	size      = "2800000000"
  preferred_tables {
      project_id  = "my-project-name"
      dataset_id  = google_bigquery_dataset.bar.dataset_id
      table_id    = google_bigquery_table.foo.table_id
  }
  preferred_tables {
      project_id  = "my-project-name"
      dataset_id  = google_bigquery_dataset.bar2.dataset_id
      table_id    = google_bigquery_table.foo2.table_id
  }
}
```
