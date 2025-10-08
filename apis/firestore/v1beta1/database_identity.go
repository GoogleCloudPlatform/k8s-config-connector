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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DatabaseIDURL is the format for the externalRef of a FirestoreDatabase.
	DatabaseIDURL = parent.ProjectURLTemplate + "/databases/{{database}}"
	ServiceDomain = "firestore.googleapis.com"
)

var _ identity.Identity = &FirestoreDatabaseIdentity{}

// FirestoreDatabaseIdentity represents the identity of a Firestore Database.
// +k8s:deepcopy-gen=false
type FirestoreDatabaseIdentity struct {
	Parent   *parent.ProjectParent
	Database string
}

func (i *FirestoreDatabaseIdentity) String() string {
	return i.Parent.String() + "/databases/" + i.Database
}

func (i *FirestoreDatabaseIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "//firestore.googleapis.com/")

	tokens := strings.Split(ref, "/databases/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of FirestoreDatabase external=%q was not known (use %s)", ref, DatabaseIDURL)
	}
	i.Parent = &parent.ProjectParent{}
	if err := i.Parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.Database = tokens[1]
	if i.Database == "" {
		return fmt.Errorf("database was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &FirestoreDatabase{}

func (obj *FirestoreDatabase) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &FirestoreDatabaseIdentity{}

	// Resolve Parent
	if err := obj.Spec.ProjectRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentRef: %w", err)
	}
	newIdentity.Parent = &parent.ProjectParent{}
	if err := newIdentity.Parent.FromExternal(obj.Spec.ProjectRef.External); err != nil {
		return nil, fmt.Errorf("parsing projectRef.external=%q: %w", obj.Spec.ProjectRef.External, err)
	}
	// Get desired ID
	newIdentity.Database = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.Database == "" {
		newIdentity.Database = obj.GetName()
	}
	if newIdentity.Database == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &FirestoreDatabaseIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
