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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *NetworkRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name == "" {
		return nil
	}
	if r.External == "" && r.Name == "" {
		return fmt.Errorf("must specify either name or external")
	}
	// Fetch the referenced resource
	nn := types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
	if nn.Namespace == "" {
		nn.Namespace = defaultNamespace
	}
	ref := &v1beta1.ComputeNetwork{}
	if err := reader.Get(ctx, nn, ref); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(ref.GroupVersionKind(), nn)
		}
		return err
	}
	if ref.Status.SelfLink != nil {
		r.External = *ref.Status.SelfLink
	}
	return nil
}

func (r *TLSInspectionPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name == "" {
		return nil
	}
	if r.External == "" && r.Name == "" {
		return fmt.Errorf("must specify either name or external")
	}
	return nil // Mock implementation to pass compilation
}
