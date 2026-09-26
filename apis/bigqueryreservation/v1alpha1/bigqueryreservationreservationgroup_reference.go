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

package v1alpha1

import (
	"context"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BigQueryReservationReservationGroupRef{}

// BigQueryReservationReservationGroupRef is a reference to a GCP BigQueryReservationReservationGroup.
type BigQueryReservationReservationGroupRef struct {
	/* A reference to an externally managed BigQueryReservationReservationGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/reservationGroups/{{reservationGroupID}}". */
	External string `json:"external,omitempty"`

	// The name of a BigQueryReservationReservationGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryReservationReservationGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *BigQueryReservationReservationGroupRef) GetGVK() schema.GroupVersionKind {
	return BigQueryReservationReservationGroupGVK
}

func (r *BigQueryReservationReservationGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *BigQueryReservationReservationGroupRef) GetExternal() string {
	return r.External
}

func (r *BigQueryReservationReservationGroupRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BigQueryReservationReservationGroupRef) ValidateExternal(ref string) error {
	id := &BigQueryReservationReservationGroupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BigQueryReservationReservationGroupRef) ParseExternalToIdentity() (interface{}, error) {
	id := &BigQueryReservationReservationGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BigQueryReservationReservationGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

func init() {
	refs.Register(&BigQueryReservationReservationGroupRef{})
}
