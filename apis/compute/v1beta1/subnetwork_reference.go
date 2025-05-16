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

var _ refsv1beta1.ExternalNormalizer = &ComputeNetworkRef{}
var ComputeSubnetworkGVK = GroupVersion.WithKind("ComputeSubnetwork")

// ComputeSubnetworkRef defines the resource reference to ComputeSubnetwork, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeSubnetworkRef struct {
	// The value of an externally managed ComputeSubnetwork resource.
	External string `json:"external,omitempty"`

	// The name of a ComputeSubnetwork resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeSubnetwork resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeSubnetwork.
// If the "External" is given in the other resource's spec.ComputeSubnetworkRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeSubnetwork object from the cluster.
func (ref *ComputeSubnetworkRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if ref.External != "" && ref.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeSubnetworkGVK.Kind)
	}
	// From given External
	// For backward compatibility, we are not validating the external format.
	// todo: validate external when it's referenced by a pure direct resource
	if ref.External != "" {
		return ref.External, nil
	}

	// From the Config Connector object
	if ref.Namespace == "" {
		ref.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: ref.Name, Namespace: ref.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeSubnetworkGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ComputeSubnetworkGVK, key, err)
	}

	externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
	if externalRef != "" {
		return externalRef, nil
	}

	selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
	if selfLink == "" {
		return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	}

	external := fixStaleExternalFormat(selfLink)
	return external, nil
}

func ParseComputeSubnetworkExternal(external string) (*SubnetworkIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeSubnetwork external value")
	}
	external = fixStaleExternalFormat(external)
	tokens := strings.Split(external, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "subnetworks" {
		return &SubnetworkIdentity{
			parent: &SubnetworkParent{ProjectID: tokens[1], Region: tokens[3]},
			id:     tokens[5],
		}, nil
	}
	return nil, fmt.Errorf("format of computeSubnetwork external=%q was not known (use projects/<project>/regions/<region>/subnetworks/<subnetworkid>)", external)
}
