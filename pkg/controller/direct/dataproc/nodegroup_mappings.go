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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DiskConfig_FromProto(mapCtx *direct.MapContext, in *pb.DiskConfig) *krm.DiskConfig {
	if in == nil {
		return nil
	}
	out := &krm.DiskConfig{}
	out.BootDiskType = direct.LazyPtr(in.GetBootDiskType())
	out.BootDiskSizeGB = direct.LazyPtr(in.GetBootDiskSizeGb())
	out.NumLocalSSDs = direct.LazyPtr(in.GetNumLocalSsds())
	out.LocalSSDInterface = direct.LazyPtr(in.GetLocalSsdInterface())
	out.BootDiskProvisionedIOPs = in.BootDiskProvisionedIops
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}
func DiskConfig_ToProto(mapCtx *direct.MapContext, in *krm.DiskConfig) *pb.DiskConfig {
	if in == nil {
		return nil
	}
	out := &pb.DiskConfig{}
	out.BootDiskType = direct.ValueOf(in.BootDiskType)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
	out.NumLocalSsds = direct.ValueOf(in.NumLocalSSDs)
	out.LocalSsdInterface = direct.ValueOf(in.LocalSSDInterface)
	out.BootDiskProvisionedIops = in.BootDiskProvisionedIOPs
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}
