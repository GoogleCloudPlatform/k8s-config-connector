// Copyright 2022 Google LLC
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

package mockserviceusage

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/serviceusage/v1beta1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/anypb"
)

type ServiceUsageV1Beta1 struct {
	*MockService
	pb.UnimplementedServiceUsageServer
}

func (s *ServiceUsageV1Beta1) GenerateServiceIdentity(ctx context.Context, req *pb.GenerateServiceIdentityRequest) (*longrunning.Operation, error) {
	name, err := s.parseServiceName(req.Parent)
	if err != nil {
		return nil, err
	}

	op, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Find a less hacky way to provide a LRO response
	response := &longrunning.Operation_Response{}
	identity := &pb.ServiceIdentity{}

	switch name.ServiceName {
	case "container.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@container-engine-robot.iam.gserviceaccount.com"
		identity.UniqueId = "123456789001"
	case "pubsub.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-pubsub.iam.gserviceaccount.com"
		identity.UniqueId = "123456789002"
	case "compute.googleapis.com":
		// compute.googleapis.com does not send the P4SA
	case "sqladmin.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-cloud-sql.iam.gserviceaccount.com"
		identity.UniqueId = "123456789003"
	case "alloydb.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-alloydb.iam.gserviceaccount.com"
		identity.UniqueId = "123456789004"
	case "aiplatform.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-aiplatform.iam.gserviceaccount.com"
		identity.UniqueId = "123456789005"
	case "secretmanager.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-secretmanager.iam.gserviceaccount.com"
		identity.UniqueId = "123456789006"
	case "bigquery.googleapis.com":
		identity.Email = "bq-" + strconv.FormatInt(name.Project.Number, 10) + "@bigquery-encryption.iam.gserviceaccount.com"
		identity.UniqueId = "123456789007"
	case "apigee.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-apigee.iam.gserviceaccount.com"
		identity.UniqueId = "123456789008"
	case "managedkafka.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-managedkafka.iam.gserviceaccount.com"
	case "workflows.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-workflows.iam.gserviceaccount.com"
	case "spanner.googleapis.com":
		identity.Email = "service-" + strconv.FormatInt(name.Project.Number, 10) + "@gcp-sa-spanner.iam.gserviceaccount.com"
		identity.UniqueId = "123456789009"
	default:
		return nil, fmt.Errorf("generating serviceIdentity for service %q not implemented in mock", name.ServiceName)
	}

	any, err := anypb.New(identity)
	if err != nil {
		return nil, err
	}
	response.Response = any

	op.Done = true
	op.Result = response

	return op, nil
}
