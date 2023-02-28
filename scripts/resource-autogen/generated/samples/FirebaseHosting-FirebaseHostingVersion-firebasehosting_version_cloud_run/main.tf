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

resource "google_cloud_run_v2_service" "default" {
  provider = google-beta
  project  = "my-project-name"
  name     = "cloud-run-service-via-hosting"
  location = "us-central1"

  # Warning: allows all public traffic
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}

resource "google_firebase_hosting_version" "default" {
  provider = google-beta
  site_id  = google_firebase_hosting_site.default.site_id
  config {
    rewrites {
      glob = "/hello/**"
      run {
        service_id = google_cloud_run_v2_service.default.name
        region = google_cloud_run_v2_service.default.location
      }
    }
  }
}

resource "google_firebase_hosting_release" "default" {
  provider     = google-beta
  site_id      = google_firebase_hosting_site.default.site_id
  version_name = google_firebase_hosting_version.default.name
  message      = "Cloud Run Integration"
}
```
