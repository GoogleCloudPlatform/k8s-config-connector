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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeAutoscalerIdentity{}
	_ identity.Resource   = &ComputeAutoscaler{}
)

var ComputeAutoscalerIdentityFormat = gcpurls.Template[ComputeAutoscalerIdentity](
	"compute.googleapis.com",
	"projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
)

// ComputeAutoscalerIdentity is the identity of a GCP ComputeAutoscaler resource.
// +k8s:deepcopy-gen=false
type ComputeAutoscalerIdentity struct {
	Project    string
	Zone       string
	Autoscaler string
}

func (i *ComputeAutoscalerIdentity) String() string {
	return ComputeAutoscalerIdentityFormat.ToString(*i)
}

func (i *ComputeAutoscalerIdentity) FromExternal(ref string) error {
	ref = refs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeAutoscalerIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeAutoscaler external=%q was not known (use %s): %w", ref, ComputeAutoscalerIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeAutoscaler external=%q was not known (use %s)", ref, ComputeAutoscalerIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeAutoscalerIdentity) Host() string {
	return ComputeAutoscalerIdentityFormat.Host()
}

func getIdentityFromComputeAutoscalerSpec(ctx context.Context, reader client.Reader, obj *ComputeAutoscaler) (*ComputeAutoscalerIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	zone := obj.Spec.Zone
	if zone == "" {
		return nil, fmt.Errorf("cannot resolve zone: spec.zone is empty")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeAutoscalerIdentity{
		Project:    projectID,
		Zone:       zone,
		Autoscaler: resourceID,
	}
	return identity, nil
}

func (obj *ComputeAutoscaler) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeAutoscalerSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
