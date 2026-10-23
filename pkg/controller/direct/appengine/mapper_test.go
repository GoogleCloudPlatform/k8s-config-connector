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

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	krmappenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
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
				t.Fatalf("FromProto mismatch:\ngot  %+v\nwant %+v", krm, tc.krmStatus)
			}

			mapCtx = &direct.MapContext{}
			proto := AppEngineFlexibleAppVersionStatus_v1alpha1_ToProto(mapCtx, tc.krmStatus)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("unexpected error in ToProto: %v", err)
			}
			if !reflect.DeepEqual(proto, tc.protoVersion) {
				t.Fatalf("ToProto mismatch:\ngot  %+v\nwant %+v", proto, tc.protoVersion)
			}
		})
	}
}
