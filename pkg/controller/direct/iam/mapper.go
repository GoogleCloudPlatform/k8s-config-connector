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

package iam

import (
	pb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	iampb "cloud.google.com/go/iam/apiv2/iampb"
	krmiamv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	krmiamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	exprpb "google.golang.org/genproto/googleapis/type/expr"
)

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func Expr_v1alpha1_FromProto(mapCtx *direct.MapContext, in *exprpb.Expr) *krmv1alpha1.Expr {
	return Expr_FromProto(mapCtx, in)
}

func Expr_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Expr) *exprpb.Expr {
	return Expr_ToProto(mapCtx, in)
}

func IAMServiceAccountKeySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmiamv1beta1.IAMServiceAccountKeySpec) *pb.ServiceAccountKey {
	return IAMServiceAccountKeySpec_ToProto(mapCtx, in)
}

func IAMServiceAccountKeySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccountKey) *krmiamv1beta1.IAMServiceAccountKeySpec {
	return IAMServiceAccountKeySpec_FromProto(mapCtx, in)
}

func DenyRule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmiamv1alpha1.DenyRule) *iampb.DenyRule {
	return DenyRule_ToProto(mapCtx, in)
}

func DenyRule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *iampb.DenyRule) *krmiamv1alpha1.DenyRule {
	return DenyRule_FromProto(mapCtx, in)
}
