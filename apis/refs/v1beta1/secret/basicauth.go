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
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

var _ SecretRef = &BasicAuthSecretRef{}

type BasicAuthSecretRef struct {
	// +required
	// The `metadata.name` field of a Kubernetes `Secret`
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a Kubernetes `Secret`.
	Namespace string `json:"namespace,omitempty"`

	// The public field with json:"-" tag is to skip the field
	// in the CRD, and bypass "the unexported field error"
	// when controller-gen parses the Unstructured object to a typed object.
	Username string `json:"-"`
	Password string `json:"-"`
}

func (b *BasicAuthSecretRef) GetName() string {
	return b.Name
}
func (b *BasicAuthSecretRef) GetNamespace() string {
	return b.Namespace
}

func (b *BasicAuthSecretRef) Set(secret *corev1.Secret) error {
	if secret.Type != corev1.SecretTypeBasicAuth {
		return fmt.Errorf("the referenced Secret %s should use type %s, got %s",
			b.Name, corev1.SecretTypeBasicAuth, secret.Type)
	}
	if secret.Data != nil {
		b.Username = string(secret.Data["username"])
		b.Password = string(secret.Data["password"])
	}
	return nil
}
