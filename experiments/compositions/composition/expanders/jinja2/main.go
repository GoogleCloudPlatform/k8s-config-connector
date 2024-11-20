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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	"google.golang.org/grpc"
	"tailscale.com/atomicfile"
)

var (
	port = flag.Int("port", 8443, "The server port")
)

// server is used to implement expander.Evaluator
type server struct {
	pb.UnimplementedExpanderServer
}

type expander struct {
	path     string
	context  []byte
	config   []byte
	value    []byte
	facade   []byte
	resource string
}

// Verify the expander config/template
func (s *server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResult, error) {
	result := &pb.ValidateResult{
		Status: pb.Status_SUCCESS,
		Error:  &pb.Error{},
	}

	//log.Printf("ValidateRequest:\n  config: %s\n  context: %s\n  facade: %s\n  values: %s\n",
	// req.Config, req.Context, req.Facade, req.Value)
	dir, err := os.MkdirTemp("", "jinja2")
	if err != nil {
		newerr := fmt.Errorf("unexpected tmp file creation failure. %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	// cleanup
	defer func() {
		err := os.RemoveAll(dir)
		log.Printf("Error removing directory: %s, %v", dir, err)
	}()

	e := expander{
		path:     dir,
		context:  req.Context,
		facade:   req.Facade,
		value:    req.Value,
		config:   req.Config,
		resource: req.Resource,
	}
	if err = e.WriteInputsToFileSystem(); err != nil {
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

	//log.Printf("EvaluateRequest:\n  config: %s\n  context: %s\n  facade: "
	//    "%s\n  values: %s\n", req.Config, req.Context, req.Facade, req.Value)
	dir, err := os.MkdirTemp("", "jinja2")
	if err != nil {
		newerr := fmt.Errorf("unexpected tmp file creation failure. %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	// cleanup
	defer func() {
		err := os.RemoveAll(dir)
		log.Printf("Error removing directory: %s, %v", dir, err)
	}()

	e := expander{
		path:     dir,
		context:  req.Context,
		facade:   req.Facade,
		value:    req.Value,
		config:   req.Config,
		resource: req.Resource,
	}
	if err = e.WriteInputsToFileSystem(); err != nil {
		newerr := fmt.Errorf("error processing inputs: %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}

	return e.Evaluate(result)
}

func (e *expander) WriteInputsToFileSystem() error {
	context := map[string]interface{}{}
	if string(e.context) != "" {
		err := json.Unmarshal(e.context, &context)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Context: %w", err)
		}
	}

	facade := map[string]interface{}{}
	if string(e.facade) != "" {
		err := json.Unmarshal(e.facade, &facade)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Facade: %w", err)
		}
	}

	getterValues := map[string]interface{}{}
	if string(e.value) != "" {
		err := json.Unmarshal(e.value, &getterValues)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Values: %w", err)
		}
	}

	valuesObj := map[string]interface{}{
		"context":  context,
		e.resource: facade,
		"values":   getterValues,
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
	err = atomicfile.WriteFile(filepath.Join(e.path, "/template"), e.config, 0644)
	if err != nil {
		return fmt.Errorf("failed to write req.Config to file: %w", err)
	}

	return nil
}

// nolint:unparam
func (e *expander) Validate(result *pb.ValidateResult) (*pb.ValidateResult, error) {
	// usage: python parse_template.py <file>
	args := []string{
		"parse_template.py",
		filepath.Join(e.path, "/template"),
	}

	op, err := exec.Command("python", args...).CombinedOutput()
	if err != nil {
		result.Error.Message = fmt.Sprintf("failed validating template:\n %s", string(op))
		result.Status = pb.Status_VALIDATE_FAILED
		log.Print(result.Error.Message)
		return result, nil
	}

	//log.Printf("Result:\n  type: %s\n  error: %s\n  status: %s\n", result.Type, result.Errors, result.Status)
	return result, nil
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

	//log.Printf("Result:\n  type: %s\n  error: %s\n  status: %s\n "
	//" manifests: %s\n  values: %s\n", result.Type, result.Errors,
	// result.Status, result.Manifests, result.Values)
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
