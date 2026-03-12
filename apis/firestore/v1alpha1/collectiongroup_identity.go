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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
)

const (
	// FieldIDURL is the format for the externalRef of a FirestoreField.
	CollectionGroupIDURL = v1beta1.DatabaseIDURL + "/collectionGroups/{{collectionGroup}}"
)

var _ identity.Identity = &FirestoreCollectionGroupIdentity{}

// FirestoreCollectionGroupIdentity represents the identity of a Firestore Collection Group.
// Note that Collection Groups are implicit and do not have a dedicated resource in the Firestore API.
// +k8s:deepcopy-gen=false
type FirestoreCollectionGroupIdentity struct {
	Parent          *v1beta1.FirestoreDatabaseIdentity
	CollectionGroup string
}

func (i *FirestoreCollectionGroupIdentity) String() string {
	return i.Parent.String() + "/collectionGroups/" + i.CollectionGroup
}

func (i *FirestoreCollectionGroupIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "//firestore.googleapis.com/")

	tokens := strings.Split(ref, "/collectionGroups/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of collection group external=%q was not known (use %s)", ref, CollectionGroupIDURL)
	}
	i.Parent = &v1beta1.FirestoreDatabaseIdentity{}
	if err := i.Parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.CollectionGroup = tokens[1]
	if i.CollectionGroup == "" {
		return fmt.Errorf("collection group was empty in external=%q", ref)
	}
	return nil
}
