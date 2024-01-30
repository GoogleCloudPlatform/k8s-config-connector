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
  display_name               = "dialogflowcx-agent"
  location                   = "global"
  default_language_code      = "en"
  supported_language_codes   = ["fr", "de", "es"]
  time_zone                  = "America/New_York"
  description                = "Example description."
  avatar_uri                 = "https://storage.cloud.google.com/dialogflow-test-host-image/cloud-logo.png"
  enable_stackdriver_logging = true
  enable_spell_correction    = true
  speech_to_text_settings {
    enable_speech_adaptation = true
  }
}

resource "google_dialogflow_cx_page" "page" {
  parent       = google_dialogflow_cx_agent.agent.start_flow
  display_name = "MyPage"

  transition_routes {
    intent = google_dialogflow_cx_intent.intent.id
    trigger_fulfillment {
      messages {
        text {
          text = ["Training phrase response"]
        }
      }
    }
  }

  event_handlers {
    event = "some-event"
    trigger_fulfillment {
      messages {
        text {
          text = ["Handling some event"]
        }
      }
    }
  }
}

resource "google_dialogflow_cx_intent" "intent" {
  parent       = google_dialogflow_cx_agent.agent.id
  display_name = "MyIntent"
  priority     = 1
  training_phrases {
    parts {
      text = "training phrase"
    }
    repeat_count = 1
  }
}

resource "google_dialogflow_cx_test_case" "basic_test_case" {
  parent       = google_dialogflow_cx_agent.agent.id
  display_name = "MyTestCase"
  tags         = ["#tag1"]
  notes        = "demonstrates a simple training phrase response"

  test_config {
    tracking_parameters = ["some_param"]
    page                = google_dialogflow_cx_page.page.id
  }

  test_case_conversation_turns {
    user_input {
      input {
        language_code = "en"
        text {
          text = "training phrase"
        }
      }
      injected_parameters       = jsonencode({ some_param = "1" })
      is_webhook_enabled        = true
      enable_sentiment_analysis = true
    }
    virtual_agent_output {
      session_parameters = jsonencode({ some_param = "1" })
      triggered_intent {
        name = google_dialogflow_cx_intent.intent.id
      }
      current_page {
        name = google_dialogflow_cx_page.page.id
      }
      text_responses {
        text = ["Training phrase response"]
      }
    }
  }

  test_case_conversation_turns {
    user_input {
      input {
        event {
          event = "some-event"
        }
      }
    }
    virtual_agent_output {
      current_page {
        name = google_dialogflow_cx_page.page.id
      }
      text_responses {
        text = ["Handling some event"]
      }
    }
  }

  test_case_conversation_turns {
    user_input {
      input {
        dtmf {
          digits       = "12"
          finish_digit = "3"
        }
      }
    }
    virtual_agent_output {
      text_responses {
        text = ["I didn't get that. Can you say it again?"]
      }
    }
  }
}
```
