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

package v1alpha1

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDiscoveryEngineAssistant_GetIdentity_With6TokenParentRef(t *testing.T) {
	ctx := context.Background()

	obj := &DiscoveryEngineAssistant{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-assistant",
			Namespace: "default",
		},
		Spec: DiscoveryEngineAssistantSpec{
			ParentRef: &DiscoveryEngineEngineRef{
				External: "projects/my-project/locations/global/engines/my-engine",
			},
		},
	}

	id, err := obj.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("obj.GetIdentity() failed: %v", err)
	}
	want := "projects/my-project/locations/global/collections/default_collection/engines/my-engine/assistants/my-assistant"
	if got := id.String(); got != want {
		t.Errorf("id.String() = %q, want %q", got, want)
	}
}

func TestDiscoveryEngineAssistant_GetIdentity_With8TokenParentRef(t *testing.T) {
	ctx := context.Background()

	obj := &DiscoveryEngineAssistant{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-assistant",
			Namespace: "default",
		},
		Spec: DiscoveryEngineAssistantSpec{
			ParentRef: &DiscoveryEngineEngineRef{
				External: "projects/my-project/locations/global/collections/my-collection/engines/my-engine",
			},
		},
	}

	id, err := obj.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("obj.GetIdentity() failed: %v", err)
	}
	want := "projects/my-project/locations/global/collections/my-collection/engines/my-engine/assistants/my-assistant"
	if got := id.String(); got != want {
		t.Errorf("id.String() = %q, want %q", got, want)
	}
}
