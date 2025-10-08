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

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

type KMSKeyRingIdentity struct {
	Parent *parent.ProjectAndLocationParent
	ID     string
}

func (i *KMSKeyRingIdentity) String() string {
	return i.Parent.String() + "/keyRings/" + i.ID
}

func ParseKMSKeyRingExternal(external string) (*KMSKeyRingIdentity, error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	// projects/{{projectId}}/locations/{{location}}/keyRings/{{keyRingId}}
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" {
		return &KMSKeyRingIdentity{Parent: &parent.ProjectAndLocationParent{
			ProjectID: tokens[1], Location: tokens[3],
		}, ID: tokens[5]}, nil
	}
	return nil, fmt.Errorf("format of KMSKeyRing external=%q was not known (use projects/{{projectId}}/locations/{{location}}/keyRings/{{keyRingId}})", external)
}
