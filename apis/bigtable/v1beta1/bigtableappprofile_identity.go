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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigtableAppProfileIdentity{}
	_ identity.Resource   = &BigtableAppProfile{}
)

var BigtableAppProfileIdentityFormat = gcpurls.Template[BigtableAppProfileIdentity]("bigtableadmin.googleapis.com", "projects/{project}/instances/{instance}/appProfiles/{appprofile}")

// +k8s:deepcopy-gen=false
type BigtableAppProfileIdentity struct {
	Project    string
	Instance   string
	AppProfile string
}

func (i *BigtableAppProfileIdentity) String() string {
	return BigtableAppProfileIdentityFormat.ToString(*i)
}

func (i *BigtableAppProfileIdentity) FromExternal(ref string) error {
	parsed, match, err := BigtableAppProfileIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableAppProfile external=%q was not known (use %s): %w", ref, BigtableAppProfileIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableAppProfile external=%q was not known (use %s)", ref, BigtableAppProfileIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigtableAppProfileIdentity) Host() string {
	return BigtableAppProfileIdentityFormat.Host()
}

func getIdentityFromBigtableAppProfileSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigtableAppProfileIdentity, error) {
	appprofile := &BigtableAppProfile{}
	if u, ok := obj.(*unstructured.Unstructured); ok {
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, appprofile); err != nil {
			return nil, fmt.Errorf("failed to convert from unstructured: %w", err)
		}
	} else if typed, ok := obj.(*BigtableAppProfile); ok {
		appprofile = typed
	} else {
		return nil, fmt.Errorf("expected BigtableAppProfile or *unstructured.Unstructured, got %T", obj)
	}

	resourceID := common.ValueOf(appprofile.Spec.ResourceID)
	if resourceID == "" {
		resourceID = appprofile.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Resolve Instance ID
	instanceRef := appprofile.Spec.InstanceRef
	if instanceRef == nil {
		return nil, fmt.Errorf("spec.instanceRef is required")
	}
	instanceExternal, err := instanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	projectParent, instanceID, err := ParseInstanceExternal(instanceExternal)
	if err != nil {
		return nil, err
	}

	return &BigtableAppProfileIdentity{
		Project:    projectParent.ProjectID,
		Instance:   instanceID,
		AppProfile: resourceID,
	}, nil
}

func (obj *BigtableAppProfile) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableAppProfileSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BigtableAppProfileIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigtableAppProfile identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
