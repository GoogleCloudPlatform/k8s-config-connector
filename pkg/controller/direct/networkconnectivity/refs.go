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

package networkconnectivity

import (
	"context"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func normalizeProjectRef(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ProjectRef) (*refs.ProjectRef, error) {
	if ref == nil {
		return nil, nil
	}

	project, err := refs.ResolveProject(ctx, reader, src, ref)
	if err != nil {
		return nil, err
	}

	return &refs.ProjectRef{
		External: "projects/" + project.ProjectID,
	}, nil
}

type refNormalizer struct {
	ctx     context.Context
	kube    client.Reader
	src     client.Object
	project refs.Project
}

func (r *refNormalizer) VisitField(path string, v any) error {
	if projectRef, ok := v.(*refs.ProjectRef); ok {
		if ref, err := normalizeProjectRef(r.ctx, r.kube, r.src, projectRef); err != nil {
			return err
		} else if ref != nil {
			*projectRef = *ref
		}
	}

	if networkRef, ok := v.(*refs.ComputeNetworkRef); ok {
		if err := networkRef.Normalize(r.ctx, r.kube, r.src); err != nil {
			return err
		}
	}

	if subnetworkRef, ok := v.(*refs.ComputeSubnetworkRef); ok {
		resolved, err := refs.ResolveComputeSubnetwork(r.ctx, r.kube, r.src, subnetworkRef)
		if err != nil {
			return err
		}
		*subnetworkRef = *resolved
	}

	if subnetworkRefs, ok := v.([]refs.ComputeSubnetworkRef); ok {
		for i := range subnetworkRefs {
			subnetworkRef := &subnetworkRefs[i]
			resolved, err := refs.ResolveComputeSubnetwork(r.ctx, r.kube, r.src, subnetworkRef)
			if err != nil {
				return err
			}
			subnetworkRefs[i] = *resolved
		}
	}

	return nil
}
