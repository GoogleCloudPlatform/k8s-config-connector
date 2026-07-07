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

package cloudsecuritycompliance

import (
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudControlDetails_FromProto(mapCtx *direct.MapContext, in *pb.CloudControlDetails) *krm.CloudControlDetails {
	if in == nil {
		return nil
	}
	out := &krm.CloudControlDetails{}
	if in.GetName() != "" {
		out.CloudControlRef = &krm.CloudSecurityComplianceCloudControlRef{External: in.GetName()}
	}
	out.MajorRevisionID = direct.LazyPtr(in.GetMajorRevisionId())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Parameter_FromProto)
	return out
}

func CloudControlDetails_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlDetails) *pb.CloudControlDetails {
	if in == nil {
		return nil
	}
	out := &pb.CloudControlDetails{}
	if in.CloudControlRef != nil {
		out.Name = in.CloudControlRef.External
	}
	out.MajorRevisionId = direct.ValueOf(in.MajorRevisionID)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Parameter_ToProto)
	return out
}
