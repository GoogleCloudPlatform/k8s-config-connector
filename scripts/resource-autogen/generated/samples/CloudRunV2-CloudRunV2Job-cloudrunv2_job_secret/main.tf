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
resource "google_cloud_run_v2_job" "default" {
  name     = "cloudrun-job"
  location = "us-central1"

  template {
    template {
      volumes {
        name = "a-volume"
        secret {
          secret = google_secret_manager_secret.secret.secret_id
          default_mode = 292 # 0444
          items {
            version = "1"
            path = "my-secret"
            mode = 256 # 0400
          }
        }
      }
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
        volume_mounts {
          name = "a-volume"
          mount_path = "/secrets"
        }
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }

  depends_on = [
    google_secret_manager_secret_version.secret-version-data,
    google_secret_manager_secret_iam_member.secret-access,
  ]
}

data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret"
  replication {
    auto {}
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}
```
