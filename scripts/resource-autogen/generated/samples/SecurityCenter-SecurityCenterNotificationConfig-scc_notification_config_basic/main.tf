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
resource "google_pubsub_topic" "scc_notification" {
  name = "my-topic"
}

resource "google_scc_notification_config" "custom_notification_config" {
  config_id    = "my-config"
  organization = "123456789"
  description  = "My custom Cloud Security Command Center Finding Notification Configuration"
  pubsub_topic =  google_pubsub_topic.scc_notification.id

  streaming_config {
    filter = "category = \"OPEN_FIREWALL\" AND state = \"ACTIVE\""
  }
}
```
