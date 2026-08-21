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

package iap

import (
	pb "cloud.google.com/go/iap/apiv1/iappb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iap/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func IAPBrandStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Brand) *krm.IAPBrandStatus {
	if in == nil {
		return nil
	}
	out := &krm.IAPBrandStatus{}
	out.OrgInternalOnly = direct.LazyPtr(in.GetOrgInternalOnly())
	return out
}

func IAPBrandStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.IAPBrandStatus) *pb.Brand {
	if in == nil {
		return nil
	}
	out := &pb.Brand{}
	out.OrgInternalOnly = direct.ValueOf(in.OrgInternalOnly)
	return out
}
