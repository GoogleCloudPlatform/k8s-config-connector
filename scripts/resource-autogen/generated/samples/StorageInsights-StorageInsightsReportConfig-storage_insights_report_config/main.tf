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
data "google_project" "project" {
}

resource "google_storage_insights_report_config" "config" {
  display_name = "Test Report Config"
  location = "us-central1"
  frequency_options {
    frequency = "WEEKLY"
    start_date {
      day = 15
      month = 3
      year = 2050
    }
    end_date {
      day = 15
      month = 4
      year = 2050
    }
  }
  csv_options {
    record_separator = "\n"
    delimiter = ","
    header_required = false
  }
  object_metadata_report_options {
    metadata_fields = ["bucket", "name", "project"]
    storage_filters {
      bucket = google_storage_bucket.report_bucket.name
    }
    storage_destination_options {
      bucket = google_storage_bucket.report_bucket.name
      destination_path = "test-insights-reports"
    }
  }
}

resource "google_storage_bucket" "report_bucket" {
  name                        = "my-bucket"
  location                    = "us-central1"
  force_destroy               = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_member" "admin" {
  bucket = google_storage_bucket.report_bucket.name
  role   = "roles/storage.admin"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-storageinsights.iam.gserviceaccount.com"
}
```
