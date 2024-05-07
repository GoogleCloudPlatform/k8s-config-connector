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
resource "google_project" "agent_project" {
  project_id = "my-project"
  name = "my-project"
  org_id = "123456789"
}

resource "google_project_service" "agent_project" {
  project = google_project.agent_project.project_id
  service = "dialogflow.googleapis.com"
  disable_dependent_services = false
}

resource "google_service_account" "dialogflow_service_account" {
  account_id = "my-account"
}

resource "google_project_iam_member" "agent_create" {
  project = google_project_service.agent_project.project
  role    = "roles/dialogflow.admin"
  member  = "serviceAccount:${google_service_account.dialogflow_service_account.email}"
}

resource "google_dialogflow_agent" "basic_agent" {
  project = google_project.agent_project.project_id
  display_name = "example_agent"
  default_language_code = "en"
  time_zone = "America/New_York"
}

resource "google_dialogflow_intent" "full_intent" {
  project = google_project.agent_project.project_id
  depends_on = [google_dialogflow_agent.basic_agent]
  display_name = "full-intent"
  webhook_state = "WEBHOOK_STATE_ENABLED"
  priority = 1
  is_fallback = false
  ml_disabled = true
  action = "some_action"
  reset_contexts = true
  input_context_names = ["projects/${google_project.agent_project.project_id}/agent/sessions/-/contexts/some_id"]
  events = ["some_event"]
  default_response_platforms = ["FACEBOOK","SLACK"]
}
```
