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
resource "google_compute_address" "scanner_static_ip" {
  provider = google-beta
  name     = "scan-basic-static-ip"
}

resource "google_security_scanner_scan_config" "scan-config" {
  provider         = google-beta
  display_name     = "terraform-scan-config"
  starting_urls    = ["http://${google_compute_address.scanner_static_ip.address}"]
  target_platforms = ["COMPUTE"]
}
```
