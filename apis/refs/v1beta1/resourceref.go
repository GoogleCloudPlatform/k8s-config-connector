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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ComputeNetworkRef_ConvertToExternal(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	// Validate the ResourceRef
	if err := ValidateResourceRef(ref); err != nil {
		return nil, err
	}

	// Validate and get external references if exist
	if ref.External != "" {
		if err := validateExternalComputeNetwork(ref); err != nil {
			return nil, err
		}
		return ref, nil
	}

	// Convert references to external references
	// Parse the ComputeNetwork resource
	computenetwork, err := ParseResourceRef(ctx, reader, src, ref, "compute.cnrm.cloud.google.com", "v1beta1", "ComputeNetwork")
	if err != nil {
		return nil, err
	}

	// Get the resourceID
	computenetworkID, err := GetResourceID(computenetwork)
	if err != nil {
		return nil, err
	}

	// Get the project ID
	computeNetworkProjectID, err := GetProjectID(ctx, reader, computenetwork)
	if err != nil {
		return nil, err
	}

	return &v1alpha1.ResourceRef{
		External: fmt.Sprintf("projects/%s/global/networks/%s", computeNetworkProjectID, computenetworkID),
	}, nil
}

func ComputeTargetHTTPProxyRef_ConvertToExternal(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	if err := ValidateResourceRef(ref); err != nil {
		return nil, err
	}

	if ref.External != "" {
		if err := validateExternalComputeTargetHTTPProxy(ref); err != nil {
			return nil, err
		}
		return ref, nil
	}

	computeTargetHTTPProxy, err := ParseResourceRef(ctx, reader, src, ref, "compute.cnrm.cloud.google.com", "v1beta1", "computeTargetHTTPProxy")
	if err != nil {
		return nil, err
	}
	computeTargetHTTPProxyID, err := GetResourceID(computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}
	computeTargetHTTPProxyProjectID, err := GetProjectID(ctx, reader, computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}
	location, err := GetLocation(computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	var external string
	if location == "global" {
		external = fmt.Sprintf("projects/%s/global/targetHttpProxies/%s", computeTargetHTTPProxyProjectID, computeTargetHTTPProxyID)
	} else {
		external = fmt.Sprintf("projects/%s/location/%s/targetHttpProxies/%s", computeTargetHTTPProxyProjectID, location, computeTargetHTTPProxyID)
	}

	return &v1alpha1.ResourceRef{
		External: external,
	}, nil
}

func validateExternalComputeNetwork(ref *v1alpha1.ResourceRef) error {
	tokens := strings.Split(ref.External, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
		return nil
	}
	return fmt.Errorf(
		"format of ComputeNetwork external=%q was not known (use projects/<projectId>/global/networks/<networkid>)",
		ref.External)
}

func validateExternalComputeTargetHTTPProxy(ref *v1alpha1.ResourceRef) error {
	tokens := strings.Split(ref.External, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "targetHttpProxies" {
		return nil
	} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "location" && tokens[4] == "targetHttpProxies" {
		return nil
	}
	return fmt.Errorf(
		"format of ComputeTargetHTTPProxy external=%q was not known "+
			"(use projects/<projectId>/global/targetHttpProxies/<proxyId> or projects/<projectId>/location/<location>/targetHttpProxies/<proxyId>)",
		ref.External)
}
