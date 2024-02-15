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

resource "google_project_iam_custom_role" "iamcustomrolesampleproject" {
  description = "This role only contains two permissions - publish and update"
  permissions = [
    "pubsub.topics.publish",
    "pubsub.topics.update",
  ]
  project = "my-project"
  role_id = "iamcustomrolesampleproject"
  stage   = "GA"
  title   = "Example Project-Level Custom Role"
}
# terraform import google_project_iam_custom_role.iamcustomrolesampleproject my-project##iamcustomrolesampleproject
