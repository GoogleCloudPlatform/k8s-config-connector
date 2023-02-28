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
resource "google_os_config_guest_policies" "guest_policies" {
  provider = google-beta
  guest_policy_id = "guest-policy"

  assignment {
    zones = ["us-east1-b", "us-east1-d"]
  }

  recipes {
    name          = "guest-policy-recipe"
    desired_state = "INSTALLED"

    artifacts {
      id = "guest-policy-artifact-id"

      gcs {
        bucket     = "my-bucket"
        object     = "executable.msi"
        generation = 1546030865175603
      }
    }

    install_steps {
      msi_installation {
        artifact_id = "guest-policy-artifact-id"
      }
    }
  }
}
```
