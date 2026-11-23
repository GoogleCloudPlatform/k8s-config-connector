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

package iam_test

import (
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/iam"
)

func TestIAMPolicySpec_EtagMapping(t *testing.T) {
	mapCtx := &direct.MapContext{}
	etag := "abcdef0123A="
	krmPolicy := &krm.IAMPolicySpec{
		Etag: etag,
	}

	protoPolicy := iam.IAMPolicySpec_ToProto(mapCtx, krmPolicy)
	if string(protoPolicy.Etag) != etag {
		t.Errorf("IAMPolicySpec_ToProto Etag = %q, want %q", string(protoPolicy.Etag), etag)
	}

	convertedKRM := iam.IAMPolicySpec_FromProto(mapCtx, protoPolicy)
	if convertedKRM.Etag != etag {
		t.Errorf("IAMPolicySpec_FromProto Etag = %q, want %q", convertedKRM.Etag, etag)
	}
}
