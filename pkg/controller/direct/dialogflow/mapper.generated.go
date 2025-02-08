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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
)
func Changelog_FromProto(mapCtx *direct.MapContext, in *pb.Changelog) *krm.Changelog {
	if in == nil {
		return nil
	}
	out := &krm.Changelog{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UserEmail = direct.LazyPtr(in.GetUserEmail())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Action = direct.LazyPtr(in.GetAction())
	out.Type = direct.LazyPtr(in.GetType())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func Changelog_ToProto(mapCtx *direct.MapContext, in *krm.Changelog) *pb.Changelog {
	if in == nil {
		return nil
	}
	out := &pb.Changelog{}
	out.Name = direct.ValueOf(in.Name)
	out.UserEmail = direct.ValueOf(in.UserEmail)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Action = direct.ValueOf(in.Action)
	out.Type = direct.ValueOf(in.Type)
	out.Resource = direct.ValueOf(in.Resource)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
