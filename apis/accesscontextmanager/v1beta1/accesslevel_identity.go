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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &AccessContextManagerAccessLevelIdentity{}
	_ identity.Resource   = &AccessContextManagerAccessLevel{}
)

var AccessContextManagerAccessLevelIdentityFormat = gcpurls.Template[AccessContextManagerAccessLevelIdentity]("accesscontextmanager.googleapis.com", "accessPolicies/{accessPolicy}/accessLevels/{accessLevel}")

// +k8s:deepcopy-gen=false
type AccessContextManagerAccessLevelIdentity struct {
	AccessPolicy string
	AccessLevel  string
}

func (i *AccessContextManagerAccessLevelIdentity) String() string {
	return AccessContextManagerAccessLevelIdentityFormat.ToString(*i)
}

func (i *AccessContextManagerAccessLevelIdentity) FromExternal(ref string) error {
	parsed, match, err := AccessContextManagerAccessLevelIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of AccessContextManagerAccessLevel external=%q was not known (use %s): %w", ref, AccessContextManagerAccessLevelIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of AccessContextManagerAccessLevel external=%q was not known (use %s)", ref, AccessContextManagerAccessLevelIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AccessContextManagerAccessLevelIdentity) Host() string {
	return AccessContextManagerAccessLevelIdentityFormat.Host()
}

func getIdentityFromAccessContextManagerAccessLevelSpec(ctx context.Context, reader client.Reader, obj client.Object) (*AccessContextManagerAccessLevelIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	var accessPolicyRef *AccessPolicyRef
	if u, ok := obj.(*unstructured.Unstructured); ok {
		refMap, found, err := unstructured.NestedMap(u.Object, "spec", "accessPolicyRef")
		if err != nil {
			return nil, fmt.Errorf("reading spec.accessPolicyRef: %w", err)
		}
		if !found {
			return nil, fmt.Errorf("spec.accessPolicyRef not found")
		}
		accessPolicyRef = &AccessPolicyRef{}
		if val, ok := refMap["external"].(string); ok {
			accessPolicyRef.External = val
		}
		if val, ok := refMap["name"].(string); ok {
			accessPolicyRef.Name = val
		}
		if val, ok := refMap["namespace"].(string); ok {
			accessPolicyRef.Namespace = val
		}
	} else {
		accessLevel := obj.(*AccessContextManagerAccessLevel)
		accessPolicyRef = accessLevel.Spec.AccessPolicyRef
	}

	if accessPolicyRef == nil {
		return nil, fmt.Errorf("spec.accessPolicyRef is required")
	}

	accessPolicyExternal, err := accessPolicyRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("resolving spec.accessPolicyRef: %w", err)
	}

	// accessPolicyExternal is in format "accessPolicies/{{accessPolicyID}}"
	tokens := strings.Split(accessPolicyExternal, "/")
	if len(tokens) != 2 || tokens[0] != "accessPolicies" {
		return nil, fmt.Errorf("invalid accessPolicy external: %q", accessPolicyExternal)
	}
	accessPolicyID := tokens[1]

	identity := &AccessContextManagerAccessLevelIdentity{
		AccessPolicy: accessPolicyID,
		AccessLevel:  resourceID,
	}
	return identity, nil
}

func (obj *AccessContextManagerAccessLevel) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAccessContextManagerAccessLevelSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AccessContextManagerAccessLevelIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change AccessContextManagerAccessLevel identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
