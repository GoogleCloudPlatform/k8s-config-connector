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

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch_RegexMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.AuthorizationPolicy_Rule_Destination_HttpHeaderMatch_RegexMatch {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationPolicy_Rule_Destination_HttpHeaderMatch_RegexMatch{}
	out.RegexMatch = direct.ValueOf(in)
	return out
}
