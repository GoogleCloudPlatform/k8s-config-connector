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

package designcenter

import (
	"testing"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	saasservicemgmtv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/saasservicemgmt/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
)

func TestApplicationTemplate_SaaSRuntimeContext_RoundTrip(t *testing.T) {
	ctx := &direct.MapContext{}

	inProto := &pb.SaaSRuntimeContext{
		SaasNames: []string{
			"projects/test-project/locations/us-central1/saas/test-saas-1",
			"projects/test-project/locations/us-central1/saas/test-saas-2",
		},
	}

	krmObj := ApplicationTemplate_SaaSRuntimeContext_FromProto(ctx, inProto)
	wantKRM := &krm.ApplicationTemplate_SaaSRuntimeContext{
		SaasRefs: []saasservicemgmtv1alpha1.SaaSServiceMgmtSaaSRef{
			{External: "projects/test-project/locations/us-central1/saas/test-saas-1"},
			{External: "projects/test-project/locations/us-central1/saas/test-saas-2"},
		},
	}

	if diff := cmp.Diff(wantKRM, krmObj); diff != "" {
		t.Fatalf("ApplicationTemplate_SaaSRuntimeContext_FromProto() mismatch (-want +got):\n%s", diff)
	}

	outProto := ApplicationTemplate_SaaSRuntimeContext_ToProto(ctx, krmObj)
	if diff := cmp.Diff(inProto.SaasNames, outProto.SaasNames); diff != "" {
		t.Fatalf("ApplicationTemplate_SaaSRuntimeContext_ToProto() mismatch (-want +got):\n%s", diff)
	}
}
