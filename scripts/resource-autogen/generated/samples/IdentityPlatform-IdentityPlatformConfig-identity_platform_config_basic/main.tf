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
resource "google_project" "default" {
  project_id = "my-project"
  name       = "my-project"
  org_id     = "123456789"
  billing_account =  "000000-0000000-0000000-000000"
  labels = {
    firebase = "enabled"
  }
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"
}

resource "google_identity_platform_config" "default" {
  project = google_project.default.project_id
  autodelete_anonymous_users = true
  sign_in {
    allow_duplicate_emails = true
   
    anonymous {
        enabled = true
    }
    email {
        enabled = true
        password_required = false
    }
    phone_number {
        enabled = true
        test_phone_numbers = {
            "+11231231234" = "000000"
        }
    }
  }
  blocking_functions {
    triggers {
      event_type = "beforeSignIn"
      function_uri = "https://us-east1-my-project.cloudfunctions.net/before-sign-in"
    }
    forward_inbound_credentials {
      refresh_token = true
      access_token = true
      id_token = true
    }
  }
  quota {
    sign_up_quota_config {
      quota = 1000
      start_time = ""
      quota_duration = "7200s"
    }
  }
  authorized_domains = [
    "localhost",
    "my-project.firebaseapp.com",
    "my-project.web.app",
  ]
}
```
