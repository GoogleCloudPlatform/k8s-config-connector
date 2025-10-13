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

package iam

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/bindings"
)

type MemberIdentityResolver = bindings.MemberIdentityResolver

func ComputePartialPolicyWithMergedBindings(partialPolicy *v1beta1.IAMPartialPolicy, livePolicy *v1beta1.IAMPolicy, resolver MemberIdentityResolver) (*v1beta1.IAMPartialPolicy, error) {
	return bindings.ComputePartialPolicyWithMergedBindings(partialPolicy, livePolicy, resolver)
}

func ComputePartialPolicyWithRemainingBindings(partialPolicy *v1beta1.IAMPartialPolicy, livePolicy *v1beta1.IAMPolicy) *v1beta1.IAMPartialPolicy {
	return bindings.ComputePartialPolicyWithRemainingBindings(partialPolicy, livePolicy)
}

func ConvertIAMPartialBindingsToIAMPolicyBindings(partialPolicy *v1beta1.IAMPartialPolicy, resolver MemberIdentityResolver) (bindings []v1beta1.IAMPolicyBinding, err error) {
	return bindings.ConvertIAMPartialBindingsToIAMPolicyBindings(partialPolicy, resolver)
}