// Copyright 2025 Google LLC
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

package direct

import (
	"context"
	"fmt"

	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//func getSecret(ctx context.Context, secretKeyRef *v1alpha1.SecretKeyRef, secretNamespace string, reader client.Reader) (secretVal string, secretVersion string, err error) {
//	nn := types.NamespacedName{
//		Name:      secretKeyRef.Name,
//		Namespace: secretNamespace,
//	}
//	secret := v1.Secret{}
//	if err := reader.Get(ctx, nn, &secret); err != nil {
//		if errors.IsNotFound(err) {
//			return "", "", k8s.NewSecretNotFoundError(nn)
//		}
//		return "", "", fmt.Errorf("error getting Secret %+v: %w", nn, err)
//	}
//	secretValBytes, ok := secret.Data[secretKeyRef.Key]
//	if !ok {
//		return "", "", k8s.NewKeyInSecretNotFoundError(secretKeyRef.Key, nn)
//	}
//	return string(secretValBytes), secret.GetResourceVersion(), nil
//}

func getSecret(ctx context.Context, secretKeyRef *v1alpha1.SecretKeyRef, secretNamespace string, reader client.Reader) (*v1.Secret, error) {
	nn := types.NamespacedName{
		Name:      secretKeyRef.Name,
		Namespace: secretNamespace,
	}
	secret := v1.Secret{}
	if err := reader.Get(ctx, nn, &secret); err != nil {
		if errors.IsNotFound(err) {
			return nil, k8s.NewSecretNotFoundError(nn)
		}
		return nil, fmt.Errorf("error getting Secret %+v: %w", nn, err)
	}
	return &secret, nil
}

func ResolveSensitiveField(ctx context.Context, field *refsv1beta1secret.Legacy, fieldPath, namespace string, reader client.Reader) (*v1.Secret, error) {
	value := field.Value
	valueFrom := field.ValueFrom
	if value != nil && valueFrom != nil {
		return nil, fmt.Errorf("only one of '%s.value' and '%s.valueFrom' "+
			"should be configured: both are configured", fieldPath, fieldPath)
	}
	if value != nil {
		return nil, nil
	}
	if valueFrom != nil {
		if valueFrom.SecretKeyRef == nil {
			return nil, fmt.Errorf("'%s.valueFrom.secretRef' "+
				"should be configured", fieldPath)
		}
		secret, err := getSecret(ctx, valueFrom.SecretKeyRef, namespace, reader)
		if err != nil {
			return nil, err
		}
		secretValBytes, ok := secret.Data[valueFrom.SecretKeyRef.Key]
		if !ok {
			return nil, k8s.NewKeyInSecretNotFoundError(valueFrom.SecretKeyRef.Key, types.NamespacedName{
				Name:      valueFrom.SecretKeyRef.Name,
				Namespace: namespace,
			})
		}
		field.Value = PtrTo(string(secretValBytes))
		// secretVersion = secret.GetResourceVersion()
		return secret, nil
	}
	return nil, fmt.Errorf("at least one of '%s.value' and '%s.valueFrom' "+
		"should be configured: neither is configured", fieldPath, fieldPath)
}
