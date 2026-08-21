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
resource "google_data_catalog_entry" "basic_entry" {
  entry_group = google_data_catalog_entry_group.entry_group.id
  entry_id = "my_entry"

  user_specified_type = "my_custom_type"
  user_specified_system = "SomethingExternal"
}

resource "google_data_catalog_entry_group" "entry_group" {
  entry_group_id = "my_group"
}
```
