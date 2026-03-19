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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO: Move this to autoscalingpolicy_types.go once it is generated/implemented
var DataprocAutoscalingPolicyGVK = schema.GroupVersionKind{
	Group:   "dataproc.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "DataprocAutoscalingPolicy",
}

var _ refsv1beta1.Ref = &DataprocAutoscalingPolicyRef{}

// DataprocAutoscalingPolicyRef is a reference to a DataprocAutoscalingPolicy resource.
type DataprocAutoscalingPolicyRef struct {
	// A reference to an externally managed DataprocAutoscalingPolicy resource.
	// Should be in the format "projects/{project}/regions/{region}/autoscalingPolicies/{autoscalingPolicy}".
	External string `json:"external,omitempty"`

	// The name of a DataprocAutoscalingPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataprocAutoscalingPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&DataprocAutoscalingPolicyRef{})
}

func (r *DataprocAutoscalingPolicyRef) GetGVK() schema.GroupVersionKind {
	return DataprocAutoscalingPolicyGVK
}

func (r *DataprocAutoscalingPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DataprocAutoscalingPolicyRef) GetExternal() string {
	return r.External
}

func (r *DataprocAutoscalingPolicyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DataprocAutoscalingPolicyRef) ValidateExternal(ref string) error {
	id := &DataprocAutoscalingPolicyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DataprocAutoscalingPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DataprocAutoscalingPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DataprocAutoscalingPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromDataprocAutoscalingPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// TODO: Move this to autoscalingpolicy_identity.go once it is generated/implemented
var (
	_ identity.IdentityV2 = &DataprocAutoscalingPolicyIdentity{}
)

// TODO: Move this to autoscalingpolicy_identity.go once it is generated/implemented
var DataprocAutoscalingPolicyIdentityFormat = gcpurls.Template[DataprocAutoscalingPolicyIdentity]("dataproc.googleapis.com", "projects/{project}/regions/{region}/autoscalingPolicies/{autoscalingPolicy}")

// TODO: Move this to autoscalingpolicy_identity.go once it is generated/implemented
// +k8s:deepcopy-gen=false
type DataprocAutoscalingPolicyIdentity struct {
	Project           string
	Region            string
	AutoscalingPolicy string
}

// TODO: Move this to autoscalingpolicy_identity.go once it is generated/implemented
func (i *DataprocAutoscalingPolicyIdentity) String() string {
	return DataprocAutoscalingPolicyIdentityFormat.ToString(*i)
}

// TODO: Move this to autoscalingpolicy_identity.go once it is generated/implemented
func (i *DataprocAutoscalingPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := DataprocAutoscalingPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataprocAutoscalingPolicy external=%q was not known (use %s): %w", ref, DataprocAutoscalingPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataprocAutoscalingPolicy external=%q was not known (use %s)", ref, DataprocAutoscalingPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

// TODO: Move this to autoscalingpolicy_identity.go once it is generated/implemented
func (i *DataprocAutoscalingPolicyIdentity) Host() string {
	return DataprocAutoscalingPolicyIdentityFormat.Host()
}

// TODO: Move this to autoscalingpolicy_identity.go once it is generated/implemented
func getIdentityFromDataprocAutoscalingPolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataprocAutoscalingPolicyIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refsv1beta1.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &DataprocAutoscalingPolicyIdentity{
		Project:           projectID,
		Region:            location,
		AutoscalingPolicy: resourceID,
	}
	return identity, nil
}
