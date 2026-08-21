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
resource "google_cloud_run_v2_service" "default" {
  provider = google-beta
  name     = "cloudrun-service"
  location = "us-central1"
  launch_stage = "BETA"
  ingress = "INGRESS_TRAFFIC_ALL"
  template {
    containers {
      name = "hello-1"
      ports {
        container_port = 8080
      }
      image = "us-docker.pkg.dev/cloudrun/container/hello"
      depends_on = ["hello-2"]
      volume_mounts {
        name = "empty-dir-volume"
	mount_path = "/mnt"
      }
    }
    containers {
      name = "hello-2"
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
    volumes {
      name = "empty-dir-volume"
      empty_dir {
        medium = "MEMORY"
        size_limit = "256Mi"
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }
}
```
