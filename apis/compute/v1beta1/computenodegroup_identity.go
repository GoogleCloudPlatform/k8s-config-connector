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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeNodeGroupIdentity{}
)

var ComputeNodeGroupIdentityFormat = gcpurls.Template[ComputeNodeGroupIdentity]("compute.googleapis.com", "projects/{project}/zones/{zone}/nodeGroups/{nodeGroup}")

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
