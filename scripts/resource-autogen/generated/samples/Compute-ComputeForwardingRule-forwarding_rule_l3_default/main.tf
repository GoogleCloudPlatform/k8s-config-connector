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
resource "google_compute_forwarding_rule" "fwd_rule" {
  provider        = google-beta
  name            = "l3-forwarding-rule"
  backend_service = google_compute_region_backend_service.service.id
  ip_protocol     = "L3_DEFAULT"
  all_ports       = true
}

resource "google_compute_region_backend_service" "service" {
  provider              = google-beta
  region                = "us-central1"
  name                  = "service"
  health_checks         = [google_compute_region_health_check.health_check.id]
  protocol              = "UNSPECIFIED"
  load_balancing_scheme = "EXTERNAL"
}

resource "google_compute_region_health_check" "health_check" {
  provider           = google-beta
  name               = "health-check"
  region             = "us-central1"

  tcp_health_check {
    port = 80
  }
}
```
