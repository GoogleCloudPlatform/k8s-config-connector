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
resource "google_dataproc_metastore_service" "dpms2" {
  service_id = "dpms2"
  location   = "us-central1"

  # DPMS 2 requires SPANNER database type, and does not require
  # a maintenance window.
  database_type = "SPANNER"

  hive_metastore_config {
    version           = "3.1.2"
  }

  scaling_config {
    instance_size = "EXTRA_SMALL"
  }
}
```
