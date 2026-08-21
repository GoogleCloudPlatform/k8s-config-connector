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
resource "google_data_loss_prevention_job_trigger" "hybrid_trigger" {
  parent = "projects/my-project-name"

  triggers {
    manual {}
  }

  inspect_job {
    inspect_template_name = "fake"
    actions {
      save_findings {
        output_config {
          table {
            project_id = "project"
            dataset_id = "dataset"
          }
        }
      }
    }
    storage_config {
      hybrid_options {
        description = "Hybrid job trigger for data from the comments field of a table that contains customer appointment bookings"
        required_finding_label_keys = [
          "appointment-bookings-comments"
        ]
        labels = {
          env = "prod"
        }
        table_options {
          identifying_fields {
            name = "booking_id"
          }
        }
      }
    }
  }
}
```
