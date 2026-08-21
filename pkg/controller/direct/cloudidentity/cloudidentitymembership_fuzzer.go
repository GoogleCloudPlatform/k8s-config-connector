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

package cloudidentity

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/cloudidentity/v1beta1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(membershipFuzzer())
}

func membershipFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.Membership{},
		CloudIdentityMembershipSpec_FromAPI, CloudIdentityMembershipSpec_ToAPI,
		CloudIdentityMembershipStatus_FromAPI, CloudIdentityMembershipStatus_ToAPI,
	)

	f.SpecField(".MemberKey")
	f.SpecField(".PreferredMemberKey")
	f.SpecField(".Roles")

	f.StatusField(".CreateTime")
	f.StatusField(".DeliverySetting")
	f.StatusField(".Type")
	f.StatusField(".UpdateTime")

	f.IdentityField(".Name")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	f.Ignore_JSONBookkeeping(".MemberKey.ForceSendFields")
	f.Ignore_JSONBookkeeping(".MemberKey.NullFields")
	f.Ignore_JSONBookkeeping(".PreferredMemberKey.ForceSendFields")
	f.Ignore_JSONBookkeeping(".PreferredMemberKey.NullFields")
	f.Ignore_JSONBookkeeping(".Roles.ForceSendFields")
	f.Ignore_JSONBookkeeping(".Roles.NullFields")
	f.Ignore_JSONBookkeeping(".Roles.ExpiryDetail.ForceSendFields")
	f.Ignore_JSONBookkeeping(".Roles.ExpiryDetail.NullFields")
	f.Ignore_JSONBookkeeping(".Roles.RestrictionEvaluations.ForceSendFields")
	f.Ignore_JSONBookkeeping(".Roles.RestrictionEvaluations.NullFields")
	f.Ignore_JSONBookkeeping(".Roles.RestrictionEvaluations.MemberRestrictionEvaluation.ForceSendFields")
	f.Ignore_JSONBookkeeping(".Roles.RestrictionEvaluations.MemberRestrictionEvaluation.NullFields")
	f.Ignore_JSONBookkeeping(".ExpiryDetail.ForceSendFields")
	f.Ignore_JSONBookkeeping(".ExpiryDetail.NullFields")
	f.Ignore_JSONBookkeeping(".RestrictionEvaluations.ForceSendFields")
	f.Ignore_JSONBookkeeping(".RestrictionEvaluations.NullFields")
	f.Ignore_JSONBookkeeping(".RestrictionEvaluations.MemberRestrictionEvaluation.ForceSendFields")
	f.Ignore_JSONBookkeeping(".RestrictionEvaluations.MemberRestrictionEvaluation.NullFields")
	f.Ignore_JSONBookkeeping(".RestrictionEvaluations.MemberRestrictionEvaluation.State.ForceSendFields")
	f.Ignore_JSONBookkeeping(".RestrictionEvaluations.MemberRestrictionEvaluation.State.NullFields")

	return f
}
