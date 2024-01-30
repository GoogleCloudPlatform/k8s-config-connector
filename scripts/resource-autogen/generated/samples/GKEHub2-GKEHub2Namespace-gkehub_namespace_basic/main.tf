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
resource "google_gke_hub_scope" "namespace" {
  scope_id = "tf-test-scope%{random_suffix}"
}


resource "google_gke_hub_namespace" "namespace" { 
  scope_namespace_id = "tf-test-namespace%{random_suffix}"
  scope_id = "tf-test-scope%{random_suffix}"
  scope = "${google_gke_hub_scope.namespace.name}"
  namespace_labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec"
  }
  labels = {
      keyb = "valueb"
      keya = "valuea"
      keyc = "valuec" 
  }
  depends_on = [google_gke_hub_scope.namespace]
}
```
