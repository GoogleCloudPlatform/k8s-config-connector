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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO: Move this to nodegroup_types.go once it is generated/implemented
var ComputeNodeGroupGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeNodeGroup",
}

var _ refsv1beta1.Ref = &ComputeNodeGroupRef{}

// ComputeNodeGroupRef is a reference to a ComputeNodeGroup resource.
type ComputeNodeGroupRef struct {
	// A reference to an externally managed ComputeNodeGroup resource.
	// Should be in the format "projects/{projectID}/zones/{zone}/nodeGroups/{nodeGroup}".
	External string `json:"external,omitempty"`

	// The name of a ComputeNodeGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeNodeGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&ComputeNodeGroupRef{})
}

func (r *ComputeNodeGroupRef) GetGVK() schema.GroupVersionKind {
	return ComputeNodeGroupGVK
}

func (r *ComputeNodeGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeNodeGroupRef) GetExternal() string {
	return r.External
}

func (r *ComputeNodeGroupRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeNodeGroupRef) ValidateExternal(ref string) error {
	id := &ComputeNodeGroupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeNodeGroupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeNodeGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeNodeGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

// ComputeNodeGroupIdentity is a stub for the identity of a ComputeNodeGroup.
// TODO: Move this to nodegroup_identity.go once it is generated/implemented
type ComputeNodeGroupIdentity struct {
	Project   string
	Zone      string
	NodeGroup string
}

var _ identity.Identity = &ComputeNodeGroupIdentity{}

func (i *ComputeNodeGroupIdentity) String() string {
	return fmt.Sprintf("projects/%s/zones/%s/nodeGroups/%s", i.Project, i.Zone, i.NodeGroup)
}

func (i *ComputeNodeGroupIdentity) FromExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("empty external reference")
	}

	trimmed := common.FixStaleComputeExternalFormat(ref)
	tokens := strings.Split(trimmed, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "nodeGroups" {
		i.Project = tokens[1]
		i.Zone = tokens[3]
		i.NodeGroup = tokens[5]
		return nil
	}

	return fmt.Errorf("invalid format for ComputeNodeGroup external reference: %q. "+
		"Expected format: projects/{projectID}/zones/{zone}/nodeGroups/{nodeGroup}", ref)
}
