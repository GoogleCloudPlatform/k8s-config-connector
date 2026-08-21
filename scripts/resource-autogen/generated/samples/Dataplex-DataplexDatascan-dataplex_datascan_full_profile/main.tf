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
resource "google_dataplex_datascan" "full_profile" {
  location     = "us-central1"
  display_name = "Full Datascan Profile"
  data_scan_id = "tf-test-datascan%{random_suffix}"
  description  = "Example resource - Full Datascan Profile"
  labels = {
    author = "billing"
  }

  data {
    resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      schedule {
        cron = "TZ=America/New_York 1 1 * * *"
      }
    }
  }

  data_profile_spec {
    sampling_percent = 80
    row_filter = "word_count > 10"
  }

  project = "my-project-name"
}
```
