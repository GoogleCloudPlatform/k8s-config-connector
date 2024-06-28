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

package util

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	multiclusterv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/api/v1alpha1"
)

func GetMultiClusterLease(ctx context.Context, c client.Client, nn types.NamespacedName) (*multiclusterv1alpha1.MultiClusterLease, error) {
	mcl := &multiclusterv1alpha1.MultiClusterLease{}
	if err := c.Get(ctx, nn, mcl); err != nil {
		return nil, fmt.Errorf("failed to get MultiClusterLease %s: %w", nn.Name, err)
	}
	return mcl, nil
}

func UpdateMultiClusterLease(ctx context.Context, c client.Client, mcl *multiclusterv1alpha1.MultiClusterLease) error {
	if err := c.Update(ctx, mcl); err != nil {
		return fmt.Errorf("failed to update MultiClusterLease %s: %w", mcl.GetName(), err)
	}
	return nil
}
