// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeNodeGroupIdentity{}
	_ identity.Resource   = &ComputeNodeGroup{}
)

var ComputeNodeGroupIdentityFormat = gcpurls.Template[ComputeNodeGroupIdentity]("compute.googleapis.com", "projects/{project}/zones/{zone}/nodeGroups/{nodeGroup}")

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ComputeNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNodeGroupSpec   `json:"spec,omitempty"`
	Status ComputeNodeGroupStatus `json:"status,omitempty"`
}

type ComputeNodeGroupSpec struct {
	Zone string `json:"zone"`

	ResourceID *string `json:"resourceID,omitempty"`
}

type ComputeNodeGroupStatus struct {
	ObservedGeneration *int64               `json:"observedGeneration,omitempty"`
	ExternalRef        *string              `json:"externalRef,omitempty"`
	Conditions         []v1alpha1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen=false
type ComputeNodeGroupIdentity struct {
	Project   string
	Zone      string
	NodeGroup string
}

func (i *ComputeNodeGroupIdentity) String() string {
	return ComputeNodeGroupIdentityFormat.ToString(*i)
}

func (i *ComputeNodeGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := ComputeNodeGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeNodeGroup external=%q was not known (use %s): %w", ref, ComputeNodeGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeNodeGroup external=%q was not known (use %s)", ref, ComputeNodeGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeNodeGroupIdentity) Host() string {
	return ComputeNodeGroupIdentityFormat.Host()
}

func getIdentityFromComputeNodeGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeNodeGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("expected *unstructured.Unstructured, got %T", obj)
	}

	zone, _, err := unstructured.NestedString(u.Object, "spec", "zone")
	if err != nil {
		return nil, fmt.Errorf("cannot resolve zone: %w", err)
	}
	if zone == "" {
		return nil, fmt.Errorf("zone is required but not found in spec")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeNodeGroupIdentity{
		Project:   projectID,
		Zone:      zone,
		NodeGroup: resourceID,
	}
	return identity, nil
}

func (obj *ComputeNodeGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	zone := obj.Spec.Zone
	if zone == "" {
		return nil, fmt.Errorf("cannot resolve zone")
	}

	specIdentity := &ComputeNodeGroupIdentity{
		Project:   projectID,
		Zone:      zone,
		NodeGroup: resourceID,
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ComputeNodeGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeNodeGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
