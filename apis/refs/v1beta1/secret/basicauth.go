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

var _ SecretRef = &BasicAuthSecret{}

type BasicAuthSecret struct {
	// +required
	// The `metadata.name` field of a Kubernetes `Secret`
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a Kubernetes `Secret`.
	Namespace string `json:"namespace,omitempty"`

	Username string `json:"-"`
	Password string `json:"-"`
}

func (b *BasicAuthSecret) GetName() string {
	return b.Name
}
func (b *BasicAuthSecret) GetNamespace() string {
	return b.Namespace
}

func (b *BasicAuthSecret) Set(secret *corev1.Secret) error {
	if secret.Type != corev1.SecretTypeBasicAuth {
		return fmt.Errorf("the referenced Secret in `spec.cloudSQL.credential.secretRef` should use type %s, got %s",
			corev1.SecretTypeBasicAuth, secret.Type)
	}
	if secret.Data != nil {
		b.Username = string(secret.Data["username"])
		b.Password = string(secret.Data["password"])
	}
	if secret.StringData != nil {
		b.Username = secret.StringData["username"]
		b.Password = secret.StringData["password"]
	}
	return nil
}
