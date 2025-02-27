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

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1beta1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VmImage_ImageName_ToProto(mapCtx *direct.MapContext, in *string) *pb.VmImage_ImageName {
	if in == nil {
		return nil
	}
	return &pb.VmImage_ImageName{ImageName: direct.ValueOf(in)}
}

func VmImage_ImageFamily_ToProto(mapCtx *direct.MapContext, in *string) *pb.VmImage_ImageFamily {
	if in == nil {
		return nil
	}
	return &pb.VmImage_ImageFamily{ImageFamily: direct.ValueOf(in)}
}
