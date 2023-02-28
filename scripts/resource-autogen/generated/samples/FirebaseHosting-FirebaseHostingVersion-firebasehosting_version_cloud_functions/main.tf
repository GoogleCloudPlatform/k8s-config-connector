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
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "my-project-name"
  site_id  = "site-id"
}

resource "google_storage_bucket" "bucket" {
  provider = google-beta
  project  = "my-project-name"
  name     = "site-id-function-source"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "object" {
  provider = google-beta
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "function-source.zip"  # Add path to the zipped function source code
}

resource "google_cloudfunctions_function" "function" {
  provider = google-beta
  project  = "my-project-name"

  name        = "cloud-function-via-hosting"
  description = "A Cloud Function connected to Firebase Hosing"
  runtime     = "nodejs16"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.object.name
  trigger_http          = true
  entry_point           = "helloHttp"
}

resource "google_firebase_hosting_version" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  config {
    rewrites {
      glob = "/hello/**"
      function = google_cloudfunctions_function.function.name
    }
  }
}

resource "google_firebase_hosting_release" "default" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.default.site_id
  version_name = google_firebase_hosting_version.default.name
  message      = "Cloud Functions Integration"
}
```
