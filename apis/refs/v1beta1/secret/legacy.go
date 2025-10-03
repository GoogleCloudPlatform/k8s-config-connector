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

package secret

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:object:generate:=true
type Legacy struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *LegacyValueFrom `json:"valueFrom,omitempty"`
}

// +kubebuilder:object:generate:=true
type LegacyValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

func NormalizedLegacySecret(ctx context.Context, r *v1alpha1.SecretKeyRef, reader client.Reader, otherNamespace string) ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	if r.Name == "" {
		return nil, fmt.Errorf("Secret `name` is required ")
	}
	nn := types.NamespacedName{
		Namespace: otherNamespace,
		Name:      r.Name,
	}

	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
	}
	if err := reader.Get(ctx, nn, secret); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced Secret %v not found", nn)
		}
		return nil, fmt.Errorf("error reading referenced Secret %v: %w", nn, err)
	}
	data, ok := secret.Data[r.Key]
	if !ok {
		return nil, fmt.Errorf("%s not found in Secret %s", r.Key, r.Name)
	}
	return data, nil
}
