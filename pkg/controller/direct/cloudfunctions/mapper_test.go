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

package cloudfunctions

import (
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudfunctions/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestCloudFunctionsFunctionSpec_RoundTrip(t *testing.T) {
	testCases := []struct {
		name string
		in   *krm.CloudFunctionsFunctionSpec
	}{
		{
			name: "Basic HTTP Function",
			in: &krm.CloudFunctionsFunctionSpec{
				Runtime:          "nodejs16",
				EntryPoint:       direct.LazyPtr("helloWorld"),
				SourceArchiveURL: direct.LazyPtr("gs://my-bucket/source.zip"),
				HTTPSTrigger: &krm.FunctionHttpsTrigger{
					SecurityLevel: direct.LazyPtr("SECURE_ALWAYS"),
				},
				AvailableMemoryMb:         direct.LazyPtr(int64(256)),
				Timeout:                   direct.LazyPtr("1m0s"),
				EnvironmentVariables:      map[string]string{},
				BuildEnvironmentVariables: map[string]string{},
			},
		},
		{
			name: "Event Driven Function",
			in: &krm.CloudFunctionsFunctionSpec{
				Runtime:    "python39",
				EntryPoint: direct.LazyPtr("handler"),
				SourceRepository: &krm.FunctionSourceRepository{
					URL: "https://source.developers.google.com/projects/p1/repos/r1/revisions/rev1/paths/p1",
				},
				EventTrigger: &krm.FunctionEventTrigger{
					EventType: "google.pubsub.topic.publish",
					ResourceRef: krm.FunctionResourceRef{
						External: "projects/p1/topics/t1",
					},
					FailurePolicy: direct.LazyPtr(true),
				},
				MaxInstances:              direct.LazyPtr(int64(10)),
				MinInstances:              direct.LazyPtr(int64(1)),
				EnvironmentVariables:      map[string]string{},
				BuildEnvironmentVariables: map[string]string{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			proto := CloudFunctionsFunctionSpec_ToProto(mapCtx, tc.in)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("ToProto failed: %v", err)
			}

			out := CloudFunctionsFunctionSpec_FromProto(mapCtx, proto)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("FromProto failed: %v", err)
			}

			if !reflect.DeepEqual(tc.in, out) {
				t.Errorf("Roundtrip failed.\nIn: %+v\nOut: %+v", tc.in, out)
			}
		})
	}
}
