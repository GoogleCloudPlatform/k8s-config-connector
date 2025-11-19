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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

const (
	// CollectionGroupIdentityURL is the format for the externalRef of a FirestoreIndex.
	CollectionGroupIdentityURL = DatabaseIDURL + "/collectionGroups/{{collectionGroup}}"
)

var _ identity.Identity = &CollectionGroupIdentity{}

// CollectionGroupIdentity represents the identity of a collection group.
// Note that collection groups are implicit resources in Firestore.
// +k8s:deepcopy-gen=false
type CollectionGroupIdentity struct {
	Parent          *FirestoreDatabaseIdentity
	CollectionGroup string
}

func (i *CollectionGroupIdentity) String() string {
	return i.Parent.String() + "/collectionGroups/" + i.CollectionGroup
}

// popURLToken pops the last token from the given URL, making sure it is of the given kind.
func popURLToken(url string, kind string) (string, string, bool) {
	lastIndex := strings.LastIndex(url, "/"+kind+"/")
	if lastIndex == -1 {
		return "", "", false
	}
	suffix := url[lastIndex+len(kind)+2:]
	if strings.Contains(suffix, "/") {
		// Not the last token
		return "", "", false
	}
	if suffix == "" {
		// No token found
		return "", "", false
	}
	return url[:lastIndex], suffix, true
}

func (i *CollectionGroupIdentity) FromExternal(ref string) error {
	prefix, collectionGroup, ok := popURLToken(ref, "collectionGroups")
	if !ok {
		return fmt.Errorf("format of FirestoreIndex external=%q was not known (use %s)", ref, CollectionGroupIdentityURL)
	}
	i.Parent = &FirestoreDatabaseIdentity{}
	if err := i.Parent.FromExternal(prefix); err != nil {
		return fmt.Errorf("format of FirestoreIndex external=%q was not known (use %s)", ref, CollectionGroupIdentityURL)
	}
	i.CollectionGroup = collectionGroup
	return nil
}
