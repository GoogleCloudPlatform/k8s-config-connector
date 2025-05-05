// Copyright 2025 Google LLC
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

var _ refsv1beta1.ExternalNormalizer = &AspectTypeRef{}
var DataplexAspectTypeGVK = GroupVersion.WithKind("DataplexAspectType")

// AspectTypeRef defines the resource reference to DataplexAspectTypeRef, which "External" field
// holds the GCP identifier for the KRM object.
type AspectTypeRef struct {
	// A reference to an externally managed DataplexAspectType resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/aspectTypes/{{aspecttypeID}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a DataplexEntryType resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataplexEntryType resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

// NormalizedExternal provision the "External" value for other resource that depends on DataplexAspectTypeRef.
// If the "External" is given in the other resource's spec.DataplexAspectTypeRefRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual DataplexAspectTypeRef object from the cluster.
func (r *AspectTypeRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External == "" {
		return "", fmt.Errorf("external should be specified on %s reference", DataplexAspectTypeGVK.Kind)
	}
	return r.External, nil
}
