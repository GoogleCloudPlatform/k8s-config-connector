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
resource "google_game_services_game_server_cluster" "default" {
    
  cluster_id = ""
  realm_id   = google_game_services_realm.default.realm_id

  connection_info {
    gke_cluster_reference {
      cluster = "locations/us-west1/clusters/%{agones_cluster}"
    }
    namespace = "default"
  }
}

resource "google_game_services_realm" "default" {
  realm_id   = "realm"
  time_zone  = "PST8PDT"

  description = "Test Game Realm"
}
```
