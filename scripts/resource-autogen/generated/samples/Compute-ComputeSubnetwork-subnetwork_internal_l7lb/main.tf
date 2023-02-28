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
resource "google_compute_subnetwork" "network-for-l7lb" {
  provider = google-beta

  name          = "l7lb-test-subnetwork"
  ip_cidr_range = "10.0.0.0/22"
  region        = "us-central1"
  purpose       = "INTERNAL_HTTPS_LOAD_BALANCER"
  role          = "ACTIVE"
  network       = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  provider = google-beta

  name                    = "l7lb-test-network"
  auto_create_subnetworks = false
}
```
