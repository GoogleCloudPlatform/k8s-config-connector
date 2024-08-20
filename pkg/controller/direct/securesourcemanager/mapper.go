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

func SecureSourceManagerInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: PrivateConfig
	// MISSING: State
	// MISSING: StateNote
	if in.GetKmsKey() != "" {
		out.KmsKeyRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKey()}
	}
	// MISSING: HostConfig
	return out
}
func SecureSourceManagerInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: PrivateConfig
	// MISSING: State
	// MISSING: StateNote
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
	}
	// MISSING: HostConfig
	return out
}

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
