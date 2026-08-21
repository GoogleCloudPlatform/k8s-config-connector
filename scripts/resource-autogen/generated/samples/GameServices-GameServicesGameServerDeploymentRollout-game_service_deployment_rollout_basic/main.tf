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
resource "google_game_services_game_server_deployment" "default" {
  deployment_id  = "tf-test-deployment"
  description = "a deployment description"
}

resource "google_game_services_game_server_config" "default" {
  config_id     = "tf-test-config"
  deployment_id = google_game_services_game_server_deployment.default.deployment_id
  description   = "a config description"

  fleet_configs {
    name       = "some-non-guid"
    fleet_spec = jsonencode({ "replicas" : 1, "scheduling" : "Packed", "template" : { "metadata" : { "name" : "tf-test-game-server-template" }, "spec" : { "ports": [{"name": "default", "portPolicy": "Dynamic", "containerPort": 7654, "protocol": "UDP"}], "template" : { "spec" : { "containers" : [{ "name" : "simple-udp-server", "image" : "gcr.io/agones-images/udp-server:0.14" }] } } } } })

    // Alternate usage:
    // fleet_spec = file(fleet_configs.json)
  }
}

resource "google_game_services_game_server_deployment_rollout" "default" {
  deployment_id              = google_game_services_game_server_deployment.default.deployment_id
  default_game_server_config = google_game_services_game_server_config.default.name
}
```
