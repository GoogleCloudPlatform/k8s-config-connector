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

package recaptchaenterprise

import (
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ReCAPTCHAEnterpriseFirewallPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState {
	if in == nil {
		return nil
	}
	return &krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState{}
}

func ReCAPTCHAEnterpriseFirewallPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	return &pb.FirewallPolicy{}
}
