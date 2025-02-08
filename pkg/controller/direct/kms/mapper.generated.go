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

package kms

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/kms/inventory/apiv1/inventorypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func KmsProtectedResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProtectedResource) *krm.KmsProtectedResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KmsProtectedResourceObservedState{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: ProjectID
	// MISSING: CloudProduct
	// MISSING: ResourceType
	// MISSING: Location
	// MISSING: Labels
	// MISSING: CryptoKeyVersion
	// MISSING: CryptoKeyVersions
	// MISSING: CreateTime
	return out
}
func KmsProtectedResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KmsProtectedResourceObservedState) *pb.ProtectedResource {
	if in == nil {
		return nil
	}
	out := &pb.ProtectedResource{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: ProjectID
	// MISSING: CloudProduct
	// MISSING: ResourceType
	// MISSING: Location
	// MISSING: Labels
	// MISSING: CryptoKeyVersion
	// MISSING: CryptoKeyVersions
	// MISSING: CreateTime
	return out
}
func KmsProtectedResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProtectedResource) *krm.KmsProtectedResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.KmsProtectedResourceSpec{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: ProjectID
	// MISSING: CloudProduct
	// MISSING: ResourceType
	// MISSING: Location
	// MISSING: Labels
	// MISSING: CryptoKeyVersion
	// MISSING: CryptoKeyVersions
	// MISSING: CreateTime
	return out
}
func KmsProtectedResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.KmsProtectedResourceSpec) *pb.ProtectedResource {
	if in == nil {
		return nil
	}
	out := &pb.ProtectedResource{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: ProjectID
	// MISSING: CloudProduct
	// MISSING: ResourceType
	// MISSING: Location
	// MISSING: Labels
	// MISSING: CryptoKeyVersion
	// MISSING: CryptoKeyVersions
	// MISSING: CreateTime
	return out
}
func ProtectedResource_FromProto(mapCtx *direct.MapContext, in *pb.ProtectedResource) *krm.ProtectedResource {
	if in == nil {
		return nil
	}
	out := &krm.ProtectedResource{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Project = direct.LazyPtr(in.GetProject())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.CloudProduct = direct.LazyPtr(in.GetCloudProduct())
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Labels = in.Labels
	out.CryptoKeyVersion = direct.LazyPtr(in.GetCryptoKeyVersion())
	out.CryptoKeyVersions = in.CryptoKeyVersions
	// MISSING: CreateTime
	return out
}
func ProtectedResource_ToProto(mapCtx *direct.MapContext, in *krm.ProtectedResource) *pb.ProtectedResource {
	if in == nil {
		return nil
	}
	out := &pb.ProtectedResource{}
	out.Name = direct.ValueOf(in.Name)
	out.Project = direct.ValueOf(in.Project)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.CloudProduct = direct.ValueOf(in.CloudProduct)
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Location = direct.ValueOf(in.Location)
	out.Labels = in.Labels
	out.CryptoKeyVersion = direct.ValueOf(in.CryptoKeyVersion)
	out.CryptoKeyVersions = in.CryptoKeyVersions
	// MISSING: CreateTime
	return out
}
func ProtectedResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProtectedResource) *krm.ProtectedResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProtectedResourceObservedState{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: ProjectID
	// MISSING: CloudProduct
	// MISSING: ResourceType
	// MISSING: Location
	// MISSING: Labels
	// MISSING: CryptoKeyVersion
	// MISSING: CryptoKeyVersions
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func ProtectedResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProtectedResourceObservedState) *pb.ProtectedResource {
	if in == nil {
		return nil
	}
	out := &pb.ProtectedResource{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: ProjectID
	// MISSING: CloudProduct
	// MISSING: ResourceType
	// MISSING: Location
	// MISSING: Labels
	// MISSING: CryptoKeyVersion
	// MISSING: CryptoKeyVersions
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
