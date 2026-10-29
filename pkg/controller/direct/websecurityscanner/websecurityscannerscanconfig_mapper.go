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

package websecurityscanner

import (
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/websecurityscanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ScanConfig_Authentication_CustomAccount_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig_Authentication_CustomAccount) *krm.ScanConfig_Authentication_CustomAccount {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfig_Authentication_CustomAccount{}
	out.Username = direct.LazyPtr(in.GetUsername())
	if in.GetPassword() != "" {
		out.Password = &refsv1beta1secret.Legacy{
			Value: direct.LazyPtr(in.GetPassword()),
		}
	}
	out.LoginURL = direct.LazyPtr(in.GetLoginUrl())
	return out
}

func ScanConfig_Authentication_CustomAccount_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfig_Authentication_CustomAccount) *pb.ScanConfig_Authentication_CustomAccount {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig_Authentication_CustomAccount{}
	out.Username = direct.ValueOf(in.Username)
	if in.Password != nil {
		out.Password = direct.ValueOf(in.Password.Value)
	}
	out.LoginUrl = direct.ValueOf(in.LoginURL)
	return out
}

func ScanConfig_Authentication_GoogleAccount_FromProto(mapCtx *direct.MapContext, in *pb.ScanConfig_Authentication_GoogleAccount) *krm.ScanConfig_Authentication_GoogleAccount {
	if in == nil {
		return nil
	}
	out := &krm.ScanConfig_Authentication_GoogleAccount{}
	out.Username = direct.LazyPtr(in.GetUsername())
	if in.GetPassword() != "" {
		out.Password = &refsv1beta1secret.Legacy{
			Value: direct.LazyPtr(in.GetPassword()),
		}
	}
	return out
}

func ScanConfig_Authentication_GoogleAccount_ToProto(mapCtx *direct.MapContext, in *krm.ScanConfig_Authentication_GoogleAccount) *pb.ScanConfig_Authentication_GoogleAccount {
	if in == nil {
		return nil
	}
	out := &pb.ScanConfig_Authentication_GoogleAccount{}
	out.Username = direct.ValueOf(in.Username)
	if in.Password != nil {
		out.Password = direct.ValueOf(in.Password.Value)
	}
	return out
}
