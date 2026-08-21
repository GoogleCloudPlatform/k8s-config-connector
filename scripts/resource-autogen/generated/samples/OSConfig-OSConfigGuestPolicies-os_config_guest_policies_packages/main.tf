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
    group_labels {
      labels = {
        color = "red",
        env = "test"
      }
    }

    group_labels {
      labels = {
        color = "blue",
        env = "test"
      }
    }
  }

  packages {
    name          = "my-package"
    desired_state = "INSTALLED"
  }

  packages {
    name          = "bad-package-1"
    desired_state = "REMOVED"
  }

  packages {
    name          = "bad-package-2"
    desired_state = "REMOVED"
    manager       = "APT"
  }

  package_repositories {
    apt {
      uri          = "https://packages.cloud.google.com/apt"
      archive_type = "DEB"
      distribution = "cloud-sdk-stretch"
      components   = ["main"]
    }
  }

  package_repositories {
    yum {
      id           = "google-cloud-sdk"
      display_name = "Google Cloud SDK"
      base_url     = "https://packages.cloud.google.com/yum/repos/cloud-sdk-el7-x86_64"
      gpg_keys     = ["https://packages.cloud.google.com/yum/doc/yum-key.gpg", "https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg"]
    }
  }
}
```
