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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &VertexAIModelDeploymentMonitoringJobRef{}

// VertexAIModelDeploymentMonitoringJobRef is a reference to a VertexAIModelDeploymentMonitoringJob.
type VertexAIModelDeploymentMonitoringJobRef struct {
	// A reference to an externally managed VertexAIModelDeploymentMonitoringJob resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/modelDeploymentMonitoringJobs/{{modeldeploymentmonitoringjob}}"
	External string `json:"external,omitempty"`

	// The name of a VertexAIModelDeploymentMonitoringJob resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIModelDeploymentMonitoringJob resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIModelDeploymentMonitoringJobRef{}, &VertexAIModelDeploymentMonitoringJob{})
}

func (r *VertexAIModelDeploymentMonitoringJobRef) GetGVK() schema.GroupVersionKind {
	return VertexAIModelDeploymentMonitoringJobGVK
}

func (r *VertexAIModelDeploymentMonitoringJobRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIModelDeploymentMonitoringJobRef) GetExternal() string {
	return r.External
}

func (r *VertexAIModelDeploymentMonitoringJobRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAIModelDeploymentMonitoringJobRef) ValidateExternal(ref string) error {
	id := &VertexAIModelDeploymentMonitoringJobIdentity{}
	return id.FromExternal(ref)
}

func (r *VertexAIModelDeploymentMonitoringJobRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAIModelDeploymentMonitoringJobIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIModelDeploymentMonitoringJobRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*VertexAIModelDeploymentMonitoringJob](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromVertexAIModelDeploymentMonitoringJobSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
