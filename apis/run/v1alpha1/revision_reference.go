// Copyright 2024 Google LLC
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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &RevisionRef{}

// Revision defines the resource reference to RunRevision, which "External" field
// holds the GCP identifier for the KRM object.
type RevisionRef struct {
	// A reference to an externally managed RunRevision resource.
	// Should be in the format "projects/{{project}}/locations/{{location}}/revisions/{{revision}}".
	External string `json:"external"`

	// The resource is not yet supported so referencing by name is not possible.
	//
	//// The name of a ComputeFirewallPolicyRule resource.
	//Name string `json:"name,omitempty"`
	//
	//// The namespace of a ComputeFirewallPolicyRule resource.
	//Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeFirewallPolicyRule.
// If the "External" is given in the other resource's spec.ComputeFirewallPolicyRuleRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeFirewallPolicyRule object from the cluster.
func (r *RevisionRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	//if r.External != "" && r.Name != "" {
	//	return "", fmt.Errorf("cannot specify both name and external on %s reference", RunRevisionGVK.Kind)
	//}
	// From given External
	if r.External == "" {
		return "", fmt.Errorf("external must be specified")
	}

	if err := parseRevisionExternal(r.External); err != nil {
		return "", err
	}
	return r.External, nil

	//// From the Config Connector object
	//if r.Namespace == "" {
	//	r.Namespace = otherNamespace
	//}
	//key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	//u := &unstructured.Unstructured{}
	//u.SetGroupVersionKind(RunRevisionGVK)
	//if err := reader.Get(ctx, key, u); err != nil {
	//	if apierrors.IsNotFound(err) {
	//		return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	//	}
	//	return "", fmt.Errorf("reading referenced %s %s: %w", RunRevisionGVK, key, err)
	//}
	//// Get external from status.externalRef. This is the most trustworthy place.
	//actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	//if err != nil {
	//	return "", fmt.Errorf("reading status.externalRef: %w", err)
	//}
	//if actualExternalRef == "" {
	//	return "", fmt.Errorf("RunRevision is not ready yet")
	//}
	//r.External = actualExternalRef
	//return r.External, nil
}

func parseRevisionExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "revisions" {
		return fmt.Errorf("format of RunRevision external=%q was not known (use projects/{{project}}/locations/{{location}}/revisions/{{revision}})", external)
	}
	return nil
}
