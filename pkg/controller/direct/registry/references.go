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

package registry

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func AdapterForReference(ctx context.Context, reader client.Reader, sourceNamespace string, resourceRef v1beta1.ResourceReference) (directbase.Adapter, error) {
	obj := &unstructured.Unstructured{}

	var gk schema.GroupKind
	switch resourceRef.Kind {
	default:
		gk = resourceRef.GroupVersionKind().GroupKind()
	}

	if gk.Group == "" {
		return nil, fmt.Errorf("cannot find group for reference %v (must set apiVersion)", resourceRef)
	}

	if resourceRef.External != "" {
		uri := ""
		if !strings.HasPrefix(uri, "//") {
			switch gk.Group {
			case "privateca.cnrm.google.com":
				uri = "//privateca.googleapis.com/" + resourceRef.External
			default:
				return nil, fmt.Errorf("unknown format for external reference for %v: %q", gk, resourceRef.External)
			}
		}

		adapter, err := AdapterForURL(ctx, uri)
		if err != nil {
			return nil, fmt.Errorf("resolving %q: %w", uri, err)
		}
		if adapter == nil {
			return nil, fmt.Errorf("unknown format for external reference for %v: %q", gk, resourceRef.External)
		}
		return adapter, nil
	}

	model, err := GetModel(gk)
	if err != nil {
		return nil, fmt.Errorf("cannot handle references to %v (in direct controller)", gk)
	}

	gvk, ok := PreferredGVK(gk)
	if !ok {
		return nil, fmt.Errorf("preferred GVK is not known for %v", gk)
	}

	obj.SetGroupVersionKind(gvk)
	nn := types.NamespacedName{
		Namespace: resourceRef.Namespace,
		Name:      resourceRef.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = sourceNamespace
	}

	if err := reader.Get(ctx, nn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(gvk, nn)
		}
		return nil, fmt.Errorf("error retrieving resource '%v' with GroupVersionKind '%v': %w", nn, gvk, err)
	}

	adapter, err := model.AdapterForObject(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("building adapter: %w", err)
	}

	return adapter, nil
}
