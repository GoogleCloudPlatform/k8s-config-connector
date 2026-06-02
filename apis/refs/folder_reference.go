// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package refs

// FolderRef is a clean resource reference to a GCP Folder that does not include the kind field.
type FolderRef struct {
	/* The 'name' field of a folder, when not managed by Config Connector. */
	// +optional
	External string `json:"external,omitempty"`
	/* The 'name' field of a 'Folder' resource. */
	// +optional
	Name string `json:"name,omitempty"`
	/* The 'namespace' field of a 'Folder' resource. */
	// +optional
	Namespace string `json:"namespace,omitempty"`
}
