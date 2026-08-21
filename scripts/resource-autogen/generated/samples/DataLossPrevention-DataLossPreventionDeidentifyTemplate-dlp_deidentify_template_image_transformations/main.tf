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
resource "google_data_loss_prevention_deidentify_template" "basic" {
  parent = "projects/my-project-name"
  description = "Description"
  display_name = "Displayname"
  
  deidentify_config {
    image_transformations {
      transforms {
        redaction_color {
          red = 0.5
          blue = 1
          green = 0.2
        }
        selected_info_types {
          info_types {
            name = "COLOR_INFO"
            version = "latest"
          }
        }
      }

      transforms {
        all_info_types {}
      }

      transforms {
        all_text {}
      }
    }
  }
}
```
