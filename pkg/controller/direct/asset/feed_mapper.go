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

package asset

import (
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	exprpb "google.golang.org/genproto/googleapis/type/expr"
)

func Expr_FromProto(mapCtx *direct.MapContext, in *exprpb.Expr) *krm.Expr {
	if in == nil {
		return nil
	}
	out := &krm.Expr{}
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}

func Expr_ToProto(mapCtx *direct.MapContext, in *krm.Expr) *exprpb.Expr {
	if in == nil {
		return nil
	}
	out := &exprpb.Expr{}
	out.Expression = direct.ValueOf(in.Expression)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Location = direct.ValueOf(in.Location)
	return out
}

func PubsubDestination_FromProto(mapCtx *direct.MapContext, in *pb.PubsubDestination) *krm.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &krm.PubsubDestination{}
	out.TopicRef = &refsv1beta1.PubSubTopicRef{External: in.GetTopic()}
	return out
}
func PubsubDestination_ToProto(mapCtx *direct.MapContext, in *krm.PubsubDestination) *pb.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &pb.PubsubDestination{}
	out.Topic = in.TopicRef.External
	return out
}
