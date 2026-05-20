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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &MemorystoreInstanceEndpointIdentity{}
	_ identity.Resource   = &MemorystoreInstanceEndpoint{}
)

var MemorystoreInstanceEndpointIdentityFormat = gcpurls.Template[MemorystoreInstanceEndpointIdentity]("memorystore.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// +k8s:deepcopy-gen=false
type MemorystoreInstanceEndpointIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *MemorystoreInstanceEndpointIdentity) String() string {
	return MemorystoreInstanceEndpointIdentityFormat.ToString(*i)
}

func (i *MemorystoreInstanceEndpointIdentity) FromExternal(ref string) error {
	parsed, match, err := MemorystoreInstanceEndpointIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MemorystoreInstanceEndpoint external=%q was not known (use %s): %w", ref, MemorystoreInstanceEndpointIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MemorystoreInstanceEndpoint external=%q was not known (use %s)", ref, MemorystoreInstanceEndpointIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MemorystoreInstanceEndpointIdentity) Host() string {
	return MemorystoreInstanceEndpointIdentityFormat.Host()
}

func getIdentityFromMemorystoreInstanceEndpointSpec(ctx context.Context, reader client.Reader, obj client.Object) (*MemorystoreInstanceEndpointIdentity, error) {
	u, ok := obj.(*unstructured.Unstructured)
	if ok {
		obj = &MemorystoreInstanceEndpoint{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
			return nil, fmt.Errorf("error converting to MemorystoreInstanceEndpoint: %w", err)
		}
	}

	memorystoreInstanceEndpoint, ok := obj.(*MemorystoreInstanceEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected MemorystoreInstanceEndpoint, got %T", obj)
	}

	if memorystoreInstanceEndpoint.Spec.InstanceRef == nil {
		return nil, fmt.Errorf("spec.instanceRef is required")
	}

	if err := memorystoreInstanceEndpoint.Spec.InstanceRef.Normalize(ctx, reader, memorystoreInstanceEndpoint.GetNamespace()); err != nil {
		return nil, fmt.Errorf("normalizing instanceRef: %w", err)
	}

	instanceID := memorystoreInstanceEndpoint.Spec.InstanceRef.GetExternal()
	if instanceID == "" {
		return nil, fmt.Errorf("cannot resolve instanceRef")
	}

	identity := &MemorystoreInstanceEndpointIdentity{}
	if err := identity.FromExternal(instanceID); err != nil {
		return nil, err
	}

	return identity, nil
}

func (obj *MemorystoreInstanceEndpoint) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMemorystoreInstanceEndpointSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	if obj.Status.ExternalRef != nil {
		statusIdentity := &MemorystoreInstanceEndpointIdentity{}
		if err := statusIdentity.FromExternal(*obj.Status.ExternalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change MemorystoreInstanceEndpoint identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
