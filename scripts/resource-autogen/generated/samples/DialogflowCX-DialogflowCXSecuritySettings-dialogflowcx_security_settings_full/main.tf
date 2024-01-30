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
resource "google_data_loss_prevention_inspect_template" "inspect" {
  parent       = "projects/my-project-name/locations/global"
  display_name = "dialogflowcx-inspect-template"
  inspect_config {
    info_types {
      name = "EMAIL_ADDRESS"
    }
  }
}

resource "google_data_loss_prevention_deidentify_template" "deidentify" {
  parent       = "projects/my-project-name/locations/global"
  display_name = "dialogflowcx-deidentify-template"
  deidentify_config {
    info_type_transformations {
      transformations {
        primitive_transformation {
          replace_config {
            new_value {
              string_value = "[REDACTED]"
            }
          }
        }
      }
    }
  }
}

resource "google_storage_bucket" "bucket" {
  name                        = "dialogflowcx-bucket"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_dialogflow_cx_security_settings" "basic_security_settings" {
  display_name        = "dialogflowcx-security-settings"
  location            = "global"
  redaction_strategy  = "REDACT_WITH_SERVICE"
  redaction_scope     = "REDACT_DISK_STORAGE"
  inspect_template    = google_data_loss_prevention_inspect_template.inspect.id
  deidentify_template = google_data_loss_prevention_deidentify_template.deidentify.id
  purge_data_types    = ["DIALOGFLOW_HISTORY"]
  audio_export_settings {
    gcs_bucket             = google_storage_bucket.bucket.id
    audio_export_pattern   = "export"
    enable_audio_redaction = true
    audio_format           = "OGG"
  }
  insights_export_settings {
    enable_insights_export = true
  }
  retention_strategy = "REMOVE_AFTER_CONVERSATION"
}
```
