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
  name     = "cloudrun-service"
  location = "us-central1"
  launch_stage = "BETA"
  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
    vpc_access{
      network_interfaces {
        network = "default"
        subnetwork = "default"
        tags = ["tag1", "tag2", "tag3"]
      }
      egress = "ALL_TRAFFIC"
    }
  }
}
```
