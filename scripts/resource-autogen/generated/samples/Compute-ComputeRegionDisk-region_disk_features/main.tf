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
resource "google_compute_region_disk" "regiondisk" {
  name                      = "my-region-features-disk"
  type                      = "pd-ssd"
  region                    = "us-central1"
  physical_block_size_bytes = 4096

  guest_os_features {
    type = "SECURE_BOOT"
  }

  guest_os_features {
    type = "MULTI_IP_SUBNET"
  }

  guest_os_features {
    type = "WINDOWS"
  }

  licenses = ["https://www.googleapis.com/compute/v1/projects/windows-cloud/global/licenses/windows-server-core"]

  replica_zones = ["us-central1-a", "us-central1-f"]
}
```
