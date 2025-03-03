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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SpannerInstanceRef struct {
	/* The SpannerInstance selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `SpannerInstance` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `SpannerInstance` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type SpannerInstance struct {
	ProjectID    string
	InstanceName string
}

func (s *SpannerInstance) String() string {
	return "projects/" + s.ProjectID + "/instances/" + s.InstanceName
}

func ResolveSpannerInstanceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *SpannerInstanceRef) (*SpannerInstance, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on instanceRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.instanceRef.name and spec.instanceRef.external")
	}

	if ref.External != "" {
		// External must be in form `projects/<projectID>/instances/<instanceName>`.
		// see https://cloud.google.com/spanner/docs/reference/rest#rest-resource:-v1.projects.instances
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instances" {
			return &SpannerInstance{
				ProjectID:    tokens[1],
				InstanceName: tokens[3],
			}, nil
		}
		return nil, fmt.Errorf("format of sqlinstance external=%q was not known (use projects/<projectID>/instances/<instanceName>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	spannerinstance := &unstructured.Unstructured{}
	spannerinstance.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "spanner.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SpannerInstance",
	})
	if err := reader.Get(ctx, key, spannerinstance); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced SpannerInstance %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced SpannerInstance %v: %w", key, err)
	}
	resource, err := k8s.NewResource(spannerinstance)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(spannerinstance.GroupVersionKind(), key)
	}

	resourceID, _, err := unstructured.NestedString(spannerinstance.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from Spanner Database %s/%s: %w", spannerinstance.GetNamespace(), spannerinstance.GetName(), err)
	}
	if resourceID == "" {
		resourceID = spannerinstance.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, spannerinstance)
	if err != nil {
		return nil, err
	}

	return &SpannerInstance{
		ProjectID:    projectID,
		InstanceName: resourceID,
	}, nil
}

func ResolveSpannerInstanceID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	instanceRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "instanceRef", "external")
	if instanceRefExternal != "" {
		instanceRef := &SpannerInstanceRef{
			External: instanceRefExternal,
		}
		instance, err := ResolveSpannerInstanceRef(ctx, reader, obj, instanceRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse instanceRef.external %q in %v %v/%v: %w", instanceRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return instance.InstanceName, nil
	}

	instanceRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "instanceRef", "name")
	if instanceRefName != "" {
		namespace, _, _ := unstructured.NestedString(obj.Object, "spec", "instanceRef", "namespace")
		instanceRef := &SpannerInstanceRef{
			Name:      instanceRefName,
			Namespace: namespace,
		}
		if instanceRef.Namespace == "" {
			instanceRef.Namespace = obj.GetNamespace()
		}
		instance, err := ResolveSpannerInstanceRef(ctx, reader, obj, instanceRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse instanceRef.name %q in %v %v/%v: %w", instanceRefName, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return instance.InstanceName, nil
	}

	return "", fmt.Errorf("cannot find instance id for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
