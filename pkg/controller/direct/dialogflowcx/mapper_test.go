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

package dialogflowcx

import (
	"testing"

	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krmdialogflowcxv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/utils/ptr"
)

func TestGenericWebServiceMapping(t *testing.T) {
	mapCtx := &direct.MapContext{}

	krmObj := &krmdialogflowcxv1alpha1.GenericWebService{
		URI: "https://example.com/webhook",
		RequestHeaders: map[string]string{
			"Header1": "Value1",
		},
		AllowedCACerts: []string{"cert1", "cert2"},
	}

	protoObj := GenericWebService_v1alpha1_ToProto(mapCtx, krmObj)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting to proto: %v", mapCtx.Err())
	}

	expectedProto := &pb.Webhook_GenericWebService{
		Uri: "https://example.com/webhook",
		RequestHeaders: map[string]string{
			"Header1": "Value1",
		},
		AllowedCaCerts: [][]byte{[]byte("cert1"), []byte("cert2")},
	}

	if diff := cmp.Diff(expectedProto, protoObj, protocmp.Transform()); diff != "" {
		t.Fatalf("unexpected diff (-want +got):\n%s", diff)
	}

	roundTrippedKRM := GenericWebService_v1alpha1_FromProto(mapCtx, protoObj)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting from proto: %v", mapCtx.Err())
	}

	if diff := cmp.Diff(krmObj, roundTrippedKRM); diff != "" {
		t.Fatalf("unexpected diff (-want +got):\n%s", diff)
	}
}

func TestDialogflowCXWebhookSpec_ToProto(t *testing.T) {
	mapCtx := &direct.MapContext{}

	krmSpec := &krmdialogflowcxv1alpha1.DialogflowCXWebhookSpec{
		DisplayName: "my-webhook",
		Disabled:    ptr.To(false),
		GenericWebService: &krmdialogflowcxv1alpha1.GenericWebService{
			URI: "https://example.com/webhook",
		},
	}

	protoObj := DialogflowCXWebhookSpec_v1alpha1_ToProto(mapCtx, krmSpec)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error: %v", mapCtx.Err())
	}

	expectedProto := &pb.Webhook{
		DisplayName: "my-webhook",
		Disabled:    false,
		Webhook: &pb.Webhook_GenericWebService_{
			GenericWebService: &pb.Webhook_GenericWebService{
				Uri: "https://example.com/webhook",
			},
		},
	}

	if diff := cmp.Diff(expectedProto, protoObj, protocmp.Transform()); diff != "" {
		t.Fatalf("unexpected diff (-want +got):\n%s", diff)
	}
}
