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

package mockdialogflow

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	pb_v2 "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type conversationDatasetsServer struct {
	*MockService
	pb_v2.UnimplementedConversationDatasetsServer
}

type conversationDatasetName struct {
	Project             *projects.ProjectData
	Location            string
	ConversationDataset string
}

func (n *conversationDatasetName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/conversationDatasets/%s", n.Project.ID, n.Location, n.ConversationDataset)
}

func (s *MockService) parseConversationDatasetName(name string) (*conversationDatasetName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "conversationDatasets" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &conversationDatasetName{
			Project:             project,
			Location:            tokens[3],
			ConversationDataset: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/conversationDatasets/{conversationDataset}", name)
}

func (s *MockService) parseConversationDatasetParent(parent string) (*conversationDatasetName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &conversationDatasetName{
			Project:  project,
			Location: tokens[3],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}", parent)
}

func (s *conversationDatasetsServer) GetConversationDataset(ctx context.Context, req *pb_v2.GetConversationDatasetRequest) (*pb_v2.ConversationDataset, error) {
	name, err := s.parseConversationDatasetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb_v2.ConversationDataset{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ConversationDataset %s not found in projects/%s/locations/%s.", name.ConversationDataset, name.Project.ID, name.Location)
		}
		return nil, err
	}

	return obj, nil
}

func (s *conversationDatasetsServer) CreateConversationDataset(ctx context.Context, req *pb_v2.CreateConversationDatasetRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseConversationDatasetParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	datasetID := ""
	if req.GetConversationDataset().GetName() != "" {
		parsedName, err := s.parseConversationDatasetName(req.GetConversationDataset().GetName())
		if err == nil && parsedName.ConversationDataset != "" {
			datasetID = parsedName.ConversationDataset
		}
	}
	if datasetID == "" {
		n, err := rand.Int(rand.Reader, big.NewInt(1<<62))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
		datasetID = base64.RawURLEncoding.EncodeToString([]byte(n.String()))
	}

	parent.ConversationDataset = datasetID
	fqn := parent.String()

	obj := proto.Clone(req.GetConversationDataset()).(*pb_v2.ConversationDataset)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.SatisfiesPzi = proto.Bool(false)
	obj.SatisfiesPzs = proto.Bool(false)
	if obj.ConversationInfo == nil {
		obj.ConversationInfo = &pb_v2.ConversationInfo{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", parent.Project.ID, parent.Location)
	metadata := &pb_v2.CreateConversationDatasetOperationMetadata{
		ConversationDataset: fqn,
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *conversationDatasetsServer) ListConversationDatasets(ctx context.Context, req *pb_v2.ListConversationDatasetsRequest) (*pb_v2.ListConversationDatasetsResponse, error) {
	parent, err := s.parseConversationDatasetParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%s/locations/%s/conversationDatasets/", parent.Project.ID, parent.Location)

	response := &pb_v2.ListConversationDatasetsResponse{}
	datasetKind := (&pb_v2.ConversationDataset{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, datasetKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb_v2.ConversationDataset)
		response.ConversationDatasets = append(response.ConversationDatasets, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *conversationDatasetsServer) DeleteConversationDataset(ctx context.Context, req *pb_v2.DeleteConversationDatasetRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConversationDatasetName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb_v2.ConversationDataset{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &longrunningpb.Operation{Done: true}, nil
}

func RegisterConversationDatasetsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := pb_v2.NewConversationDatasetsClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	for _, prefix := range []string{"/v2", "/v2beta1"} {
		// POST {prefix}/{parent=projects/*/locations/*}/conversationDatasets
		createPath := prefix + "/{parent=projects/*/locations/*}/conversationDatasets"
		if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.CreateConversationDatasetRequest{}
			protoReq.Parent = pathParams["parent"]

			newReader, berr := utilities.IOReaderFactory(r.Body)
			if berr != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
				return
			}
			protoReq.ConversationDataset = &pb_v2.ConversationDataset{}
			if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.ConversationDataset); err != nil && err != io.EOF {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
				return
			}

			op, err := client.CreateConversationDataset(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// GET {prefix}/{name=projects/*/locations/*/conversationDatasets/*}
		getPath := prefix + "/{name=projects/*/locations/*/conversationDatasets/*}"
		if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.GetConversationDatasetRequest{
				Name: pathParams["name"],
			}

			obj, err := client.GetConversationDataset(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// GET {prefix}/{parent=projects/*/locations/*}/conversationDatasets
		listPath := prefix + "/{parent=projects/*/locations/*}/conversationDatasets"
		if err := mux.HandlePath("GET", listPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.ListConversationDatasetsRequest{
				Parent: pathParams["parent"],
			}

			resp, err := client.ListConversationDatasets(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, resp, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// DELETE {prefix}/{name=projects/*/locations/*/conversationDatasets/*}
		deletePath := prefix + "/{name=projects/*/locations/*/conversationDatasets/*}"
		if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.DeleteConversationDatasetRequest{
				Name: pathParams["name"],
			}

			op, err := client.DeleteConversationDataset(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
		}); err != nil {
			return err
		}
	}

	return nil
}
