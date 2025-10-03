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

resource "google_pubsub_subscription" "pubsubsubscription_sample" {
  ack_deadline_seconds = 15

  expiration_policy {
    ttl = "2678400s"
  }

  labels = {
    cnrm-lease-expiration = "1603984859"
    cnrm-lease-holder-id  = "btpp498colih6qs1pe5g"
    label-one             = "value-one"
    managed-by-cnrm       = "true"
  }

  message_retention_duration = "86400s"
  name                       = "pubsubsubscription-sample"
  project                    = "my-project"
  topic                      = "projects/my-project/topics/pubsubsubscription-dep"
}
# terraform import google_pubsub_subscription.pubsubsubscription_sample projects/my-project/subscriptions/pubsubsubscription-sample
