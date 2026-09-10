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

package bigquerydataset

import (
	"slices"
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
)

func TestBigQueryDataset_AccessDiffAlone(t *testing.T) {
	desiredKRM := &krm.BigQueryDataset{
		Spec: krm.BigQueryDatasetSpec{
			Description:  directPtr("test dataset"),
			FriendlyName: directPtr("Test Dataset"),
		},
	}

	actual := &bigquery.DatasetMetadata{
		Description: "test dataset",
		Name:        "Test Dataset",
		Access: []*bigquery.AccessEntry{
			{
				Role:       bigquery.OwnerRole,
				Entity:     "projectOwners",
				EntityType: bigquery.SpecialGroupEntity,
			},
		},
	}

	// desired has only an access change (an additional reader)
	desired := cloneBigQueryDatasetMetadate(actual)
	desired.Access = []*bigquery.AccessEntry{
		{
			Role:       bigquery.OwnerRole,
			Entity:     "projectOwners",
			EntityType: bigquery.SpecialGroupEntity,
		},
		{
			Role:       bigquery.ReaderRole,
			Entity:     "google.com",
			EntityType: bigquery.DomainEntity,
		},
	}

	resource := cloneBigQueryDatasetMetadate(actual)

	updateMask, report := computeDiff(desiredKRM, desired, resource)

	// Verify that access alone is captured in the updateMask
	if len(updateMask.Paths) != 1 {
		t.Fatalf("expected updateMask.Paths to have 1 entry, got %d: %v", len(updateMask.Paths), updateMask.Paths)
	}
	if !slices.Contains(updateMask.Paths, "access") {
		t.Fatalf("expected updateMask.Paths to contain 'access', got: %v", updateMask.Paths)
	}

	// Verify that diff was reported
	if len(report.Fields) != 1 {
		t.Fatalf("expected report.Fields to have 1 entry, got %d: %v", len(report.Fields), report.Fields)
	}
	if report.Fields[0].ID != "access" {
		t.Fatalf("expected diff field ID 'access', got: %s", report.Fields[0].ID)
	}

	// Verify that resource.Access matches desired entries
	if len(resource.Access) != 2 {
		t.Fatalf("expected resource.Access to have 2 entries, got %d", len(resource.Access))
	}
}

func TestBigQueryDataset_NoDiff(t *testing.T) {
	desiredKRM := &krm.BigQueryDataset{
		Spec: krm.BigQueryDatasetSpec{
			Description: directPtr("test dataset"),
		},
	}

	actual := &bigquery.DatasetMetadata{
		Description: "test dataset",
		Access: []*bigquery.AccessEntry{
			{
				Role:       bigquery.OwnerRole,
				Entity:     "projectOwners",
				EntityType: bigquery.SpecialGroupEntity,
			},
		},
	}

	desired := cloneBigQueryDatasetMetadate(actual)
	resource := cloneBigQueryDatasetMetadate(actual)

	updateMask, report := computeDiff(desiredKRM, desired, resource)

	if len(updateMask.Paths) != 0 {
		t.Fatalf("expected empty updateMask.Paths, got: %v", updateMask.Paths)
	}
	if len(report.Fields) != 0 {
		t.Fatalf("expected empty report.Fields, got: %v", report.Fields)
	}
}

func TestBigQueryDataset_AccessDiffWithOtherFields(t *testing.T) {
	desiredKRM := &krm.BigQueryDataset{
		Spec: krm.BigQueryDatasetSpec{
			Description: directPtr("updated dataset description"),
		},
	}

	actual := &bigquery.DatasetMetadata{
		Description:            "initial description",
		DefaultTableExpiration: 3600 * time.Second,
		Access: []*bigquery.AccessEntry{
			{
				Role:       bigquery.OwnerRole,
				Entity:     "projectOwners",
				EntityType: bigquery.SpecialGroupEntity,
			},
		},
	}

	desired := cloneBigQueryDatasetMetadate(actual)
	desired.Description = "updated dataset description"
	desired.Access = []*bigquery.AccessEntry{
		{
			Role:       bigquery.OwnerRole,
			Entity:     "projectOwners",
			EntityType: bigquery.SpecialGroupEntity,
		},
		{
			Role:       bigquery.ReaderRole,
			Entity:     "google.com",
			EntityType: bigquery.DomainEntity,
		},
	}

	resource := cloneBigQueryDatasetMetadate(actual)

	updateMask, report := computeDiff(desiredKRM, desired, resource)

	if len(updateMask.Paths) != 2 {
		t.Fatalf("expected 2 updateMask paths, got %d: %v", len(updateMask.Paths), updateMask.Paths)
	}
	if !slices.Contains(updateMask.Paths, "description") {
		t.Errorf("expected updateMask to contain 'description'")
	}
	if !slices.Contains(updateMask.Paths, "access") {
		t.Errorf("expected updateMask to contain 'access'")
	}

	if len(report.Fields) != 2 {
		t.Fatalf("expected 2 diff fields, got %d", len(report.Fields))
	}
}

func directPtr[T any](v T) *T {
	return &v
}
