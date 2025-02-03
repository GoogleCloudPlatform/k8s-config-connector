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

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Repository_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.Errorf("Repository_CreateTime_FromProto not implemented")
	return nil
}

func Repository_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.Errorf("Repository_CreateTime_ToProto not implemented")
	return nil
}

func Repository_UpdateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.Errorf("Repository_UpdateTime_FromProto not implemented")
	return nil
}

func Repository_UpdateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.Errorf("Repository_UpdateTime_ToProto not implemented")
	return nil
}

func OperationMetadata_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.Errorf("OperationMetadata_CreateTime_FromProto not implemented")
	return nil
}

func OperationMetadata_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.Errorf("OperationMetadata_CreateTime_ToProto not implemented")
	return nil
}

func OperationMetadata_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.Errorf("OperationMetadata_EndTime_FromProto not implemented")
	return nil
}

func OperationMetadata_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.Errorf("OperationMetadata_EndTime_ToProto not implemented")
	return nil
}
func Instance_PrivateConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PrivateConfig) *krm.Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PrivateConfig{}
	out.IsPrivate = direct.LazyPtr(in.GetIsPrivate())
	if in.GetCaPool() != "" {
		out.CaPoolRef = &refs.PrivateCACAPoolRef{External: in.GetCaPool()}
	}
	out.HTTPServiceAttachment = direct.LazyPtr(in.GetHttpServiceAttachment())
	out.SSHServiceAttachment = direct.LazyPtr(in.GetSshServiceAttachment())
	return out
}
func Instance_PrivateConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PrivateConfig) *pb.Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PrivateConfig{}
	out.IsPrivate = direct.ValueOf(in.IsPrivate)
	if in.CaPoolRef != nil {
		out.CaPool = in.CaPoolRef.External
	}
	out.HttpServiceAttachment = direct.ValueOf(in.HTTPServiceAttachment)
	out.SshServiceAttachment = direct.ValueOf(in.SSHServiceAttachment)
	return out
}
func SecureSourceManagerRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.SecureSourceManagerRepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerRepositoryObservedState{}
	out.Uid = direct.LazyPtr(in.Uid)
	out.Etag = direct.LazyPtr(in.Etag)
	out.URIs = Repository_URIs_FromProto(mapCtx, in.GetUris())
	return out
}
func SecureSourceManagerRepositoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerRepositoryObservedState) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.Uid = direct.ValueOf(in.Uid)
	out.Etag = direct.ValueOf(in.Etag)
	out.Uris = Repository_URIs_ToProto(mapCtx, in.URIs)
	return out
}
func SecureSourceManagerRepositorySpec_InstanceRef_FromProto(mapCtx *direct.MapContext, in string) *krm.SecureSourceManagerInstanceRef {
	if in == "" {
		return nil
	}
	return &krm.SecureSourceManagerInstanceRef{External: in}
}
