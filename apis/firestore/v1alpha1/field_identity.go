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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// FieldIDURL is the format for the externalRef of a FirestoreField.
	FieldIDURL = parent.ProjectURLTemplate + "/fields/{{field}}"
)

var _ identity.Identity = &FirestoreFieldIdentity{}

// FirestoreFieldIdentity represents the identity of a Firestore Field.
// +k8s:deepcopy-gen=false
type FirestoreFieldIdentity struct {
	Parent *FirestoreCollectionGroupIdentity
	Field  string
}

func (i *FirestoreFieldIdentity) String() string {
	return i.Parent.String() + "/fields/" + i.Field
}

func (i *FirestoreFieldIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "//firestore.googleapis.com/")

	tokens := strings.Split(ref, "/fields/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of FirestoreField external=%q was not known (use %s)", ref, FieldIDURL)
	}
	i.Parent = &FirestoreCollectionGroupIdentity{}
	if err := i.Parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.Field = tokens[1]
	if i.Field == "" {
		return fmt.Errorf("field was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &FirestoreField{}

func (obj *FirestoreField) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	db := &v1beta1.FirestoreDatabaseIdentity{}
	if err := db.FromExternal(obj.Spec.DatabaseRef.External); err != nil {
		return nil, fmt.Errorf("parsing databaseRef.external=%q: %w", obj.Spec.DatabaseRef.External, err)
	}

	collectionGroup := common.ValueOf(obj.Spec.CollectionGroup)
	if collectionGroup == "" {
		return nil, fmt.Errorf("spec.collectionGroup is required")
	}

	// Get desired ID
	field := common.ValueOf(obj.Spec.ResourceID)
	if field == "" {
		field = obj.GetName()
	}
	if field == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	newIdentity := &FirestoreFieldIdentity{
		Parent: &FirestoreCollectionGroupIdentity{
			Parent:          db,
			CollectionGroup: collectionGroup,
		},
		Field: field,
	}

	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &FirestoreFieldIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
