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
resource "google_compute_node_template" "soletenant-tmpl" {
  name      = "soletenant-tmpl"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}

resource "google_compute_node_group" "nodes" {
  name        = "soletenant-group"
  zone        = "us-central1-f"
  description = "example google_compute_node_group for Terraform Google Provider"
  maintenance_policy = "RESTART_IN_PLACE"
  maintenance_window {
    start_time = "08:00"
  }
  initial_size  = 1
  node_template = google_compute_node_template.soletenant-tmpl.id
  autoscaling_policy {
    mode      = "ONLY_SCALE_OUT"
    min_nodes = 1
    max_nodes = 10
  }
}
```
