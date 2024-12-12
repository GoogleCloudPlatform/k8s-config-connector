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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ComputeBackendServiceRef{}
var ComputeBackendServiceGVK = GroupVersion.WithKind("ComputeBackendService")

// ComputeBackendServiceRef defines the resource reference to ComputeBackendService, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeBackendServiceRef struct {
	// The value of an externally managed ComputeBackendService resource.
	// Should be in the format "projects/{{project}}/global/backendServices/{{backendService}}"
	// or "projects/{{project}}/regions/{{region}}/backendServices/{{backendService}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeBackendService resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeBackendService resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeBackendService.
// If the "External" is given in the other resource's spec.ComputeBackendServiceRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeBackendService object from the cluster.
func (r *ComputeBackendServiceRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	referenceContext := refsv1beta1.ReferenceContext{IsDirectOnly: false, TargetField: "status.selfLink"}
	// Get value from spec.ComputeBackendServiceRef.external
	if r.External != "" {
		if r.Name != "" {
			return "", fmt.Errorf("cannot specify both name and external on reference")
		}
		if referenceContext.IsDirectOnly {
			if _, err := ParseBackendServiceExternal(r.External); err != nil {
				return "", err
			}
		}
		// To ensure backward compatibility for existing users, we do not enforce external format
		// validation for non-DirectOnly resources
		return r.External, nil
	}

	if r.Name == "" {
		return "", fmt.Errorf("must specify either name or external on reference")
	}

	// Get value from the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}

	u, err := refsv1beta1.ResolveResourceName(ctx, reader, key, ComputeBackendServiceGVK)
	if err != nil {
		return "", err
	}

	actualExternalRef, found, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("error getting status.externalRef for %s %s/%s: %w", u.GetKind(), u.GetNamespace(), u.GetName(), err)
	}
	// If object is DirectOnly, it is created by direct controller.
	// Get value from status.externalRef, which is the most trustworthy source.
	if referenceContext.IsDirectOnly {
		if !found || actualExternalRef == "" {
			return "", fmt.Errorf("status.externalRef is required but is missing or empty for %s %s/%s", u.GetKind(), u.GetNamespace(), u.GetName())
		}
		r.External = actualExternalRef
	}

	// If object not DirectOnly, it can be created by either direct controller or legacy controller, depends on user's settings.
	// If status.externalRef does not exist, it's created by legacy controller. Get values from target field.
	if !found {
		targetField := referenceContext.TargetField
		tokens := strings.Split(targetField, ".")
		targetField, found, err := unstructured.NestedString(u.Object, tokens...)
		if err != nil {
			return "", fmt.Errorf("error getting target field %s for %s %s/%s: %w", targetField, u.GetKind(), u.GetNamespace(), u.GetName(), err)
		}
		if !found || targetField == "" {
			return "", fmt.Errorf("target field %s is required but is missing or empty for %s %s/%s", targetField, u.GetKind(), u.GetNamespace(), u.GetName())
		}
		r.External = targetField
	} else {
		// If status.externalRef exists, it's created by direct controller. Get value from status.externalRef.
		r.External = actualExternalRef
	}
	return r.External, nil
}

func ParseBackendServiceExternal(external string) (identity *BackendServiceIdentity, err error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[3] == "backendServices" {
		return &BackendServiceIdentity{
			parent: &BackendServiceParent{ProjectID: tokens[1], Region: tokens[2]},
			id:     tokens[4],
		}, nil
	} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "backendServices" {
		return &BackendServiceIdentity{
			parent: &BackendServiceParent{ProjectID: tokens[1], Region: tokens[3]},
			id:     tokens[5],
		}, nil
	}
	acceptedFormat := "projects/{{project}}/global/backendServices/{{backendService}} or projects/{{project}}/regions/{{region}}/backendServices/{{backendService}}"
	return nil, fmt.Errorf("nvalid format: %s, allowed format(s): %s", external, acceptedFormat)
}
