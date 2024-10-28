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

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	celconfigurationv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/expander/cel-expander/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/expander/cel-expander/pkg/cel"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/expander/cel-expander/pkg/resource"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

var (
	port = flag.Int("port", 8443, "The server port")
)

type Expander struct {
	resource      string
	config        *celconfigurationv1alpha1.CELConfiguration
	facade        *unstructured.Unstructured
	context       *unstructured.Unstructured
	fetchedValues *map[string]interface{}
	cel           *cel.Engine
	resources     []*resource.Resource
	ctx           context.Context
}

func NewExpander(ctx context.Context, req *pb.EvaluateRequest, vreq *pb.ValidateRequest) (*Expander, string, error) {
	var configBytes []byte
	var facadeBytes []byte
	var contextBytes []byte
	var fetchedValueBytes []byte
	var resource string

	failedMessage := ""

	e := &Expander{
		ctx: ctx,
	}

	if req != nil {
		configBytes = req.Config
		facadeBytes = req.Facade
		contextBytes = req.Context
		fetchedValueBytes = req.Value
		resource = req.Resource
		if string(facadeBytes) == "" {
			return nil, failedMessage, fmt.Errorf("Empty Facade for an Evaluate call")
		}
	}
	if vreq != nil {
		configBytes = vreq.Config
		facadeBytes = vreq.Facade
		contextBytes = vreq.Context
		fetchedValueBytes = vreq.Value
		resource = vreq.Resource
	}

	e.config = &celconfigurationv1alpha1.CELConfiguration{}
	e.resource = resource

	if string(configBytes) == "" {
		failedMessage = fmt.Sprintf("empty Config passed")
		return nil, failedMessage, nil
	}
	err := json.Unmarshal(configBytes, e.config)
	if err != nil {
		return nil, failedMessage, fmt.Errorf("error unmarshalling req.Config into CELConfiguration CR: %w", err)
	}

	if string(contextBytes) != "" {
		e.context = &unstructured.Unstructured{}
		err := e.context.UnmarshalJSON(contextBytes)
		if err != nil {
			return nil, failedMessage, fmt.Errorf("error unmarshalling req.Context: %w", err)
		}
	}

	if string(fetchedValueBytes) != "" {
		e.fetchedValues = &map[string]interface{}{}
		err := json.Unmarshal(fetchedValueBytes, e.fetchedValues)
		if err != nil {
			return nil, failedMessage, fmt.Errorf("error unmarshalling req.Context: %w", err)
		}
	}

	if string(facadeBytes) != "" {
		e.facade = &unstructured.Unstructured{}
		err := e.facade.UnmarshalJSON(facadeBytes)
		if err != nil {
			return nil, failedMessage, fmt.Errorf("error unmarshalling req.Facade into Unstructured: %w", err)
		}
		if e.facade.GetNamespace() == "" {
			//e.facade.SetNamespace("default")
			failedMessage = fmt.Sprintf("missing namespace in req.Facade object")
			return nil, failedMessage, nil
		}
	}

	return e, failedMessage, nil
}

func (e *Expander) ProcessInputs() error {
	var err error

	values := map[string]interface{}{}
	// Setup values
	if e.facade != nil {
		values[e.resource] = e.facade.Object
	}
	if e.context != nil {
		values["context"] = e.context.Object
	}
	if e.fetchedValues != nil {
		values["fetched"] = *e.fetchedValues
	}

	// Create a CEL engine
	e.cel, err = cel.NewEngine(e.resource, values)
	if err != nil {
		return fmt.Errorf("error creating CEL engine: %w", err)
	}

	// load resources
	for _, rsrc := range e.config.Spec.Resources {
		yamlContent, err := yaml.JSONToYAML(rsrc.Definition.Raw)
		if err != nil {
			return fmt.Errorf("failed to marshall resource:%s file to yaml: %w", rsrc.Name, err)
		}
		r, err := resource.NewResourceFromRaw(rsrc.Name, yamlContent)
		if err != nil {
			return fmt.Errorf("error creating resource for %s, %w", rsrc.Name, err)
		}
		e.resources = append(e.resources, r)
	}
	return nil

}

func (e *Expander) Validate(result *pb.ValidateResult) (*pb.ValidateResult, error) {
	return result, nil
}

func (e *Expander) Evaluate(result *pb.EvaluateResult) (*pb.EvaluateResult, error) {
	// Loop through resources
	manifests := []byte{}
	for _, rsrc := range e.resources {
		// loop through variables
		for vindex := range rsrc.Variables {
			// cel.eval()
			result, err := e.cel.Eval(rsrc.Variables[vindex].CELExpression)
			if err != nil {
				// TODO: consume the error and mark result failed ?
				return nil, fmt.Errorf("error Evaluating expression: %s, %w", rsrc.Variables[vindex].Expression, err)
			}
			rsrc.Variables[vindex].ResolvedValue = result
		}
		// Replace variables
		err := rsrc.ApplyResolvedVariables()
		if err != nil {
			return nil, fmt.Errorf("error applying resolved variables for %s, %w", rsrc.Name, err)
		}
		manifests = append(manifests, []byte("\n---\n")...)
		manifests = append(manifests, rsrc.Raw...)
	}

	// return objects
	result.Manifests = manifests
	return result, nil
}

// ------------- GRPC Server implementation ----------------

// server is used to implement expander.Evaluator
type server struct {
	pb.UnimplementedExpanderServer
}

// Verify the expander config/template
func (s *server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResult, error) {
	result := &pb.ValidateResult{
		Status: pb.Status_SUCCESS,
		Error:  &pb.Error{},
	}

	log.Printf("Validate called")
	e, failedMessage, err := NewExpander(ctx, nil, req)
	if err != nil {
		return nil, err
	}
	if failedMessage != "" {
		result.Error.Message = failedMessage
		result.Status = pb.Status_VALIDATE_FAILED
		return result, nil
	}

	if err = e.ProcessInputs(); err != nil {
		newerr := fmt.Errorf("error processing inputs: %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}

	return e.Validate(result)
}

// Evaluate the expander config in context of inputs and return manifests
func (s *server) Evaluate(ctx context.Context, req *pb.EvaluateRequest) (*pb.EvaluateResult, error) {
	result := &pb.EvaluateResult{
		Status: pb.Status_SUCCESS,
		Type:   pb.ResultType_MANIFESTS,
		Error:  &pb.Error{},
	}

	log.Printf("Evaluate called")
	e, failedMessage, err := NewExpander(ctx, req, nil)
	if err != nil {
		return nil, err
	}
	if failedMessage != "" {
		result.Error.Message = failedMessage
		result.Status = pb.Status_EVALUATE_FAILED
		return result, nil
	}

	if err = e.ProcessInputs(); err != nil {
		newerr := fmt.Errorf("error processing inputs: %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}

	return e.Evaluate(result)
}

// -------------------- MAIN -------------------------------

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExpanderServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
