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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getIdentityFromComposerUserWorkloadsSecretSpec(ctx context.Context, reader client.Reader, obj *ComposerUserWorkloadsSecret) (*ComposerUserWorkloadsSecretIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	if obj.Spec.ParentRef == nil {
		return nil, fmt.Errorf("spec.parentRef is required")
	}
	parentRef := *obj.Spec.ParentRef
	if err := parentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentRef: %w", err)
	}
	parsed := &ComposerUserWorkloadsSecretIdentity{}
	if err := parsed.FromExternal(strings.TrimSuffix(parentRef.GetExternal(), "/") + "/userWorkloadsSecrets" + "/_"); err != nil {
		return nil, fmt.Errorf("parsing spec.parentRef external=%q: %w", parentRef.GetExternal(), err)
	}
	parsed.UserWorkloadsSecret = resourceID
	return parsed, nil
}
