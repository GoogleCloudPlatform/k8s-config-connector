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
resource "google_cloud_identity_group" "group" {
  display_name = "my-identity-group"

  parent = "customers/A01b123xz"

  group_key {
  	id = "my-identity-group@example.com"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group" "child-group" {
  display_name = "my-identity-group-child"

  parent = "customers/A01b123xz"

  group_key {
  	id = "my-identity-group-child@example.com"
  }

  labels = {
    "cloudidentity.googleapis.com/groups.discussion_forum" = ""
  }
}

resource "google_cloud_identity_group_membership" "cloud_identity_group_membership_basic" {
  group    = google_cloud_identity_group.group.id

  preferred_member_key {
    id = google_cloud_identity_group.child-group.group_key[0].id
  }

  roles {
  	name = "MEMBER"
  }
}
```
