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

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

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

	if networkRef, ok := v.(*computev1beta1.ComputeNetworkRef); ok {
		external, err := networkRef.NormalizedExternal(r.ctx, r.kube, r.src.GetNamespace())
		if err != nil {
			return err
		}
		networkRef.External = external
	}

	if subnetworkRef, ok := v.(*computev1beta1.ComputeSubnetworkRef); ok {
		external, err := subnetworkRef.NormalizedExternal(r.ctx, r.kube, r.src.GetNamespace())
		if err != nil {
			return err
		}
		refined, err := refineComputeSubnetworkRef(external)
		if err != nil {
			return err
		}
		subnetworkRef.External = refined
	}

	if subnetworkRefs, ok := v.([]computev1beta1.ComputeSubnetworkRef); ok {
		for i := range subnetworkRefs {
			subnetworkRef := &subnetworkRefs[i]
			external, err := subnetworkRef.NormalizedExternal(r.ctx, r.kube, r.src.GetNamespace())
			if err != nil {
				return err
			}
			refined, err := refineComputeSubnetworkRef(external)
			if err != nil {
				return err
			}
			subnetworkRefs[i].External = refined
		}
	}

	return nil
}

// refineComputeSubnetworkRef refine the subnetwork format because DataflowFlexTemplateJob has a specific format requirement:
// "You can specify a subnetwork using either a complete URL or an abbreviated path.
//
//	Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK"
//	or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network,
//	you must use the complete URL.
func refineComputeSubnetworkRef(external string) (string, error) {
	if external == "" {
		return "", fmt.Errorf("ComputeSubnetworkRef's external is empty")
	}
	// Validate if the non shared-VPC network format id abbreviated path. This is required by GCP service.
	tokens := strings.Split(external, "/")
	if len(tokens) == 4 && tokens[0] == "regions" && tokens[2] == "subnetworks" {
		return external, nil
	}

	// Validate and refine the shared-VPC network format to full URL. This is required by GCP service.
	fullURLPrefix := "https://www.googleapis.com/compute/v1/"
	trimmed := strings.TrimPrefix(external, fullURLPrefix)
	tokens = strings.Split(trimmed, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "subnetworks" {
		return fullURLPrefix + trimmed, nil
	}
	return "", fmt.Errorf("format of subnetwork external=%q was not known, use regions/<region>/subnetworks/<subnetwork>, projects/<projectID>/regions/<region>/subnetworks/<subnetwork> or  https://www.googleapis.com/compute/v1/projects/<projectID>/regions/<region>/subnetworks/<subnetwork>", external)
}
