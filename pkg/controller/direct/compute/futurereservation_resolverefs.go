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
	"fmt"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveComputeFutureReservationRefs(ctx context.Context, reader client.Reader, projectMapper *projects.ProjectMapper, obj *krm.ComputeFutureReservation) error {
	if obj.Spec.ShareSettings == nil {
		return nil
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return err
	}

	for i := range obj.Spec.ShareSettings.ProjectMap {
		entry := &obj.Spec.ShareSettings.ProjectMap[i]
		if entry.KeyRef != nil && (entry.KeyRef.Kind == "" || entry.KeyRef.Kind == "Project") {
			if err := convertToProjectNumber(ctx, projectMapper, &entry.KeyRef.External); err != nil {
				return err
			}
		}
		if entry.Value != nil && entry.Value.ProjectIDRef != nil {
			if err := convertToProjectNumber(ctx, projectMapper, &entry.Value.ProjectIDRef.External); err != nil {
				return err
			}
		}
	}

	// Sort by KeyRef.External to ensure stability in the fixtures test log
	sort.Slice(obj.Spec.ShareSettings.ProjectMap, func(i, j int) bool {
		iKeyRef := obj.Spec.ShareSettings.ProjectMap[i].KeyRef
		jKeyRef := obj.Spec.ShareSettings.ProjectMap[j].KeyRef

		if iKeyRef == nil && jKeyRef == nil {
			return false
		}
		if iKeyRef == nil {
			return true
		}
		if jKeyRef == nil {
			return false
		}
		return iKeyRef.External < jKeyRef.External
	})

	return nil
}

func convertToProjectNumber(ctx context.Context, projectMapper *projects.ProjectMapper, external *string) error {
	id := &refs.ProjectIdentity{}
	if err := id.FromExternal(*external); err != nil {
		return err
	}
	projectNumber, err := projectMapper.LookupProjectNumber(ctx, id.ProjectID)
	if err != nil {
		return err
	}
	*external = fmt.Sprintf("%d", projectNumber)
	return nil
}
