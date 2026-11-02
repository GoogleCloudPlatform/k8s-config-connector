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
	"reflect"
	"testing"
	"time"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	krmappenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"k8s.io/utils/ptr"
)

func TestAppEngineFirewallRuleSpecMapping(t *testing.T) {
	tests := []struct {
		name      string
		protoRule *pb.FirewallRule
		krmSpec   *krmappenginev1alpha1.AppEngineFirewallRuleSpec
	}{
		{
			name:      "nil",
			protoRule: nil,
			krmSpec:   nil,
		},
		{
			name: "full fields allow",
			protoRule: &pb.FirewallRule{
				Priority:    1000,
				Action:      pb.FirewallRule_ALLOW,
				SourceRange: "192.168.1.0/24",
				Description: "Allow local subnet",
			},
			krmSpec: &krmappenginev1alpha1.AppEngineFirewallRuleSpec{
				ResourceID:  ptr.To("1000"),
				Action:      "ALLOW",
				SourceRange: "192.168.1.0/24",
				Description: ptr.To("Allow local subnet"),
			},
		},
		{
			name: "full fields deny",
			protoRule: &pb.FirewallRule{
				Priority:    2147483647,
				Action:      pb.FirewallRule_DENY,
				SourceRange: "*",
				Description: "Default deny",
			},
			krmSpec: &krmappenginev1alpha1.AppEngineFirewallRuleSpec{
				ResourceID:  ptr.To("2147483647"),
				Action:      "DENY",
				SourceRange: "*",
				Description: ptr.To("Default deny"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineFirewallRuleSpec_v1alpha1_FromProto(mapCtx, tc.protoRule)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmSpec) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmSpec)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineFirewallRuleSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoRule) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoRule)
			}
		})
	}
}

func TestAppEngineServiceSplitTrafficSpecMapping(t *testing.T) {
	tests := []struct {
		name         string
		protoService *pb.Service
		krmSpec      *krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec
	}{
		{
			name:         "nil",
			protoService: nil,
			krmSpec:      nil,
		},
		{
			name: "full fields IP shard",
			protoService: &pb.Service{
				Id: "default",
				Split: &pb.TrafficSplit{
					ShardBy: pb.TrafficSplit_IP,
					Allocations: map[string]float64{
						"v1": 0.5,
						"v2": 0.5,
					},
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec{
				ResourceID: ptr.To("default"),
				Split: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSplit{
					ShardBy: ptr.To("IP"),
					Allocations: map[string]string{
						"v1": "0.5",
						"v2": "0.5",
					},
				},
			},
		},
		{
			name: "cookie shard with 3 decimal precision",
			protoService: &pb.Service{
				Id: "backend",
				Split: &pb.TrafficSplit{
					ShardBy: pb.TrafficSplit_COOKIE,
					Allocations: map[string]float64{
						"v1": 0.8,
						"v2": 0.2,
					},
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec{
				ResourceID: ptr.To("backend"),
				Split: &krmappenginev1alpha1.AppEngineServiceSplitTrafficSplit{
					ShardBy: ptr.To("COOKIE"),
					Allocations: map[string]string{
						"v1": "0.8",
						"v2": "0.2",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto(mapCtx, tc.protoService)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmSpec) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmSpec)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoService) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoService)
			}
		})
	}
}

func TestAppEngineStandardAppVersionSpecMapping(t *testing.T) {
	tests := []struct {
		name         string
		protoVersion *pb.Version
		krmSpec      *krmappenginev1alpha1.AppEngineStandardAppVersionSpec
	}{
		{
			name:         "nil",
			protoVersion: nil,
			krmSpec:      nil,
		},
		{
			name: "automatic scaling and handlers",
			protoVersion: &pb.Version{
				Runtime:           "nodejs18",
				InstanceClass:     "F2",
				Threadsafe:        true,
				AppEngineApis:     true,
				RuntimeApiVersion: "1",
				ServiceAccount:    "sa@my-project.iam.gserviceaccount.com",
				Scaling: &pb.Version_AutomaticScaling{
					AutomaticScaling: &pb.AutomaticScaling{
						MaxConcurrentRequests: 10,
						MaxIdleInstances:      5,
						MinIdleInstances:      1,
						MaxPendingLatency:     durationpb.New(3500 * time.Millisecond),
						MinPendingLatency:     durationpb.New(1000 * time.Millisecond),
						StandardSchedulerSettings: &pb.StandardSchedulerSettings{
							MaxInstances:                10,
							MinInstances:                2,
							TargetCpuUtilization:        0.8,
							TargetThroughputUtilization: 0.9,
						},
					},
				},
				InboundServices: []pb.InboundServiceType{
					pb.InboundServiceType_INBOUND_SERVICE_WARMUP,
				},
				Entrypoint: &pb.Entrypoint{
					Command: &pb.Entrypoint_Shell{
						Shell: "node server.js",
					},
				},
				Handlers: []*pb.UrlMap{
					{
						UrlRegex:       "/static/.*",
						SecurityLevel:  pb.SecurityLevel_SECURE_ALWAYS,
						Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
						AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
						HandlerType: &pb.UrlMap_StaticFiles{
							StaticFiles: &pb.StaticFilesHandler{
								Path:                "public",
								UploadPathRegex:     "public/.*",
								HttpHeaders:         map[string]string{"Cache-Control": "public, max-age=3600"},
								MimeType:            "text/html",
								Expiration:          durationpb.New(3600 * time.Second),
								RequireMatchingFile: true,
								ApplicationReadable: true,
							},
						},
					},
					{
						UrlRegex:                 "/api/.*",
						RedirectHttpResponseCode: pb.UrlMap_REDIRECT_HTTP_RESPONSE_CODE_301,
						HandlerType: &pb.UrlMap_Script{
							Script: &pb.ScriptHandler{
								ScriptPath: "auto",
							},
						},
					},
				},
				Libraries: []*pb.Library{
					{
						Name:    "django",
						Version: "1.11",
					},
				},
				EnvVariables: map[string]string{
					"NODE_ENV": "production",
				},
				Deployment: &pb.Deployment{
					Files: map[string]*pb.FileInfo{
						"app.js": {
							SourceUrl: "https://storage.googleapis.com/my-bucket/app.js",
							Sha1Sum:   "abc12345",
						},
					},
					Zip: &pb.ZipInfo{
						SourceUrl:  "https://storage.googleapis.com/my-bucket/app.zip",
						FilesCount: 5,
					},
				},
				VpcAccessConnector: &pb.VpcAccessConnector{
					Name:          "projects/my-project/locations/us-central1/connectors/c1",
					EgressSetting: pb.VpcAccessConnector_ALL_TRAFFIC,
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineStandardAppVersionSpec{
				Runtime:           "nodejs18",
				InstanceClass:     ptr.To("F2"),
				Threadsafe:        ptr.To(true),
				AppEngineApis:     ptr.To(true),
				RuntimeAPIVersion: ptr.To("1"),
				ServiceAccount:    ptr.To("sa@my-project.iam.gserviceaccount.com"),
				AutomaticScaling: &krmappenginev1alpha1.AutomaticScaling{
					MaxConcurrentRequests: ptr.To(10),
					MaxIdleInstances:      ptr.To(5),
					MinIdleInstances:      ptr.To(1),
					MaxPendingLatency:     ptr.To("3.5s"),
					MinPendingLatency:     ptr.To("1s"),
					StandardSchedulerSettings: &krmappenginev1alpha1.StandardSchedulerSettings{
						MaxInstances:                ptr.To(10),
						MinInstances:                ptr.To(2),
						TargetCPUUtilization:        ptr.To(0.8),
						TargetThroughputUtilization: ptr.To(0.9),
					},
				},
				InboundServices: []string{"INBOUND_SERVICE_WARMUP"},
				Entrypoint: &krmappenginev1alpha1.Entrypoint{
					Shell: "node server.js",
				},
				Handlers: []krmappenginev1alpha1.URLMap{
					{
						URLRegex:       ptr.To("/static/.*"),
						SecurityLevel:  ptr.To("SECURE_ALWAYS"),
						Login:          ptr.To("LOGIN_OPTIONAL"),
						AuthFailAction: ptr.To("AUTH_FAIL_ACTION_REDIRECT"),
						StaticFiles: &krmappenginev1alpha1.StaticFilesHandler{
							Path:                ptr.To("public"),
							UploadPathRegex:     ptr.To("public/.*"),
							HTTPHeaders:         map[string]string{"Cache-Control": "public, max-age=3600"},
							MimeType:            ptr.To("text/html"),
							Expiration:          ptr.To("1h0m0s"),
							RequireMatchingFile: ptr.To(true),
							ApplicationReadable: ptr.To(true),
						},
					},
					{
						URLRegex:                 ptr.To("/api/.*"),
						RedirectHTTPResponseCode: ptr.To("REDIRECT_HTTP_RESPONSE_CODE_301"),
						Script: &krmappenginev1alpha1.ScriptHandler{
							ScriptPath: "auto",
						},
					},
				},
				Libraries: []krmappenginev1alpha1.Library{
					{
						Name:    ptr.To("django"),
						Version: ptr.To("1.11"),
					},
				},
				EnvVariables: map[string]string{
					"NODE_ENV": "production",
				},
				Deployment: &krmappenginev1alpha1.Deployment{
					Files: []krmappenginev1alpha1.DeploymentFile{
						{
							Name:      "app.js",
							SourceUrl: "https://storage.googleapis.com/my-bucket/app.js",
							Sha1Sum:   ptr.To("abc12345"),
						},
					},
					Zip: &krmappenginev1alpha1.DeploymentZip{
						SourceUrl:  "https://storage.googleapis.com/my-bucket/app.zip",
						FilesCount: ptr.To(5),
					},
				},
				VPCAccessConnector: &krmappenginev1alpha1.VPCAccessConnector{
					Name:          "projects/my-project/locations/us-central1/connectors/c1",
					EgressSetting: ptr.To("ALL_TRAFFIC"),
				},
			},
		},
		{
			name: "basic scaling",
			protoVersion: &pb.Version{
				Runtime: "python39",
				Scaling: &pb.Version_BasicScaling{
					BasicScaling: &pb.BasicScaling{
						IdleTimeout:  durationpb.New(900 * time.Second),
						MaxInstances: 10,
					},
				},
				Entrypoint: &pb.Entrypoint{
					Command: &pb.Entrypoint_Shell{
						Shell: "gunicorn -b :$PORT main:app",
					},
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineStandardAppVersionSpec{
				Runtime: "python39",
				BasicScaling: &krmappenginev1alpha1.BasicScaling{
					IdleTimeout:  ptr.To("15m0s"),
					MaxInstances: 10,
				},
				Entrypoint: &krmappenginev1alpha1.Entrypoint{
					Shell: "gunicorn -b :$PORT main:app",
				},
			},
		},
		{
			name: "manual scaling",
			protoVersion: &pb.Version{
				Runtime: "java17",
				Scaling: &pb.Version_ManualScaling{
					ManualScaling: &pb.ManualScaling{
						Instances: 3,
					},
				},
				Entrypoint: &pb.Entrypoint{
					Command: &pb.Entrypoint_Shell{
						Shell: "java -jar target/app.jar",
					},
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineStandardAppVersionSpec{
				Runtime: "java17",
				ManualScaling: &krmappenginev1alpha1.ManualScaling{
					Instances: 3,
				},
				Entrypoint: &krmappenginev1alpha1.Entrypoint{
					Shell: "java -jar target/app.jar",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineStandardAppVersionSpec_v1alpha1_FromProto(mapCtx, tc.protoVersion)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if diff := cmp.Diff(tc.krmSpec, krm); diff != "" {
				t.Fatalf("FromProto mismatch (-want +got):\n%s", diff)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineStandardAppVersionSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if diff := cmp.Diff(tc.protoVersion, proto, protocmp.Transform()); diff != "" {
				t.Fatalf("ToProto mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAppEngineFlexibleAppVersionSpecMapping(t *testing.T) {
	tests := []struct {
		name         string
		protoVersion *pb.Version
		krmSpec      *krmappenginev1alpha1.AppEngineFlexibleAppVersionSpec
	}{
		{
			name:         "nil",
			protoVersion: nil,
			krmSpec:      nil,
		},
		{
			name: "basic fields",
			protoVersion: &pb.Version{
				Id:                        "v1",
				Runtime:                   "nodejs",
				RuntimeChannel:            "custom",
				RuntimeApiVersion:         "1",
				RuntimeMainExecutablePath: "server.js",
				InstanceClass:             "B1",
				ServiceAccount:            "test@appspot.gserviceaccount.com",
				ServingStatus:             pb.ServingStatus_SERVING,
				InboundServices:           []pb.InboundServiceType{pb.InboundServiceType_INBOUND_SERVICE_WARMUP},
				BetaSettings:              map[string]string{"foo": "bar"},
				EnvVariables:              map[string]string{"ENV": "prod"},
				DefaultExpiration:         &durationpb.Duration{Seconds: 300},
				NobuildFilesRegex:         "\\.git.*",
				Scaling: &pb.Version_ManualScaling{
					ManualScaling: &pb.ManualScaling{
						Instances: 2,
					},
				},
				Network: &pb.Network{
					Name:           "default",
					SubnetworkName: "sub-1",
				},
				Resources: &pb.Resources{
					Cpu:      2,
					DiskGb:   10,
					MemoryGb: 4,
				},
				ReadinessCheck: &pb.ReadinessCheck{
					Path:             "/ready",
					FailureThreshold: 2,
					SuccessThreshold: 2,
				},
				LivenessCheck: &pb.LivenessCheck{
					Path:             "/live",
					FailureThreshold: 4,
					SuccessThreshold: 2,
				},
				Entrypoint: &pb.Entrypoint{
					Command: &pb.Entrypoint_Shell{
						Shell: "node server.js",
					},
				},
				VpcAccessConnector: &pb.VpcAccessConnector{
					Name: "projects/test/locations/us-central1/connectors/conn1",
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineFlexibleAppVersionSpec{
				ResourceID:                ptr.To("v1"),
				Runtime:                   "nodejs",
				RuntimeChannel:            ptr.To("custom"),
				RuntimeApiVersion:         ptr.To("1"),
				RuntimeMainExecutablePath: ptr.To("server.js"),
				InstanceClass:             ptr.To("B1"),
				ServiceAccount:            ptr.To("test@appspot.gserviceaccount.com"),
				ServingStatus:             ptr.To("SERVING"),
				InboundServices:           []string{"INBOUND_SERVICE_WARMUP"},
				BetaSettings:              map[string]string{"foo": "bar"},
				EnvVariables:              map[string]string{"ENV": "prod"},
				DefaultExpiration:         ptr.To("5m0s"),
				NobuildFilesRegex:         ptr.To("\\.git.*"),
				ManualScaling: &krmappenginev1alpha1.AppEngineFlexibleAppVersionManualScaling{
					Instances: 2,
				},
				Network: &krmappenginev1alpha1.AppEngineFlexibleAppVersionNetwork{
					Name:       "default",
					Subnetwork: ptr.To("sub-1"),
				},
				Resources: &krmappenginev1alpha1.AppEngineFlexibleAppVersionResources{
					Cpu:      ptr.To(2),
					DiskGb:   ptr.To(10),
					MemoryGb: ptr.To(float64(4)),
				},
				ReadinessCheck: &krmappenginev1alpha1.AppEngineFlexibleAppVersionReadinessCheck{
					Path:             "/ready",
					FailureThreshold: ptr.To(float64(2)),
					SuccessThreshold: ptr.To(float64(2)),
				},
				LivenessCheck: &krmappenginev1alpha1.AppEngineFlexibleAppVersionLivenessCheck{
					Path:             "/live",
					FailureThreshold: ptr.To(float64(4)),
					SuccessThreshold: ptr.To(float64(2)),
				},
				Entrypoint: &krmappenginev1alpha1.AppEngineFlexibleAppVersionEntrypoint{
					Shell: "node server.js",
				},
				VpcAccessConnector: &krmappenginev1alpha1.AppEngineFlexibleAppVersionVpcAccessConnector{
					Name: "projects/test/locations/us-central1/connectors/conn1",
				},
			},
		},
		{
			name: "automatic scaling and deployment",
			protoVersion: &pb.Version{
				Runtime: "custom",
				Scaling: &pb.Version_AutomaticScaling{
					AutomaticScaling: &pb.AutomaticScaling{
						MaxTotalInstances: 10,
						MinTotalInstances: 1,
						CpuUtilization: &pb.CpuUtilization{
							TargetUtilization: 0.75,
						},
						DiskUtilization: &pb.DiskUtilization{
							TargetReadBytesPerSecond: 1024,
						},
						NetworkUtilization: &pb.NetworkUtilization{
							TargetSentBytesPerSecond: 2048,
						},
						RequestUtilization: &pb.RequestUtilization{
							TargetRequestCountPerSecond: 100,
							TargetConcurrentRequests:    50,
						},
					},
				},
				Deployment: &pb.Deployment{
					Container: &pb.ContainerInfo{
						Image: "gcr.io/test/image:tag",
					},
					Zip: &pb.ZipInfo{
						SourceUrl:  "https://storage.googleapis.com/test/app.zip",
						FilesCount: 5,
					},
					Files: map[string]*pb.FileInfo{
						"file1.txt": {SourceUrl: "https://storage.googleapis.com/test/file1.txt", Sha1Sum: "abc123"},
					},
				},
				ApiConfig: &pb.ApiConfigHandler{
					Script:         "main.py",
					AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
					Login:          pb.LoginRequirement_LOGIN_REQUIRED,
					SecurityLevel:  pb.SecurityLevel_SECURE_ALWAYS,
				},
				EndpointsApiService: &pb.EndpointsApiService{
					Name:            "api.endpoints.test.cloud.goog",
					RolloutStrategy: pb.EndpointsApiService_MANAGED,
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineFlexibleAppVersionSpec{
				Runtime: "custom",
				AutomaticScaling: &krmappenginev1alpha1.AppEngineFlexibleAppVersionAutomaticScaling{
					MaxTotalInstances: ptr.To(10),
					MinTotalInstances: ptr.To(1),
					CpuUtilization: &krmappenginev1alpha1.AppEngineFlexibleAppVersionCpuUtilization{
						TargetUtilization: 0.75,
					},
					DiskUtilization: &krmappenginev1alpha1.AppEngineFlexibleAppVersionDiskUtilization{
						TargetReadBytesPerSecond: ptr.To(1024),
					},
					NetworkUtilization: &krmappenginev1alpha1.AppEngineFlexibleAppVersionNetworkUtilization{
						TargetSentBytesPerSecond: ptr.To(2048),
					},
					RequestUtilization: &krmappenginev1alpha1.AppEngineFlexibleAppVersionRequestUtilization{
						TargetRequestCountPerSecond: ptr.To("100"),
						TargetConcurrentRequests:    ptr.To(float64(50)),
					},
				},
				Deployment: &krmappenginev1alpha1.AppEngineFlexibleAppVersionDeployment{
					Container: &krmappenginev1alpha1.AppEngineFlexibleAppVersionContainer{
						Image: "gcr.io/test/image:tag",
					},
					Zip: &krmappenginev1alpha1.AppEngineFlexibleAppVersionZip{
						SourceUrl:  "https://storage.googleapis.com/test/app.zip",
						FilesCount: ptr.To(5),
					},
					Files: []krmappenginev1alpha1.AppEngineFlexibleAppVersionFiles{
						{
							Name:      "file1.txt",
							SourceUrl: "https://storage.googleapis.com/test/file1.txt",
							Sha1Sum:   ptr.To("abc123"),
						},
					},
				},
				ApiConfig: &krmappenginev1alpha1.AppEngineFlexibleAppVersionApiConfig{
					Script:         "main.py",
					AuthFailAction: ptr.To("AUTH_FAIL_ACTION_REDIRECT"),
					Login:          ptr.To("LOGIN_REQUIRED"),
					SecurityLevel:  ptr.To("SECURE_ALWAYS"),
				},
				EndpointsApiService: &krmappenginev1alpha1.AppEngineFlexibleAppVersionEndpointsApiService{
					Name:            "api.endpoints.test.cloud.goog",
					RolloutStrategy: ptr.To("MANAGED"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineFlexibleAppVersionSpec_v1alpha1_FromProto(mapCtx, tc.protoVersion)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if diff := cmp.Diff(tc.krmSpec, krm); diff != "" {
				t.Fatalf("FromProto mismatch (-want +got):\n%s", diff)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoVersion) {
				t.Fatalf("ToProto mismatch:\ngot  %+v\nwant %+v", proto, tc.protoVersion)
			}
		})
	}
}

func TestAppEngineFlexibleAppVersionStatusMapping(t *testing.T) {
	tests := []struct {
		name         string
		protoVersion *pb.Version
		krmStatus    *krmappenginev1alpha1.AppEngineFlexibleAppVersionStatus
	}{
		{
			name:         "nil",
			protoVersion: nil,
			krmStatus:    nil,
		},
		{
			name: "name populated",
			protoVersion: &pb.Version{
				Name: "apps/my-app/services/default/versions/v1",
			},
			krmStatus: &krmappenginev1alpha1.AppEngineFlexibleAppVersionStatus{
				Name: ptr.To("apps/my-app/services/default/versions/v1"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineFlexibleAppVersionStatus_v1alpha1_FromProto(mapCtx, tc.protoVersion)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmStatus) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmStatus)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineFlexibleAppVersionStatus_v1alpha1_ToProto(mapCtx, tc.krmStatus)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoVersion) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoVersion)
			}
		})
	}
}

func TestAppEngineDomainMappingSpecMapping(t *testing.T) {
	tests := []struct {
		name        string
		protoDomain *pb.DomainMapping
		krmSpec     *krmappenginev1alpha1.AppEngineDomainMappingSpec
	}{
		{
			name:        "nil",
			protoDomain: nil,
			krmSpec:     nil,
		},
		{
			name: "ssl settings automatic",
			protoDomain: &pb.DomainMapping{
				SslSettings: &pb.SslSettings{
					SslManagementType:           pb.SslSettings_AUTOMATIC,
					PendingManagedCertificateId: "cert-pending-1",
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineDomainMappingSpec{
				SSLSettings: &krmappenginev1alpha1.DomainmappingSslSettings{
					SSLManagementType:           "AUTOMATIC",
					PendingManagedCertificateID: ptr.To("cert-pending-1"),
				},
			},
		},
		{
			name: "ssl settings manual",
			protoDomain: &pb.DomainMapping{
				SslSettings: &pb.SslSettings{
					CertificateId:     "cert-12345",
					SslManagementType: pb.SslSettings_MANUAL,
				},
			},
			krmSpec: &krmappenginev1alpha1.AppEngineDomainMappingSpec{
				SSLSettings: &krmappenginev1alpha1.DomainmappingSslSettings{
					CertificateID:     ptr.To("cert-12345"),
					SSLManagementType: "MANUAL",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineDomainMappingSpec_v1alpha1_FromProto(mapCtx, tc.protoDomain)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmSpec) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmSpec)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineDomainMappingSpec_v1alpha1_ToProto(mapCtx, tc.krmSpec)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoDomain) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoDomain)
			}
		})
	}
}

func TestAppEngineDomainMappingStatusMapping(t *testing.T) {
	tests := []struct {
		name        string
		protoDomain *pb.DomainMapping
		krmStatus   *krmappenginev1alpha1.AppEngineDomainMappingStatus
	}{
		{
			name:        "nil",
			protoDomain: nil,
			krmStatus:   nil,
		},
		{
			name: "name and resource records",
			protoDomain: &pb.DomainMapping{
				Name: "apps/my-project/domainMappings/example.com",
				ResourceRecords: []*pb.ResourceRecord{
					{
						Name:   "www",
						Rrdata: "example.com",
						Type:   pb.ResourceRecord_CNAME,
					},
					{
						Name:   "@",
						Rrdata: "192.0.2.1",
						Type:   pb.ResourceRecord_A,
					},
				},
			},
			krmStatus: &krmappenginev1alpha1.AppEngineDomainMappingStatus{
				Name: ptr.To("apps/my-project/domainMappings/example.com"),
				ResourceRecords: []krmappenginev1alpha1.DomainmappingResourceRecords{
					{
						Name:   ptr.To("www"),
						Rrdata: ptr.To("example.com"),
						Type:   ptr.To("CNAME"),
					},
					{
						Name:   ptr.To("@"),
						Rrdata: ptr.To("192.0.2.1"),
						Type:   ptr.To("A"),
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			krm := AppEngineDomainMappingStatus_v1alpha1_FromProto(mapCtx, tc.protoDomain)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in FromProto: %v", err)
			}
			if !reflect.DeepEqual(krm, tc.krmStatus) {
				t.Fatalf("FromProto mismatch: got %+v, want %+v", krm, tc.krmStatus)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineDomainMappingStatus_v1alpha1_ToProto(mapCtx, tc.krmStatus)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoDomain) {
				t.Fatalf("ToProto mismatch: got %+v, want %+v", proto, tc.protoDomain)
			}
		})
	}
}
