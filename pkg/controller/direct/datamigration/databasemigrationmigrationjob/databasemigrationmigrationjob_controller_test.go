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

package databasemigrationmigrationjob

import (
	"context"
	"testing"

	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/google/go-cmp/cmp"
)

func TestCompareMigrationJob_ExcludesImmutableFields(t *testing.T) {
	ctx := context.Background()

	actual := &pb.MigrationJob{
		Name:                "projects/p1/locations/l1/migrationJobs/mj1",
		DisplayName:         "Old Display Name",
		Source:              "projects/p1/locations/l1/connectionProfiles/cp-src",
		Destination:         "projects/p1/locations/l1/connectionProfiles/cp-dst",
		Type:                pb.MigrationJob_ONE_TIME,
		SourceDatabase:      &pb.DatabaseType{Engine: pb.DatabaseEngine_MYSQL},
		DestinationDatabase: &pb.DatabaseType{Engine: pb.DatabaseEngine_MYSQL},
	}

	desired := &pb.MigrationJob{
		Name:                "projects/p1/locations/l1/migrationJobs/mj1",
		DisplayName:         "New Display Name",
		Source:              "projects/p1/locations/l1/connectionProfiles/cp-src-new",
		Destination:         "projects/p1/locations/l1/connectionProfiles/cp-dst-new",
		Type:                pb.MigrationJob_CONTINUOUS,
		SourceDatabase:      &pb.DatabaseType{Engine: pb.DatabaseEngine_POSTGRESQL},
		DestinationDatabase: &pb.DatabaseType{Engine: pb.DatabaseEngine_POSTGRESQL},
	}

	diffs, updateMask, err := compareMigrationJob(ctx, actual, desired)
	if err != nil {
		t.Fatalf("compareMigrationJob failed: %v", err)
	}

	if !diffs.HasDiff() {
		t.Errorf("expected diffs, got none")
	}

	expectedPaths := []string{"display_name"}
	if diff := cmp.Diff(expectedPaths, updateMask.GetPaths()); diff != "" {
		t.Errorf("updateMask mismatch (-want +got):\n%s", diff)
	}
}

func TestCompareMigrationJob_NoDiff(t *testing.T) {
	ctx := context.Background()

	actual := &pb.MigrationJob{
		Name:        "projects/p1/locations/l1/migrationJobs/mj1",
		DisplayName: "Same Display Name",
	}

	desired := &pb.MigrationJob{
		Name:        "projects/p1/locations/l1/migrationJobs/mj1",
		DisplayName: "Same Display Name",
	}

	diffs, updateMask, err := compareMigrationJob(ctx, actual, desired)
	if err != nil {
		t.Fatalf("compareMigrationJob failed: %v", err)
	}

	if diffs.HasDiff() {
		t.Errorf("expected no diffs, got diffs")
	}

	if len(updateMask.GetPaths()) != 0 {
		t.Errorf("expected empty updateMask, got %v", updateMask.GetPaths())
	}
}
