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
resource "google_project" "project" {
	project_id = "project_id"
	name       = "project_id"
	org_id     = "123456789"
}

resource "google_tags_tag_key" "key" {
	parent = "organizations/123456789"
	short_name = "keyname"
	description = "For keyname resources."
}

resource "google_tags_tag_value" "value" {
	parent = "tagKeys/${google_tags_tag_key.key.name}"
	short_name = "valuename"
	description = "For valuename resources."
}

resource "google_tags_tag_binding" "binding" {
	parent = "//cloudresourcemanager.googleapis.com/projects/${google_project.project.number}"
	tag_value = "tagValues/${google_tags_tag_value.value.name}"
}
```
