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

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.AddResourcePoliciesDiskRequest
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeDiskResourcePolicyAttachmentFuzzer())
}

func computeDiskResourcePolicyAttachmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.AddResourcePoliciesDiskRequest{},
		ComputeDiskResourcePolicyAttachmentSpec_FromProto, ComputeDiskResourcePolicyAttachmentSpec_ToProto,
	)

	// Field comparison of KRM ComputeDiskResourcePolicyAttachmentSpec to proto fields:
	// - diskRef                               => .disk (SpecField)
	// - projectRef                            => .project (SpecField)
	// - zone                                  => .zone (SpecField)
	// - resourceID                            => .disks_add_resource_policies_request_resource.resource_policies[0] (SpecField)

	// Spec fields
	f.SpecField(".disk")
	f.SpecField(".project")
	f.SpecField(".zone")
	f.SpecField(".disks_add_resource_policies_request_resource")
	f.SpecField(".disks_add_resource_policies_request_resource.resource_policies")

	// Unimplemented / Request-scoped fields
	f.Unimplemented_NotYetTriaged(".request_id")

	f.FilterSpec = func(in *pb.AddResourcePoliciesDiskRequest) {
		resProp := in.DisksAddResourcePoliciesRequestResource
		if resProp != nil {
			if len(resProp.ResourcePolicies) == 0 || resProp.ResourcePolicies[0] == "" {
				in.DisksAddResourcePoliciesRequestResource = nil
			} else if len(resProp.ResourcePolicies) > 1 {
				resProp.ResourcePolicies = resProp.ResourcePolicies[:1]
			}
		}
	}

	return f
}

func ComputeDiskResourcePolicyAttachmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.AddResourcePoliciesDiskRequest) *computev1alpha1.ComputeDiskResourcePolicyAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &computev1alpha1.ComputeDiskResourcePolicyAttachmentSpec{}
	out.Zone = in.GetZone()
	if in.GetProject() != "" {
		out.ProjectRef = refs.ProjectRef{
			External: in.GetProject(),
		}
	}
	if in.GetDisk() != "" {
		out.DiskRef = computev1beta1.ComputeDiskRef{
			External: in.GetDisk(),
		}
	}
	resProp := in.GetDisksAddResourcePoliciesRequestResource()
	if resProp != nil && len(resProp.ResourcePolicies) > 0 && resProp.ResourcePolicies[0] != "" {
		out.ResourceID = direct.PtrTo(resProp.ResourcePolicies[0])
	}
	return out
}

func ComputeDiskResourcePolicyAttachmentSpec_ToProto(mapCtx *direct.MapContext, in *computev1alpha1.ComputeDiskResourcePolicyAttachmentSpec) *pb.AddResourcePoliciesDiskRequest {
	if in == nil {
		return nil
	}
	out := &pb.AddResourcePoliciesDiskRequest{}
	out.Zone = in.Zone
	if in.ProjectRef.External != "" {
		out.Project = in.ProjectRef.External
	}
	if in.DiskRef.External != "" {
		out.Disk = in.DiskRef.External
	}
	if in.ResourceID != nil && *in.ResourceID != "" {
		out.DisksAddResourcePoliciesRequestResource = &pb.DisksAddResourcePoliciesRequest{
			ResourcePolicies: []string{*in.ResourceID},
		}
	}
	return out
}
