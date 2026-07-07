// Copyright 2026 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigLakeLockIdentity{}
	_ identity.Resource   = &BigLakeLock{}
)

type DatabaseParent struct {
	Project  string
	Location string
	Catalog  string
	Database string
}

var DatabaseIdentityFormat = gcpurls.Template[DatabaseParent]("biglake.googleapis.com", "projects/{project}/locations/{location}/catalogs/{catalog}/databases/{database}")

var BigLakeLockIdentityFormat = gcpurls.Template[BigLakeLockIdentity]("biglake.googleapis.com", "projects/{project}/locations/{location}/catalogs/{catalog}/databases/{database}/locks/{lock}")

// BigLakeLockIdentity is the identity of a BigLakeLock.
// +k8s:deepcopy-gen=false
type BigLakeLockIdentity struct {
	Project  string
	Location string
	Catalog  string
	Database string
	Lock     string
}

func (i *BigLakeLockIdentity) String() string {
	return BigLakeLockIdentityFormat.ToString(*i)
}

func (i *BigLakeLockIdentity) FromExternal(ref string) error {
	parsed, match, err := BigLakeLockIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigLakeLock external=%q was not known (use %s): %w", ref, BigLakeLockIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigLakeLock external=%q was not known (use %s)", ref, BigLakeLockIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigLakeLockIdentity) Host() string {
	return BigLakeLockIdentityFormat.Host()
}

func getIdentityFromBigLakeLockSpec(ctx context.Context, reader client.Reader, obj *BigLakeLock) (*BigLakeLockIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	if obj.Spec.ParentRef == nil {
		return nil, fmt.Errorf("spec.parentDatabaseRef is required")
	}

	if err := obj.Spec.ParentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentDatabaseRef: %w", err)
	}

	externalDB := obj.Spec.ParentRef.GetExternal()
	if externalDB == "" {
		return nil, fmt.Errorf("resolved spec.parentDatabaseRef external identifier is empty")
	}

	parsedDB, match, err := DatabaseIdentityFormat.Parse(externalDB)
	if err != nil || !match {
		return nil, fmt.Errorf("format of parent BigLakeDatabase external=%q was not known: %w", externalDB, err)
	}

	identity := &BigLakeLockIdentity{
		Project:  parsedDB.Project,
		Location: parsedDB.Location,
		Catalog:  parsedDB.Catalog,
		Database: parsedDB.Database,
		Lock:     resourceID,
	}
	return identity, nil
}

func (obj *BigLakeLock) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigLakeLockSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Validate against the ID stored in status.externalRef, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BigLakeLockIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
	}
	return specIdentity, nil
}
