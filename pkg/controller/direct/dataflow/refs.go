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

package dataflow

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func normalizeProjectRef(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ProjectRef) (*refs.ProjectRef, error) {
	if ref == nil {
		return nil, nil
	}

	project, err := refs.ResolveProject(ctx, reader, src.GetNamespace(), ref)
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
		resolved, err := RefineComputeSubnetworkRef(r.ctx, r.kube, r.src, subnetworkRef)
		if err != nil {
			return err
		}
		*subnetworkRef = *resolved
	}

	if subnetworkRefs, ok := v.([]refs.ComputeSubnetworkRef); ok {
		for i := range subnetworkRefs {
			subnetworkRef := &subnetworkRefs[i]
			resolved, err := RefineComputeSubnetworkRef(r.ctx, r.kube, r.src, subnetworkRef)
			if err != nil {
				return err
			}
			subnetworkRefs[i] = *resolved
		}
	}

	return nil
}

// RefineComputeSubnetworkRef refine the subnetwork format because DataflowFlexTemplateJob has a specific format requirement:
// "You can specify a subnetwork using either a complete URL or an abbreviated path.
//
//	Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK"
//	or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network,
//	you must use the complete URL.
func RefineComputeSubnetworkRef(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeSubnetworkRef) (*refs.ComputeSubnetworkRef, error) {
	if ref == nil {
		return nil, nil
	}
	// Use common ComputeSubnetwork resolver
	if ref.External == "" {
		var err error
		ref, err = refs.ResolveComputeSubnetwork(ctx, reader, src, ref)
		if err != nil {
			return nil, err
		}
	}

	// Validate non-shared-VPC network format. This is not allowed in the common ComputeSubnetwork resolver
	tokens := strings.Split(ref.External, "/")
	if len(tokens) == 4 && tokens[0] == "regions" && tokens[2] == "subnetworks" {
		return &refs.ComputeSubnetworkRef{
			External: ref.External,
		}, nil
	}

	// Validate and refine the shared-VPC network format to full URL. This is required by GCP service.
	fullURLPrefix := "https://www.googleapis.com/compute/v1/"
	ref.External = strings.TrimPrefix(ref.External, fullURLPrefix)
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "subnetworks" {
		return &refs.ComputeSubnetworkRef{
			External: fullURLPrefix + "projects/" + tokens[1] + "/regions/" + tokens[3] + "/subnetworks/" + tokens[5],
		}, nil
	}
	return nil, fmt.Errorf("format of subnetwork external=%q was not known, use regions/<region>/subnetworks/<subnetwork>, projects/<projectID>/regions/<region>/subnetworks/<subnetwork> or  https://www.googleapis.com/compute/v1/projects/<projectID>/regions/<region>/subnetworks/<subnetwork>", ref.External)
}
