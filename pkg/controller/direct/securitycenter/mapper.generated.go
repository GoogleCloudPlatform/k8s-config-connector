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

package securitycenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1beta1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Asset_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.Asset {
	if in == nil {
		return nil
	}
	out := &krm.Asset{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SecurityCenterProperties = Asset_SecurityCenterProperties_FromProto(mapCtx, in.GetSecurityCenterProperties())
	// MISSING: ResourceProperties
	out.SecurityMarks = SecurityMarks_FromProto(mapCtx, in.GetSecurityMarks())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Asset_ToProto(mapCtx *direct.MapContext, in *krm.Asset) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	out.Name = direct.ValueOf(in.Name)
	out.SecurityCenterProperties = Asset_SecurityCenterProperties_ToProto(mapCtx, in.SecurityCenterProperties)
	// MISSING: ResourceProperties
	out.SecurityMarks = SecurityMarks_ToProto(mapCtx, in.SecurityMarks)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Asset_SecurityCenterProperties_FromProto(mapCtx *direct.MapContext, in *pb.Asset_SecurityCenterProperties) *krm.Asset_SecurityCenterProperties {
	if in == nil {
		return nil
	}
	out := &krm.Asset_SecurityCenterProperties{}
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.ResourceParent = direct.LazyPtr(in.GetResourceParent())
	out.ResourceProject = direct.LazyPtr(in.GetResourceProject())
	out.ResourceOwners = in.ResourceOwners
	return out
}
func Asset_SecurityCenterProperties_ToProto(mapCtx *direct.MapContext, in *krm.Asset_SecurityCenterProperties) *pb.Asset_SecurityCenterProperties {
	if in == nil {
		return nil
	}
	out := &pb.Asset_SecurityCenterProperties{}
	out.ResourceName = direct.ValueOf(in.ResourceName)
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.ResourceParent = direct.ValueOf(in.ResourceParent)
	out.ResourceProject = direct.ValueOf(in.ResourceProject)
	out.ResourceOwners = in.ResourceOwners
	return out
}
func SecurityMarks_FromProto(mapCtx *direct.MapContext, in *pb.SecurityMarks) *krm.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &krm.SecurityMarks{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Marks = in.Marks
	return out
}
func SecurityMarks_ToProto(mapCtx *direct.MapContext, in *krm.SecurityMarks) *pb.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &pb.SecurityMarks{}
	out.Name = direct.ValueOf(in.Name)
	out.Marks = in.Marks
	return out
}
