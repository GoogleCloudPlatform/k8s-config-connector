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
resource "random_id" "url_signature" {
  byte_length = 16
}

resource "google_compute_backend_service_signed_url_key" "backend_key" {
  name            = "test-key"
  key_value       = random_id.url_signature.b64_url
  backend_service = google_compute_backend_service.example_backend.name
}

resource "google_compute_backend_service" "example_backend" {
  name        = "my-backend-service"
  description = "Our company website"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10
  enable_cdn  = true

  backend {
    group = google_compute_instance_group_manager.webservers.instance_group
  }

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_instance_group_manager" "webservers" {
  name               = "my-webservers"

  version {
    instance_template  = google_compute_instance_template.webserver.id
    name               = "primary"
  }

  base_instance_name = "webserver"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "webserver" {
  name         = "standard-webserver"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = "debian-cloud/debian-11"
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "test"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
```
