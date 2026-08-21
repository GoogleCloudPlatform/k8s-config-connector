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
resource "google_looker_instance" "looker-instance" {
  name               = "my-instance"
  platform_edition   = "LOOKER_CORE_STANDARD"
  region             = "us-central1"
  public_ip_enabled  = true
  admin_settings {
    allowed_email_domains = ["google.com"]
  }
  // User metadata config is only available when platform edition is LOOKER_CORE_STANDARD.
  user_metadata {
    additional_developer_user_count = 10 
    additional_standard_user_count  = 10
    additional_viewer_user_count    = 10
  }
  maintenance_window {
    day_of_week = "THURSDAY"
    start_time {
      hours   = 22
      minutes = 0
      seconds = 0
      nanos   = 0
    }
  }
  deny_maintenance_period {    
    start_date {
      year = 2050
      month = 1
      day = 1
    }
    end_date {
      year = 2050
      month = 2
      day = 1
    }
    time {
      hours = 10
      minutes = 0
      seconds = 0
      nanos = 0
    }
  }
  oauth_config {
    client_id = "my-client-id"
    client_secret = "my-client-secret"
  }  
}
```
