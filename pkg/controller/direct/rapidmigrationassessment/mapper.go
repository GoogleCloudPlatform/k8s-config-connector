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

package rapidmigrationassessment

import (
	pb "cloud.google.com/go/rapidmigrationassessment/apiv1/rapidmigrationassessmentpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/rapidmigrationassessment/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GuestOSScan_FromProto(mapCtx *direct.MapContext, in *pb.GuestOsScan) *krm.GuestOSScan {
	if in == nil {
		return nil
	}
	out := &krm.GuestOSScan{}
	out.CoreSource = direct.LazyPtr(in.GetCoreSource())
	return out
}

func GuestOSScan_ToProto(mapCtx *direct.MapContext, in *krm.GuestOSScan) *pb.GuestOsScan {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsScan{}
	out.CoreSource = direct.ValueOf(in.CoreSource)
	return out
}

func VSphereScan_FromProto(mapCtx *direct.MapContext, in *pb.VSphereScan) *krm.VSphereScan {
	if in == nil {
		return nil
	}
	out := &krm.VSphereScan{}
	out.CoreSource = direct.LazyPtr(in.GetCoreSource())
	return out
}

func VSphereScan_ToProto(mapCtx *direct.MapContext, in *krm.VSphereScan) *pb.VSphereScan {
	if in == nil {
		return nil
	}
	out := &pb.VSphereScan{}
	out.CoreSource = direct.ValueOf(in.CoreSource)
	return out
}
