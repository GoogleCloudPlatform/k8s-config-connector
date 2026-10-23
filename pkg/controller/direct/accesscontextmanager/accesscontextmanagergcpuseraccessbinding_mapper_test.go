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

package accesscontextmanager

import (
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accesscontextmanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestAccessContextManagerGCPUserAccessBindingSpecMappers(t *testing.T) {
	groupKey := "01d520gv4vjcrht"
	accessLevels := []string{"accessPolicies/9522/accessLevels/device_trusted"}
	orgExternal := "organizations/123456789"

	spec := &krm.AccessContextManagerGCPUserAccessBindingSpec{
		GroupKey:     groupKey,
		AccessLevels: accessLevels,
		OrganizationRef: &krm.OrganizationRef{
			External: orgExternal,
		},
	}

	mapCtx := &direct.MapContext{}
	proto := AccessContextManagerGCPUserAccessBindingSpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting to proto: %v", mapCtx.Err())
	}

	if proto.GetGroupKey() != groupKey {
		t.Errorf("GetGroupKey() = %q, want %q", proto.GetGroupKey(), groupKey)
	}
	if !reflect.DeepEqual(proto.GetAccessLevels(), accessLevels) {
		t.Errorf("GetAccessLevels() = %v, want %v", proto.GetAccessLevels(), accessLevels)
	}

	// Test FromProto
	roundtripSpec := AccessContextManagerGCPUserAccessBindingSpec_v1alpha1_FromProto(mapCtx, proto)
	if mapCtx.Err() != nil {
		t.Fatalf("unexpected error converting from proto: %v", mapCtx.Err())
	}
	if roundtripSpec.GroupKey != groupKey {
		t.Errorf("roundtripSpec.GroupKey = %q, want %q", roundtripSpec.GroupKey, groupKey)
	}
	if !reflect.DeepEqual(roundtripSpec.AccessLevels, accessLevels) {
		t.Errorf("roundtripSpec.AccessLevels = %v, want %v", roundtripSpec.AccessLevels, accessLevels)
	}

	// Test nil inputs
	if got := AccessContextManagerGCPUserAccessBindingSpec_v1alpha1_ToProto(mapCtx, nil); got != nil {
		t.Errorf("ToProto(nil) = %v, want nil", got)
	}
	if got := AccessContextManagerGCPUserAccessBindingSpec_v1alpha1_FromProto(mapCtx, nil); got != nil {
		t.Errorf("FromProto(nil) = %v, want nil", got)
	}
}
