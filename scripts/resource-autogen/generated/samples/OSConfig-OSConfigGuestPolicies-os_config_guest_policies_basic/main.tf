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
data "google_compute_image" "my_image" {
  provider = google-beta
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_instance" "foobar" {
  provider = google-beta
  name           = "guest-policy-inst"
  machine_type   = "e2-medium"
  zone           = "us-central1-a"
  can_ip_forward = false
  tags           = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }
}

resource "google_os_config_guest_policies" "guest_policies" {
  provider = google-beta
  guest_policy_id = "guest-policy"

  assignment {
    instances = [google_compute_instance.foobar.id]
  }

  packages {
    name = "my-package"
    desired_state = "UPDATED"
  }
}
```
