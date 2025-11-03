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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// FirestoreIndexIdentityURL is the format for the externalRef of a FirestoreIndex.
	FirestoreIndexIdentityURL = CollectionGroupIdentityURL + "/indexes/{{index}}"
)

var _ identity.Identity = &FirestoreIndexIdentity{}

// FirestoreIndexIdentity represents the identity of a Firestore Database.
// +k8s:deepcopy-gen=false
type FirestoreIndexIdentity struct {
	Parent *CollectionGroupIdentity
	Index  string
}

func (i *FirestoreIndexIdentity) String() string {
	return i.Parent.String() + "/indexes/" + i.Index
}

func (i *FirestoreIndexIdentity) FromExternal(ref string) error {
	prefix, index, ok := popURLToken(ref, "indexes")
	if !ok {
		return fmt.Errorf("format of FirestoreIndex external=%q was not known (use %s)", ref, FirestoreIndexIdentityURL)
	}
	i.Parent = &CollectionGroupIdentity{}
	if err := i.Parent.FromExternal(prefix); err != nil {
		return fmt.Errorf("format of FirestoreIndex external=%q was not known (use %s)", ref, FirestoreIndexIdentityURL)
	}
	i.Index = index
	return nil
}

var _ identity.Resource = &FirestoreIndex{}

func (obj *FirestoreIndex) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	statusName := direct.ValueOf(obj.Status.Name)
	if statusName == "" {
		return nil, nil
	}

	i := &FirestoreIndexIdentity{}
	if err := i.FromExternal(statusName); err != nil {
		return nil, fmt.Errorf("getting identity from status.name=%q: %w", statusName, err)
	}
	return i, nil
}
