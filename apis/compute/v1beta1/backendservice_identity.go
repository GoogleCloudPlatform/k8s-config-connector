// Copyright 2024 Google LLC
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

package v1beta1

type BackendServiceIdentity struct {
	id     string
	parent *BackendServiceParent
}

func (i *BackendServiceIdentity) String() string {
	return i.parent.String() + "/backendServices/" + i.id
}

func (i *BackendServiceIdentity) Parent() *BackendServiceParent {
	return i.parent
}

func (i *BackendServiceIdentity) ID() string {
	return i.id
}

type BackendServiceParent struct {
	ProjectID string
	Location  string
}

func (p *BackendServiceParent) String() string {
	if p.Location == "global" {
		return "projects/" + p.ProjectID + "/global"
	} else {
		return "projects/" + p.ProjectID + "/regions/" + p.Location
	}
}
