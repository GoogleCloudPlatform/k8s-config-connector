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

package designcenter

import (
	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	saasservicemgmtv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/saasservicemgmt/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ApplicationTemplate_SaaSRuntimeContext_FromProto(mapCtx *direct.MapContext, in *pb.SaaSRuntimeContext) *krm.ApplicationTemplate_SaaSRuntimeContext {
	if in == nil {
		return nil
	}
	out := &krm.ApplicationTemplate_SaaSRuntimeContext{}
	for _, name := range in.GetSaasNames() {
		out.SaasRefs = append(out.SaasRefs, saasservicemgmtv1alpha1.SaaSServiceMgmtSaaSRef{External: name})
	}
	return out
}

func ApplicationTemplate_SaaSRuntimeContext_ToProto(mapCtx *direct.MapContext, in *krm.ApplicationTemplate_SaaSRuntimeContext) *pb.SaaSRuntimeContext {
	if in == nil {
		return nil
	}
	out := &pb.SaaSRuntimeContext{}
	for _, ref := range in.SaasRefs {
		if ref.External != "" {
			out.SaasNames = append(out.SaasNames, ref.External)
		}
	}
	return out
}
