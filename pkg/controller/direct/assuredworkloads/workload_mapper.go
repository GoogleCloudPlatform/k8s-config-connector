// Copyright 2025 Google LLC
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

// +generated:mapper
// krm.group: assuredworkloads.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.assuredworkloads.v1

package assuredworkloads

import (
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AssuredWorkloadsWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.AssuredWorkloadsWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.AssuredWorkloadsWorkloadSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ComplianceRegime = direct.Enum_FromProto(mapCtx, in.GetComplianceRegime())
	if in.GetBillingAccount() != "" {
		out.BillingAccountRef = &billingv1alpha1.BillingAccountRef{External: in.GetBillingAccount()}
	}
	// MISSING: Etag
	out.Labels = in.Labels
	if in.GetProvisionedResourcesParent() != "" {
		out.ProvisionedResourcesParent = &refs.FolderRef{External: in.GetProvisionedResourcesParent()}
	}
	out.KMSSettings = Workload_KMSSettings_FromProto(mapCtx, in.GetKmsSettings())
	out.ResourceSettings = direct.Slice_FromProto(mapCtx, in.ResourceSettings, Workload_ResourceSettings_FromProto)
	out.EnableSovereignControls = direct.LazyPtr(in.GetEnableSovereignControls())
	out.Partner = direct.Enum_FromProto(mapCtx, in.GetPartner())
	return out
}

func AssuredWorkloadsWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.AssuredWorkloadsWorkloadSpec) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ComplianceRegime = direct.Enum_ToProto[pb.Workload_ComplianceRegime](mapCtx, in.ComplianceRegime)
	if in.BillingAccountRef != nil {
		out.BillingAccount = in.BillingAccountRef.External
	}
	// MISSING: Etag
	out.Labels = in.Labels
	// ProvisionedResourcesParent is handled manually in the controller
	out.KmsSettings = Workload_KMSSettings_ToProto(mapCtx, in.KMSSettings)
	out.ResourceSettings = direct.Slice_ToProto(mapCtx, in.ResourceSettings, Workload_ResourceSettings_ToProto)
	out.EnableSovereignControls = direct.ValueOf(in.EnableSovereignControls)
	out.Partner = direct.Enum_ToProto[pb.Workload_Partner](mapCtx, in.Partner)
	return out
}

func Workload_SaaEnrollmentResponse_FromProto(mapCtx *direct.MapContext, in *pb.Workload_SaaEnrollmentResponse) *krm.Workload_SaaEnrollmentResponse {
	if in == nil {
		return nil
	}
	out := &krm.Workload_SaaEnrollmentResponse{}
	out.SetupErrors = direct.EnumSlice_FromProto(mapCtx, in.SetupErrors)
	if in.SetupStatus != nil {
		out.SetupStatus = direct.ZeroBasedEnum_FromProto(mapCtx, in.GetSetupStatus())
	}
	return out
}

func Workload_SaaEnrollmentResponse_ToProto(mapCtx *direct.MapContext, in *krm.Workload_SaaEnrollmentResponse) *pb.Workload_SaaEnrollmentResponse {
	if in == nil {
		return nil
	}
	out := &pb.Workload_SaaEnrollmentResponse{}
	out.SetupErrors = direct.EnumSlice_ToProto[pb.Workload_SaaEnrollmentResponse_SetupError](mapCtx, in.SetupErrors)
	if in.SetupStatus != nil {
		out.SetupStatus = direct.PtrTo(direct.Enum_ToProto[pb.Workload_SaaEnrollmentResponse_SetupState](mapCtx, in.SetupStatus))
	}
	return out
}
