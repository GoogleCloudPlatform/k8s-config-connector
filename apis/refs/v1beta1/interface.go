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
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Resolver define the required methods to resolve the external identifier of a ConfigConnector object.
type Resolver interface {
	// Resolve assigns value to the "External" field, based on the Unstructured ConfigConnector object.
	// In general, `status.externalRef` shall be used if present, otherwise it shall be built from the `spec` parent fields
	// and metadata.name (or `spec.resourceID` if applicable).
	Resolve(context.Context, client.Reader, *unstructured.Unstructured) (string, error)

	// Validate the value of the "External" field
	ValidateExternal() error
}

func NormalizeExternal(ctx context.Context, reader client.Reader, gvk schema.GroupVersionKind, r Resolver) (string, error) {
	external, name, namespace := "", "", ""
	refV := reflect.ValueOf(r).Elem()

	for i := 0; i < refV.NumField(); i++ {
		field := refV.Type().Field(i)
		val := refV.Field(i).String()
		if field.Name == "External" {
			external = val
		}
		if field.Name == "Name" {
			name = val
		}
		if field.Name == "Namespace" {
			namespace = val
		}
	}

	if external != "" && name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", gvk.Kind)
	}
	if external != "" {
		if err := r.ValidateExternal(); err != nil {
			return "", err
		}
		return external, nil
	}
	key := types.NamespacedName{Name: name, Namespace: namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", gvk, key, err)
	}

	external, err := r.Resolve(ctx, reader, u)
	if err != nil {
		return "", fmt.Errorf("resolve external reference from %s %s: %w", gvk, key, err)
	}
	return external, nil
}
