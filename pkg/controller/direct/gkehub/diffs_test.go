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
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
)

var (
	dummyString1 = "fakeString1"
	dummyString2 = "fakeString2"
	dummyBool1   = true
	dummyBool2   = false
)

func Test_ListFieldDiffs(t *testing.T) {
	testCases := []struct {
		name          string
		left          krm.GKEHubFeatureMembershipSpec
		right         krm.GKEHubFeatureMembershipSpec
		expectedDiffs []diff
		expectedError bool
	}{
		{
			name: "left is a subset of right, no diff",
			left: krm.GKEHubFeatureMembershipSpec{
				Configmanagement: &krm.FeaturemembershipConfigmanagement{
					ConfigSync: &krm.FeaturemembershipConfigSync{
						Git: &krm.FeaturemembershipGit{
							HttpsProxy: &dummyString1,
							SyncRev:    &dummyString1,
						},
						PreventDrift: &dummyBool1,
					},
					Version: &dummyString1,
				},
			},
			right: krm.GKEHubFeatureMembershipSpec{
				Configmanagement: &krm.FeaturemembershipConfigmanagement{
					ConfigSync: &krm.FeaturemembershipConfigSync{
						Git: &krm.FeaturemembershipGit{
							HttpsProxy: &dummyString1,
							SyncRev:    &dummyString1,
							SecretType: &dummyString1,
							SyncBranch: &dummyString1,
						},
						PreventDrift: &dummyBool1,
						SourceFormat: &dummyString1,
					},
					Version: &dummyString1,
				},
			},
		},
		{
			name: "left has some fields with their values different from right, diff exists",
			left: krm.GKEHubFeatureMembershipSpec{
				Configmanagement: &krm.FeaturemembershipConfigmanagement{
					ConfigSync: &krm.FeaturemembershipConfigSync{
						Git: &krm.FeaturemembershipGit{
							HttpsProxy: &dummyString1,
							SyncRev:    &dummyString1,
						},
						PreventDrift: &dummyBool1,
					},
					Version: &dummyString1,
				},
			},
			right: krm.GKEHubFeatureMembershipSpec{
				Configmanagement: &krm.FeaturemembershipConfigmanagement{
					ConfigSync: &krm.FeaturemembershipConfigSync{
						Git: &krm.FeaturemembershipGit{
							HttpsProxy: &dummyString2,
							SyncRev:    &dummyString1,
							SecretType: &dummyString1,
							SyncBranch: &dummyString1,
						},
						PreventDrift: &dummyBool2,
						SourceFormat: &dummyString1,
					},
					Version: &dummyString2,
				},
			},
			expectedDiffs: []diff{
				{
					path:       "configmanagement.configSync.git.httpsProxy",
					desiredVal: dummyString1,
				},
				{
					path:       "configmanagement.version",
					desiredVal: dummyString1,
				},
				{
					path:       "configmanagement.configSync.preventDrift",
					desiredVal: dummyBool1,
				},
			},
		},
		{
			name: "left has some fields that right doesn't have, diff exists",
			left: krm.GKEHubFeatureMembershipSpec{
				Configmanagement: &krm.FeaturemembershipConfigmanagement{
					ConfigSync: &krm.FeaturemembershipConfigSync{
						Git: &krm.FeaturemembershipGit{
							HttpsProxy: &dummyString1,
							SyncRev:    &dummyString1,
							SecretType: &dummyString1,
							SyncBranch: &dummyString1,
						},
						PreventDrift: &dummyBool1,
						SourceFormat: &dummyString1,
					},
					Version: &dummyString1,
				},
			},
			right: krm.GKEHubFeatureMembershipSpec{
				Configmanagement: &krm.FeaturemembershipConfigmanagement{
					ConfigSync: &krm.FeaturemembershipConfigSync{
						Git: &krm.FeaturemembershipGit{
							HttpsProxy: &dummyString1,
							SecretType: &dummyString1,
						},
						PreventDrift: &dummyBool1,
					},
					Version: &dummyString1,
				},
			},
			expectedDiffs: []diff{
				{
					path:       "configmanagement.configSync.git.syncRev",
					desiredVal: dummyString1,
				}, {
					path:       "configmanagement.configSync.git.syncBranch",
					desiredVal: dummyString1,
				},
				{
					path:       "configmanagement.configSync.sourceFormat",
					desiredVal: dummyString1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			diffs, err := ListFieldDiffs(tc.left, &tc.right)
			if err != nil != tc.expectedError {
				t.Errorf("ListFieldDiffs returns error %v; expected %t", err, tc.expectedError)
			}
			if !diffsEqual(diffs, tc.expectedDiffs) {
				t.Errorf("ListFieldDiffs returns diffs %+v; expected %+v", diffs, tc.expectedDiffs)
			}
		})
	}
}

func diffsEqual(l, r []diff) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	}
	if len(l) != len(r) {
		return false
	}
	checkedTimes := 0
	for _, lDiff := range l {
		for _, rDiff := range r {
			if !reflect.DeepEqual(lDiff, rDiff) {
				continue
			}
			checkedTimes++
			break
		}
	}
	return checkedTimes == len(l)
}

func Test_SetObjWithDiffs(t *testing.T) {
	testCases := []struct {
		name             string
		resource         map[string]interface{}
		diffs            []diff
		expectedResource map[string]interface{}
		expectedError    bool
	}{
		{
			name: "two equal objects with no diff",
			resource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"git": map[string]interface{}{
								"httpsProxy": dummyString1,
								"syncRev":    dummyString1,
							},
							"preventDrift": dummyBool1,
						},
						"version": dummyString1,
					},
				},
			},
			expectedResource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"git": map[string]interface{}{
								"httpsProxy": dummyString1,
								"syncRev":    dummyString1,
							},
							"preventDrift": dummyBool1,
						},
						"version": dummyString1,
					},
				},
			},
		},
		{
			name: "unrecognized field type",
			resource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"git": map[string]interface{}{
								"httpsProxy": dummyString1,
								"syncRev":    dummyString1,
							},
							"preventDrift": dummyBool1,
						},
						"version": dummyString1,
					},
				},
			},
			diffs: []diff{
				{
					path: "spec.configmanagement.configSync.fieldWithUnexpectedType",
					desiredVal: map[int]interface{}{
						1: 1,
						2: 2,
					},
				},
			},
			expectedError: true,
			expectedResource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"git": map[string]interface{}{
								"httpsProxy": dummyString1,
								"syncRev":    dummyString1,
							},
							"preventDrift": dummyBool1,
						},
						"version": dummyString1,
					},
				},
			},
		},
		{
			name: "set some null fields",
			resource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"preventDrift": dummyBool1,
						},
						"version": dummyString1,
					},
				},
			},
			diffs: []diff{
				{
					path: "spec.configmanagement.configSync.git",
					desiredVal: map[string]interface{}{
						"httpsProxy": dummyString1,
						"syncRev":    dummyString1,
						"secretType": dummyString1,
						"syncBranch": dummyString1,
					},
				},
				{
					path:       "spec.configmanagement.configSync.sourceFormat",
					desiredVal: dummyString1,
				},
				{
					path: "spec.configmanagement.policyController.exemptableNamespaces",
					desiredVal: []interface{}{
						dummyString1,
					},
				},
			},

			expectedResource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"git": map[string]interface{}{
								"httpsProxy": dummyString1,
								"syncRev":    dummyString1,
								"secretType": dummyString1,
								"syncBranch": dummyString1,
							},
							"preventDrift": dummyBool1,
							"sourceFormat": dummyString1,
						},
						"policyController": map[string]interface{}{
							"exemptableNamespaces": []interface{}{
								dummyString1,
							},
						},
						"version": dummyString1,
					},
				},
			},
		},
		{
			name: "overrride field values",
			resource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"git": map[string]interface{}{
								"httpsProxy": dummyString1,
								"syncRev":    dummyString1,
								"secretType": dummyString1,
								"syncBranch": dummyString1,
							},
							"preventDrift": dummyBool1,
							"sourceFormat": dummyString1,
						},
						"version": dummyString1,
					},
				},
			},
			expectedResource: map[string]interface{}{
				"spec": map[string]interface{}{
					"configmanagement": map[string]interface{}{
						"configSync": map[string]interface{}{
							"git": map[string]interface{}{
								"httpsProxy": dummyString2,
								"syncRev":    dummyString1,
								"secretType": dummyString1,
								"syncBranch": dummyString1,
							},
							"preventDrift": dummyBool2,
							"sourceFormat": dummyString1,
						},
						"version": dummyString2,
					},
				},
			},
			diffs: []diff{
				{
					path:       "spec.configmanagement.configSync.git.httpsProxy",
					desiredVal: dummyString2,
				},
				{
					path:       "spec.configmanagement.version",
					desiredVal: dummyString2,
				},
				{
					path:       "spec.configmanagement.configSync.preventDrift",
					desiredVal: dummyBool2,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetObjWithDiffs(&tc.resource, tc.diffs)
			if err != nil != tc.expectedError {
				t.Errorf("setFeatureMembershipKRMWithDiffs returns error %v; expected %t", err, tc.expectedError)
			}
			if !reflect.DeepEqual(tc.resource, tc.expectedResource) {
				t.Errorf("setFeatureMembershipKRMWithDiffs sets obj as %+v, expected: %+v", tc.resource, tc.expectedResource)
			}
		})
	}
}
