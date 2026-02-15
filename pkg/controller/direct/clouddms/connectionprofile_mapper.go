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

// +generated:mapper
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1

package clouddms

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
)

func Status_ToProto(mapCtx *direct.MapContext, in *krm.Status) *statuspb.Status {
	if in == nil {
		return nil
	}
	out := &statuspb.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func Status_FromProto(mapCtx *direct.MapContext, in *statuspb.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
