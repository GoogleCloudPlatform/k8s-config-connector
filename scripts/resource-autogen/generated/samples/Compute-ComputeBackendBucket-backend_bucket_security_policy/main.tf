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
resource "google_compute_backend_bucket" "image_backend" {
  name        = "image-backend-bucket"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_backend.name
  enable_cdn  = true
  edge_security_policy = google_compute_security_policy.policy.id
}

resource "google_storage_bucket" "image_backend" {
  name     = "image-store-bucket"
  location = "EU"
}

resource "google_compute_security_policy" "policy" {
  name        = "image-store-bucket"
  description = "basic security policy"
	type = "CLOUD_ARMOR_EDGE"
}
```
