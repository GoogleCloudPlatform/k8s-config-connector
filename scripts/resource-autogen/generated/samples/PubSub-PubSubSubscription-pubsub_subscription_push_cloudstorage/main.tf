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
resource "google_storage_bucket" "example" {
  name  = "example-bucket"
  location = "US"
  uniform_bucket_level_access = true
}

resource "google_pubsub_topic" "example" {
  name = "example-topic"
}

resource "google_pubsub_subscription" "example" {
  name  = "example-subscription"
  topic = google_pubsub_topic.example.name

  cloud_storage_config {
    bucket = google_storage_bucket.example.name

    filename_prefix = "pre-"
    filename_suffix = "-%{random_suffix}"
  
    max_bytes = 1000
    max_duration = "300s"
  }
  depends_on = [ 
    google_storage_bucket.example,
    google_storage_bucket_iam_member.admin,
  ]
}

data "google_project" "project" {
}

resource "google_storage_bucket_iam_member" "admin" {
  bucket = google_storage_bucket.example.name
  role   = "roles/storage.admin"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}
```
