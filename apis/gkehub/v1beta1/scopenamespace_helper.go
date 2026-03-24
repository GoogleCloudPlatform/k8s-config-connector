// Copyright 2026 Google LLC
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

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveScopeID(ctx context.Context, reader client.Reader, namespace string, scopeRef *GKEHubScopeRef) (string, error) {
	if scopeRef.External != nil && *scopeRef.External != "" {
		id := &GKEHubScopeIdentity{}
		if err := id.FromExternal(*scopeRef.External); err != nil {
			return "", err
		}
		return id.ID(), nil
	}

	name := ""
	if scopeRef.Name != nil {
		name = *scopeRef.Name
	}
	if name == "" {
		return "", fmt.Errorf("scopeRef must have external or name")
	}

	ns := namespace
	if scopeRef.Namespace != nil && *scopeRef.Namespace != "" {
		ns = *scopeRef.Namespace
	}

	key := types.NamespacedName{Namespace: ns, Name: name}
	scope := &GKEHubScope{}
	if err := reader.Get(ctx, key, scope); err != nil {
		return "", err
	}

	if scope.Spec.ResourceID != nil && *scope.Spec.ResourceID != "" {
		return *scope.Spec.ResourceID, nil
	}
	return scope.GetName(), nil
}
