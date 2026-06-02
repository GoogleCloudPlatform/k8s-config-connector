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
	endpoint := &MemorystoreInstanceEndpoint{}
	if u, ok := obj.(*unstructured.Unstructured); ok {
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, endpoint); err != nil {
			return nil, fmt.Errorf("failed to convert from unstructured: %w", err)
		}
	} else if typed, ok := obj.(*MemorystoreInstanceEndpoint); ok {
		endpoint = typed
	} else {
		return nil, fmt.Errorf("expected MemorystoreInstanceEndpoint or *unstructured.Unstructured, got %T", obj)
	}

	if endpoint.Spec.InstanceRef == nil {
		return nil, fmt.Errorf("spec.instanceRef is required")
	}

	if err := endpoint.Spec.InstanceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}

	external := endpoint.Spec.InstanceRef.GetExternal()
	identity := &MemorystoreInstanceEndpointIdentity{}
	if err := identity.FromExternal(external); err != nil {
		return nil, err
	}
	return identity, nil
}

func (obj *MemorystoreInstanceEndpoint) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromMemorystoreInstanceEndpointSpec(ctx, reader, obj)
}
