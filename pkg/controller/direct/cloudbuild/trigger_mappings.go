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

package cloudbuild

import (
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Hash_FromProto(mapCtx *direct.MapContext, in *pb.Hash) *krm.Hash {
	if in == nil {
		return nil
	}
	out := &krm.Hash{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Value = in.GetValue()
	return out
}
func Hash_ToProto(mapCtx *direct.MapContext, in *krm.Hash) *pb.Hash {
	if in == nil {
		return nil
	}
	out := &pb.Hash{}
	out.Type = direct.Enum_ToProto[pb.Hash_HashType](mapCtx, in.Type)
	out.Value = in.Value
	return out
}
