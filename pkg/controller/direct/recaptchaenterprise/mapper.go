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

package recaptchaenterprise

import (
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KeyTestingOptions_FromProto(mapCtx *direct.MapContext, in *pb.TestingOptions) *krm.KeyTestingOptions {
	if in == nil {
		return nil
	}
	out := &krm.KeyTestingOptions{}
	out.TestingChallenge = direct.Enum_FromProto(mapCtx, in.GetTestingChallenge())
	if in.GetTestingScore() != 0 {
		v := float64(in.GetTestingScore())
		out.TestingScore = &v
	}
	return out
}

func KeyTestingOptions_ToProto(mapCtx *direct.MapContext, in *krm.KeyTestingOptions) *pb.TestingOptions {
	if in == nil {
		return nil
	}
	out := &pb.TestingOptions{}
	out.TestingChallenge = direct.Enum_ToProto[pb.TestingOptions_TestingChallenge](mapCtx, in.TestingChallenge)
	if in.TestingScore != nil {
		out.TestingScore = float32(*in.TestingScore)
	}
	return out
}
