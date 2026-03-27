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

// resolveSQLUserPasswordRef resolves the password secret reference, if any.
func resolveSQLUserPasswordRef(ctx context.Context, kube client.Reader, obj *krm.SQLUser) error {
	if obj.Spec.Password == nil {
		return nil
	}

	if obj.Spec.Password.Value != nil && obj.Spec.Password.ValueFrom != nil {
		return fmt.Errorf("cannot specify both spec.password.value and spec.password.valueFrom")
	}

	if obj.Spec.Password.Value != nil {
		return nil
	}

	if obj.Spec.Password.ValueFrom != nil {
		if obj.Spec.Password.ValueFrom.SecretKeyRef == nil {
			return fmt.Errorf("spec.password.valueFrom.secretKeyRef must be set when valueFrom is specified")
		}

		key := types.NamespacedName{
			Namespace: obj.Namespace,
			Name:      obj.Spec.Password.ValueFrom.SecretKeyRef.Name,
		}

		secret := &corev1.Secret{}
		if err := kube.Get(ctx, key, secret); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewSecretNotFoundError(key)
			}
			return fmt.Errorf("error reading referenced Secret %v: %w", key, err)
		}

		password := string(secret.Data[obj.Spec.Password.ValueFrom.SecretKeyRef.Key])
		obj.Spec.Password.Value = direct.PtrTo(password)
	}

	return nil
}

// ResolveSQLUserInstanceRef resolves the instanceRef to get the Cloud SQL instance name.
func ResolveSQLUserInstanceRef(ctx context.Context, kube client.Reader, obj *krm.SQLUser) (string, error) {
	instanceRef := &obj.Spec.InstanceRef

	if instanceRef.External != "" && instanceRef.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external for instanceRef")
	}

	if instanceRef.External != "" {
		// Handle full URI format: projects/{project}/instances/{instance}
		if parts := strings.Split(instanceRef.External, "/"); len(parts) == 4 && parts[0] == "projects" && parts[2] == "instances" {
			return parts[3], nil
		}
		return instanceRef.External, nil
	}

	if instanceRef.Name == "" {
		return "", fmt.Errorf("instanceRef.name or instanceRef.external must be specified")
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
		return "", fmt.Errorf("resolving instanceRef %s/%s: %w", nn.Namespace, nn.Name, err)
	}

	// Use spec.resourceID if set, otherwise fall back to metadata.name.
	instanceName := resource.GetName()
	specResourceID, _, _ := unstructured.NestedString(resource.Object, "spec", "resourceID")
	if specResourceID != "" {
		instanceName = specResourceID
	}

	return instanceName, nil
}
