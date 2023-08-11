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
resource "google_access_context_manager_access_level" "access-level-service-account" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/accessLevels/chromeos_no_lock"
  title  = "chromeos_no_lock"
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

  lifecycle {
    ignore_changes = [basic.0.conditions]
  }
}

resource "google_service_account" "created-later" {
  account_id = "my-account-id"
}

resource "google_access_context_manager_access_level_condition" "access-level-conditions" {
  access_level = google_access_context_manager_access_level.access-level-service-account.name
  ip_subnetworks = ["192.0.4.0/24"]
  members = ["user:test@google.com", "user:test2@google.com", "serviceAccount:${google_service_account.created-later.email}"]
  negate = false
  device_policy {
    require_screen_lock = false
    require_admin_approval = false
    require_corp_owned = true
    os_constraints {
      os_type = "DESKTOP_CHROME_OS"
    }
  }
  regions = [
    "IT",
    "US",
  ]
}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "my policy"
}
```
