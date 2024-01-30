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
resource "google_compute_region_backend_service" "default" {
  provider                        = google-beta
  name                            = "region-service"
  region                          = "us-central1"
  health_checks                   = [google_compute_region_health_check.health_check.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
  protocol                        = "TCP"
  load_balancing_scheme           = "EXTERNAL"
  connection_tracking_policy {
    tracking_mode                                = "PER_SESSION"
    connection_persistence_on_unhealthy_backends = "NEVER_PERSIST"
    idle_timeout_sec                             = 60
    enable_strong_affinity                       = true
  }
}

resource "google_compute_region_health_check" "health_check" {
  provider           = google-beta
  name               = "rbs-health-check"
  region             = "us-central1"

  tcp_health_check {
    port = 22
  }
}
```
