// Copyright 2022 Google LLC
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

package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetSecretVal(secretKeyRef *corekccv1alpha1.SecretKeyReference, secretNamespace string, kubeClient client.Client) (secretVal string, secretVersion string, err error) {
	nn := types.NamespacedName{
		Name:      secretKeyRef.Name,
		Namespace: secretNamespace,
	}
	secret := v1.Secret{}
	if err := kubeClient.Get(context.TODO(), nn, &secret); err != nil {
		if errors.IsNotFound(err) {
			return "", "", NewSecretNotFoundError(nn)
		}
		return "", "", fmt.Errorf("error getting Secret %v: %w", nn, err)
	}
	secretValBytes, ok := secret.Data[secretKeyRef.Key]
	if !ok {
		return "", "", NewKeyInSecretNotFoundError(secretKeyRef.Key, nn)
	}
	return string(secretValBytes), secret.GetResourceVersion(), nil
}

func GetSecretVersionsFromAnnotations(resource *Resource) (map[string]string, error) {
	annotationVal, ok := GetAnnotation(ObservedSecretVersionsAnnotation, resource)
	if !ok {
		return nil, nil
	}
	secretVersions := make(map[string]string)
	if err := json.Unmarshal([]byte(annotationVal), &secretVersions); err != nil {
		return nil, fmt.Errorf("error unmarshalling value of %v: %w", ObservedSecretVersionsAnnotation, err)
	}
	return secretVersions, nil
}

func UpdateOrRemoveObservedSecretVersionsAnnotation(resource *Resource, secretVersions map[string]string, hasSensitiveFields bool) error {
	// The annotation should only be set for resources with sensitive fields.
	if !hasSensitiveFields {
		RemoveAnnotation(ObservedSecretVersionsAnnotation, resource)
		return nil
	}
	b, err := json.Marshal(secretVersions)
	if err != nil {
		return err
	}
	SetAnnotation(ObservedSecretVersionsAnnotation, string(b), resource)
	return nil
}
