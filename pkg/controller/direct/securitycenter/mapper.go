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

package securitycenter

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/type/expr"
)

func Expr_FromProto(mapCtx *direct.MapContext, in *expr.Expr) *krm.Expr {
	if in == nil {
		return nil
	}
	return &krm.Expr{
		Expression:  direct.LazyPtr(in.GetExpression()),
		Title:       direct.LazyPtr(in.GetTitle()),
		Description: direct.LazyPtr(in.GetDescription()),
		Location:    direct.LazyPtr(in.GetLocation()),
	}
}

func Expr_ToProto(mapCtx *direct.MapContext, in *krm.Expr) *expr.Expr {
	if in == nil {
		return nil
	}
	return &expr.Expr{
		Expression:  direct.ValueOf(in.Expression),
		Title:       direct.ValueOf(in.Title),
		Description: direct.ValueOf(in.Description),
		Location:    direct.ValueOf(in.Location),
	}
}
