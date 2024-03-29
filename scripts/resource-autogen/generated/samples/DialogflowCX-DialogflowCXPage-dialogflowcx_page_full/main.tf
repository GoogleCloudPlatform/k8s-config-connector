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
resource "google_dialogflow_cx_agent" "agent" {
  display_name = "dialogflowcx-agent"
  location = "global"
  default_language_code = "en"
  supported_language_codes = ["fr","de","es"]
  time_zone = "America/New_York"
  description = "Example description."
  avatar_uri = "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"
  enable_stackdriver_logging = true
  enable_spell_correction    = true
	speech_to_text_settings {
		enable_speech_adaptation = true
	}
}


resource "google_dialogflow_cx_page" "basic_page" {
  parent       = google_dialogflow_cx_agent.agent.start_flow
  display_name = "MyPage"

  entry_fulfillment {
		messages {
			text {
				text = ["Welcome to page"]
			}
		}
   }

   form {
		parameters {
			display_name = "param1"
			entity_type  = "projects/-/locations/-/agents/-/entityTypes/sys.date"
			fill_behavior {
				initial_prompt_fulfillment {
					messages {
						text {
							text = ["Please provide param1"]
						}
					}
				}
			}
			required = "true"
			redact   = "true"
		}
	}

    transition_routes {
		condition = "$page.params.status = 'FINAL'"
		trigger_fulfillment {
			messages {
				text {
					text = ["information completed, navigating to page 2"]
				}
			}
		}
		target_page = google_dialogflow_cx_page.my_page2.id
	}
} 

resource "google_dialogflow_cx_page" "my_page2" {
    parent       = google_dialogflow_cx_agent.agent.start_flow
    display_name  = "MyPage2"
}
```
