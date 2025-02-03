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

package documentai

import (
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ProcessorVersionAlias_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersionAlias) *krm.ProcessorVersionAlias {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersionAlias{}
	out.Alias = direct.LazyPtr(in.GetAlias())
	out.ProcessorVersion = direct.LazyPtr(in.GetProcessorVersion())
	return out
}
func ProcessorVersionAlias_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersionAlias) *pb.ProcessorVersionAlias {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersionAlias{}
	out.Alias = direct.ValueOf(in.Alias)
	out.ProcessorVersion = direct.ValueOf(in.ProcessorVersion)
	return out
}
