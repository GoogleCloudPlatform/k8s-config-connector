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
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}

resource "google_access_context_manager_access_level" "test-access" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/accessLevels/tf_test_chromeos_no_lock%{random_suffix}"
  title  = "tf_test_chromeos_no_lock%{random_suffix}"
  basic {
    conditions {
      device_policy {
        require_screen_lock = true
        os_constraints {
          os_type = "DESKTOP_CHROME_OS"
        }
      }
      regions = [
        "CH",
        "IT",
        "US",
      ]
    }
  }
}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/${google_project.project.org_id}"
  title  = "my policy"
}

resource "google_iam_access_boundary_policy" "example" {
  parent   = urlencode("cloudresourcemanager.googleapis.com/projects/${google_project.project.project_id}")
  name     = "my-ab-policy"
  display_name = "My AB policy"
  rules {
    description = "AB rule"
    access_boundary_rule {
      available_resource = "*"
      available_permissions = ["*"]
      availability_condition {
        title = "Access level expr"
        expression = "request.matchAccessLevels('${google_project.project.org_id}', ['${google_access_context_manager_access_level.test-access.name}'])"
      }
    }
  }
}
```
