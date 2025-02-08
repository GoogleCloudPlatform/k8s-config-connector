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

package gsuiteaddons

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gsuiteaddons/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb"
)
func Authorization_FromProto(mapCtx *direct.MapContext, in *pb.Authorization) *krm.Authorization {
	if in == nil {
		return nil
	}
	out := &krm.Authorization{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.OauthClientID = direct.LazyPtr(in.GetOauthClientId())
	return out
}
func Authorization_ToProto(mapCtx *direct.MapContext, in *krm.Authorization) *pb.Authorization {
	if in == nil {
		return nil
	}
	out := &pb.Authorization{}
	out.Name = direct.ValueOf(in.Name)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.OauthClientId = direct.ValueOf(in.OauthClientID)
	return out
}
func GsuiteaddonsAuthorizationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Authorization) *krm.GsuiteaddonsAuthorizationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GsuiteaddonsAuthorizationObservedState{}
	// MISSING: Name
	// MISSING: ServiceAccountEmail
	// MISSING: OauthClientID
	return out
}
func GsuiteaddonsAuthorizationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GsuiteaddonsAuthorizationObservedState) *pb.Authorization {
	if in == nil {
		return nil
	}
	out := &pb.Authorization{}
	// MISSING: Name
	// MISSING: ServiceAccountEmail
	// MISSING: OauthClientID
	return out
}
func GsuiteaddonsAuthorizationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Authorization) *krm.GsuiteaddonsAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.GsuiteaddonsAuthorizationSpec{}
	// MISSING: Name
	// MISSING: ServiceAccountEmail
	// MISSING: OauthClientID
	return out
}
func GsuiteaddonsAuthorizationSpec_ToProto(mapCtx *direct.MapContext, in *krm.GsuiteaddonsAuthorizationSpec) *pb.Authorization {
	if in == nil {
		return nil
	}
	out := &pb.Authorization{}
	// MISSING: Name
	// MISSING: ServiceAccountEmail
	// MISSING: OauthClientID
	return out
}
