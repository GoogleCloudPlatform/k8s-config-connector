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

package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

const (
	ConnectionProfileIDURL = parent.ProjectAndLocationURL + "/connectionProfiles/{{connectionProfileID}}"
)

// +k8s:deepcopy-gen=false
type ConnectionProfileIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *ConnectionProfileIdentity) String() string {
	return i.parent.String() + "/connectionProfiles/" + i.id
}

func (i *ConnectionProfileIdentity) Parent() *parent.ProjectAndLocationParent {
	return i.parent
}

func (i *ConnectionProfileIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/connectionProfiles/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of ConnectionProfile external=%q was not known (use %s)", external, ConnectionProfileIDURL)
	}
	i.parent = &parent.ProjectAndLocationParent{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("catalogID was empty in external=%q", external)
	}
	return nil
}

var _ identity.Identity = &ConnectionProfileIdentity{}
