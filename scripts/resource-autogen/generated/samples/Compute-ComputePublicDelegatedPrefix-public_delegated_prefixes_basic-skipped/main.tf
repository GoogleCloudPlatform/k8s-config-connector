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
resource "google_compute_public_advertised_prefix" "advertised" {
  name = "my-prefix"
  description = "description"
  dns_verification_ip = "127.127.0.0"
  ip_cidr_range = "127.127.0.0/16"
}

resource "google_compute_public_delegated_prefix" "prefixes" {
  name = "my-prefix"
  region = "us-central1"
  description = "my description"
  ip_cidr_range = "127.127.0.0/24"
  parent_prefix = google_compute_public_advertised_prefix.advertised.id
}
```
