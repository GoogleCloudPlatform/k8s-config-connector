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

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDBUserSpec_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func AlloyDBUserSpec_Password_FromProto(mapCtx *direct.MapContext, in string) *refsv1beta1secret.Legacy {
	return nil
}

func AlloyDBUserSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBUserSpec) *pb.User {
	if in == nil {
		return nil
	}
	out := &pb.User{}
	out.Password = AlloyDBUserSpec_Password_ToProto(mapCtx, in.Password)
	out.DatabaseRoles = in.DatabaseRoles
	out.UserType = direct.Enum_ToProto[pb.User_UserType](mapCtx, in.UserType)
	return out
}

func AlloyDBUserSpec_FromProto(mapCtx *direct.MapContext, in *pb.User) *krm.AlloyDBUserSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBUserSpec{}
	out.Password = AlloyDBUserSpec_Password_FromProto(mapCtx, in.GetPassword())
	out.DatabaseRoles = in.DatabaseRoles
	out.UserType = direct.Enum_FromProto(mapCtx, in.GetUserType())
	return out
}
