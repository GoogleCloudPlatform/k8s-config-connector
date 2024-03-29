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

resource "google_filestore_instance" "instance" {
  name = "tf-fs-inst"
  location = "us-central1-b"
  tier = "BASIC_SSD"

  file_shares {
    capacity_gb = 2560
    name        = "share1"
  }

  networks {
    network = "default"
    modes   = ["MODE_IPV4"]
    connect_mode = "DIRECT_PEERING"
  }
}

resource "google_filestore_backup" "backup" {
  name        = "tf-fs-bkup"
  location    = "us-central1"
  source_instance   = google_filestore_instance.instance.id
  source_file_share = "share1"

  description = "This is a filestore backup for the test instance"
  labels = {
    "files":"label1",
    "other-label": "label2"
  }
}
```
