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

package mocksql

import (
	"context"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type staticProjectStore struct{}

func (s *staticProjectStore) GetProject(p *projects.ProjectName) (*projects.ProjectData, error) {
	return &projects.ProjectData{ID: p.ProjectID, Number: 123456789}, nil
}

func (s *staticProjectStore) GetProjectByID(projectID string) (*projects.ProjectData, error) {
	return &projects.ProjectData{ID: projectID, Number: 123456789}, nil
}

func (s *staticProjectStore) GetProjectByNumber(projectNumber string) (*projects.ProjectData, error) {
	return &projects.ProjectData{ID: "test-project", Number: 123456789}, nil
}

func (s *staticProjectStore) GetProjectByIDOrNumber(idOrNum string) (*projects.ProjectData, error) {
	return &projects.ProjectData{ID: "test-project", Number: 123456789}, nil
}

func newTestSQLInstancesService() *sqlInstancesService {
	mem := storage.NewInMemoryStorage()
	ms := &MockService{
		storage:    mem,
		projects:   &staticProjectStore{},
		operations: &operations{storage: mem},
	}
	ms.users = &sqlUsersService{MockService: ms}
	return &sqlInstancesService{MockService: ms}
}

func TestSQLInstancePSC_LifecycleAndValidation(t *testing.T) {
	ctx := context.Background()
	svc := newTestSQLInstancesService()

	const (
		project  = "test-project"
		instance = "psc-inst"
		region   = "europe-north2"
		naURI    = "projects/test-project/regions/europe-north2/networkAttachments/my-na"
	)

	// 1. Insert with networkAttachmentUri set must fail with InvalidArgument (matching live Cloud SQL).
	_, err := svc.Insert(ctx, &pb.SqlInstancesInsertRequest{
		Project: project,
		Body: &pb.DatabaseInstance{
			Name:            instance,
			Region:          region,
			DatabaseVersion: pb.SqlDatabaseVersion_POSTGRES_16,
			Settings: &pb.Settings{
				Tier: "db-custom-1-3840",
				IpConfiguration: &pb.IpConfiguration{
					Ipv4Enabled: wrapperspb.Bool(false),
					PscConfig: &pb.PscConfig{
						PscEnabled:           PtrTo(true),
						NetworkAttachmentUri: PtrTo(naURI),
					},
				},
			},
		},
	})
	if status.Code(err) != codes.InvalidArgument || !strings.Contains(err.Error(), "can not be assigned with instance creation") {
		t.Fatalf("expected InvalidArgument on creation with networkAttachmentUri, got: %v", err)
	}

	// 2. Insert PSC-enabled instance without networkAttachmentUri succeeds and populates PSC output fields.
	if _, err := svc.Insert(ctx, &pb.SqlInstancesInsertRequest{
		Project: project,
		Body: &pb.DatabaseInstance{
			Name:            instance,
			Region:          region,
			DatabaseVersion: pb.SqlDatabaseVersion_POSTGRES_16,
			Settings: &pb.Settings{
				Tier:    "db-custom-1-3840",
				Edition: pb.Settings_ENTERPRISE,
				IpConfiguration: &pb.IpConfiguration{
					Ipv4Enabled: wrapperspb.Bool(false),
					PscConfig: &pb.PscConfig{
						PscEnabled:              PtrTo(true),
						AllowedConsumerProjects: []string{project},
					},
				},
			},
		},
	}); err != nil {
		t.Fatalf("Insert PSC instance failed: %v", err)
	}

	got, err := svc.Get(ctx, &pb.SqlInstancesGetRequest{Project: project, Instance: instance})
	if err != nil {
		t.Fatalf("Get after Insert failed: %v", err)
	}
	if got.GetPscServiceAttachmentLink() == "" || got.GetDnsName() == "" {
		t.Fatalf("expected non-empty pscServiceAttachmentLink and dnsName, got link=%q dns=%q", got.GetPscServiceAttachmentLink(), got.GetDnsName())
	}

	// 3. Update (PUT) attaching networkAttachmentUri succeeds.
	if _, err := svc.Update(ctx, &pb.SqlInstancesUpdateRequest{
		Project:  project,
		Instance: instance,
		Body: &pb.DatabaseInstance{
			Name:            instance,
			Region:          region,
			InstanceType:    pb.SqlInstanceType_CLOUD_SQL_INSTANCE,
			DatabaseVersion: pb.SqlDatabaseVersion_POSTGRES_16,
			Settings: &pb.Settings{
				Tier:    "db-custom-1-3840",
				Edition: pb.Settings_ENTERPRISE,
				IpConfiguration: &pb.IpConfiguration{
					Ipv4Enabled: wrapperspb.Bool(false),
					PscConfig: &pb.PscConfig{
						PscEnabled:              PtrTo(true),
						AllowedConsumerProjects: []string{project},
						NetworkAttachmentUri:    PtrTo(naURI),
					},
				},
			},
		},
	}); err != nil {
		t.Fatalf("Update attach networkAttachmentUri failed: %v", err)
	}

	gotAttached, err := svc.Get(ctx, &pb.SqlInstancesGetRequest{Project: project, Instance: instance})
	if err != nil {
		t.Fatalf("Get after attach failed: %v", err)
	}
	if gotAttached.GetSettings().GetIpConfiguration().GetPscConfig().GetNetworkAttachmentUri() != naURI {
		t.Fatalf("expected networkAttachmentUri %q, got %q", naURI, gotAttached.GetSettings().GetIpConfiguration().GetPscConfig().GetNetworkAttachmentUri())
	}

	// 4. Cross-region networkAttachmentUri on Update is rejected with InvalidArgument.
	_, err = svc.Update(ctx, &pb.SqlInstancesUpdateRequest{
		Project:  project,
		Instance: instance,
		Body: &pb.DatabaseInstance{
			Name:            instance,
			Region:          region,
			InstanceType:    pb.SqlInstanceType_CLOUD_SQL_INSTANCE,
			DatabaseVersion: pb.SqlDatabaseVersion_POSTGRES_16,
			Settings: &pb.Settings{
				Tier:    "db-custom-1-3840",
				Edition: pb.Settings_ENTERPRISE,
				IpConfiguration: &pb.IpConfiguration{
					Ipv4Enabled: wrapperspb.Bool(false),
					PscConfig: &pb.PscConfig{
						PscEnabled:           PtrTo(true),
						NetworkAttachmentUri: PtrTo("projects/test-project/regions/europe-west1/networkAttachments/wrong-region"),
					},
				},
			},
		},
	})
	if status.Code(err) != codes.InvalidArgument || !strings.Contains(err.Error(), "same region") {
		t.Fatalf("expected cross-region InvalidArgument error, got: %v", err)
	}

	// 5. Disabling pscEnabled on a PSC-only instance is rejected with InvalidArgument.
	_, err = svc.Update(ctx, &pb.SqlInstancesUpdateRequest{
		Project:  project,
		Instance: instance,
		Body: &pb.DatabaseInstance{
			Name:            instance,
			Region:          region,
			InstanceType:    pb.SqlInstanceType_CLOUD_SQL_INSTANCE,
			DatabaseVersion: pb.SqlDatabaseVersion_POSTGRES_16,
			Settings: &pb.Settings{
				Tier:    "db-custom-1-3840",
				Edition: pb.Settings_ENTERPRISE,
				IpConfiguration: &pb.IpConfiguration{
					Ipv4Enabled: wrapperspb.Bool(false),
					PscConfig: &pb.PscConfig{
						PscEnabled: PtrTo(false),
					},
				},
			},
		},
	})
	if status.Code(err) != codes.InvalidArgument || !strings.Contains(err.Error(), "psc connectivity can not be disabled") {
		t.Fatalf("expected disabling PSC on PSC-only instance to fail with InvalidArgument, got: %v", err)
	}

	// 6. Update (PUT) detaching networkAttachmentUri (nil / JSON null) clears it cleanly.
	if _, err := svc.Update(ctx, &pb.SqlInstancesUpdateRequest{
		Project:  project,
		Instance: instance,
		Body: &pb.DatabaseInstance{
			Name:            instance,
			Region:          region,
			InstanceType:    pb.SqlInstanceType_CLOUD_SQL_INSTANCE,
			DatabaseVersion: pb.SqlDatabaseVersion_POSTGRES_16,
			Settings: &pb.Settings{
				Tier:    "db-custom-1-3840",
				Edition: pb.Settings_ENTERPRISE,
				IpConfiguration: &pb.IpConfiguration{
					Ipv4Enabled: wrapperspb.Bool(false),
					PscConfig: &pb.PscConfig{
						PscEnabled:              PtrTo(true),
						AllowedConsumerProjects: []string{project},
						NetworkAttachmentUri:    nil,
					},
				},
			},
		},
	}); err != nil {
		t.Fatalf("Update detach networkAttachmentUri failed: %v", err)
	}

	gotDetached, err := svc.Get(ctx, &pb.SqlInstancesGetRequest{Project: project, Instance: instance})
	if err != nil {
		t.Fatalf("Get after detach failed: %v", err)
	}
	if gotDetached.GetSettings().GetIpConfiguration().GetPscConfig().GetNetworkAttachmentUri() != "" {
		t.Fatalf("expected cleared networkAttachmentUri after detach, got %q", gotDetached.GetSettings().GetIpConfiguration().GetPscConfig().GetNetworkAttachmentUri())
	}
}
