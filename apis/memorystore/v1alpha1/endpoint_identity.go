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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EndpointIdentity defines the resource reference to MemorystoreInstance, which "External" field
// holds the GCP identifier for the KRM object.
type EndpointIdentity struct {
	parent string
}

func (i *EndpointIdentity) String() string {
	return i.parent
}

func (i *EndpointIdentity) ID() string {
	return i.parent
}

func (i *EndpointIdentity) Parent() string {
	return i.parent
}

// New builds a EndpointIdentity from the Config Connector Instance object.
func NewEndpointIdentity(ctx context.Context, reader client.Reader, obj *MemorystoreInstanceEndpoint) (*EndpointIdentity, error) {

	if obj.Spec.InstanceRef == nil {
		return nil, fmt.Errorf("spec.instanceRef is required")
	}
	instanceRef, err := refsv1beta1.ResolveMemorystoreInstance(ctx, reader, obj, obj.Spec.InstanceRef)
	if err != nil {
		return nil, err
	}
	return &EndpointIdentity{
		parent: instanceRef.External,
	}, nil
}
