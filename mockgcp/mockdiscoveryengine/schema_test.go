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

package mockdiscoveryengine

import (
	"context"
	"testing"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setupMockSchemaService() (*MockService, *dataStoreService, *schemaService) {
	store := &mockProjectStore{
		projectsByID: map[string]*projects.ProjectData{
			"test-project": {
				Number: 1234567890,
				ID:     "test-project",
			},
		},
		projectsByNumber: map[string]*projects.ProjectData{
			"1234567890": {
				Number: 1234567890,
				ID:     "test-project",
			},
		},
	}
	env := &common.MockEnvironment{
		Projects: store,
	}
	st := storage.NewInMemoryStorage()
	s := &MockService{
		MockEnvironment: env,
		storage:         st,
		operations:      operations.NewOperationsService(st),
	}
	return s, &dataStoreService{MockService: s}, &schemaService{MockService: s}
}

func TestDataStoreSchemaLifecycle(t *testing.T) {
	ctx := context.Background()
	_, dsSvc, schemaSvc := setupMockSchemaService()

	parent := "projects/test-project/locations/global/collections/default_collection"
	ds1Name := parent + "/dataStores/ds-1"

	// 1. CreateDataStore (with SkipDefaultSchemaCreation: false) provisions default_schema and sets DefaultSchemaId = "default_schema".
	createDSOp, err := dsSvc.CreateDataStore(ctx, &pb.CreateDataStoreRequest{
		Parent:      parent,
		DataStoreId: "ds-1",
		DataStore: &pb.DataStore{
			DisplayName: "Data Store 1",
		},
		SkipDefaultSchemaCreation: false,
	})
	if err != nil {
		t.Fatalf("CreateDataStore failed: %v", err)
	}
	if !createDSOp.GetDone() {
		t.Errorf("expected CreateDataStore operation to be done")
	}

	ds, err := dsSvc.GetDataStore(ctx, &pb.GetDataStoreRequest{Name: ds1Name})
	if err != nil {
		t.Fatalf("GetDataStore failed: %v", err)
	}
	if got, want := ds.GetDefaultSchemaId(), "default_schema"; got != want {
		t.Errorf("GetDataStore defaultSchemaId = %q, want %q", got, want)
	}

	defaultSchemaName := ds1Name + "/schemas/default_schema"
	if _, err := schemaSvc.GetSchema(ctx, &pb.GetSchemaRequest{Name: defaultSchemaName}); err != nil {
		t.Fatalf("GetSchema default_schema failed: %v", err)
	}

	// 2. Calling CreateSchema with SchemaId: "custom_schema" while default_schema exists fails with codes.AlreadyExists.
	_, err = schemaSvc.CreateSchema(ctx, &pb.CreateSchemaRequest{
		Parent:   ds1Name,
		SchemaId: "custom_schema",
		Schema: &pb.Schema{
			Schema: &pb.Schema_JsonSchema{
				JsonSchema: `{"type":"object"}`,
			},
		},
	})
	if status.Code(err) != codes.AlreadyExists {
		t.Errorf("CreateSchema with existing default_schema got status %v, want codes.AlreadyExists", status.Code(err))
	}

	// 3. Calling DeleteSchema on default_schema clears DefaultSchemaId ("") on the parent DataStore.
	delOp, err := schemaSvc.DeleteSchema(ctx, &pb.DeleteSchemaRequest{Name: defaultSchemaName})
	if err != nil {
		t.Fatalf("DeleteSchema failed: %v", err)
	}
	if !delOp.GetDone() {
		t.Errorf("expected DeleteSchema operation to be done")
	}

	dsAfterDelete, err := dsSvc.GetDataStore(ctx, &pb.GetDataStoreRequest{Name: ds1Name})
	if err != nil {
		t.Fatalf("GetDataStore after DeleteSchema failed: %v", err)
	}
	if got := dsAfterDelete.GetDefaultSchemaId(); got != "" {
		t.Errorf("GetDataStore defaultSchemaId after delete = %q, want empty", got)
	}

	// 4. Calling CreateSchema with SchemaId: "custom_schema" after default_schema is deleted succeeds,
	// and a second CreateSchema call fails with codes.AlreadyExists.
	createSchemaOp, err := schemaSvc.CreateSchema(ctx, &pb.CreateSchemaRequest{
		Parent:   ds1Name,
		SchemaId: "custom_schema",
		Schema: &pb.Schema{
			Schema: &pb.Schema_JsonSchema{
				JsonSchema: `{"type":"object"}`,
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateSchema after default_schema deleted failed: %v", err)
	}
	if !createSchemaOp.GetDone() {
		t.Errorf("expected CreateSchema operation to be done")
	}

	customSchemaName := ds1Name + "/schemas/custom_schema"
	if _, err := schemaSvc.GetSchema(ctx, &pb.GetSchemaRequest{Name: customSchemaName}); err != nil {
		t.Fatalf("GetSchema custom_schema failed: %v", err)
	}

	_, err = schemaSvc.CreateSchema(ctx, &pb.CreateSchemaRequest{
		Parent:   ds1Name,
		SchemaId: "another_schema",
		Schema: &pb.Schema{
			Schema: &pb.Schema_JsonSchema{
				JsonSchema: `{"type":"object"}`,
			},
		},
	})
	if status.Code(err) != codes.AlreadyExists {
		t.Errorf("second CreateSchema got status %v, want codes.AlreadyExists", status.Code(err))
	}

	// 5. Test on a DataStore created with SkipDefaultSchemaCreation: true.
	ds2Name := parent + "/dataStores/ds-2"
	_, err = dsSvc.CreateDataStore(ctx, &pb.CreateDataStoreRequest{
		Parent:      parent,
		DataStoreId: "ds-2",
		DataStore: &pb.DataStore{
			DisplayName: "Data Store 2",
		},
		SkipDefaultSchemaCreation: true,
	})
	if err != nil {
		t.Fatalf("CreateDataStore ds-2 failed: %v", err)
	}

	ds2, err := dsSvc.GetDataStore(ctx, &pb.GetDataStoreRequest{Name: ds2Name})
	if err != nil {
		t.Fatalf("GetDataStore ds-2 failed: %v", err)
	}
	if got := ds2.GetDefaultSchemaId(); got != "" {
		t.Errorf("GetDataStore ds-2 defaultSchemaId = %q, want empty", got)
	}

	// Calling CreateSchema on ds-2 succeeds.
	_, err = schemaSvc.CreateSchema(ctx, &pb.CreateSchemaRequest{
		Parent:   ds2Name,
		SchemaId: "custom_schema",
		Schema: &pb.Schema{
			Schema: &pb.Schema_JsonSchema{
				JsonSchema: `{"type":"object"}`,
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateSchema on ds-2 failed: %v", err)
	}

	// Second CreateSchema call on ds-2 fails with codes.AlreadyExists.
	_, err = schemaSvc.CreateSchema(ctx, &pb.CreateSchemaRequest{
		Parent:   ds2Name,
		SchemaId: "another_schema",
		Schema: &pb.Schema{
			Schema: &pb.Schema_JsonSchema{
				JsonSchema: `{"type":"object"}`,
			},
		},
	})
	if status.Code(err) != codes.AlreadyExists {
		t.Errorf("second CreateSchema on ds-2 got status %v, want codes.AlreadyExists", status.Code(err))
	}
}
