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
resource "google_data_loss_prevention_job_trigger" "with_trigger_id" {
  parent = "projects/my-project-name"
  description = "Starting description"
  display_name = "display"
  trigger_id = "id-"

  triggers {
    schedule {
      recurrence_period_duration = "86400s"
    }
  }

  inspect_job {
    inspect_template_name = "fake"
    actions {
      save_findings {
        output_config {
          table {
            project_id = "project"
            dataset_id = "dataset123"
          }
        }
      }
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
