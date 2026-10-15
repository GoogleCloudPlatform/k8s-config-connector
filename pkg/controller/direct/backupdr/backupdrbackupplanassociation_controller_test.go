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

package backupdr

import (
	"context"
	"net/http"
	"testing"
	"time"

	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"google.golang.org/api/googleapi"
	"google.golang.org/protobuf/proto"
)

func TestNormalizeActualReferences(t *testing.T) {
	ctx := context.Background()

	cache := projects.NewProjectCache(nil, time.Hour)
	cache.InsertForTest("my-project", 123456789)
	projectMapper := projects.NewProjectMapper(cache)

	const (
		desiredBackupPlan  = "projects/my-project/locations/us-central1/backupPlans/my-plan"
		desiredInstance    = "projects/my-project/zones/us-central1-a/instances/my-instance"
		desiredSQLInstance = "projects/my-project/instances/my-sql-instance"
	)

	tests := []struct {
		name             string
		resourceType     string
		desiredResource  string
		actualBackupPlan string
		actualResource   string
		// instanceIDs are the numeric IDs of the instances known by the Compute Engine API, by name.
		instanceIDs     map[string]uint64
		instanceIDError error
		wantBackupPlan  string
		wantResource    string
		wantLookups     int
		wantErr         bool
	}{
		{
			name:             "identical references",
			actualBackupPlan: desiredBackupPlan,
			actualResource:   desiredInstance,
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     desiredInstance,
		},
		{
			name:             "project number and numeric instance ID",
			actualBackupPlan: "projects/123456789/locations/us-central1/backupPlans/my-plan",
			actualResource:   "projects/123456789/zones/us-central1-a/instances/987654321",
			instanceIDs:      map[string]uint64{"my-instance": 987654321},
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     desiredInstance,
			wantLookups:      1,
		},
		{
			name:             "project number and instance name",
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/zones/us-central1-a/instances/my-instance",
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     desiredInstance,
		},
		{
			name:             "different backup plan",
			actualBackupPlan: "projects/123456789/locations/us-central1/backupPlans/other-plan",
			actualResource:   desiredInstance,
			wantBackupPlan:   "projects/123456789/locations/us-central1/backupPlans/other-plan",
			wantResource:     desiredInstance,
		},
		{
			name:             "different backup plan location",
			actualBackupPlan: "projects/123456789/locations/us-east1/backupPlans/my-plan",
			actualResource:   desiredInstance,
			wantBackupPlan:   "projects/123456789/locations/us-east1/backupPlans/my-plan",
			wantResource:     desiredInstance,
		},
		{
			name:             "different numeric instance ID",
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/zones/us-central1-a/instances/111111111",
			instanceIDs:      map[string]uint64{"my-instance": 987654321},
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     "projects/123456789/zones/us-central1-a/instances/111111111",
			wantLookups:      1,
		},
		{
			name:             "different instance name",
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/zones/us-central1-a/instances/other-instance",
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     "projects/123456789/zones/us-central1-a/instances/other-instance",
		},
		{
			name:             "different zone",
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/zones/us-central1-b/instances/987654321",
			instanceIDs:      map[string]uint64{"my-instance": 987654321},
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     "projects/123456789/zones/us-central1-b/instances/987654321",
		},
		{
			name:             "desired instance does not exist",
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/zones/us-central1-a/instances/987654321",
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     "projects/123456789/zones/us-central1-a/instances/987654321",
			wantLookups:      1,
		},
		{
			name:             "error getting the desired instance",
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/zones/us-central1-a/instances/987654321",
			instanceIDError:  &googleapi.Error{Code: http.StatusForbidden},
			wantLookups:      1,
			wantErr:          true,
		},
		{
			name:             "unknown project number",
			actualBackupPlan: "projects/111111111/locations/us-central1/backupPlans/my-plan",
			actualResource:   desiredInstance,
			wantErr:          true,
		},
		{
			name:             "Cloud SQL instance with project number",
			resourceType:     krm.ResourceType_SQLInstance,
			desiredResource:  desiredSQLInstance,
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/instances/my-sql-instance",
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     desiredSQLInstance,
		},
		{
			name:             "different Cloud SQL instance",
			resourceType:     krm.ResourceType_SQLInstance,
			desiredResource:  desiredSQLInstance,
			actualBackupPlan: desiredBackupPlan,
			actualResource:   "projects/123456789/instances/other-sql-instance",
			wantBackupPlan:   desiredBackupPlan,
			wantResource:     "projects/123456789/instances/other-sql-instance",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resourceType := tc.resourceType
			if resourceType == "" {
				resourceType = krm.ResourceType_ComputeInstance
			}
			desiredResource := tc.desiredResource
			if desiredResource == "" {
				desiredResource = desiredInstance
			}

			desired := &pb.BackupPlanAssociation{
				ResourceType: resourceType,
				Resource:     desiredResource,
				BackupPlan:   desiredBackupPlan,
			}
			actual := &pb.BackupPlanAssociation{
				Name:         "projects/my-project/locations/us-central1/backupPlanAssociations/my-association",
				ResourceType: resourceType,
				Resource:     tc.actualResource,
				BackupPlan:   tc.actualBackupPlan,
				State:        pb.BackupPlanAssociation_ACTIVE,
			}
			original := proto.CloneOf(actual)

			lookups := 0
			getComputeInstanceID := func(ctx context.Context, id *computev1beta1.ComputeInstanceIdentity) (uint64, error) {
				lookups++
				if tc.instanceIDError != nil {
					return 0, tc.instanceIDError
				}
				instanceID, ok := tc.instanceIDs[id.Instance]
				if !ok {
					return 0, &googleapi.Error{Code: http.StatusNotFound}
				}
				return instanceID, nil
			}

			got, err := normalizeActualReferences(ctx, projectMapper, getComputeInstanceID, desired, actual)
			if lookups != tc.wantLookups {
				t.Errorf("normalizeActualReferences() looked up the instance ID %d times, want %d", lookups, tc.wantLookups)
			}
			if !proto.Equal(actual, original) {
				t.Errorf("normalizeActualReferences() modified actual, got %v, want %v", actual, original)
			}
			if tc.wantErr {
				if err == nil {
					t.Fatalf("normalizeActualReferences() expected an error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("normalizeActualReferences() unexpected error: %v", err)
			}
			if got.GetBackupPlan() != tc.wantBackupPlan {
				t.Errorf("normalizeActualReferences() BackupPlan = %q, want %q", got.GetBackupPlan(), tc.wantBackupPlan)
			}
			if got.GetResource() != tc.wantResource {
				t.Errorf("normalizeActualReferences() Resource = %q, want %q", got.GetResource(), tc.wantResource)
			}
			if got.GetName() != actual.GetName() || got.GetState() != actual.GetState() {
				t.Errorf("normalizeActualReferences() changed other fields, got %v, want %v", got, actual)
			}
		})
	}
}
