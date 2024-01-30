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
resource "google_service_account" "service_account" {
  account_id   = "my-account"
  display_name = "Service Account"
}

resource "google_data_pipeline_pipeline" "primary" {
  name         = "my-pipeline"
  display_name = "my-pipeline"
  type         = "PIPELINE_TYPE_BATCH"
  state        = "STATE_ACTIVE"
  region       = "us-central1"

  workload {
    dataflow_launch_template_request {
      project_id = "my-project"
      gcs_path   = "gs://my-bucket/path"
      launch_parameters {
        job_name = "my-job"
        parameters = {
          "name" : "wrench"
        }
        environment {
          num_workers                = 5
          max_workers                = 5
          zone                       = "us-centra1-a"
          service_account_email      = google_service_account.service_account.email
          network                    = "default"
          temp_location              = "gs://my-bucket/tmp_dir"
          bypass_temp_dir_validation = false
          machine_type               = "E2"
          additional_user_labels = {
            "context" : "test"
          }
          worker_region    = "us-central1"
          worker_zone      = "us-central1-a"

          enable_streaming_engine = "false"
        }
        update                 = false
        transform_name_mapping = { "name" : "wrench" }
      }
      location = "us-central1"
    }
  }
  schedule_info {
    schedule = "* */2 * * *"
  }
}
```
