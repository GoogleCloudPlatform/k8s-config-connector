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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var BigQueryReservationReservationGroupGVK = GroupVersion.WithKind("BigQueryReservationReservationGroup")

var _ refsv1beta1.ExternalNormalizer = &ReservationGroupRef{}

// ReservationGroupRef is a reference to a BigQueryReservationReservationGroup.
type ReservationGroupRef struct {
	// A reference to an externally managed BigQueryReservationReservationGroup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/reservationGroups/{{reservationGroupID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryReservationReservationGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryReservationReservationGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provisions the "External" value for other resource that depends on BigQueryReservationReservationGroup.
func (r *ReservationGroupRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigQueryReservationReservationGroupGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseReservationGroupExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(BigQueryReservationReservationGroupGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigQueryReservationReservationGroupGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

type ReservationGroupParent struct {
	ProjectID string
	Location  string
}

func ParseReservationGroupExternal(external string) (parent *ReservationGroupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "reservationGroups" {
		return nil, "", fmt.Errorf("format of BigQueryReservationReservationGroup external=%q was not known (use projects/{{projectID}}/locations/{{location}}/reservationGroups/{{reservationGroupID}})", external)
	}
	parent = &ReservationGroupParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
