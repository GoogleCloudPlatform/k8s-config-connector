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

package dialogflow

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ContinuousTestResult_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousTestResult) *krm.ContinuousTestResult {
	if in == nil {
		return nil
	}
	out := &krm.ContinuousTestResult{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Result = direct.Enum_FromProto(mapCtx, in.GetResult())
	out.TestCaseResults = in.TestCaseResults
	out.RunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRunTime())
	return out
}
func ContinuousTestResult_ToProto(mapCtx *direct.MapContext, in *krm.ContinuousTestResult) *pb.ContinuousTestResult {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousTestResult{}
	out.Name = direct.ValueOf(in.Name)
	out.Result = direct.Enum_ToProto[pb.ContinuousTestResult_AggregatedTestResult](mapCtx, in.Result)
	out.TestCaseResults = in.TestCaseResults
	out.RunTime = direct.StringTimestamp_ToProto(mapCtx, in.RunTime)
	return out
}
