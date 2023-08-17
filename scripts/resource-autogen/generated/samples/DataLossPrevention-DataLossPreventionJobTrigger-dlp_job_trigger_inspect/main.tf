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
resource "google_data_loss_prevention_job_trigger" "inspect" {
  parent = "projects/my-project-name"
  description = "Description"
  display_name = "Displayname"

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
            dataset_id = "dataset"
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
    inspect_config {
      custom_info_types {
        info_type {
          name = "MY_CUSTOM_TYPE"
        }
  
        likelihood = "UNLIKELY"
  
        regex {
          pattern = "test*"
        }
      }
  
      info_types {
        name = "EMAIL_ADDRESS"
      }
  
      min_likelihood = "UNLIKELY"
      rule_set {
        info_types {
          name = "EMAIL_ADDRESS"
        }
        rules {
          exclusion_rule {
            regex {
              pattern = ".+@example.com"
            }
            matching_type = "MATCHING_TYPE_FULL_MATCH"
          }
        }
      }
  
      rule_set {
        info_types {
          name = "MY_CUSTOM_TYPE"
        }
        rules {
          hotword_rule {
            hotword_regex {
              pattern = "example*"
            }
            proximity {
              window_before = 50
            }
            likelihood_adjustment {
              fixed_likelihood = "VERY_LIKELY"
            }
          }
        }
      }
  
      limits {
        max_findings_per_item    = 10
        max_findings_per_request = 50
      }
    }
  }
}
```
