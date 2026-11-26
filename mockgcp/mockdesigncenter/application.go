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

package mockdesigncenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apphubpb "cloud.google.com/go/apphub/apiv1/apphubpb"
	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
)

func (s *DesignCenterServer) GetApplication(ctx context.Context, req *pb.GetApplicationRequest) (*pb.Application, error) {
	name, err := s.parseApplicationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populateApplication(obj), nil
}

func (s *DesignCenterServer) CreateApplication(ctx context.Context, req *pb.CreateApplicationRequest) (*pb.Application, error) {
	parent, err := s.parseSpaceName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent space does not exist
	parentSpace := &pb.Space{}
	if err := s.storage.Get(ctx, parent.String(), parentSpace); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.ApplicationId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "applicationId is required")
	}

	if req.Application == nil {
		return nil, status.Errorf(codes.InvalidArgument, "application is required")
	}

	applicationID := req.ApplicationId
	fqn := fmt.Sprintf("%s/applications/%s", parent.String(), applicationID)

	name, err := s.parseApplicationName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.Application).(*pb.Application)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj = s.populateApplication(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DesignCenterServer) UpdateApplication(ctx context.Context, req *pb.UpdateApplicationRequest) (*pb.Application, error) {
	if req.Application == nil {
		return nil, status.Errorf(codes.InvalidArgument, "application is required")
	}

	name, err := s.parseApplicationName(req.Application.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Application)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.Application, paths); err != nil {
		return nil, err
	}

	updated = s.populateApplication(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DesignCenterServer) DeleteApplication(ctx context.Context, req *pb.DeleteApplicationRequest) (*longrunning.Operation, error) {
	name, err := s.parseApplicationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Application{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		Target:     fqn,
		Verb:       "delete",
		CreateTime: timestamppb.Now(),
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		metadata.ResourceMetadata = &pb.OperationMetadata_ApplicationMetadata{
			ApplicationMetadata: &pb.ApplicationOperationMetadata{
				Step: "SUCCEEDED",
			},
		}
		if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	})
}

func (s *DesignCenterServer) populateApplication(obj *pb.Application) *pb.Application {
	name, err := s.parseApplicationName(obj.Name)
	if err == nil {
		if obj.ApphubApplication == "" {
			obj.ApphubApplication = fmt.Sprintf("projects/%s/locations/%s/applications/%s", name.Project.ID, name.Location, name.Application)
		}
		if len(obj.ProjectParameters) == 0 {
			obj.ProjectParameters = []*pb.ProjectParameters{
				{
					ProjectId: name.Project.ID,
					Roles: []string{
						"roles/apphub.editor",
						"roles/config.agent",
					},
				},
			}
		}
		if len(obj.RolesApisConfig) == 0 {
			obj.RolesApisConfig = []*pb.RolesAndApisConfig{
				{
					ProjectId: name.Project.ID,
					Roles: []*pb.RoleInfo{
						{
							Role:           "roles/apphub.editor",
							SystemProvided: true,
						},
						{
							Role:           "roles/config.agent",
							SystemProvided: true,
						},
					},
				},
			}
		}
	}

	if obj.State == pb.Application_STATE_UNSPECIFIED {
		obj.State = pb.Application_DRAFT
	}
	if obj.Type == pb.ApplicationType_APPLICATION_TYPE_UNSPECIFIED {
		obj.Type = pb.ApplicationType_TERRAFORM_APP
	}
	if obj.CompositionType == pb.ApplicationCompositionType_APPLICATION_COMPOSITION_TYPE_UNSPECIFIED {
		obj.CompositionType = pb.ApplicationCompositionType_STANDARD
	}
	if obj.Scope == nil {
		obj.Scope = &apphubpb.Scope{
			Type: apphubpb.Scope_REGIONAL,
		}
	}

	if len(obj.ComponentParameters) == 0 {
		emptyListVal, _ := structpb.NewValue([]any{})
		obj.ComponentParameters = []*pb.ComponentParameters{
			{
				Component: "projects/gcpdesigncenter/locations/us-central1/spaces/googlespace/applicationTemplates/simple-agent-runtime/components/agent-runtime-1",
				ComponentParameterSchema: []*pb.ComponentParameterSchema{
					{IsRequired: true, Key: "display_name", Type: "string"},
					{IsRequired: true, Key: "project_id", Type: "string"},
					{IsRequired: true, Key: "region", Type: "string"},
					{Key: "description", Type: "string"},
					{Key: "kms_key_name", Type: "string"},
					{Key: "spec", Type: "object({\n    agent_framework = optional(string)\n    class_methods   = optional(list(any))\n    deployment_spec = optional(object({\n      container_concurrency = optional(number, 9)\n      env                   = optional(map(string), {})\n      max_instances         = optional(number, 100)\n      min_instances         = optional(number, 1)\n      psc_interface_config = optional(object({\n        network_attachment = optional(string)\n        dns_peering_configs = optional(list(object({\n          domain         = string\n          target_project = string\n          target_network = string\n        })), [])\n      }))\n      resource_limits = optional(map(string))\n      secret_env = optional(list(object({\n        name = string\n        secret_ref = object({\n          secret  = string\n          version = optional(string)\n        })\n      })), [])\n    }))\n    package_spec = optional(object({\n      dependency_files_gcs_uri = optional(string)\n      pickle_object_gcs_uri    = optional(string)\n      python_version           = optional(string)\n      requirements_gcs_uri     = optional(string)\n    }))\n    container_spec = optional(object({\n      image_uri = string\n    }))\n    source_code_spec = optional(object({\n      inline_source = optional(object({\n        source_archive = optional(string)\n      }))\n      image_spec = optional(object({\n        build_args = optional(map(string))\n      }))\n      python_spec = optional(object({\n        entrypoint_module = optional(string)\n        entrypoint_object = optional(string)\n        requirements_file = optional(string)\n        version           = optional(string)\n      }))\n      developer_connect_source = optional(object({\n        config = object({\n          git_repository_link = string\n          dir                 = string\n          revision            = string\n        })\n      }))\n    }))\n    service_account = optional(string)\n    identity_type   = optional(string, \"SERVICE_ACCOUNT\")\n  })"},
					{DefaultValue: emptyListVal, Key: "google_managed_agent_gateway_config", Type: "list(object({\n    gateway_type = string\n    gateway_id   = string\n  }))"},
					{DefaultValue: emptyListVal, Key: "service_account_roles", Type: "list(string)"},
					{DefaultValue: emptyListVal, Key: "effective_identity_roles", Type: "list(string)"},
					{DefaultValue: emptyListVal, Key: "module_depends_on", Type: "list(string)"},
				},
				Parameters: []*pb.Parameter{
					{Key: "region", Value: structpb.NewStringValue("us-central1")},
					{Key: "display_name", Value: structpb.NewStringValue("Agent Runtime")},
				},
				State: pb.ComponentParameters_DRAFT,
			},
		}
	}

	if obj.SerializedApplicationTemplate == nil {
		emptyListVal, _ := structpb.NewValue([]any{})
		obj.SerializedApplicationTemplate = &pb.SerializedApplicationTemplate{
			Components: []*pb.SerializedComponent{
				{
					Apis: []string{
						"cloudresourcemanager.googleapis.com",
						"serviceusage.googleapis.com",
						"iamcredentials.googleapis.com",
						"iam.googleapis.com",
						"cloudbilling.googleapis.com",
						"cloudkms.googleapis.com",
						"servicenetworking.googleapis.com",
						"secretmanager.googleapis.com",
						"notebooks.googleapis.com",
						"aiplatform.googleapis.com",
						"iap.googleapis.com",
						"modelarmor.googleapis.com",
						"dlp.googleapis.com",
					},
					ComponentParameterSchema: []*pb.ComponentParameterSchema{
						{IsRequired: true, Key: "display_name", Type: "string"},
						{IsRequired: true, Key: "project_id", Type: "string"},
						{IsRequired: true, Key: "region", Type: "string"},
						{Key: "description", Type: "string"},
						{Key: "kms_key_name", Type: "string"},
						{Key: "spec", Type: "object({\n    agent_framework = optional(string)\n    class_methods   = optional(list(any))\n    deployment_spec = optional(object({\n      container_concurrency = optional(number, 9)\n      env                   = optional(map(string), {})\n      max_instances         = optional(number, 100)\n      min_instances         = optional(number, 1)\n      psc_interface_config = optional(object({\n        network_attachment = optional(string)\n        dns_peering_configs = optional(list(object({\n          domain         = string\n          target_project = string\n          target_network = string\n        })), [])\n      }))\n      resource_limits = optional(map(string))\n      secret_env = optional(list(object({\n        name = string\n        secret_ref = object({\n          secret  = string\n          version = optional(string)\n        })\n      })), [])\n    }))\n    package_spec = optional(object({\n      dependency_files_gcs_uri = optional(string)\n      pickle_object_gcs_uri    = optional(string)\n      python_version           = optional(string)\n      requirements_gcs_uri     = optional(string)\n    }))\n    container_spec = optional(object({\n      image_uri = string\n    }))\n    source_code_spec = optional(object({\n      inline_source = optional(object({\n        source_archive = optional(string)\n      }))\n      image_spec = optional(object({\n        build_args = optional(map(string))\n      }))\n      python_spec = optional(object({\n        entrypoint_module = optional(string)\n        entrypoint_object = optional(string)\n        requirements_file = optional(string)\n        version           = optional(string)\n      }))\n      developer_connect_source = optional(object({\n        config = object({\n          git_repository_link = string\n          dir                 = string\n          revision            = string\n        })\n      }))\n    }))\n    service_account = optional(string)\n    identity_type   = optional(string, \"SERVICE_ACCOUNT\")\n  })"},
						{DefaultValue: emptyListVal, Key: "google_managed_agent_gateway_config", Type: "list(object({\n    gateway_type = string\n    gateway_id   = string\n  }))"},
						{DefaultValue: emptyListVal, Key: "service_account_roles", Type: "list(string)"},
						{DefaultValue: emptyListVal, Key: "effective_identity_roles", Type: "list(string)"},
						{DefaultValue: emptyListVal, Key: "module_depends_on", Type: "list(string)"},
					},
					Parameters: []*pb.Parameter{
						{Key: "region", Value: structpb.NewStringValue("us-central1")},
						{Key: "display_name", Value: structpb.NewStringValue("Agent Runtime")},
					},
					Roles: []string{
						"roles/aiplatform.admin",
						"roles/compute.admin",
						"roles/compute.networkAdmin",
						"roles/notebooks.admin",
						"roles/iam.securityAdmin",
						"roles/iam.serviceAccountAdmin",
						"roles/iam.serviceAccountUser",
						"roles/cloudkms.admin",
						"roles/iam.roleAdmin",
						"roles/storage.admin",
						"roles/cloudkms.cryptoKeyEncrypterDecrypter",
						"roles/modelarmor.admin",
						"roles/modelarmor.floorSettingsAdmin",
						"roles/artifactregistry.reader",
						"roles/serviceusage.serviceUsageAdmin",
					},
					SharedTemplateRevisionUri: "projects/gcpdesigncenter/locations/us-central1/spaces/googlespace/sharedTemplates/googlespace-vertex-ai-agent-engine-1sx3p0y/revisions/r-10",
					Uri:                       "projects/gcpdesigncenter/locations/us-central1/spaces/googlespace/applicationTemplates/simple-agent-runtime/components/agent-runtime-1",
				},
			},
			CompositionType: pb.ApplicationCompositionType_STANDARD,
			Description:     "Google opinionated template that configures a single Agent Runtime component.",
			DisplayName:     "Simple Agent Runtime",
			IacFormat:       pb.ApplicationTemplate_TERRAFORM,
			Uri:             "projects/gcpdesigncenter/locations/us-central1/spaces/googlespace/applicationTemplates/simple-agent-runtime",
		}
	}

	return obj
}
