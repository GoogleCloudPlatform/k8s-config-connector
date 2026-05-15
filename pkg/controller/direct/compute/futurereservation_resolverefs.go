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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveComputeFutureReservationRefs(ctx context.Context, reader client.Reader, projectMapper *projects.ProjectMapper, obj *krm.ComputeFutureReservation) error {
	if obj.Spec.ShareSettings == nil {
		return nil
	}
	for i := range obj.Spec.ShareSettings.ProjectMap {
		entry := &obj.Spec.ShareSettings.ProjectMap[i]
		if entry.KeyRef != nil {
			kind := entry.KeyRef.Kind
			switch kind {
			// If Kind is not specified, default to "Project"
			case "Project", "":
				projectRef := &refs.ProjectRef{
					External:  entry.KeyRef.External,
					Name:      entry.KeyRef.Name,
					Namespace: entry.KeyRef.Namespace,
				}
				project, err := refs.ResolveProject(ctx, reader, obj.Namespace, projectRef)
				if err != nil {
					return err
				}
				projectID := project.ProjectID
				projectNumber, err := projectMapper.LookupProjectNumber(ctx, projectID)
				if err != nil {
					return err
				}
				entry.KeyRef.External = fmt.Sprintf("%d", projectNumber)

			default:
				return fmt.Errorf("unsupported kind %q for ExtendedProjectRef", kind)
			}
		}
		if entry.Value != nil && entry.Value.ProjectIDRef != nil {
			project, err := refs.ResolveProject(ctx, reader, obj.Namespace, entry.Value.ProjectIDRef)
			if err != nil {
				return err
			}
			projectID := project.ProjectID
			projectNumber, err := projectMapper.LookupProjectNumber(ctx, projectID)
			if err != nil {
				return err
			}
			entry.Value.ProjectIDRef.External = fmt.Sprintf("%d", projectNumber)
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
