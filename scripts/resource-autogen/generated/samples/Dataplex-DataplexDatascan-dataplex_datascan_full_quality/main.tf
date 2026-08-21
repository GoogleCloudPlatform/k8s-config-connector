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
resource "google_dataplex_datascan" "full_quality" {
  location = "us-central1"
  display_name = "Full Datascan Quality"
  data_scan_id = "tf-test-datascan%{random_suffix}"
  description = "Example resource - Full Datascan Quality"
  labels = {
    author = "billing"
  }

  data {
    resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/austin_bikeshare/tables/bikeshare_stations"
  }

  execution_spec {
    trigger {
      schedule {
        cron = "TZ=America/New_York 1 1 * * *"
      }
    }
    field = "modified_date"
  }

  data_quality_spec {
    sampling_percent = 5
    row_filter = "station_id > 1000"
    rules {
      column = "address"
      dimension = "VALIDITY"
      threshold = 0.99
      non_null_expectation {}
    }

    rules {
      column = "council_district"
      dimension = "VALIDITY"
      ignore_null = true
      threshold = 0.9
      range_expectation {
        min_value = 1
        max_value = 10
        strict_min_enabled = true
        strict_max_enabled = false
      }
    }

    rules {
      column = "power_type"
      dimension = "VALIDITY"
      ignore_null = false
      regex_expectation {
        regex = ".*solar.*"
      }
    }

    rules {
      column = "property_type"
      dimension = "VALIDITY"
      ignore_null = false
      set_expectation {
        values = ["sidewalk", "parkland"]
      }
    }


    rules {
      column = "address"
      dimension = "UNIQUENESS"
      uniqueness_expectation {}
    }

    rules {
      column = "number_of_docks"
      dimension = "VALIDITY"
      statistic_range_expectation {
        statistic = "MEAN"
        min_value = 5
        max_value = 15
        strict_min_enabled = true
        strict_max_enabled = true
      }
    }

    rules {
      column = "footprint_length"
      dimension = "VALIDITY"
      row_condition_expectation {
        sql_expression = "footprint_length > 0 AND footprint_length <= 10"
      }
    }

    rules {
      dimension = "VALIDITY"
      table_condition_expectation {
        sql_expression = "COUNT(*) > 0"
      }
    }
  }


  project = "my-project-name"
}
```
