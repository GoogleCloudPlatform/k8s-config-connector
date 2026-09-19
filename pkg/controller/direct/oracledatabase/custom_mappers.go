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

package oracledatabase

import (
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
)

func TimeZone_FromProto(mapCtx *direct.MapContext, in *datetime.TimeZone) *krm.TimeZone {
	if in == nil {
		return nil
	}
	out := &krm.TimeZone{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

func TimeZone_ToProto(mapCtx *direct.MapContext, in *krm.TimeZone) *datetime.TimeZone {
	if in == nil {
		return nil
	}
	out := &datetime.TimeZone{}
	out.Id = direct.ValueOf(in.ID)
	out.Version = direct.ValueOf(in.Version)
	return out
}

func ExadbVMClusterStorageDetails_FromProto(mapCtx *direct.MapContext, in *pb.ExadbVmClusterStorageDetails) *krm.ExadbVMClusterStorageDetails {
	if in == nil {
		return nil
	}
	out := &krm.ExadbVMClusterStorageDetails{}
	out.SizeInGBsPerNode = direct.LazyPtr(in.GetSizeInGbsPerNode())
	return out
}

func ExadbVMClusterStorageDetails_ToProto(mapCtx *direct.MapContext, in *krm.ExadbVMClusterStorageDetails) *pb.ExadbVmClusterStorageDetails {
	if in == nil {
		return nil
	}
	out := &pb.ExadbVmClusterStorageDetails{}
	out.SizeInGbsPerNode = direct.ValueOf(in.SizeInGBsPerNode)
	return out
}
