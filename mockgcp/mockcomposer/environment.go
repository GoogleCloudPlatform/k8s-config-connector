// Copyright 2024 Google LLC
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

// +tool:mockgcp-support
// proto.service: google.cloud.orchestration.airflow.service.v1.Environments
// proto.message: google.cloud.orchestration.airflow.service.v1.Environment

package mockcomposer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *MockService) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *MockService) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunningpb.Operation, error) {
	// The parent field is projects/{projectId}/locations/{locationId}.
	name, err := s.parseParentEnvironment(req.Parent)
	if err != nil {
		return nil, err
	}
	obj := proto.Clone(req.GetEnvironment()).(*pb.Environment)
	// The resource name of the environment.
	obj.Name = fmt.Sprintf("projects/%s/locations/%s/environments/%s", name.Project.ID, name.Location, "TEMP")

	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	s.populateDefaultsForEnvironment(obj)

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s/environments", name.Project.ID, name.Location)
	lroRet := proto.Clone(obj).(*pb.Environment)
	lroRet.CreateTime = nil
	lroRet.UpdateTime = nil
	lroRet.Uuid = ""
	// State can only be set to running with UpdateEnvironment
	lroRet.State = pb.Environment_CREATING
	return s.operations.StartLRO(ctx, lroPrefix, nil, func() (proto.Message, error) {
		// Environment ID is parsed from environment.Name, projects/*/locations/*/environments/*
		tokens := strings.Split(lroRet.Name, "/")
		environmentId := tokens[5]
		// The resource name of the environment.
		obj.Name = fmt.Sprintf("projects/%s/locations/%s/environments/%s", name.Project.ID, name.Location, environmentId)
		fqn := obj.Name
		obj.Uuid = "123e4567-e89b-12d3-a456-426614174000" // this doesn't seem to be configurable

		// Create needs to set state
		obj.State = pb.Environment_CREATING

		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *MockService) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: support update mask

	// TODO:  only labels are updateable with a mask
	proto.Merge(obj, req.GetEnvironment())

	s.populateDefaultsForEnvironment(obj)

	obj.State = pb.Environment_UPDATING
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	updatedObj := proto.Clone(obj).(*pb.Environment)
	updatedObj.CreateTime = nil
	updatedObj.UpdateTime = nil
	updatedObj.Uuid = ""
	prefix := fmt.Sprintf("projects/%s/locations/%s/environments", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		// State can only be set to running with UpdateEnvironment
		updatedObj.State = pb.Environment_RUNNING
		obj.State = pb.Environment_RUNNING
		obj.UpdateTime = timestamppb.New(time.Now())
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return updatedObj, nil
	})
}

func (s *MockService) DeleteEnvironment(ctx context.Context, req *pb.DeleteEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Environment{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s/environments", name.Project.ID, name.Location)
	deletedObj.CreateTime = nil
	deletedObj.UpdateTime = nil
	deletedObj.Uuid = ""
	// State can only be set to running with UpdateEnvironment
	deletedObj.State = pb.Environment_DELETING
	return s.operations.StartLRO(ctx, prefix, nil, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *MockService) populateDefaultsForEnvironment(obj *pb.Environment) {
	if obj.Config == nil {
		obj.Config = &pb.EnvironmentConfig{}
	}

	s.populateDefaultsForEnvironmentConfig(obj.Config)
}

func (s *MockService) populateDefaultsForEnvironmentConfig(config *pb.EnvironmentConfig) {
	if config.NodeCount == 0 {
		config.NodeCount = 3
	}

	if config.SoftwareConfig == nil {
		config.SoftwareConfig = &pb.SoftwareConfig{}
	}

	s.populateDefaultsForSoftwareConfig(config.SoftwareConfig)

	if config.PrivateEnvironmentConfig == nil {
		config.PrivateEnvironmentConfig = &pb.PrivateEnvironmentConfig{}
	}
	// TODO: more populating

}

func (s *MockService) populateDefaultsForSoftwareConfig(config *pb.SoftwareConfig) {
	// TODO:
}

type environmentName struct {
	Project         *projects.ProjectData
	Location        string
	EnvironmentName string
}

func (n *environmentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/environments/" + n.EnvironmentName
}

// parseEnvironmentName parses a string into a environmentName.
// The expected form is `projects/*/locations/*/environments/*`.
func (s *MockService) parseEnvironmentName(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "environments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &environmentName{
			Project:         project,
			Location:        tokens[3],
			EnvironmentName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// parseParentEnvironment parses a string into project and location.
// The expected form is `projects/*/locations/*`.
func (s *MockService) parseParentEnvironment(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &environmentName{
			Project:  project,
			Location: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


</example>




<example>
in.proto.message: google.cloud.security.privateca.v1.CaPool
in.proto.message.definition: message CaPool {
  option (google.api.resource) = {
    type: "privateca.googleapis.com/CaPool"
    pattern: "projects/{project}/locations/{location}/caPools/{ca_pool}"
  };

  // The Tier of this [CaPool][google.cloud.security.privateca.v1.CaPool].
  enum Tier {
    // Not specified.
    TIER_UNSPECIFIED = 0;

    // Enterprise tier.
    ENTERPRISE = 1;

    // DevOps tier.
    DEVOPS = 2;
  }

  // Output only. The resource name for this
  // [CaPool][google.cloud.security.privateca.v1.CaPool] in the format
  // `projects/*/locations/*/caPools/*`.
  string name = 1 [(google.api.field_behavior) = OUTPUT_ONLY];

  // Required. Immutable. The [Tier][google.cloud.security.privateca.v1.CaPool.Tier]
  // of this [CaPool][google.cloud.security.privateca.v1.CaPool].
  Tier tier = 2 [
    (google.api.field_behavior) = IMMUTABLE,
    (google.api.field_behavior) = REQUIRED
  ];

  // Optional. The
  // [IssuancePolicy][google.cloud.security.privateca.v1.CaPool.IssuancePolicy]
  // to control how
  // [Certificates][google.cloud.security.privateca.v1.Certificate] will be
  // issued from this [CaPool][google.cloud.security.privateca.v1.CaPool].
  IssuancePolicy issuance_policy = 3
      [(google.api.field_behavior) = OPTIONAL];

  // Optional. The
  // [PublishingOptions][google.cloud.security.privateca.v1.CaPool.PublishingOptions]
  // to follow when issuing
  // [Certificates][google.cloud.security.privateca.v1.Certificate] from any
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
  // in this [CaPool][google.cloud.security.privateca.v1.CaPool].
  PublishingOptions publishing_options = 4
      [(google.api.field_behavior) = OPTIONAL];

  // Optional. Labels with user-defined metadata.
  map<string, string> labels = 5 [(google.api.field_behavior) = OPTIONAL];
}
in.proto.service: google.cloud.security.privateca.v1.CertificateAuthorityService
in.proto.service.definition: service CertificateAuthorityService {
  option (google.api.default_host) = "privateca.googleapis.com";
  option (google.api.oauth_scopes) = "https://www.googleapis.com/auth/cloud-platform";

  // Create a new
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
  // in a given Project and Location.
  rpc CreateCertificateAuthority(CreateCertificateAuthorityRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*}/certificateAuthorities"
      body: "certificate_authority"
    };
    option (google.api.method_signature) = "parent,certificate_authority,certificate_authority_id";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateAuthority"
      metadata_type: "OperationMetadata"
    };
  }

  // Returns a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority].
  rpc GetCertificateAuthority(GetCertificateAuthorityRequest) returns (CertificateAuthority) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/certificateAuthorities/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists
  // [CertificateAuthorities][google.cloud.security.privateca.v1.CertificateAuthority].
  rpc ListCertificateAuthorities(ListCertificateAuthoritiesRequest) returns (ListCertificateAuthoritiesResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*}/certificateAuthorities"
    };
    option (google.api.method_signature) = "parent";
  }

  // Update a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority].
  rpc UpdateCertificateAuthority(UpdateCertificateAuthorityRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{certificate_authority.name=projects/*/locations/*/certificateAuthorities/*}"
      body: "certificate_authority"
    };
    option (google.api.method_signature) = "certificate_authority,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateAuthority"
      metadata_type: "OperationMetadata"
    };
  }

  // Fetch a certificate signing request (CSR) from a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
  // that is in state
  // [AWAITING_USER_ACTIVATION][google.cloud.security.privateca.v1.CertificateAuthority.State.AWAITING_USER_ACTIVATION]
  // and is of type
  // [SUBORDINATE][google.cloud.security.privateca.v1.CertificateAuthority.Type.SUBORDINATE].
  rpc FetchCaCerts(FetchCaCertsRequest) returns (FetchCaCertsResponse) {
    option (google.api.http) = {
      post: "/v1/{ca_pool=projects/*/locations/*/caPools/*}:fetchCaCerts"
      body: "*"
    };
  }

  // Returns a
  // [CertificateRevocationList][google.cloud.security.privateca.v1.CertificateRevocationList].
  rpc GetCertificateRevocationList(GetCertificateRevocationListRequest) returns (CertificateRevocationList) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/caPools/*/certificateAuthorities/*/certificateRevocationLists/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists
  // [CertificateRevocationLists][google.cloud.security.privateca.v1.CertificateRevocationList].
  rpc ListCertificateRevocationLists(ListCertificateRevocationListsRequest) returns (ListCertificateRevocationListsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/caPools/*/certificateAuthorities/*}/certificateRevocationLists"
    };
    option (google.api.method_signature) = "parent";
  }

  // Update a
  // [CertificateRevocationList][google.cloud.security.privateca.v1.CertificateRevocationList].
  rpc UpdateCertificateRevocationList(UpdateCertificateRevocationListRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{certificate_revocation_list.name=projects/*/locations/*/caPools/*/certificateAuthorities/*/certificateRevocationLists/*}"
      body: "certificate_revocation_list"
    };
    option (google.api.method_signature) = "certificate_revocation_list,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateRevocationList"
      metadata_type: "OperationMetadata"
    };
  }

  // Create a new
  // [Certificate][google.cloud.security.privateca.v1.Certificate] in a given
  // Project, Location from a particular
  // [CaPool][google.cloud.security.privateca.v1.CaPool].
  rpc CreateCertificate(CreateCertificateRequest) returns (Certificate) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*/caPools/*}/certificates"
      body: "certificate"
    };
    option (google.api.method_signature) = "parent,certificate,certificate_id";
  }

  // Returns a [Certificate][google.cloud.security.privateca.v1.Certificate].
  rpc GetCertificate(GetCertificateRequest) returns (Certificate) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/caPools/*/certificates/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists
  // [Certificates][google.cloud.security.privateca.v1.Certificate].
  rpc ListCertificates(ListCertificatesRequest) returns (ListCertificatesResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*/caPools/*}/certificates"
    };
    option (google.api.method_signature) = "parent";
  }

  // Revoke a
  // [Certificate][google.cloud.security.privateca.v1.Certificate].
  rpc RevokeCertificate(RevokeCertificateRequest) returns (Certificate) {
    option (google.api.http) = {
      post: "/v1/{name=projects/*/locations/*/caPools/*/certificates/*}:revoke"
      body: "*"
    };
    option (google.api.method_signature) = "name";
  }

  // Update a [Certificate][google.cloud.security.privateca.v1.Certificate].
  rpc UpdateCertificate(UpdateCertificateRequest) returns (Certificate) {
    option (google.api.http) = {
      patch: "/v1/{certificate.name=projects/*/locations/*/caPools/*/certificates/*}"
      body: "certificate"
    };
    option (google.api.method_signature) = "certificate,update_mask";
  }

  // Activate a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
  // that is in state
  // [AWAITING_USER_ACTIVATION][google.cloud.security.privateca.v1.CertificateAuthority.State.AWAITING_USER_ACTIVATION]
  // and is of type
  // [SUBORDINATE][google.cloud.security.privateca.v1.CertificateAuthority.Type.SUBORDINATE].
  // After the parent CA has been issued a CA certificate using
  // [FetchCaCerts][google.cloud.security.privateca.v1.CertificateAuthorityService.FetchCaCerts],
  // this method can complete the activation process.
  rpc ActivateCertificateAuthority(ActivateCertificateAuthorityRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=projects/*/locations/*/certificateAuthorities/*}:activate"
      body: "*"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateAuthority"
      metadata_type: "OperationMetadata"
    };
  }

  // Create a new [CaPool][google.cloud.security.privateca.v1.CaPool] in a
  // given Project and Location.
  rpc CreateCaPool(CreateCaPoolRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*}/caPools"
      body: "ca_pool"
    };
    option (google.api.method_signature) = "parent,ca_pool,ca_pool_id";
    option (google.longrunning.operation_info) = {
      response_type: "CaPool"
      metadata_type: "OperationMetadata"
    };
  }

  // Update a [CaPool][google.cloud.security.privateca.v1.CaPool].
  rpc UpdateCaPool(UpdateCaPoolRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{ca_pool.name=projects/*/locations/*/caPools/*}"
      body: "ca_pool"
    };
    option (google.api.method_signature) = "ca_pool,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "CaPool"
      metadata_type: "OperationMetadata"
    };
  }

  // Returns a [CaPool][google.cloud.security.privateca.v1.CaPool].
  rpc GetCaPool(GetCaPoolRequest) returns (CaPool) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/caPools/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists [CaPools][google.cloud.security.privateca.v1.CaPool].
  rpc ListCaPools(ListCaPoolsRequest) returns (ListCaPoolsResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*}/caPools"
    };
    option (google.api.method_signature) = "parent";
  }

  // Delete a [CaPool][google.cloud.security.privateca.v1.CaPool].
  rpc DeleteCaPool(DeleteCaPoolRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/caPools/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "OperationMetadata"
    };
  }

  // Disable a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority].
  rpc DisableCertificateAuthority(DisableCertificateAuthorityRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=projects/*/locations/*/certificateAuthorities/*}:disable"
      body: "*"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateAuthority"
      metadata_type: "OperationMetadata"
    };
  }

  // Enable a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority].
  rpc EnableCertificateAuthority(EnableCertificateAuthorityRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=projects/*/locations/*/certificateAuthorities/*}:enable"
      body: "*"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateAuthority"
      metadata_type: "OperationMetadata"
    };
  }

  // Undelete a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority]
  // that has been deleted.
  rpc UndeleteCertificateAuthority(UndeleteCertificateAuthorityRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=projects/*/locations/*/certificateAuthorities/*}:undelete"
      body: "*"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateAuthority"
      metadata_type: "OperationMetadata"
    };
  }

  // Delete a
  // [CertificateAuthority][google.cloud.security.privateca.v1.CertificateAuthority].
  rpc DeleteCertificateAuthority(DeleteCertificateAuthorityRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/certificateAuthorities/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateAuthority"
      metadata_type: "OperationMetadata"
    };
  }

  // Create a new
  // [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate]
  // in a given Project and Location.
  rpc CreateCertificateTemplate(CreateCertificateTemplateRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{parent=projects/*/locations/*}/certificateTemplates"
      body: "certificate_template"
    };
    option (google.api.method_signature) = "parent,certificate_template,certificate_template_id";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateTemplate"
      metadata_type: "OperationMetadata"
    };
  }

  // Delete a
  // [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate].
  rpc DeleteCertificateTemplate(DeleteCertificateTemplateRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      delete: "/v1/{name=projects/*/locations/*/certificateTemplates/*}"
    };
    option (google.api.method_signature) = "name";
    option (google.longrunning.operation_info) = {
      response_type: "google.protobuf.Empty"
      metadata_type: "OperationMetadata"
    };
  }

  // Returns a
  // [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate].
  rpc GetCertificateTemplate(GetCertificateTemplateRequest) returns (CertificateTemplate) {
    option (google.api.http) = {
      get: "/v1/{name=projects/*/locations/*/certificateTemplates/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists
  // [CertificateTemplates][google.cloud.security.privateca.v1.CertificateTemplate].
  rpc ListCertificateTemplates(ListCertificateTemplatesRequest) returns (ListCertificateTemplatesResponse) {
    option (google.api.http) = {
      get: "/v1/{parent=projects/*/locations/*}/certificateTemplates"
    };
    option (google.api.method_signature) = "parent";
  }

  // Update a
  // [CertificateTemplate][google.cloud.security.privateca.v1.CertificateTemplate].
  rpc UpdateCertificateTemplate(UpdateCertificateTemplateRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v1/{certificate_template.name=projects/*/locations/*/certificateTemplates/*}"
      body: "certificate_template"
    };
    option (google.api.method_signature) = "certificate_template,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "CertificateTemplate"
      metadata_type: "OperationMetadata"
    };
  }
}
out: // Copyright 2024 Google LLC
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

// +tool:mockgcp-support
// proto.service: google.cloud.security.privateca.v1.CertificateAuthorityService
// proto.message: google.cloud.security.privateca.v1.CaPool

package mockprivateca

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/security/privateca/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type privateCAServer struct {
	*MockService
	pb.UnimplementedCertificateAuthorityServiceServer
}

func (s *privateCAServer) GetCaPool(ctx context.Context, req *pb.GetCaPoolRequest) (*pb.CaPool, error) {
	name, err := s.parseCaPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CaPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "CaPool %s not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *privateCAServer) CreateCaPool(ctx context.Context, req *pb.CreateCaPoolRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/caPools/%s", req.GetParent(), req.GetCaPoolId())
	name, err := s.parseCaPoolName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := proto.Clone(req.GetCaPool()).(*pb.CaPool)
	obj.Name = fqn

	s.populateDefaultsForCAPool(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroName := fmt.Sprintf("%s/operations/operation-%s", name.Project.ID, name.Location)

	lro, err := s.operations.StartLRO(ctx, lroName, &pb.OperationMetadata{}, func() (proto.Message, error) {
		return obj, nil
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in starting LRO: %w", err)
	}

	return lro, nil
}

func (s *privateCAServer) populateDefaultsForCAPool(obj *pb.CaPool) {
	if obj.Tier == pb.CaPool_TIER_UNSPECIFIED {
		obj.Tier = pb.CaPool_ENTERPRISE
	}
}

func (s *privateCAServer) UpdateCaPool(ctx context.Context, req *pb.UpdateCaPoolRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetCaPool().GetName()
	name, err := s.parseCaPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CaPool{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "capool %q not found", name)
		}
		return nil, err
	}

	// We will update the parent config with the new capool as we
	// dont have an update_mask implementation.
	proto.Merge(obj, req.GetCaPool())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroName := fmt.Sprintf("%s/operations/operation-%s", name.Project.ID, name.Location)

	lro, err := s.operations.StartLRO(ctx,

