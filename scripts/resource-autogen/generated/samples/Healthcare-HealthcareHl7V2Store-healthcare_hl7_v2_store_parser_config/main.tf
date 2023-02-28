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
resource "google_healthcare_hl7_v2_store" "store" {
  provider = google-beta
  name    = "example-hl7-v2-store"
  dataset = google_healthcare_dataset.dataset.id

  parser_config {
    allow_null_header  = false
    segment_terminator = "Jw=="
    schema = <<EOF
{
  "schemas": [{
    "messageSchemaConfigs": {
      "ADT_A01": {
        "name": "ADT_A01",
        "minOccurs": 1,
        "maxOccurs": 1,
        "members": [{
            "segment": {
              "type": "MSH",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "EVN",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "PID",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "ZPD",
              "minOccurs": 1,
              "maxOccurs": 1
            }
          },
          {
            "segment": {
              "type": "OBX"
            }
          },
          {
            "group": {
              "name": "PROCEDURE",
              "members": [{
                  "segment": {
                    "type": "PR1",
                    "minOccurs": 1,
                    "maxOccurs": 1
                  }
                },
                {
                  "segment": {
                    "type": "ROL"
                  }
                }
              ]
            }
          },
          {
            "segment": {
              "type": "PDA",
              "maxOccurs": 1
            }
          }
        ]
      }
    }
  }],
  "types": [{
    "type": [{
        "name": "ZPD",
        "primitive": "VARIES"
      }

    ]
  }],
  "ignoreMinOccurs": true
}
EOF
  }
}

resource "google_healthcare_dataset" "dataset" {
  provider = google-beta
  name     = "example-dataset"
  location = "us-central1"
}
```
