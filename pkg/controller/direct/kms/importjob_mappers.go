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
	"strings"

	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KMSImportJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.ImportJob) *krm.KMSImportJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSImportJobSpec{}
	// MISSING: Name
	out.ImportMethod = direct.Enum_FromProto(mapCtx, in.GetImportMethod())
	out.ProtectionLevel = direct.Enum_FromProto(mapCtx, in.GetProtectionLevel())
	return out
}
func KMSImportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSImportJobSpec) *pb.ImportJob {
	if in == nil {
		return nil
	}
	out := &pb.ImportJob{}
	// MISSING: Name

	importMethodUpper := strings.ToUpper(*in.ImportMethod)
	out.ImportMethod = direct.Enum_ToProto[pb.ImportJob_ImportMethod](mapCtx, &importMethodUpper)
	protectionLevelUpper := strings.ToUpper(*in.ProtectionLevel)
	out.ProtectionLevel = direct.Enum_ToProto[pb.ProtectionLevel](mapCtx, &protectionLevelUpper)
	return out
}
