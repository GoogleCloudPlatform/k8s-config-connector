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
resource "google_data_loss_prevention_job_trigger" "job_notification_emails" {
  parent       = "projects/my-project-name"
  description  = "Description for the job_trigger created by terraform"
  display_name = "TerraformDisplayName"
  
  triggers {
    schedule {
      recurrence_period_duration = "86400s"
    }
  }
  
  inspect_job {
    inspect_template_name = "sample-inspect-template"
    actions {
      job_notification_emails {}
    }
    storage_config {
      cloud_storage_options {
        file_set {
          url = "gs://mybucket/directory/"
        }
      }
    }
  }
}
```
