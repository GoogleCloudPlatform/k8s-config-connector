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
// Cloud Functions Example
resource "google_compute_region_network_endpoint_group" "function_neg" {
  name                  = "function-neg"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  cloud_function {
    function = google_cloudfunctions_function.function_neg.name
  }
}

resource "google_cloudfunctions_function" "function_neg" {
  name        = "function-neg"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = true
  timeout               = 60
  entry_point           = "helloGET"
}

resource "google_storage_bucket" "bucket" {
  name     = "cloudfunctions-function-example-bucket"
  location = "US"
}

resource "google_storage_bucket_object" "archive" { 
  name   = "index.zip"
  bucket = google_storage_bucket.bucket.name
  source = "path/to/index.zip"
}
```
