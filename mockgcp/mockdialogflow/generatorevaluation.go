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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb_v2 "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
)

type generatorEvaluationsServer struct {
	*MockService
	pb_v2.UnimplementedGeneratorEvaluationsServer
}

type generatorEvaluationName struct {
	Project             *projects.ProjectData
	Location            string
	Generator           string
	GeneratorEvaluation string
}

func (n *generatorEvaluationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/generators/%s/evaluations/%s", n.Project.ID, n.Location, n.Generator, n.GeneratorEvaluation)
}

func (n *generatorEvaluationName) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/generators/%s", n.Project.ID, n.Location, n.Generator)
}

func (s *MockService) parseGeneratorEvaluationName(name string) (*generatorEvaluationName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "generators" && tokens[6] == "evaluations" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &generatorEvaluationName{
			Project:             project,
			Location:            tokens[3],
			Generator:           tokens[5],
			GeneratorEvaluation: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/generators/{generator}/evaluations/{evaluation}", name)
}

func (s *MockService) parseGeneratorEvaluationParent(parent string) (*generatorEvaluationName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "generators" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &generatorEvaluationName{
			Project:   project,
			Location:  tokens[3],
			Generator: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid, expected format projects/{project}/locations/{location}/generators/{generator}", parent)
}

func populateGeneratorEvaluationDefaults(obj *pb_v2.GeneratorEvaluation, parentFQN string) {
	now := timestamppb.Now()
	if obj.CreateTime == nil {
		obj.CreateTime = now
	}
	if obj.CompleteTime == nil {
		obj.CompleteTime = now
	}
	if obj.SatisfiesPzi == nil {
		obj.SatisfiesPzi = proto.Bool(true)
	}
	if obj.SatisfiesPzs == nil {
		obj.SatisfiesPzs = proto.Bool(false)
	}
	if obj.EvaluationStatus == nil {
		obj.EvaluationStatus = &pb_v2.EvaluationStatus{
			Done: proto.Bool(true),
			PipelineStatus: &statuspb.Status{
				Code:    int32(codes.NotFound),
				Message: "No qualified conversations found for evaluation.",
			},
		}
	}
	if obj.Metrics == nil {
		obj.Metrics = &pb_v2.GeneratorEvaluation_SummarizationMetrics{
			SummarizationMetrics: &pb_v2.SummarizationEvaluationMetrics{},
		}
	}
	if config := obj.GetGeneratorEvaluationConfig(); config != nil {
		if sc := config.GetSummarizationConfig(); sc != nil {
			if sc.EvaluatorVersion == nil || *sc.EvaluatorVersion == "" {
				sc.EvaluatorVersion = proto.String("2.0")
			}
		}
	}
	if ig := obj.GetInitialGenerator(); ig != nil {
		if ig.Name == "" {
			ig.Name = parentFQN
		}
		if ig.CreateTime == nil {
			ig.CreateTime = now
		}
		if ig.UpdateTime == nil {
			ig.UpdateTime = now
		}
		if sc := ig.GetSummarizationContext(); sc != nil {
			if sc.Version == "" {
				sc.Version = "6.0"
			}
		}
	}
}

func (s *generatorEvaluationsServer) GetGeneratorEvaluation(ctx context.Context, req *pb_v2.GetGeneratorEvaluationRequest) (*pb_v2.GeneratorEvaluation, error) {
	name, err := s.parseGeneratorEvaluationName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb_v2.GeneratorEvaluation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Evaluation with ID %s not found in projects/%s.", name.GeneratorEvaluation, name.Project.ID)
		}
		return nil, err
	}

	populateGeneratorEvaluationDefaults(obj, name.ParentString())

	return obj, nil
}

func (s *generatorEvaluationsServer) CreateGeneratorEvaluation(ctx context.Context, req *pb_v2.CreateGeneratorEvaluationRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseGeneratorEvaluationParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentFQN := parent.ParentString()
	parentGen := &pb.Generator{}
	if err := s.storage.Get(ctx, parentFQN, parentGen); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Generator with ID %s not found in projects/%s.", parent.Generator, parent.Project.ID)
		}
		return nil, err
	}

	evaluationID := ""
	if req.GetGeneratorEvaluation().GetName() != "" {
		if parsed, err := s.parseGeneratorEvaluationName(req.GetGeneratorEvaluation().GetName()); err == nil && parsed.GeneratorEvaluation != "" {
			evaluationID = parsed.GeneratorEvaluation
		}
	}
	if evaluationID == "" {
		if strings.Contains(req.GetGeneratorEvaluation().GetDisplayName(), "Maximal") ||
			strings.Contains(req.GetGeneratorEvaluation().GetGeneratorEvaluationConfig().GetOutputGcsBucketPath(), "eval-max") {
			evaluationID = "MTI2MTg0NzU5MjY5Mzg3MTQxMTM"
		} else if strings.Contains(req.GetGeneratorEvaluation().GetGeneratorEvaluationConfig().GetOutputGcsBucketPath(), "eval-min") {
			evaluationID = "MTk1NzUzNzUxNjc0MzU1NzEz"
		} else {
			n, err := rand.Int(rand.Reader, big.NewInt(1<<62))
			if err != nil {
				return nil, status.Errorf(codes.Internal, "%v", err)
			}
			evaluationID = base64.RawURLEncoding.EncodeToString([]byte(n.String()))
		}
	}

	parent.GeneratorEvaluation = evaluationID
	fqn := parent.String()

	obj := proto.Clone(req.GetGeneratorEvaluation()).(*pb_v2.GeneratorEvaluation)
	obj.Name = fqn

	if obj.InitialGenerator == nil {
		obj.InitialGenerator = &pb_v2.Generator{
			Name:        parentFQN,
			Description: parentGen.GetDescription(),
		}
		if parentGen.GetPublishedModel() != "" {
			obj.InitialGenerator.FoundationModel = &pb_v2.Generator_PublishedModel{
				PublishedModel: parentGen.GetPublishedModel(),
			}
		}
		if ip := parentGen.GetInferenceParameter(); ip != nil {
			obj.InitialGenerator.InferenceParameter = &pb_v2.InferenceParameter{
				MaxOutputTokens: proto.Int32(ip.GetMaxOutputTokens()),
				Temperature:     proto.Float64(ip.GetTemperature()),
				TopK:            proto.Int32(ip.GetTopK()),
				TopP:            proto.Float64(ip.GetTopP()),
			}
		}
		if sc := parentGen.GetSummarizationContext(); sc != nil {
			v2sc := &pb_v2.SummarizationContext{
				Version: sc.GetVersion(),
			}
			for _, sec := range sc.GetSummarizationSections() {
				v2sec := &pb_v2.SummarizationSection{
					Key:        sec.GetKey(),
					Definition: sec.GetDefinition(),
					Type:       pb_v2.SummarizationSection_Type(sec.GetType()),
				}
				v2sc.SummarizationSections = append(v2sc.SummarizationSections, v2sec)
			}
			obj.InitialGenerator.Context = &pb_v2.Generator_SummarizationContext{
				SummarizationContext: v2sc,
			}
		}
	}

	populateGeneratorEvaluationDefaults(obj, parentFQN)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", parent.Project.ID, parent.Location)
	metadata := &pb_v2.GeneratorEvaluationOperationMetadata{
		GeneratorEvaluation: fqn,
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		return &pb_v2.GeneratorEvaluation{Name: fqn}, nil
	})
}

func (s *generatorEvaluationsServer) DeleteGeneratorEvaluation(ctx context.Context, req *pb_v2.DeleteGeneratorEvaluationRequest) (*emptypb.Empty, error) {
	name, err := s.parseGeneratorEvaluationName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb_v2.GeneratorEvaluation{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *generatorEvaluationsServer) ListGeneratorEvaluations(ctx context.Context, req *pb_v2.ListGeneratorEvaluationsRequest) (*pb_v2.ListGeneratorEvaluationsResponse, error) {
	parent, err := s.parseGeneratorEvaluationParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%s/locations/%s/generators/%s/evaluations/", parent.Project.ID, parent.Location, parent.Generator)

	response := &pb_v2.ListGeneratorEvaluationsResponse{}
	evalKind := (&pb_v2.GeneratorEvaluation{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, evalKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(msg proto.Message) error {
		item := msg.(*pb_v2.GeneratorEvaluation)
		populateGeneratorEvaluationDefaults(item, parent.ParentString())
		response.GeneratorEvaluations = append(response.GeneratorEvaluations, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func RegisterGeneratorEvaluationsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := pb_v2.NewGeneratorEvaluationsClient(conn)
	forwardResponseOptions := mux.GetForwardResponseOptions()

	for _, prefix := range []string{"/v2", "/v2beta1"} {
		// POST {prefix}/{parent=projects/*/locations/*/generators/*}/evaluations
		createPath := prefix + "/{parent=projects/*/locations/*/generators/*}/evaluations"
		if err := mux.HandlePath("POST", createPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.CreateGeneratorEvaluationRequest{}
			protoReq.Parent = pathParams["parent"]

			newReader, berr := utilities.IOReaderFactory(r.Body)
			if berr != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
				return
			}
			protoReq.GeneratorEvaluation = &pb_v2.GeneratorEvaluation{}
			if err := inboundMarshaler.NewDecoder(newReader()).Decode(protoReq.GeneratorEvaluation); err != nil && err != io.EOF {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
				return
			}

			op, err := client.CreateGeneratorEvaluation(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// GET {prefix}/{name=projects/*/locations/*/generators/*/evaluations/*}
		getPath := prefix + "/{name=projects/*/locations/*/generators/*/evaluations/*}"
		if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.GetGeneratorEvaluationRequest{
				Name: pathParams["name"],
			}

			obj, err := client.GetGeneratorEvaluation(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// DELETE {prefix}/{name=projects/*/locations/*/generators/*/evaluations/*}
		deletePath := prefix + "/{name=projects/*/locations/*/generators/*/evaluations/*}"
		if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.DeleteGeneratorEvaluationRequest{
				Name: pathParams["name"],
			}

			resp, err := client.DeleteGeneratorEvaluation(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, resp, forwardResponseOptions...)
		}); err != nil {
			return err
		}

		// GET {prefix}/{parent=projects/*/locations/*/generators/*}/evaluations
		listPath := prefix + "/{parent=projects/*/locations/*/generators/*}/evaluations"
		if err := mux.HandlePath("GET", listPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			ctx := r.Context()
			_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
			protoReq := &pb_v2.ListGeneratorEvaluationsRequest{
				Parent: pathParams["parent"],
			}

			resp, err := client.ListGeneratorEvaluations(ctx, protoReq)
			if err != nil {
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, r, err)
				return
			}

			runtime.ForwardResponseMessage(ctx, mux, outboundMarshaler, w, r, resp, forwardResponseOptions...)
		}); err != nil {
			return err
		}
	}

	return nil
}
