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
resource "google_deployment_manager_deployment" "deployment" {
  name = "my-deployment"
  target {
    config {
      content = <<EOF
imports:
- path: service_account.jinja
- path: vm.jinja

resources:
- name: &SA_NAME my-vm-account
  type: service_account.jinja
- name: my-vm
  type: vm.jinja
  properties:
    serviceAccountId: *SA_NAME
EOF

    imports {
      name = "vm.jinja"
      content = file("path/to/vm.jinja")
    }

    imports {
      name = "service_account.jinja"
      content = file("path/to/service_account.jinja")
    }
  }
}
```
