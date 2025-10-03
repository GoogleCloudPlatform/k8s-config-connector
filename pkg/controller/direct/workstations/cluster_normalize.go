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

package workstations

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
)

func NormalizeWorkstationCluster(ctx context.Context, kube client.Reader, obj *krm.WorkstationCluster) error {
	// Resolve network.
	if err := obj.Spec.NetworkRef.Normalize(ctx, kube, obj); err != nil {
		return err
	}

	// Resolve subnetwork.
	subnet, err := refs.ResolveComputeSubnetwork(ctx, kube, obj, &obj.Spec.SubnetworkRef)
	if err != nil {
		return err
	}
	obj.Spec.SubnetworkRef.External = subnet.External

	// Resolve projects (in private cluster config).
	if obj.Spec.PrivateClusterConfig != nil && obj.Spec.PrivateClusterConfig.AllowedProjects != nil {
		var resolvedProjects []refs.ProjectRef
		for _, projectRef := range obj.Spec.PrivateClusterConfig.AllowedProjects {
			resolvedProject, err := refs.ResolveProject(ctx, kube, obj.GetNamespace(), &projectRef)
			if err != nil {
				return err
			}
			resolvedProjects = append(resolvedProjects, refs.ProjectRef{
				External: resolvedProject.ProjectID,
			})
		}
		obj.Spec.PrivateClusterConfig.AllowedProjects = resolvedProjects
	}
	return nil
}
