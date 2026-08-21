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
resource "google_pubsub_lite_reservation" "example" {
  name = "example-reservation"
  project = data.google_project.project.number
  throughput_capacity = 2
}

resource "google_pubsub_lite_topic" "example" {
  name = "example-topic"
  project = data.google_project.project.number

  partition_config {
    count = 1
    capacity {
      publish_mib_per_sec = 4
      subscribe_mib_per_sec = 8
    }
  }

  retention_config {
    per_partition_bytes = 32212254720
  }

  reservation_config {
    throughput_reservation = google_pubsub_lite_reservation.example.name
  }
}

data "google_project" "project" {
}
```
