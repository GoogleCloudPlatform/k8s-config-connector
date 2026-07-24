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

package sql

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// resolveSQLUserPasswordRef resolves the password from spec.password.value or
// spec.password.valueFrom.secretKeyRef, returning the resolved password string.
// It does NOT modify the input obj.
func resolveSQLUserPasswordRef(ctx context.Context, kube client.Reader, obj *krm.SQLUser) (string, error) {
	if obj.Spec.Password == nil {
		return "", nil
	}

	if obj.Spec.Password.Value != nil && obj.Spec.Password.ValueFrom != nil {
		return "", fmt.Errorf("cannot specify both spec.password.value and spec.password.valueFrom")
	}

	if obj.Spec.Password.Value != nil {
		return direct.ValueOf(obj.Spec.Password.Value), nil
	}

	if obj.Spec.Password.ValueFrom != nil {
		if obj.Spec.Password.ValueFrom.SecretKeyRef == nil {
			return "", fmt.Errorf("spec.password.valueFrom.secretKeyRef must be set when valueFrom is specified")
		}

		key := types.NamespacedName{
			Namespace: obj.Namespace,
			Name:      obj.Spec.Password.ValueFrom.SecretKeyRef.Name,
		}

		secret := &corev1.Secret{}
		if err := kube.Get(ctx, key, secret); err != nil {
			if apierrors.IsNotFound(err) {
				return "", k8s.NewSecretNotFoundError(key)
			}
			return "", fmt.Errorf("error reading referenced Secret %v: %w", key, err)
		}

		passwordBytes, ok := secret.Data[obj.Spec.Password.ValueFrom.SecretKeyRef.Key]
		if !ok {
			return "", fmt.Errorf("key %q not found in Secret %v", obj.Spec.Password.ValueFrom.SecretKeyRef.Key, key)
		}
		return string(passwordBytes), nil
	}

	return "", nil
}

// ResolveSQLUserInstanceRef resolves the instanceRef to get the Cloud SQL instance name
// and optionally the project ID (for cross-project references).
// Returns (instanceID, projectID, error). projectID may be empty if not determinable.
func ResolveSQLUserInstanceRef(ctx context.Context, kube client.Reader, obj *krm.SQLUser) (string, string, error) {
	instanceRef := &obj.Spec.InstanceRef

	if instanceRef.External != "" && instanceRef.Name != "" {
		return "", "", fmt.Errorf("cannot specify both name and external for instanceRef")
	}

	if instanceRef.External != "" {
		instanceID, projectID := parseInstanceExternal(instanceRef.External)
		return instanceID, projectID, nil
	}

	if instanceRef.Name == "" {
		return "", "", fmt.Errorf("instanceRef.name or instanceRef.external must be specified")
	}

	nn := types.NamespacedName{
		Namespace: instanceRef.Namespace,
		Name:      instanceRef.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = obj.GetNamespace()
	}

	resource := &unstructured.Unstructured{}
	resource.SetGroupVersionKind(krm.SQLInstanceGVK)
	if err := kube.Get(ctx, nn, resource); err != nil {
		return "", "", fmt.Errorf("resolving instanceRef %s/%s: %w", nn.Namespace, nn.Name, err)
	}

	// Use spec.resourceID if set, otherwise fall back to metadata.name.
	instanceName := resource.GetName()
	specResourceID, _, _ := unstructured.NestedString(resource.Object, "spec", "resourceID")
	if specResourceID != "" {
		instanceName = specResourceID
	}

	// Read the instance's project annotation for cross-project support.
	instanceProjectID, _ := resource.GetAnnotations()[k8s.ProjectIDAnnotation]

	return instanceName, instanceProjectID, nil
}

// parseInstanceExternal parses an instance external reference, which can be:
// - A plain instance name: "my-instance"
// - A relative path: "projects/{project}/instances/{instance}"
// - A full selfLink URL: "https://sqladmin.googleapis.com/.../projects/{project}/instances/{instance}"
// Returns (instanceID, projectID). projectID is empty if not present in the reference.
func parseInstanceExternal(external string) (string, string) {
	parts := strings.Split(external, "/")
	// Walk backwards looking for "instances" to handle any prefix length.
	for i := len(parts) - 2; i >= 0; i-- {
		if parts[i] == "instances" {
			instanceID := parts[i+1]
			// Look for "projects" before "instances".
			for j := i - 1; j >= 0; j-- {
				if parts[j] == "projects" && j+1 < i {
					return instanceID, parts[j+1]
				}
			}
			return instanceID, ""
		}
	}
	// Plain instance name.
	return external, ""
}
