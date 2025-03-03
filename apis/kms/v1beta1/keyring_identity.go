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

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

type KMSKeyRingIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *KMSKeyRingIdentity) String() string {
	return i.parent.String() + "/keyRings/" + i.id
}

func ParseKMSKeyRingExternal(external string) (*KMSKeyRingIdentity, error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	p, err := parent.ParseProjectAndLocationParent(strings.Join(tokens[:len(tokens)-2], "/"))
	if err != nil {
		return nil, err
	}
	if tokens[len(tokens)-2] == "keyRings" {
		return &KMSKeyRingIdentity{parent: p, id: tokens[len(tokens)-1]}, nil
	}
	return nil, fmt.Errorf("format of KMSKeyRing external=%q was not known (use %s/keyRings/{{key_ring_id}})", external, p)

}
