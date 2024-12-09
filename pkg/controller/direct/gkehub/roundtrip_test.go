// Copyright 2024 Google LLC
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

package gkehub

import (
	"encoding/json"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	// "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting" // todo acpana in the future
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
)

// KRM ->ToProto-> proto ->FromProto-> KRM
func FuzzGKEHubFeatureMembership(f *testing.F) {
	// spec
	f.Fuzz(func(t *testing.T, seed int64) {
		filler := fuzz.NewRandomFiller(seed, nil, nil)

		// krm1
		k1 := &krm.GKEHubFeatureMembershipSpec{}

		// fill
		filler.Fill(t, k1)

		// ToProto
		p, err := featureMembershipSpecKRMtoMembershipFeatureSpecAPI(k1)
		if err != nil {
			t.Fatalf("error converting KRM to proto: %v \n KRM: %s", err, prettyPrint(t, k1))
		}

		// krm2 : FromProto
		mapCtx := &direct.MapContext{}
		k2 := GKEHubFeatureMembershipSpec_FromProto(mapCtx, p)
		if mapCtx.Err() != nil {
			t.Fatalf("error mapping from proto to KRM: %v", mapCtx.Err())
		}

		// Using cmpopts.IgnoreFields to ignore specific fields in ProjectRef
		opts := cmp.Options{
			// project ref
			cmpopts.IgnoreFields(krm.FeatureProjectRef{}, "External", "Name", "Namespace"),

			// other refs
			cmpopts.IgnoreFields(refs.IAMServiceAccountRef{}, "Name", "Namespace"),
			cmpopts.IgnoreFields(refs.MetricsGcpServiceAccountRef{}, "Name", "Namespace"),
			cmpopts.IgnoreFields(krm.FeatureRef{}, "External", "Name", "Namespace"),    // todo acpana double check external resolution
			cmpopts.IgnoreFields(krm.MembershipRef{}, "External", "Name", "Namespace"), // todo acpana double check external resolution

			// deprecated fields
			cmpopts.IgnoreFields(krm.FeaturemembershipConfigmanagement{}, "HierarchyController"),
			cmpopts.IgnoreFields(krm.FeaturemembershipConfigmanagement{}, "PolicyController"),

			// unroundtrippable fields (for now)
			cmpopts.IgnoreFields(krm.FeaturemembershipConfigSync{}, "PreventDrift"),                           // todo acpana double check
			cmpopts.IgnoreFields(krm.FeaturemembershipPolicyControllerHubConfig{}, "LogDeniesEnabled"),        // todo acpana double check
			cmpopts.IgnoreFields(krm.FeaturemembershipPolicyControllerHubConfig{}, "ReferentialRulesEnabled"), // todo acpana double check
			cmpopts.IgnoreFields(krm.FeaturemembershipPolicyControllerHubConfig{}, "MutationEnabled"),         // todo acpana double check

			cmpopts.IgnoreFields(krm.GKEHubFeatureMembershipSpec{}, "Location"),           // todo acpana double check that we can't set this
			cmpopts.IgnoreFields(krm.GKEHubFeatureMembershipSpec{}, "MembershipLocation"), // todo acpana double check that we can't set this
		}

		// compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}

func prettyPrint(t *testing.T, k *krm.GKEHubFeatureMembershipSpec) string {
	encoded, err := json.MarshalIndent(k, "", " ")
	if err != nil {
		t.Fatalf("error encoding to json: %v", err)
	}

	return string(encoded)
}
