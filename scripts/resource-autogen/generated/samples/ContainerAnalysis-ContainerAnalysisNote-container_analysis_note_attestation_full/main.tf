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
resource "google_container_analysis_note" "note" {
  name = "attestor-note"

  short_description = "test note"
  long_description = "a longer description of test note"
  expiration_time = "2120-10-02T15:01:23.045123456Z"

  related_url {
    url = "some.url"
    label = "foo"
  }

  related_url {
    url = "google.com"
  }

  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}
```
