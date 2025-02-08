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

package talent

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/talent/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/talent/apiv4beta1/talentpb"
)
func Tenant_FromProto(mapCtx *direct.MapContext, in *pb.Tenant) *krm.Tenant {
	if in == nil {
		return nil
	}
	out := &krm.Tenant{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExternalID = direct.LazyPtr(in.GetExternalId())
	out.UsageType = direct.Enum_FromProto(mapCtx, in.GetUsageType())
	out.KeywordSearchableProfileCustomAttributes = in.KeywordSearchableProfileCustomAttributes
	return out
}
func Tenant_ToProto(mapCtx *direct.MapContext, in *krm.Tenant) *pb.Tenant {
	if in == nil {
		return nil
	}
	out := &pb.Tenant{}
	out.Name = direct.ValueOf(in.Name)
	out.ExternalId = direct.ValueOf(in.ExternalID)
	out.UsageType = direct.Enum_ToProto[pb.Tenant_DataUsageType](mapCtx, in.UsageType)
	out.KeywordSearchableProfileCustomAttributes = in.KeywordSearchableProfileCustomAttributes
	return out
}
