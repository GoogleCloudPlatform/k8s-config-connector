// Copyright 2025 Google LLC
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

type SubnetworkIdentity struct {
	id     string
	parent *SubnetworkParent
}

func (i *SubnetworkIdentity) String() string {
	return i.parent.String() + "/subnetworks/" + i.id
}

func (i *SubnetworkIdentity) Parent() *SubnetworkParent {
	return i.parent
}

func (i *SubnetworkIdentity) ID() string {
	return i.id
}

type SubnetworkParent struct {
	ProjectID string
	Region    string
}

func (p *SubnetworkParent) String() string {
	return "projects/" + p.ProjectID + "/regions/" + p.Region
}
