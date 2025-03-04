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

var _ refsv1beta1.ExternalNormalizer = &ComputeNetworkAttachmentRef{}
var ComputeNetworkAttachmentGVK = GroupVersion.WithKind("ComputeNetworkAttachment")

// ComputeNetworkAttachmentRef defines the resource reference to ComputeNetworkAttachment, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeNetworkAttachmentRef struct {

	// The value of an externally managed ComputeNetworkAttachment resource in
	// the form "projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}".
	External string `json:"external,omitempty"`

	// ComputeNetworkAttachment not yet supported in Config Connector, users
	// should only use 'external' field to reference existing resources.
	/* NOTYET
	// The name of a ComputeNetworkAttachment resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeNetworkAttachment resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeNetworkAttachment.
// If the "External" is given in the other resource's spec.ComputeNetworkAttachmentRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeNetworkAttachment object from the cluster.
func (r *ComputeNetworkAttachmentRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External == "" {
		return "", fmt.Errorf("external should be specified on %s reference", ComputeNetworkAttachmentGVK.Kind)
	}
	return r.External, nil
}
