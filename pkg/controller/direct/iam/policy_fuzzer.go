// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iam

import (
	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(IAMPolicyFuzzer())
}

func IAMPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&iampb.Policy{},
		IAMPolicySpec_FromProto, IAMPolicySpec_ToProto, IAMPolicyObservedState_FromProto, IAMPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".bindings")
	f.SpecFields.Insert(".audit_configs")

	f.UnimplementedFields.Insert(".etag")                          // spec field set at controller level instead of mapper ?
	f.UnimplementedFields.Insert(".bindings[].condition.location") // not yet supported in the KRM representation
	f.UnimplementedFields.Insert(".version")                       // status field ?

	return f
}
