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

package mockfinancialservices_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockfinancialservices"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func TestServiceCreation(t *testing.T) {
	env := &common.MockEnvironment{}
	st := storage.NewInMemoryStorage()
	svc := mockfinancialservices.New(env, st)
	if svc == nil {
		t.Fatalf("expected New() to return a non-nil MockService")
	}
	expectedHosts := svc.ExpectedHosts()
	if len(expectedHosts) == 0 || expectedHosts[0] != "financialservices.googleapis.com" {
		t.Errorf("unexpected expected hosts: %v", expectedHosts)
	}
}
