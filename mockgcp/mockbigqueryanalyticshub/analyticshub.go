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

package mockbigqueryanalyticshub

import (
	"context"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/analyticshub/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type analyticsHubServer struct {
	*MockService
	pb.UnimplementedAnalyticsHubServiceServer
}

func (a *analyticsHubServer) GetDataExchange(ctx context.Context, request *pb.GetDataExchangeRequest) (*pb.DataExchange, error) {
	name, err := a.parseDataExchangeID(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DataExchange{}
	if err := a.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) CreateDataExchange(ctx context.Context, request *pb.CreateDataExchangeRequest) (*pb.DataExchange, error) {
	reqName := request.Parent + "/dataExchanges/" + request.DataExchangeId
	name, err := a.parseDataExchangeID(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(request.DataExchange).(*pb.DataExchange)
	obj.Name = fqn

	if err := a.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) UpdateDataExchange(ctx context.Context, request *pb.UpdateDataExchangeRequest) (*pb.DataExchange, error) {
	name, err := a.parseDataExchangeID(request.GetDataExchange().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DataExchange{}
	if err := a.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range request.GetUpdateMask().Paths {
		switch path {
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := a.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) DeleteDataExchange(ctx context.Context, request *pb.DeleteDataExchangeRequest) (*empty.Empty, error) {
	name, err := a.parseDataExchangeID(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.DataExchange{}
	if err := a.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)
	return &empty.Empty{}, nil
}

func (a *analyticsHubServer) GetListing(ctx context.Context, request *pb.GetListingRequest) (*pb.Listing, error) {
	name, err := a.parseListingID(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Listing{}
	if err := a.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) CreateListing(ctx context.Context, request *pb.CreateListingRequest) (*pb.Listing, error) {
	reqName := request.Parent + "/listings/" + request.ListingId
	name, err := a.parseListingID(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(request.Listing).(*pb.Listing)
	obj.Name = fqn

	if err := a.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) UpdateListing(ctx context.Context, request *pb.UpdateListingRequest) (*pb.Listing, error) {
	name, err := a.parseListingID(request.GetListing().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Listing{}
	if err := a.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range request.GetUpdateMask().Paths {
		switch path {
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := a.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *analyticsHubServer) DeleteListing(ctx context.Context, request *pb.DeleteListingRequest) (*empty.Empty, error) {
	name, err := a.parseListingID(request.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.Listing{}
	if err := a.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)
	return &empty.Empty{}, nil
}

// e.g. 'projects/${projectId}/locations/us-west2/dataExchanges/dataExchange-${uniqueId}'
type dataExchangeName struct {
	Project        *projects.ProjectData
	Location       string
	DataExchangeID string
}

func (n *dataExchangeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/dataExchanges/" + n.DataExchangeID
}

// parseDataExchangeID parses a string into a dataExchangeName.
// The expected form is projects/<projectID>/locations/<region>/dataExchanges/<dataExchangesID>
func (s *MockService) parseDataExchangeID(name string) (*dataExchangeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataExchanges" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &dataExchangeName{
			Project:        project,
			Location:       tokens[3],
			DataExchangeID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type listingName struct {
	Project        *projects.ProjectData
	Location       string
	DataExchangeID string
	ListingID      string
}

func (n *listingName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/dataExchanges/" + n.DataExchangeID + "/listings/" + n.ListingID
}

// parseDataExchangeID parses a string into a dataExchangeName.
// The expected form is projects/<projectID>/locations/<region>/dataExchanges/<dataExchangesID>/listings/<listingID>
func (s *MockService) parseListingID(name string) (*listingName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataExchanges" && tokens[6] == "listings" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &listingName{
			Project:        project,
			Location:       tokens[3],
			DataExchangeID: tokens[5],
			ListingID:      tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
