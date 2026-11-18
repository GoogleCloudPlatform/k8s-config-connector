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

package sqladmin

import (
	"context"
	"testing"

	pb "cloud.google.com/go/sql/apiv1/sqlpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sqladmin/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestBackupRunAdapter_Export(t *testing.T) {
	ctx := context.Background()

	adapter := &BackupRunAdapter{
		id: &krm.SQLAdminBackupIdentity{
			Project: "test-project",
			Backup:  "1234567890",
		},
		actual: &pb.BackupRun{
			Id:          1234567890,
			Instance:    "test-instance",
			Location:    "us-central1",
			Description: "automated test backup",
		},
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() failed: %v", err)
	}

	instanceRefExternal, found, err := unstructured.NestedString(u.Object, "spec", "instanceRef", "external")
	if err != nil || !found {
		t.Fatalf("spec.instanceRef.external not found: %v", err)
	}

	expectedInstanceRef := "projects/test-project/instances/test-instance"
	if instanceRefExternal != expectedInstanceRef {
		t.Errorf("expected spec.instanceRef.external = %q, got %q", expectedInstanceRef, instanceRefExternal)
	}

	// Verify the exported reference format can be resolved by ResolveSQLInstanceRef
	resolved, err := refsv1beta1.ResolveSQLInstanceRef(ctx, nil, nil, &refsv1beta1.SQLInstanceRef{External: instanceRefExternal})
	if err != nil {
		t.Fatalf("ResolveSQLInstanceRef failed on exported reference: %v", err)
	}
	if resolved.ProjectID != "test-project" || resolved.SQLInstanceName != "test-instance" {
		t.Errorf("unexpected resolved SQLInstance: %+v", resolved)
	}
}
