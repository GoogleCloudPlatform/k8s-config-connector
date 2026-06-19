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

package compute

import (
	"context"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveComputeNodeGroupRefs(ctx context.Context, reader client.Reader, projectMapper *projects.ProjectMapper, obj *krm.ComputeNodeGroup) error {
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return err
	}

	if obj.Spec.ShareSettings == nil {
		return nil
	}

	for i := range obj.Spec.ShareSettings.ProjectMap {
		entry := &obj.Spec.ShareSettings.ProjectMap[i]
		if entry.IDRef.External != "" {
			if err := convertToProjectNumber(ctx, projectMapper, &entry.IDRef.External); err != nil {
				return err
			}
		}
		if entry.ProjectIDRef.External != "" {
			if err := convertToProjectNumber(ctx, projectMapper, &entry.ProjectIDRef.External); err != nil {
				return err
			}
		}
	}

	// Sort by IDRef.External to ensure stability in the fixtures test log
	sort.Slice(obj.Spec.ShareSettings.ProjectMap, func(i, j int) bool {
		return obj.Spec.ShareSettings.ProjectMap[i].IDRef.External < obj.Spec.ShareSettings.ProjectMap[j].IDRef.External
	})

	return nil
}
