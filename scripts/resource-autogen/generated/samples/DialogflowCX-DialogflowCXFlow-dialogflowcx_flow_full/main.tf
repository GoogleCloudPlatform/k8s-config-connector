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


resource "google_dialogflow_cx_flow" "basic_flow" {
  parent       = google_dialogflow_cx_agent.agent.id
  display_name = "MyFlow"
  description  = "Test Flow"

  nlu_settings {
		classification_threshold = 0.3 
		model_type               = "MODEL_TYPE_STANDARD"
	}

  event_handlers {
		   event                    = "custom-event"
		   trigger_fulfillment {
			    return_partial_responses = false
				messages {
					text {
						text  = ["I didn't get that. Can you say it again?"]
					}
				}
		    }
		}

		event_handlers {
			event                    = "sys.no-match-default"
			trigger_fulfillment {
				 return_partial_responses = false
				 messages {
					 text {
						 text  = ["Sorry, could you say that again?"]
					 }
				 }
			 }
		 }

		 event_handlers {
			event                    = "sys.no-input-default"
			trigger_fulfillment {
				 return_partial_responses = false
				 messages {
					 text {
						 text  = ["One more time?"]
					 }
				 }
			 }
		 }
} 
```
