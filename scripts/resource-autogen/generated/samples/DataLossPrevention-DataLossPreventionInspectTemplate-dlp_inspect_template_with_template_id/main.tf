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
resource "google_data_loss_prevention_inspect_template" "with_template_id" {
  parent = "projects/my-project-name"
  template_id = "my-template"

  inspect_config {
    info_types {
      name = "EMAIL_ADDRESS"
    }
    info_types {
      name = "PERSON_NAME"
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
        name = "PERSON_NAME"
      }
      rules {
        hotword_rule {
          hotword_regex {
            pattern = "patient"
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
  }
}
```
