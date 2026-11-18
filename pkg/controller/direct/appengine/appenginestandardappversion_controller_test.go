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

package appengine

import (
	"context"
	"reflect"
	"sort"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
)

func TestStandardAppVersion_CompareSpecifiedSpec(t *testing.T) {
	ctx := context.Background()

	actual := &pb.Version{
		Scaling: &pb.Version_AutomaticScaling{
			AutomaticScaling: &pb.AutomaticScaling{
				MinIdleInstances:      1,
				MaxIdleInstances:      3,
				MaxConcurrentRequests: 10, // unmanaged in desired KRM
				StandardSchedulerSettings: &pb.StandardSchedulerSettings{
					TargetCpuUtilization: 0.5,
					MinInstances:         2,
					MaxInstances:         10,
				},
			},
		},
		InstanceClass: "F1",
		Entrypoint: &pb.Entrypoint{
			Command: &pb.Entrypoint_Shell{
				Shell: "gunicorn -b :$PORT main:app",
			},
		},
		Deployment: &pb.Deployment{
			Files: map[string]*pb.FileInfo{
				"main.py": {
					SourceUrl: "https://storage.googleapis.com/staging.mock-project.appspot.com/main.py",
					Sha1Sum:   "abc123sha1",
				},
			},
		},
	}

	desired := &krm.AppEngineStandardAppVersionSpec{
		AutomaticScaling: &krm.AutomaticScaling{
			MinIdleInstances: direct.LazyPtr(2), // changed from 1 to 2
			MaxIdleInstances: direct.LazyPtr(4), // changed from 3 to 4
			// MaxConcurrentRequests is omitted/unmanaged
			StandardSchedulerSettings: &krm.StandardSchedulerSettings{
				TargetCPUUtilization: direct.LazyPtr(0.5), // unchanged
				MinInstances:         direct.LazyPtr(2),   // unchanged
				MaxInstances:         direct.LazyPtr(10),  // unchanged
			},
		},
		InstanceClass: direct.LazyPtr("F2"), // changed from F1 to F2
		Entrypoint: &krm.Entrypoint{
			Shell: "gunicorn -b :$PORT main:app", // server doesn't return this; normalizer ignores diff
		},
		Deployment: &krm.Deployment{
			Files: []krm.DeploymentFile{
				{
					Name:      "main.py",
					SourceUrl: "https://storage.googleapis.com/staging.mock-project.appspot.com/main.py",
					// Sha1Sum omitted; normalizer ignores diff
				},
			},
		},
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		desired,
		actual,
		AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
		AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
		normalizeStandardAppVersionForDiff,
		true, // fineGrained
	)
	if err != nil {
		t.Fatalf("unexpected error from CompareSpecifiedSpec: %v", err)
	}

	if diffRes.Empty() {
		t.Fatalf("expected non-empty diff")
	}

	gotPaths := diffRes.FieldMask.GetPaths()
	sort.Strings(gotPaths)
	wantPaths := []string{
		"automatic_scaling.max_idle_instances",
		"automatic_scaling.min_idle_instances",
		"instance_class",
	}
	sort.Strings(wantPaths)

	if !reflect.DeepEqual(gotPaths, wantPaths) {
		t.Fatalf("unexpected paths in fieldmask: got %v, want %v", gotPaths, wantPaths)
	}

	// Verify unmanaged field is preserved in MergedDesired
	if diffRes.MergedDesired.GetAutomaticScaling().GetMaxConcurrentRequests() != 10 {
		t.Fatalf("expected MergedDesired to preserve unmanaged max_concurrent_requests=10, got %d",
			diffRes.MergedDesired.GetAutomaticScaling().GetMaxConcurrentRequests())
	}
}

func TestChooseStandardUpdateStrategy(t *testing.T) {
	allPatchable := []string{
		"instance_class",
		"automatic_scaling.min_idle_instances",
		"automatic_scaling.max_idle_instances",
		"automatic_scaling.standard_scheduler_settings.min_instances",
		"automatic_scaling.standard_scheduler_settings.max_instances",
		"automatic_scaling.standard_scheduler_settings.target_cpu_utilization",
		"automatic_scaling.standard_scheduler_settings.target_throughput_utilization",
		"manual_scaling.instances",
	}

	tests := []struct {
		name               string
		paths              []string
		scalingTypeChanged bool
		wantStrategy       standardUpdateStrategy
		wantReasons        []string
	}{
		{
			name:         "no paths and no scaling change is a no-op",
			paths:        nil,
			wantStrategy: standardUpdateNone,
		},
		{
			name:         "all 8 patchable paths together use patch",
			paths:        allPatchable,
			wantStrategy: standardUpdatePatch,
		},
		{
			name:         "mixed patchable and non-patchable paths redeploy with non-patchable reasons",
			paths:        []string{"instance_class", "env_variables"},
			wantStrategy: standardUpdateRedeploy,
			wantReasons:  []string{"env_variables"},
		},
		{
			name:         "non-leaf automatic_scaling path redeploys",
			paths:        []string{"automatic_scaling"},
			wantStrategy: standardUpdateRedeploy,
			wantReasons:  []string{"automatic_scaling"},
		},
		{
			name:         "serving_status is not patchable from KRM spec",
			paths:        []string{"serving_status"},
			wantStrategy: standardUpdateRedeploy,
			wantReasons:  []string{"serving_status"},
		},
		{
			name:         "flex-only automatic_scaling.max_total_instances redeploys",
			paths:        []string{"automatic_scaling.max_total_instances"},
			wantStrategy: standardUpdateRedeploy,
			wantReasons:  []string{"automatic_scaling.max_total_instances"},
		},
		{
			name:               "no paths with scaling type change redeploys",
			paths:              nil,
			scalingTypeChanged: true,
			wantStrategy:       standardUpdateRedeploy,
			wantReasons:        []string{"scaling type"},
		},
		{
			name:               "patchable paths with scaling type change redeploys",
			paths:              []string{"instance_class"},
			scalingTypeChanged: true,
			wantStrategy:       standardUpdateRedeploy,
			wantReasons:        []string{"scaling type"},
		},
	}

	for _, p := range allPatchable {
		tests = append(tests, struct {
			name               string
			paths              []string
			scalingTypeChanged bool
			wantStrategy       standardUpdateStrategy
			wantReasons        []string
		}{
			name:         "single patchable path " + p,
			paths:        []string{p},
			wantStrategy: standardUpdatePatch,
		})
	}

	redeployPaths := []string{
		"env_variables",
		"deployment.files",
		"handlers",
		"runtime",
		"inbound_services",
		"threadsafe",
		"service_account",
		"vpc_access_connector.name",
		"basic_scaling.max_instances",
		"automatic_scaling.max_concurrent_requests",
		"automatic_scaling.max_pending_latency",
	}
	for _, p := range redeployPaths {
		tests = append(tests, struct {
			name               string
			paths              []string
			scalingTypeChanged bool
			wantStrategy       standardUpdateStrategy
			wantReasons        []string
		}{
			name:         "non-patchable path " + p + " redeploys",
			paths:        []string{p},
			wantStrategy: standardUpdateRedeploy,
			wantReasons:  []string{p},
		})
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotStrategy, gotReasons := chooseStandardUpdateStrategy(tc.paths, tc.scalingTypeChanged)

			if gotStrategy != tc.wantStrategy {
				t.Errorf("strategy = %v, want %v", gotStrategy, tc.wantStrategy)
			}
			if !reflect.DeepEqual(gotReasons, tc.wantReasons) {
				t.Errorf("reasons = %v, want %v", gotReasons, tc.wantReasons)
			}
		})
	}
}

func TestStandardAppVersionAdapter_ScalingTypeChange(t *testing.T) {
	tests := []struct {
		name        string
		spec        krm.AppEngineStandardAppVersionSpec
		actual      *pb.Version
		wantActual  standardScalingType
		wantDesired standardScalingType
		wantChanged bool
	}{
		{
			name: "spec none and live automatic has no change",
			spec: krm.AppEngineStandardAppVersionSpec{},
			actual: &pb.Version{
				Scaling: &pb.Version_AutomaticScaling{AutomaticScaling: &pb.AutomaticScaling{}},
			},
			wantActual:  standardScalingAutomatic,
			wantDesired: standardScalingUnspecified,
			wantChanged: false,
		},
		{
			name: "spec automatic and live none defaults to automatic with no change",
			spec: krm.AppEngineStandardAppVersionSpec{
				AutomaticScaling: &krm.AutomaticScaling{},
			},
			actual:      &pb.Version{},
			wantActual:  standardScalingAutomatic,
			wantDesired: standardScalingAutomatic,
			wantChanged: false,
		},
		{
			name: "automatic to manual detects change",
			spec: krm.AppEngineStandardAppVersionSpec{
				ManualScaling: &krm.ManualScaling{Instances: 1},
			},
			actual: &pb.Version{
				Scaling: &pb.Version_AutomaticScaling{AutomaticScaling: &pb.AutomaticScaling{}},
			},
			wantActual:  standardScalingAutomatic,
			wantDesired: standardScalingManual,
			wantChanged: true,
		},
		{
			name: "manual to automatic detects scaling switch",
			spec: krm.AppEngineStandardAppVersionSpec{
				AutomaticScaling: &krm.AutomaticScaling{MaxConcurrentRequests: direct.LazyPtr(20)},
			},
			actual: &pb.Version{
				Scaling: &pb.Version_ManualScaling{ManualScaling: &pb.ManualScaling{Instances: 1}},
			},
			wantActual:  standardScalingManual,
			wantDesired: standardScalingAutomatic,
			wantChanged: true,
		},
		{
			name: "basic to automatic detects change",
			spec: krm.AppEngineStandardAppVersionSpec{
				AutomaticScaling: &krm.AutomaticScaling{},
			},
			actual: &pb.Version{
				Scaling: &pb.Version_BasicScaling{BasicScaling: &pb.BasicScaling{MaxInstances: 2}},
			},
			wantActual:  standardScalingBasic,
			wantDesired: standardScalingAutomatic,
			wantChanged: true,
		},
		{
			name: "manual to basic detects change",
			spec: krm.AppEngineStandardAppVersionSpec{
				BasicScaling: &krm.BasicScaling{MaxInstances: 2},
			},
			actual: &pb.Version{
				Scaling: &pb.Version_ManualScaling{ManualScaling: &pb.ManualScaling{Instances: 1}},
			},
			wantActual:  standardScalingManual,
			wantDesired: standardScalingBasic,
			wantChanged: true,
		},
		{
			name: "automatic to basic detects change",
			spec: krm.AppEngineStandardAppVersionSpec{
				BasicScaling: &krm.BasicScaling{MaxInstances: 2},
			},
			actual: &pb.Version{
				Scaling: &pb.Version_AutomaticScaling{AutomaticScaling: &pb.AutomaticScaling{}},
			},
			wantActual:  standardScalingAutomatic,
			wantDesired: standardScalingBasic,
			wantChanged: true,
		},
		{
			name: "same manual type has no change",
			spec: krm.AppEngineStandardAppVersionSpec{
				ManualScaling: &krm.ManualScaling{Instances: 2},
			},
			actual: &pb.Version{
				Scaling: &pb.Version_ManualScaling{ManualScaling: &pb.ManualScaling{Instances: 1}},
			},
			wantActual:  standardScalingManual,
			wantDesired: standardScalingManual,
			wantChanged: false,
		},
		{
			name: "spec with manual and automatic set uses manual ToProto precedence",
			spec: krm.AppEngineStandardAppVersionSpec{
				AutomaticScaling: &krm.AutomaticScaling{},
				ManualScaling:    &krm.ManualScaling{Instances: 1},
			},
			actual: &pb.Version{
				Scaling: &pb.Version_AutomaticScaling{AutomaticScaling: &pb.AutomaticScaling{}},
			},
			wantActual:  standardScalingAutomatic,
			wantDesired: standardScalingManual,
			wantChanged: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &standardAppVersionAdapter{
				desired: &krm.AppEngineStandardAppVersion{Spec: tc.spec},
				actual:  tc.actual,
			}

			gotActual, gotDesired, gotChanged := adapter.scalingTypeChange()

			if gotActual != tc.wantActual || gotDesired != tc.wantDesired || gotChanged != tc.wantChanged {
				t.Errorf("scalingTypeChange() = (%q, %q, %v), want (%q, %q, %v)",
					gotActual, gotDesired, gotChanged, tc.wantActual, tc.wantDesired, tc.wantChanged)
			}
		})
	}
}

func TestNormalizeHandlersForDiff(t *testing.T) {
	tests := []struct {
		name string
		in   []*pb.UrlMap
		want []*pb.UrlMap
	}{
		{
			name: "fills login security and authFail defaults and keeps explicit non-defaults",
			in: []*pb.UrlMap{
				{
					UrlRegex: "/default/.*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex:       "/admin/.*",
					Login:          pb.LoginRequirement_LOGIN_ADMIN,
					SecurityLevel:  pb.SecurityLevel_SECURE_ALWAYS,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_UNAUTHORIZED,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       "/default/.*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex:       "/admin/.*",
					Login:          pb.LoginRequirement_LOGIN_ADMIN,
					SecurityLevel:  pb.SecurityLevel_SECURE_ALWAYS,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_UNAUTHORIZED,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "normalizes SECURE_DEFAULT to SECURE_OPTIONAL",
			in: []*pb.UrlMap{
				{
					UrlRegex:      "/sec/.*",
					SecurityLevel: pb.SecurityLevel_SECURE_DEFAULT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       "/sec/.*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "strips one and several trailing runtime default handlers",
			in: []*pb.UrlMap{
				{
					UrlRegex:      "/api/.*",
					SecurityLevel: pb.SecurityLevel_SECURE_ALWAYS,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex: ".*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex:      ".*",
					Login:         pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel: pb.SecurityLevel_SECURE_OPTIONAL,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       "/api/.*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_ALWAYS,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "keeps non-trailing catch-all handler",
			in: []*pb.UrlMap{
				{
					UrlRegex: ".*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex: "/after/.*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       ".*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex:       "/after/.*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "keeps trailing catch-all when scriptPath is non-default",
			in: []*pb.UrlMap{
				{
					UrlRegex: ".*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "main.app"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       ".*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "main.app"},
					},
				},
			},
		},
		{
			name: "keeps trailing catch-all when handler is staticFiles",
			in: []*pb.UrlMap{
				{
					UrlRegex: ".*",
					HandlerType: &pb.UrlMap_StaticFiles{
						StaticFiles: &pb.StaticFilesHandler{Path: "static"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       ".*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_StaticFiles{
						StaticFiles: &pb.StaticFilesHandler{Path: "static"},
					},
				},
			},
		},
		{
			name: "keeps trailing catch-all when securityLevel is non-default",
			in: []*pb.UrlMap{
				{
					UrlRegex:      ".*",
					SecurityLevel: pb.SecurityLevel_SECURE_ALWAYS,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       ".*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_ALWAYS,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "keeps trailing catch-all when authFailAction is non-default",
			in: []*pb.UrlMap{
				{
					UrlRegex:       ".*",
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_UNAUTHORIZED,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       ".*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_UNAUTHORIZED,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "keeps trailing catch-all when redirectHttpResponseCode is non-default",
			in: []*pb.UrlMap{
				{
					UrlRegex:                 ".*",
					RedirectHttpResponseCode: pb.UrlMap_REDIRECT_HTTP_RESPONSE_CODE_301,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:                 ".*",
					Login:                    pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:            pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction:           pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					RedirectHttpResponseCode: pb.UrlMap_REDIRECT_HTTP_RESPONSE_CODE_301,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "keeps trailing catch-all when login is non-default",
			in: []*pb.UrlMap{
				{
					UrlRegex: ".*",
					Login:    pb.LoginRequirement_LOGIN_REQUIRED,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       ".*",
					Login:          pb.LoginRequirement_LOGIN_REQUIRED,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
		{
			name: "clears 0s static expiration and keeps non-zero expiration and skips nil entries",
			in: []*pb.UrlMap{
				nil,
				{
					UrlRegex: "/static0/.*",
					HandlerType: &pb.UrlMap_StaticFiles{
						StaticFiles: &pb.StaticFilesHandler{
							Path:       "static",
							Expiration: durationpb.New(0),
						},
					},
				},
				{
					UrlRegex: "/static60/.*",
					HandlerType: &pb.UrlMap_StaticFiles{
						StaticFiles: &pb.StaticFilesHandler{
							Path:       "static",
							Expiration: &durationpb.Duration{Seconds: 60},
						},
					},
				},
			},
			want: []*pb.UrlMap{
				{
					UrlRegex:       "/static0/.*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_StaticFiles{
						StaticFiles: &pb.StaticFilesHandler{
							Path: "static",
						},
					},
				},
				{
					UrlRegex:       "/static60/.*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_StaticFiles{
						StaticFiles: &pb.StaticFilesHandler{
							Path:       "static",
							Expiration: &durationpb.Duration{Seconds: 60},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normalizeHandlersForDiff(tc.in)

			if len(got) != len(tc.want) {
				t.Fatalf("len(got) = %d, want %d\ngot:  %v\nwant: %v", len(got), len(tc.want), got, tc.want)
			}
			for i := range tc.want {
				if !proto.Equal(got[i], tc.want[i]) {
					t.Errorf("handler[%d] mismatch:\ngot:  %v\nwant: %v", i, got[i], tc.want[i])
				}
			}

			cloned := make([]*pb.UrlMap, len(got))
			for i, h := range got {
				cloned[i] = proto.Clone(h).(*pb.UrlMap)
			}
			gotSecond := normalizeHandlersForDiff(cloned)

			if len(gotSecond) != len(tc.want) {
				t.Fatalf("not idempotent: len(gotSecond) = %d, want %d", len(gotSecond), len(tc.want))
			}
			for i := range tc.want {
				if !proto.Equal(gotSecond[i], tc.want[i]) {
					t.Errorf("not idempotent at [%d]:\ngot:  %v\nwant: %v", i, gotSecond[i], tc.want[i])
				}
			}
		})
	}
}

func TestStandardAppVersion_DiffIgnoresServerDefaults(t *testing.T) {
	ctx := context.Background()

	makeCreateSpecAndActual := func() (krm.AppEngineStandardAppVersionSpec, *pb.Version) {
		spec := krm.AppEngineStandardAppVersionSpec{
			Runtime:        "python311",
			ServiceAccount: direct.LazyPtr("test-project@appspot.gserviceaccount.com"),
			Threadsafe:     direct.LazyPtr(true),
			InstanceClass:  direct.LazyPtr("B2"),
			ManualScaling: &krm.ManualScaling{
				Instances: 1,
			},
			Entrypoint: &krm.Entrypoint{
				Shell: "gunicorn -b :$PORT main:app",
			},
			Deployment: &krm.Deployment{
				Files: []krm.DeploymentFile{
					{
						Name:      "main.py",
						SourceUrl: "https://storage.googleapis.com/staging.test-project.appspot.com/05e6871e9161335ea72fe540fb798803633b43ba",
					},
					{
						Name:      "app.yaml",
						SourceUrl: "https://storage.googleapis.com/staging.test-project.appspot.com/fccc0c6d2e2fcb05c84f69f40c7a613d2338e8d8",
					},
				},
			},
			EnvVariables: map[string]string{
				"PORT":  "8080",
				"STAGE": "create",
			},
			InboundServices: []string{"INBOUND_SERVICE_WARMUP"},
			Handlers: []krm.URLMap{
				{
					URLRegex:      direct.LazyPtr("/api/.*"),
					SecurityLevel: direct.LazyPtr("SECURE_ALWAYS"),
					Script: &krm.ScriptHandler{
						ScriptPath: "auto",
					},
				},
				{
					URLRegex: direct.LazyPtr(".*"),
					Script: &krm.ScriptHandler{
						ScriptPath: "auto",
					},
				},
			},
		}

		actual := &pb.Version{
			Runtime:        "python311",
			ServiceAccount: "test-project@appspot.gserviceaccount.com",
			Threadsafe:     true,
			InstanceClass:  "B2",
			Scaling: &pb.Version_ManualScaling{
				ManualScaling: &pb.ManualScaling{Instances: 1},
			},
			Deployment: &pb.Deployment{
				Files: map[string]*pb.FileInfo{
					"main.py": {
						SourceUrl: "https://storage.googleapis.com/staging.test-project.appspot.com/05e6871e9161335ea72fe540fb798803633b43ba",
						Sha1Sum:   "05e6871e9161335ea72fe540fb798803633b43ba",
					},
					"app.yaml": {
						SourceUrl: "https://storage.googleapis.com/staging.test-project.appspot.com/fccc0c6d2e2fcb05c84f69f40c7a613d2338e8d8",
						Sha1Sum:   "fccc0c6d2e2fcb05c84f69f40c7a613d2338e8d8",
					},
				},
			},
			EnvVariables: map[string]string{
				"PORT":  "8080",
				"STAGE": "create",
			},
			InboundServices: []pb.InboundServiceType{pb.InboundServiceType_INBOUND_SERVICE_WARMUP},
			Handlers: []*pb.UrlMap{
				{
					UrlRegex:       "/api/.*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_ALWAYS,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex:       ".*",
					Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
				{
					UrlRegex:      ".*",
					Login:         pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel: pb.SecurityLevel_SECURE_OPTIONAL,
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		}
		return spec, actual
	}

	t.Run("spec with handlers and server defaults produces empty diff", func(t *testing.T) {
		spec, actual := makeCreateSpecAndActual()
		adapter := &standardAppVersionAdapter{
			desired: &krm.AppEngineStandardAppVersion{Spec: spec},
			actual:  actual,
		}

		diffRes, err := common.CompareSpecifiedSpec(
			ctx,
			&adapter.desired.Spec,
			adapter.actual,
			AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
			AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
			adapter.diffNormalizer(),
			true,
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !diffRes.Empty() {
			t.Fatalf("expected empty diff, got paths: %v", diffRes.FieldMask.GetPaths())
		}
	})

	t.Run("spec with handlers and changed env variable diffs only env_variables", func(t *testing.T) {
		spec, actual := makeCreateSpecAndActual()
		spec.EnvVariables["STAGE"] = "update"
		adapter := &standardAppVersionAdapter{
			desired: &krm.AppEngineStandardAppVersion{Spec: spec},
			actual:  actual,
		}

		diffRes, err := common.CompareSpecifiedSpec(
			ctx,
			&adapter.desired.Spec,
			adapter.actual,
			AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
			AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
			adapter.diffNormalizer(),
			true,
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		wantPaths := []string{"env_variables"}
		if !reflect.DeepEqual(diffRes.FieldMask.GetPaths(), wantPaths) {
			t.Fatalf("paths = %v, want %v", diffRes.FieldMask.GetPaths(), wantPaths)
		}
	})

	t.Run("spec with deployment zip vs actual without zip produces empty diff", func(t *testing.T) {
		spec := krm.AppEngineStandardAppVersionSpec{
			Runtime: "python311",
			Deployment: &krm.Deployment{
				Zip: &krm.DeploymentZip{
					SourceUrl: "https://storage.googleapis.com/staging.test-project.appspot.com/app.zip",
				},
			},
		}
		actual := &pb.Version{
			Runtime: "python311",
		}
		adapter := &standardAppVersionAdapter{
			desired: &krm.AppEngineStandardAppVersion{Spec: spec},
			actual:  actual,
		}

		diffRes, err := common.CompareSpecifiedSpec(
			ctx,
			&adapter.desired.Spec,
			adapter.actual,
			AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
			AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
			adapter.diffNormalizer(),
			true,
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !diffRes.Empty() {
			t.Fatalf("expected empty diff, got paths: %v", diffRes.FieldMask.GetPaths())
		}
	})

	t.Run("spec without handlers preserves live default handler in MergedDesired", func(t *testing.T) {
		liveHandler := &pb.UrlMap{
			Login:         pb.LoginRequirement_LOGIN_OPTIONAL,
			SecurityLevel: pb.SecurityLevel_SECURE_OPTIONAL,
			UrlRegex:      ".*",
			HandlerType: &pb.UrlMap_Script{
				Script: &pb.ScriptHandler{ScriptPath: "auto"},
			},
		}
		spec := krm.AppEngineStandardAppVersionSpec{
			Runtime: "python311",
		}
		actual := &pb.Version{
			Runtime:  "python311",
			Handlers: []*pb.UrlMap{proto.Clone(liveHandler).(*pb.UrlMap)},
		}
		adapter := &standardAppVersionAdapter{
			desired: &krm.AppEngineStandardAppVersion{Spec: spec},
			actual:  actual,
		}

		diffRes, err := common.CompareSpecifiedSpec(
			ctx,
			&adapter.desired.Spec,
			adapter.actual,
			AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
			AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
			adapter.diffNormalizer(),
			true,
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !diffRes.Empty() {
			t.Fatalf("expected empty diff, got paths: %v", diffRes.FieldMask.GetPaths())
		}
		if len(diffRes.MergedDesired.GetHandlers()) != 1 || !proto.Equal(diffRes.MergedDesired.GetHandlers()[0], liveHandler) {
			t.Fatalf("MergedDesired.Handlers = %v, want [%v]", diffRes.MergedDesired.GetHandlers(), liveHandler)
		}
	})
}

func TestStandardAppVersion_BasicFixtureStillPatches(t *testing.T) {
	ctx := context.Background()

	const realGetJSON = `{
	  "automaticScaling": {
	    "maxConcurrentRequests": 10,
	    "maxIdleInstances": 3,
	    "maxPendingLatency": "5s",
	    "minIdleInstances": 1,
	    "minPendingLatency": "1s",
	    "standardSchedulerSettings": {
	      "maxInstances": 10,
	      "minInstances": 2,
	      "targetCpuUtilization": 0.5,
	      "targetThroughputUtilization": 0.75
	    }
	  },
	  "createTime": "2024-04-01T12:34:56.123456Z",
	  "createdBy": "overseer-kcc-tester@${projectId}.iam.gserviceaccount.com",
	  "deployment": {
	    "files": {
	      "app.yaml": {
	        "sha1Sum": "fccc0c6d2e2fcb05c84f69f40c7a613d2338e8d8",
	        "sourceUrl": "https://storage.googleapis.com/staging.${projectId}.appspot.com/fccc0c6d2e2fcb05c84f69f40c7a613d2338e8d8"
	      },
	      "main.py": {
	        "sha1Sum": "05e6871e9161335ea72fe540fb798803633b43ba",
	        "sourceUrl": "https://storage.googleapis.com/staging.${projectId}.appspot.com/05e6871e9161335ea72fe540fb798803633b43ba"
	      }
	    }
	  },
	  "diskUsageBytes": "2978869",
	  "env": "standard",
	  "envVariables": {
	    "PORT": "8080"
	  },
	  "handlers": [
	    {
	      "login": 1,
	      "script": {
	        "scriptPath": "auto"
	      },
	      "securityLevel": 2,
	      "urlRegex": ".*"
	    }
	  ],
	  "id": "000000000000000000000",
	  "instanceClass": "F1",
	  "name": "apps/${projectId}/services/default/versions/v${uniqueId}",
	  "network": {},
	  "runtime": "python311",
	  "runtimeChannel": "default",
	  "serviceAccount": "${projectId}@appspot.gserviceaccount.com",
	  "servingStatus": 1,
	  "threadsafe": true,
	  "versionUrl": "https://v${uniqueId}-dot-${projectId}.uc.r.appspot.com"
	}`

	const realPatchJSON = `{
	  "automaticScaling": {
	    "maxConcurrentRequests": 10,
	    "maxIdleInstances": 4,
	    "maxPendingLatency": "5s",
	    "minIdleInstances": 2,
	    "minPendingLatency": "1s",
	    "standardSchedulerSettings": {
	      "maxInstances": 10,
	      "minInstances": 2,
	      "targetCpuUtilization": 0.5,
	      "targetThroughputUtilization": 0.75
	    }
	  },
	  "deployment": {
	    "files": {
	      "app.yaml": {
	        "sourceUrl": "https://storage.googleapis.com/staging.${projectId}.appspot.com/fccc0c6d2e2fcb05c84f69f40c7a613d2338e8d8"
	      },
	      "main.py": {
	        "sourceUrl": "https://storage.googleapis.com/staging.${projectId}.appspot.com/05e6871e9161335ea72fe540fb798803633b43ba"
	      }
	    }
	  },
	  "envVariables": {
	    "PORT": "8080"
	  },
	  "handlers": [
	    {
	      "login": 1,
	      "script": {
	        "scriptPath": "auto"
	      },
	      "securityLevel": 2,
	      "urlRegex": ".*"
	    }
	  ],
	  "instanceClass": "F2",
	  "runtime": "python311",
	  "serviceAccount": "${projectId}@appspot.gserviceaccount.com",
	  "threadsafe": true
	}`

	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	actual := &pb.Version{}
	if err := unmarshaler.Unmarshal([]byte(realGetJSON), actual); err != nil {
		t.Fatalf("unmarshaling real GET JSON: %v", err)
	}
	wantPatchBody := &pb.Version{}
	if err := unmarshaler.Unmarshal([]byte(realPatchJSON), wantPatchBody); err != nil {
		t.Fatalf("unmarshaling real PATCH JSON: %v", err)
	}

	desiredSpec := krm.AppEngineStandardAppVersionSpec{
		Runtime:       "python311",
		InstanceClass: direct.LazyPtr("F2"),
		Entrypoint: &krm.Entrypoint{
			Shell: "gunicorn -b :$PORT main:app",
		},
		Deployment: &krm.Deployment{
			Files: []krm.DeploymentFile{
				{
					Name:      "main.py",
					SourceUrl: "https://storage.googleapis.com/staging.${projectId}.appspot.com/05e6871e9161335ea72fe540fb798803633b43ba",
				},
				{
					Name:      "app.yaml",
					SourceUrl: "https://storage.googleapis.com/staging.${projectId}.appspot.com/fccc0c6d2e2fcb05c84f69f40c7a613d2338e8d8",
				},
			},
		},
		EnvVariables: map[string]string{
			"PORT": "8080",
		},
		AutomaticScaling: &krm.AutomaticScaling{
			MaxConcurrentRequests: direct.LazyPtr(10),
			MinIdleInstances:      direct.LazyPtr(2),
			MaxIdleInstances:      direct.LazyPtr(4),
			MinPendingLatency:     direct.LazyPtr("1s"),
			MaxPendingLatency:     direct.LazyPtr("5s"),
			StandardSchedulerSettings: &krm.StandardSchedulerSettings{
				TargetCPUUtilization:        direct.LazyPtr(0.5),
				TargetThroughputUtilization: direct.LazyPtr(0.75),
				MinInstances:                direct.LazyPtr(2),
				MaxInstances:                direct.LazyPtr(10),
			},
		},
	}

	adapter := &standardAppVersionAdapter{
		desired: &krm.AppEngineStandardAppVersion{Spec: desiredSpec},
		actual:  actual,
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&adapter.desired.Spec,
		adapter.actual,
		AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
		AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
		adapter.diffNormalizer(),
		true,
	)
	if err != nil {
		t.Fatalf("CompareSpecifiedSpec failed: %v", err)
	}

	_, _, scalingChanged := adapter.scalingTypeChange()
	strategy, reasons := chooseStandardUpdateStrategy(diffRes.FieldMask.GetPaths(), scalingChanged)

	if strategy != standardUpdatePatch || len(reasons) != 0 {
		t.Fatalf("chooseStandardUpdateStrategy = (%v, %v), want (standardUpdatePatch, nil)", strategy, reasons)
	}
	wantMask := []string{
		"automatic_scaling.max_idle_instances",
		"automatic_scaling.min_idle_instances",
		"instance_class",
	}
	if !reflect.DeepEqual(diffRes.FieldMask.GetPaths(), wantMask) {
		t.Fatalf("FieldMask.Paths = %v, want %v", diffRes.FieldMask.GetPaths(), wantMask)
	}
	if !proto.Equal(diffRes.MergedDesired, wantPatchBody) {
		t.Errorf("MergedDesired mismatch vs real PATCH body:\ngot:  %v\nwant: %v", diffRes.MergedDesired, wantPatchBody)
	}
}

func TestStandardAppVersionAdapter_NewCreateVersionRequest(t *testing.T) {
	adapter := &standardAppVersionAdapter{
		id: &krm.AppEngineStandardAppVersionIdentity{
			App:     "p",
			Service: "s",
			Version: "v1",
		},
		actual: &pb.Version{
			Scaling: &pb.Version_ManualScaling{
				ManualScaling: &pb.ManualScaling{Instances: 1},
			},
			Threadsafe:     true,
			ServiceAccount: "p@appspot.gserviceaccount.com",
		},
		desired: &krm.AppEngineStandardAppVersion{
			Spec: krm.AppEngineStandardAppVersionSpec{
				Runtime:       "python311",
				InstanceClass: direct.LazyPtr("F2"),
				AutomaticScaling: &krm.AutomaticScaling{
					MaxConcurrentRequests: direct.LazyPtr(20),
				},
				Entrypoint: &krm.Entrypoint{
					Shell: "gunicorn -b :$PORT --workers 2 main:app",
				},
				Deployment: &krm.Deployment{
					Files: []krm.DeploymentFile{
						{
							Name:      "main.py",
							SourceUrl: "https://storage.googleapis.com/staging.p.appspot.com/main.py",
						},
					},
					Zip: &krm.DeploymentZip{
						SourceUrl: "https://storage.googleapis.com/staging.p.appspot.com/app.zip",
					},
				},
				Handlers: []krm.URLMap{
					{
						URLRegex: direct.LazyPtr("/z/.*"),
						Script:   &krm.ScriptHandler{ScriptPath: "auto"},
					},
					{
						URLRegex: direct.LazyPtr("/a/.*"),
						Script:   &krm.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
	}

	req, err := adapter.newCreateVersionRequest()

	if err != nil {
		t.Fatalf("newCreateVersionRequest() failed: %v", err)
	}
	if req.GetParent() != "apps/p/services/s" {
		t.Errorf("Parent = %q, want %q", req.GetParent(), "apps/p/services/s")
	}
	v := req.GetVersion()
	if v.GetId() != "v1" {
		t.Errorf("Version.Id = %q, want %q", v.GetId(), "v1")
	}
	if v.GetEntrypoint().GetShell() != "gunicorn -b :$PORT --workers 2 main:app" {
		t.Errorf("Entrypoint.Shell = %q, want write-only shell preserved", v.GetEntrypoint().GetShell())
	}
	if v.GetDeployment().GetFiles()["main.py"].GetSourceUrl() != "https://storage.googleapis.com/staging.p.appspot.com/main.py" {
		t.Errorf("Deployment.Files[main.py] missing or wrong: %v", v.GetDeployment().GetFiles())
	}
	if v.GetDeployment().GetFiles()["main.py"].GetSha1Sum() != "" {
		t.Errorf("expected empty Sha1Sum, got %q", v.GetDeployment().GetFiles()["main.py"].GetSha1Sum())
	}
	if v.GetDeployment().GetZip().GetSourceUrl() != "https://storage.googleapis.com/staging.p.appspot.com/app.zip" {
		t.Errorf("Deployment.Zip.SourceUrl = %q, want zip preserved", v.GetDeployment().GetZip().GetSourceUrl())
	}
	if len(v.GetHandlers()) != 2 || v.GetHandlers()[0].GetUrlRegex() != "/z/.*" || v.GetHandlers()[1].GetUrlRegex() != "/a/.*" {
		t.Errorf("Handlers order/content = %v, want [/z/.*, /a/.*] in spec order", v.GetHandlers())
	}
	if v.GetHandlers()[0].GetLogin() != pb.LoginRequirement_LOGIN_UNSPECIFIED ||
		v.GetHandlers()[0].GetSecurityLevel() != pb.SecurityLevel_SECURE_UNSPECIFIED ||
		v.GetHandlers()[0].GetAuthFailAction() != pb.AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED {
		t.Errorf("expected no injected handler defaults in create payload, got %v", v.GetHandlers()[0])
	}
	if v.GetManualScaling() != nil || v.GetAutomaticScaling().GetMaxConcurrentRequests() != 20 {
		t.Errorf("expected only spec's AutomaticScaling in payload, got Scaling=%v", v.GetScaling())
	}
	if v.GetThreadsafe() || v.GetServiceAccount() != "" {
		t.Errorf("expected unset threadsafe/serviceAccount to stay zero, got threadsafe=%v serviceAccount=%q", v.GetThreadsafe(), v.GetServiceAccount())
	}
}

func TestNormalizeStandardAppVersionForDiff_ClearsZip(t *testing.T) {
	ctx := context.Background()
	v := &pb.Version{
		Deployment: &pb.Deployment{
			Files: map[string]*pb.FileInfo{
				"main.py": {
					SourceUrl: "https://storage.googleapis.com/staging.p.appspot.com/main.py",
					Sha1Sum:   "12345",
				},
			},
			Zip: &pb.ZipInfo{
				SourceUrl:  "https://storage.googleapis.com/staging.p.appspot.com/app.zip",
				FilesCount: 3,
			},
		},
	}

	err := normalizeStandardAppVersionForDiff(ctx, v)

	if err != nil {
		t.Fatalf("normalizeStandardAppVersionForDiff failed: %v", err)
	}
	if v.GetDeployment().GetZip() != nil {
		t.Errorf("expected Deployment.Zip to be nil after normalization, got %v", v.GetDeployment().GetZip())
	}
	if v.GetDeployment().GetFiles()["main.py"].GetSourceUrl() == "" {
		t.Errorf("expected Deployment.Files to be preserved")
	}
	if v.GetDeployment().GetFiles()["main.py"].GetSha1Sum() != "" {
		t.Errorf("expected Sha1Sum to be cleared, got %q", v.GetDeployment().GetFiles()["main.py"].GetSha1Sum())
	}
}
