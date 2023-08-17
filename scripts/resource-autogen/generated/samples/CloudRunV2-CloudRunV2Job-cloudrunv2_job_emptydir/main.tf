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
  provider = google-beta
  name     = "cloudrun-job"
  location = "us-central1"
  launch_stage = "BETA"
  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
	volume_mounts {
	  name = "empty-dir-volume"
	  mount_path = "/mnt"
	}
      }
      volumes {
        name = "empty-dir-volume"
	empty_dir {
	  medium = "MEMORY"
	  size_limit = "128Mi"
	}
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
