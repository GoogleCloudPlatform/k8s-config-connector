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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &MonitoringMonitoredProjectRef{}

// MonitoringMonitoredProjectRef is a reference to a GCP MonitoringMonitoredProject.
type MonitoringMonitoredProjectRef struct {
	// A reference to an externally managed MonitoringMonitoredProject resource.
	// Should be in the format "locations/global/metricsScopes/{{metricsScope}}/projects/{{project}}".
	External string `json:"external,omitempty"`

	// The name of a MonitoringMonitoredProject resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MonitoringMonitoredProject resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&MonitoringMonitoredProjectRef{})
}

func (r *MonitoringMonitoredProjectRef) GetGVK() schema.GroupVersionKind {
	return MonitoringMonitoredProjectGVK
}

func (r *MonitoringMonitoredProjectRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *MonitoringMonitoredProjectRef) GetExternal() string {
	return r.External
}

func (r *MonitoringMonitoredProjectRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *MonitoringMonitoredProjectRef) ValidateExternal(ref string) error {
	id := &MonitoringMonitoredProjectIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *MonitoringMonitoredProjectRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &MonitoringMonitoredProjectIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *MonitoringMonitoredProjectRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*MonitoringMonitoredProject](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromMonitoringMonitoredProjectSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
