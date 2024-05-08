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
	"os"
	"os/exec"
	"path/filepath"

	pb "google.com/composition/proto"
	"google.golang.org/grpc"
	"tailscale.com/atomicfile"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement exander.Evaluator
type server struct {
	pb.UnimplementedExpanderServer
}

type expander struct {
	path string
	req  *pb.EvaluateRequest
}

// Verify the expander config/template
func (s *server) Validate(context.Context, *pb.ValidateRequest) (*pb.ValidateResult, error) {
	return &pb.ValidateResult{Status: pb.Status_SUCCESS}, nil
}

// Evaluate the expander config in context of inputs and return manifests
func (s *server) Evaluate(ctx context.Context, req *pb.EvaluateRequest) (*pb.EvaluateResult, error) {
	result := &pb.EvaluateResult{
		Status: pb.Status_SUCCESS,
		Type:   pb.ResultType_MANIFESTS,
		Error:  &pb.Error{},
	}

	//log.Printf("EvaluateRequest:\n  config: %s\n  context: %s\n  facade: %s\n  values: %s\n", req.Config, req.Context, req.Facade, req.Value)
	dir, err := os.MkdirTemp("", "jinja2")
	if err != nil {
		newerr := fmt.Errorf("unexpected tmp file creation failure. %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	// cleanup
	defer os.RemoveAll(dir)

	e := expander{path: dir, req: req}
	if err = e.WriteInputsToFileSystem(); err != nil {
		newerr := fmt.Errorf("error processing inputs: %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}

	return e.Evaluate(result)
}

func (e *expander) WriteInputsToFileSystem() error {
	context := map[string]interface{}{}
	if string(e.req.Context) != "" {
		err := json.Unmarshal(e.req.Context, &context)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Context: %w", err)
		}
	}

	facade := map[string]interface{}{}
	err := json.Unmarshal(e.req.Facade, &facade)
	if err != nil {
		return fmt.Errorf("error unmarshalling req.Facade: %w", err)
	}

	getterValues := map[string]interface{}{}
	if string(e.req.Value) != "" {
		err := json.Unmarshal(e.req.Value, &getterValues)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Values: %w", err)
		}
	}

	valuesObj := map[string]interface{}{
		"context":      context,
		e.req.Resource: facade,
		"values":       getterValues,
	}

	// marshall values
	values, err := json.Marshal(valuesObj)
	if err != nil {
		return fmt.Errorf("failed to marshal values: %w", err)
	}

	// Write values to file
	err = atomicfile.WriteFile(filepath.Join(e.path, "/values"), values, 0644)
	if err != nil {
		return fmt.Errorf("failed to write req.[context,facade,values] to file: %w", err)
	}

	// Write template to file
	err = atomicfile.WriteFile(filepath.Join(e.path, "/template"), e.req.Config, 0644)
	if err != nil {
		return fmt.Errorf("failed to write req.Config to file: %w", err)
	}

	return nil
}

func (e *expander) Evaluate(result *pb.EvaluateResult) (*pb.EvaluateResult, error) {
	// wait for /expanded/expanded to be created and then read it
	args := []string{
		filepath.Join(e.path, "/template"),
		filepath.Join(e.path, "/values"),
		"--format=json",
		"-o",
		filepath.Join(e.path, "/expanded"),
	}

	op, err := exec.Command("jinja2", args...).CombinedOutput()
	if err != nil {
		result.Error.Message = fmt.Sprintf("failed evaluating template:\n %s", string(op))
		result.Status = pb.Status_EVALUATE_FAILED
		log.Print(result.Error.Message)
		return result, nil
	}

	manifests, err := os.ReadFile(filepath.Join(e.path, "/expanded"))
	if err != nil {
		newerr := fmt.Errorf("failed reading expanded file: %w", err)
		log.Print(newerr)
		return nil, newerr
	}

	result.Manifests = manifests

	//log.Printf("Result:\n  type: %s\n  error: %s\n  status: %s\n  manifests: %s\n  values: %s\n", result.Type, result.Errors, result.Status, result.Manifests, result.Values)
	return result, nil
}

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
