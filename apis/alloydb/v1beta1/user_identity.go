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

package v1beta1

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
	_ identity.IdentityV2 = &AlloyDBUserIdentity{}
	_ identity.Resource   = &AlloyDBUser{}
)

var AlloyDBUserIdentityFormat = gcpurls.Template[AlloyDBUserIdentity]("alloydb.googleapis.com", "projects/{project}/locations/{location}/clusters/{cluster}/users/{user}")

// +k8s:deepcopy-gen=false
type AlloyDBUserIdentity struct {
	Project  string
	Location string
	Cluster  string
	User     string
}

func (i *AlloyDBUserIdentity) String() string {
	return AlloyDBUserIdentityFormat.ToString(*i)
}

func (i *AlloyDBUserIdentity) FromExternal(ref string) error {
	parsed, match, err := AlloyDBUserIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AlloyDBUser external=%q was not known (use %s): %w", ref, AlloyDBUserIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AlloyDBUser external=%q was not known (use %s)", ref, AlloyDBUserIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AlloyDBUserIdentity) Host() string {
	return AlloyDBUserIdentityFormat.Host()
}

func (obj *AlloyDBUser) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	clusterPath, err := obj.Spec.ClusterRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}

	parent, clusterID, err := ParseClusterExternal(clusterPath)
	if err != nil {
		return nil, err
	}

	identity := &AlloyDBUserIdentity{
		Project:  parent.ProjectID,
		Location: parent.Location,
		Cluster:  clusterID,
		User:     resourceID,
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AlloyDBUserIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != identity.String() {
			return nil, fmt.Errorf("cannot change AlloyDBUser identity (old=%q, new=%q)", statusIdentity.String(), identity.String())
		}
	}

	return identity, nil
}
