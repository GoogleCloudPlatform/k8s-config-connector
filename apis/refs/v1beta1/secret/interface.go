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
)

type SecretRef interface {
	GetName() string
	GetNamespace() string
	Set(*corev1.Secret) error
}

func NormalizedSecret(ctx context.Context, r SecretRef, reader client.Reader, otherNamespace string) error {
	if r == nil {
		return nil
	}
	if r.GetName() == "" {
		return fmt.Errorf("Secret `name` is required ")
	}
	nn := types.NamespacedName{
		Namespace: r.GetNamespace(),
		Name:      r.GetName(),
	}
	if nn.Namespace == "" {
		nn.Namespace = otherNamespace
	}

	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
	}
	if err := reader.Get(ctx, nn, secret); err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("referenced Secret %v not found", nn)
		}
		return fmt.Errorf("error reading referenced Secret %v: %w", nn, err)
	}
	return r.Set(secret)
}
