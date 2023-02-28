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
resource "google_compute_network_endpoint_group" "neg" {
  name                  = "my-lb-neg"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "NON_GCP_PRIVATE_IP_PORT"
}

resource "google_compute_network_endpoint" "default-endpoint" {
  network_endpoint_group = google_compute_network_endpoint_group.neg.name
  port = google_compute_network_endpoint_group.neg.default_port
  ip_address = "127.0.0.1"
}

resource "google_compute_network" "default" {
  name = "neg-network"
}
```
