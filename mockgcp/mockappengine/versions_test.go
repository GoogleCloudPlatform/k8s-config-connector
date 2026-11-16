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

package mockappengine

import (
	"context"
	"strings"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func setupMockVersionsServer() (*MockService, *versionsServer) {
	st := storage.NewInMemoryStorage()
	s := &MockService{
		MockEnvironment: &common.MockEnvironment{},
		storage:         st,
		operations:      operations.NewOperationsService(st),
	}
	return s, &versionsServer{MockService: s}
}

func TestUpdateVersion_RejectsDisallowedMaskPath(t *testing.T) {
	ctx := context.Background()
	_, srv := setupMockVersionsServer()

	_, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:            "v1",
			Runtime:       "python311",
			InstanceClass: "F1",
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion failed: %v", err)
	}

	_, err = srv.UpdateVersion(ctx, &pb.UpdateVersionRequest{
		Name: "apps/test-app/services/default/versions/v1",
		Version: &pb.Version{
			EnvVariables: map[string]string{"FOO": "bar"},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"envVariables"},
		},
	})

	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument, got %v", err)
	}
	if !strings.Contains(err.Error(), "but got field(s): [env_variables] from the update request.") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestUpdateVersion_RejectsEmptyMask(t *testing.T) {
	ctx := context.Background()
	_, srv := setupMockVersionsServer()

	_, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:      "v1",
			Runtime: "python311",
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion failed: %v", err)
	}

	_, err = srv.UpdateVersion(ctx, &pb.UpdateVersionRequest{
		Name:       "apps/test-app/services/default/versions/v1",
		Version:    &pb.Version{InstanceClass: "F2"},
		UpdateMask: &field_mask.FieldMask{},
	})

	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument, got %v", err)
	}
	if !strings.Contains(err.Error(), "At least one field must be specified for this operation.") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestUpdateVersion_AppliesOnlyMaskedFieldsIncludingZero(t *testing.T) {
	ctx := context.Background()
	_, srv := setupMockVersionsServer()

	_, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:            "v1",
			Runtime:       "python311",
			InstanceClass: "F1",
			EnvVariables:  map[string]string{"STAGE": "create"},
			Scaling: &pb.Version_AutomaticScaling{
				AutomaticScaling: &pb.AutomaticScaling{
					MinIdleInstances: 2,
					MaxIdleInstances: 4,
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion failed: %v", err)
	}

	_, err = srv.UpdateVersion(ctx, &pb.UpdateVersionRequest{
		Name: "apps/test-app/services/default/versions/v1",
		Version: &pb.Version{
			InstanceClass: "F2",
			EnvVariables:  map[string]string{"STAGE": "unmasked-change"},
			Scaling: &pb.Version_AutomaticScaling{
				AutomaticScaling: &pb.AutomaticScaling{
					MinIdleInstances: 0,
					MaxIdleInstances: 9,
				},
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"automaticScaling.minIdleInstances"},
		},
	})
	if err != nil {
		t.Fatalf("UpdateVersion failed: %v", err)
	}

	got, err := srv.GetVersion(ctx, &pb.GetVersionRequest{
		Name: "apps/test-app/services/default/versions/v1",
		View: pb.VersionView_FULL,
	})
	if err != nil {
		t.Fatalf("GetVersion failed: %v", err)
	}

	if got.GetAutomaticScaling().GetMinIdleInstances() != 0 {
		t.Errorf("expected masked minIdleInstances=0, got %d", got.GetAutomaticScaling().GetMinIdleInstances())
	}
	if got.GetAutomaticScaling().GetMaxIdleInstances() != 4 {
		t.Errorf("expected unmasked maxIdleInstances=4, got %d", got.GetAutomaticScaling().GetMaxIdleInstances())
	}
	if got.GetInstanceClass() != "F1" {
		t.Errorf("expected unmasked instanceClass=F1, got %q", got.GetInstanceClass())
	}
	if got.GetEnvVariables()["STAGE"] != "create" {
		t.Errorf("expected unmasked envVariables[STAGE]=create, got %q", got.GetEnvVariables()["STAGE"])
	}
}

func TestUpdateVersion_PerFieldValidation(t *testing.T) {
	ctx := context.Background()
	_, srv := setupMockVersionsServer()

	_, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:            "v-manual",
			Runtime:       "python311",
			InstanceClass: "B2",
			Scaling: &pb.Version_ManualScaling{
				ManualScaling: &pb.ManualScaling{Instances: 1},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion manual failed: %v", err)
	}

	_, err = srv.UpdateVersion(ctx, &pb.UpdateVersionRequest{
		Name: "apps/test-app/services/default/versions/v-manual",
		Version: &pb.Version{
			Scaling: &pb.Version_AutomaticScaling{
				AutomaticScaling: &pb.AutomaticScaling{MinIdleInstances: 1},
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"automaticScaling.minIdleInstances"},
		},
	})

	if status.Code(err) != codes.FailedPrecondition {
		t.Fatalf("expected FailedPrecondition for automatic_scaling path on manual version, got %v", err)
	}

	_, err = srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:            "v-auto",
			Runtime:       "python311",
			InstanceClass: "F1",
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion auto failed: %v", err)
	}

	_, err = srv.UpdateVersion(ctx, &pb.UpdateVersionRequest{
		Name: "apps/test-app/services/default/versions/v-auto",
		Version: &pb.Version{
			InstanceClass: "B2",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"instanceClass"},
		},
	})

	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for B2 instance class on automatic version, got %v", err)
	}
}

func TestCreateVersion_ReplacesExistingVersionHandlers(t *testing.T) {
	ctx := context.Background()
	_, srv := setupMockVersionsServer()

	_, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:      "v1",
			Runtime: "python311",
			Handlers: []*pb.UrlMap{
				{
					UrlRegex: "/a/.*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("first CreateVersion failed: %v", err)
	}

	_, err = srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:      "v1",
			Runtime: "python311",
			Handlers: []*pb.UrlMap{
				{
					UrlRegex: "/b/.*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("second CreateVersion failed: %v", err)
	}

	got, err := srv.GetVersion(ctx, &pb.GetVersionRequest{
		Name: "apps/test-app/services/default/versions/v1",
		View: pb.VersionView_FULL,
	})
	if err != nil {
		t.Fatalf("GetVersion failed: %v", err)
	}

	if len(got.GetHandlers()) != 2 {
		t.Fatalf("expected 2 handlers ([/b/.*, .*]), got %d: %v", len(got.GetHandlers()), got.GetHandlers())
	}
	if got.GetHandlers()[0].GetUrlRegex() != "/b/.*" {
		t.Errorf("expected first handler /b/.*, got %q", got.GetHandlers()[0].GetUrlRegex())
	}
	if got.GetHandlers()[1].GetUrlRegex() != ".*" {
		t.Errorf("expected second handler .*, got %q", got.GetHandlers()[1].GetUrlRegex())
	}
}

func TestCreateVersion_PopulatesHandlerDefaultsAndAppendsRuntimeDefault(t *testing.T) {
	ctx := context.Background()
	_, srv := setupMockVersionsServer()

	_, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:      "v1",
			Runtime: "python311",
			Handlers: []*pb.UrlMap{
				{
					UrlRegex: "/api/.*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{ScriptPath: "auto"},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion failed: %v", err)
	}

	got, err := srv.GetVersion(ctx, &pb.GetVersionRequest{
		Name: "apps/test-app/services/default/versions/v1",
		View: pb.VersionView_FULL,
	})
	if err != nil {
		t.Fatalf("GetVersion failed: %v", err)
	}

	want := []*pb.UrlMap{
		{
			UrlRegex:       "/api/.*",
			Login:          pb.LoginRequirement_LOGIN_OPTIONAL,
			SecurityLevel:  pb.SecurityLevel_SECURE_OPTIONAL,
			AuthFailAction: pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT,
			HandlerType: &pb.UrlMap_Script{
				Script: &pb.ScriptHandler{ScriptPath: "auto"},
			},
		},
		{
			UrlRegex:      ".*",
			Login:         pb.LoginRequirement_LOGIN_OPTIONAL,
			SecurityLevel: pb.SecurityLevel_SECURE_OPTIONAL,
			HandlerType: &pb.UrlMap_Script{
				Script: &pb.ScriptHandler{ScriptPath: "auto"},
			},
		},
	}
	if len(got.GetHandlers()) != len(want) {
		t.Fatalf("expected %d handlers, got %d: %v", len(want), len(got.GetHandlers()), got.GetHandlers())
	}
	for i := range want {
		if !proto.Equal(got.GetHandlers()[i], want[i]) {
			t.Errorf("handler[%d] mismatch:\ngot:  %v\nwant: %v", i, got.GetHandlers()[i], want[i])
		}
	}
}

func TestVersionOperations_StartLROAndGetOperationResponse(t *testing.T) {
	ctx := context.Background()
	svc, srv := setupMockVersionsServer()

	createOp, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:            "v1",
			Runtime:       "python311",
			InstanceClass: "F1",
			Entrypoint:    &pb.Entrypoint{Command: &pb.Entrypoint_Shell{Shell: "gunicorn main:app"}},
			Deployment: &pb.Deployment{
				Files: map[string]*pb.FileInfo{
					"main.py": {SourceUrl: "https://storage.googleapis.com/staging.test-app.appspot.com/05e6871e9161335ea72fe540fb798803633b43ba"},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion failed: %v", err)
	}

	if createOp.GetDone() {
		t.Errorf("expected initial CreateVersion op Done=false, got Done=true")
	}
	polledCreate, err := svc.operations.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: createOp.GetName()})
	if err != nil {
		t.Fatalf("GetOperation(create) failed: %v", err)
	}
	if !polledCreate.GetDone() {
		t.Fatalf("expected polled create op Done=true")
	}
	createdVer := &pb.Version{}
	if err := polledCreate.GetResponse().UnmarshalTo(createdVer); err != nil {
		t.Fatalf("unmarshaling create op response: %v", err)
	}
	if len(createdVer.GetHandlers()) != 0 || createdVer.GetDeployment() != nil || createdVer.GetEntrypoint() != nil || createdVer.GetCreatedBy() != "" {
		t.Errorf("expected BASIC view without CreatedBy in LRO response, got: %v", createdVer)
	}

	updateOp, err := srv.UpdateVersion(ctx, &pb.UpdateVersionRequest{
		Name:       "apps/test-app/services/default/versions/v1",
		Version:    &pb.Version{InstanceClass: "F2"},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"instanceClass"}},
	})
	if err != nil {
		t.Fatalf("UpdateVersion failed: %v", err)
	}

	if updateOp.GetDone() {
		t.Errorf("expected initial UpdateVersion op Done=false, got Done=true")
	}
	polledUpdate, err := svc.operations.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: updateOp.GetName()})
	if err != nil {
		t.Fatalf("GetOperation(update) failed: %v", err)
	}
	updatedVer := &pb.Version{}
	if err := polledUpdate.GetResponse().UnmarshalTo(updatedVer); err != nil {
		t.Fatalf("unmarshaling update op response: %v", err)
	}
	if updatedVer.GetInstanceClass() != "F2" || len(updatedVer.GetHandlers()) != 0 || updatedVer.GetCreatedBy() != "" {
		t.Errorf("unexpected update op response version: %v", updatedVer)
	}

	deleteOp, err := srv.DeleteVersion(ctx, &pb.DeleteVersionRequest{
		Name: "apps/test-app/services/default/versions/v1",
	})
	if err != nil {
		t.Fatalf("DeleteVersion failed: %v", err)
	}

	if deleteOp.GetDone() {
		t.Errorf("expected initial DeleteVersion op Done=false, got Done=true")
	}
	polledDelete, err := svc.operations.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: deleteOp.GetName()})
	if err != nil {
		t.Fatalf("GetOperation(delete) failed: %v", err)
	}
	emptyResp := &emptypb.Empty{}
	if err := polledDelete.GetResponse().UnmarshalTo(emptyResp); err != nil {
		t.Fatalf("unmarshaling delete op response to Empty: %v", err)
	}
}

func TestGetVersion_OmitsDeploymentZip(t *testing.T) {
	ctx := context.Background()
	_, srv := setupMockVersionsServer()

	_, err := srv.CreateVersion(ctx, &pb.CreateVersionRequest{
		Parent: "apps/test-app/services/default",
		Version: &pb.Version{
			Id:      "v1",
			Runtime: "python311",
			Deployment: &pb.Deployment{
				Zip: &pb.ZipInfo{
					SourceUrl: "https://storage.googleapis.com/staging.test-app.appspot.com/app.zip",
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateVersion failed: %v", err)
	}

	got, err := srv.GetVersion(ctx, &pb.GetVersionRequest{
		Name: "apps/test-app/services/default/versions/v1",
		View: pb.VersionView_FULL,
	})
	if err != nil {
		t.Fatalf("GetVersion failed: %v", err)
	}

	if got.GetDeployment().GetZip() != nil {
		t.Errorf("expected GetVersion(FULL) to omit deployment.zip, got %v", got.GetDeployment().GetZip())
	}
}
