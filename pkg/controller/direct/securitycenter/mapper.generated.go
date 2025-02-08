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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1p1beta1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
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
	out.IamPolicy = Asset_IamPolicy_FromProto(mapCtx, in.GetIamPolicy())
	out.CanonicalName = direct.LazyPtr(in.GetCanonicalName())
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
	out.IamPolicy = Asset_IamPolicy_ToProto(mapCtx, in.IamPolicy)
	out.CanonicalName = direct.ValueOf(in.CanonicalName)
	return out
}
func Asset_IamPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Asset_IamPolicy) *krm.Asset_IamPolicy {
	if in == nil {
		return nil
	}
	out := &krm.Asset_IamPolicy{}
	out.PolicyBlob = direct.LazyPtr(in.GetPolicyBlob())
	return out
}
func Asset_IamPolicy_ToProto(mapCtx *direct.MapContext, in *krm.Asset_IamPolicy) *pb.Asset_IamPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Asset_IamPolicy{}
	out.PolicyBlob = direct.ValueOf(in.PolicyBlob)
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
	out.ResourceDisplayName = direct.LazyPtr(in.GetResourceDisplayName())
	out.ResourceParentDisplayName = direct.LazyPtr(in.GetResourceParentDisplayName())
	out.ResourceProjectDisplayName = direct.LazyPtr(in.GetResourceProjectDisplayName())
	out.Folders = direct.Slice_FromProto(mapCtx, in.Folders, Folder_FromProto)
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
	out.ResourceDisplayName = direct.ValueOf(in.ResourceDisplayName)
	out.ResourceParentDisplayName = direct.ValueOf(in.ResourceParentDisplayName)
	out.ResourceProjectDisplayName = direct.ValueOf(in.ResourceProjectDisplayName)
	out.Folders = direct.Slice_ToProto(mapCtx, in.Folders, Folder_ToProto)
	return out
}
func Folder_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.Folder {
	if in == nil {
		return nil
	}
	out := &krm.Folder{}
	out.ResourceFolder = direct.LazyPtr(in.GetResourceFolder())
	out.ResourceFolderDisplayName = direct.LazyPtr(in.GetResourceFolderDisplayName())
	return out
}
func Folder_ToProto(mapCtx *direct.MapContext, in *krm.Folder) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	out.ResourceFolder = direct.ValueOf(in.ResourceFolder)
	out.ResourceFolderDisplayName = direct.ValueOf(in.ResourceFolderDisplayName)
	return out
}
func SecurityMarks_FromProto(mapCtx *direct.MapContext, in *pb.SecurityMarks) *krm.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &krm.SecurityMarks{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Marks = in.Marks
	out.CanonicalName = direct.LazyPtr(in.GetCanonicalName())
	return out
}
func SecurityMarks_ToProto(mapCtx *direct.MapContext, in *krm.SecurityMarks) *pb.SecurityMarks {
	if in == nil {
		return nil
	}
	out := &pb.SecurityMarks{}
	out.Name = direct.ValueOf(in.Name)
	out.Marks = in.Marks
	out.CanonicalName = direct.ValueOf(in.CanonicalName)
	return out
}
