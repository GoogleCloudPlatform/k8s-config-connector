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
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_instance_template" "igm-basic" {
  name           = "my-template"
  machine_type   = "e2-medium"
  can_ip_forward = false
  tags           = ["foo", "bar"]

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }

  network_interface {
    network = "default"
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_instance_group_manager" "igm-no-tp" {
  description = "Terraform test instance group manager"
  name        = "my-igm"

  version {
    name              = "prod"
    instance_template = google_compute_instance_template.igm-basic.self_link
  }

  base_instance_name = "igm-no-tp"
  zone               = "us-central1-c"
  target_size        = 2
}

resource "google_compute_disk" "default" {
  name  = "my-disk-name"
  type  = "pd-ssd"
  zone  = google_compute_instance_group_manager.igm.zone
  image = "debian-11-bullseye-v20220719"
  physical_block_size_bytes = 4096
}

resource "google_compute_per_instance_config" "with_disk" {
  zone = google_compute_instance_group_manager.igm.zone
  instance_group_manager = google_compute_instance_group_manager.igm.name
  name = "instance-1"
  preserved_state {
    metadata = {
      foo = "bar"
      // Adding a reference to the instance template used causes the stateful instance to update
      // if the instance template changes. Otherwise there is no explicit dependency and template
      // changes may not occur on the stateful instance
      instance_template = google_compute_instance_template.igm-basic.self_link
    }

    disk {
      device_name = "my-stateful-disk"
      source      = google_compute_disk.default.id
      mode        = "READ_ONLY"
    }
  }
}
```
