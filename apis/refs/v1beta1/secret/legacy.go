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

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
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

// NormalizeSecret normalizes the Legacy secret reference, populating the Value if ValueFrom is set.
// ReadSecretValue is normally safer to use, as it does not leave the secret value in memory,
// but this is easier to use with our mapping framework.
func (r *Legacy) NormalizeSecret(ctx context.Context, fieldPath string, namespace string, reader client.Reader) error {
	value, err := r.ReadSecretValue(ctx, fieldPath, namespace, reader)
	if err != nil {
		return err
	}
	r.Value = value
	r.ValueFrom = nil
	return nil
}

// ReadSecretValue returns the secret value, resolving the value if needed.
func (r *Legacy) ReadSecretValue(ctx context.Context, fieldPath string, namespace string, reader client.Reader) (*string, error) {
	value := r.Value
	valueFrom := r.ValueFrom
	if value != nil && valueFrom != nil {
		return nil, fmt.Errorf("only one of '%s.value' and '%s.valueFrom' "+
			"should be configured: both are configured", fieldPath, fieldPath)
	}
	if value != nil {
		return value, nil
	}
	if valueFrom != nil {
		if valueFrom.SecretKeyRef == nil {
			return nil, fmt.Errorf("'%s.valueFrom.secretRef' "+
				"should be configured", fieldPath)
		}
		secretValue, err := valueFrom.SecretKeyRef.ReadSecretValue(ctx, reader, namespace)
		if err != nil {
			return nil, err
		}
		value := string(secretValue)
		return &value, nil
	}
	return nil, fmt.Errorf("at least one of '%s.value' and '%s.valueFrom' "+
		"should be configured: neither is configured", fieldPath, fieldPath)
}

// +kubebuilder:object:generate:=true
type LegacyValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *LegacyValueFromSecretKeyRef `json:"secretKeyRef,omitempty"`
}

// +kubebuilder:object:generate:=true
type LegacyValueFromSecretKeyRef struct {
	/* Key that identifies the value to be extracted. */
	Key string `json:"key"`

	/* Name of the Secret to extract a value from. */
	Name string `json:"name"`
}

func (r *LegacyValueFromSecretKeyRef) ReadSecretValue(ctx context.Context, reader client.Reader, defaultNamespace string) ([]byte, error) {
	if r == nil {
		return nil, fmt.Errorf("SecretKeyRef is nil")
	}
	if r.Name == "" {
		return nil, fmt.Errorf("Secret `name` is required ")
	}
	nn := types.NamespacedName{
		Namespace: defaultNamespace,
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
			return nil, k8s.NewSecretNotFoundError(nn)
		}
		return nil, fmt.Errorf("error reading referenced Secret %v: %w", nn, err)
	}
	data, ok := secret.Data[r.Key]
	if !ok {
		return nil, fmt.Errorf("%s not found in Secret %s", r.Key, r.Name)
	}
	return data, nil
}
