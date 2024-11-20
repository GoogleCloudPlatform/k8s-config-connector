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
	"strings"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	helmconfigurationv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/expander/helm-expander/api/v1alpha1"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
	"tailscale.com/atomicfile"
)

var (
	port = flag.Int("port", 8443, "The server port")
)

type Expander struct {
	config        *helmconfigurationv1alpha1.HelmConfiguration
	facade        *unstructured.Unstructured
	context       *unstructured.Unstructured
	fetchedValues *map[string]interface{}
	inputResource string
	ctx           context.Context
	path          string
}

func NewExpander(ctx context.Context, req *pb.EvaluateRequest, vreq *pb.ValidateRequest) (*Expander, string, error) {
	var configBytes []byte
	var facadeBytes []byte
	var contextBytes []byte
	var fetchedValueBytes []byte
	var inputResource string

	failedMessage := ""

	e := &Expander{
		ctx: ctx,
	}

	if req != nil {
		configBytes = req.Config
		facadeBytes = req.Facade
		contextBytes = req.Context
		fetchedValueBytes = req.Value
		inputResource = req.Resource
		if string(facadeBytes) == "" {
			return nil, failedMessage, fmt.Errorf("Empty Facade for an Evaluate call")
		}
	}
	if vreq != nil {
		configBytes = vreq.Config
		facadeBytes = vreq.Facade
		contextBytes = vreq.Context
		fetchedValueBytes = vreq.Value
		inputResource = vreq.Resource
	}

	e.inputResource = inputResource
	e.config = &helmconfigurationv1alpha1.HelmConfiguration{}

	if string(configBytes) == "" {
		failedMessage = fmt.Sprintf("empty Config passed")
		return nil, failedMessage, nil
	}
	err := json.Unmarshal(configBytes, e.config)
	if err != nil {
		return nil, failedMessage, fmt.Errorf("error unmarshalling req.Config into HelmConfiguration CR: %w", err)
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

func (e *Expander) WriteInputsToFileSystem() error {
	// Write Chart.yaml to file
	yamlContent, err := yaml.JSONToYAML(e.config.Spec.Chart.Raw)
	if err != nil {
		return fmt.Errorf("failed to marshall chart.yaml: %w", err)
	}
	err = atomicfile.WriteFile(filepath.Join(e.path, "Chart.yaml"), yamlContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write Chart file: %w", err)
	}

	// Write values.yaml to file
	if len(e.config.Spec.DefaultValues.Raw) != 0 {
		yamlContent, err := yaml.JSONToYAML(e.config.Spec.DefaultValues.Raw)
		if err != nil {
			return fmt.Errorf("failed to marshall values.yaml: %w", err)
		}
		err = atomicfile.WriteFile(filepath.Join(e.path, "values.yaml"), yamlContent, 0644)
		if err != nil {
			return fmt.Errorf("failed to write values.yaml file: %w", err)
		}
	}

	// Write CRDs
	for _, crd := range e.config.Spec.CRDs {
		yamlContent, err := yaml.JSONToYAML(crd.Content.Raw)
		if err != nil {
			return fmt.Errorf("failed to marshall crds/%s file to yaml: %w", crd.FileName, err)
		}
		err = atomicfile.WriteFile(filepath.Join(e.path, "crds", crd.FileName), yamlContent, 0644)
		if err != nil {
			return fmt.Errorf("failed to write crds/%s file: %w", crd.FileName, err)
		}
	}

	// Write template files
	for _, template := range e.config.Spec.Templates {
		content := []byte{}
		if len(template.Content.Raw) != 0 {
			yamlContent, err := yaml.JSONToYAML(template.Content.Raw)
			if err != nil {
				return fmt.Errorf("failed to marshall crds/%s file to yaml: %w", template.FileName, err)
			}
			content = yamlContent
		} else if template.Template != "" {
			content = []byte(template.Template)
		}
		err = atomicfile.WriteFile(filepath.Join(e.path, "templates", template.FileName), content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write templates/%s file: %w", template.FileName, err)
		}
	}

	valuesObj := map[string]interface{}{}

	// Write Facade as values ?
	if e.context != nil {
		spec, found, err := unstructured.NestedMap(e.context.Object, "spec")
		if err != nil {
			return fmt.Errorf("failed to get .spec from Context: %w", err)
		}
		if !found {
			return fmt.Errorf("Missing .spec from Context")
		}
		valuesObj["context"] = spec
	}
	if e.fetchedValues != nil {
		valuesObj["fetched"] = *e.fetchedValues
	}
	if e.facade != nil {
		valuesObj[e.inputResource] = e.facade.Object
	}
	valuesBytes, err := json.Marshal(valuesObj)
	if err != nil {
		return fmt.Errorf("Unable to marshall values object")
	}
	err = atomicfile.WriteFile(filepath.Join(e.path, "facade-values.yaml"), valuesBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write facade-values.yaml file: %w", err)
	}
	return nil
}

func (e *Expander) Validate(result *pb.ValidateResult) (*pb.ValidateResult, error) {
	// https://helm.sh/docs/helm/helm_lint/
	// Usage:
	//  helm lint PATH [flags]
	//
	args := []string{"lint", e.path}
	if e.facade != nil {
		facadeFile := filepath.Join(e.path, "facade-values.yaml")
		args = append(args, "-f", facadeFile)
	}

	op, err := exec.Command("helm", args...).CombinedOutput()
	if err != nil {
		result.Error.Message = fmt.Sprintf("failed validating template:\n %s", string(op))
		result.Status = pb.Status_VALIDATE_FAILED
		log.Print(result.Error.Message)
		return result, nil
	}

	//log.Printf("Result:\n  type: %s\n  error: %s\n  status: %s\n", result.Type, result.Errors, result.Status)
	return result, nil
}

func (e *Expander) Evaluate(result *pb.EvaluateResult) (*pb.EvaluateResult, error) {
	// https://helm.sh/docs/helm/helm_template/
	// Usage:
	//  helm template [NAME] [CHART] [flags]
	//
	name := "unknown"
	args := []string{"template", name, e.path}
	if e.facade != nil {
		name = e.facade.GetName()
		facadeFile := filepath.Join(e.path, "facade-values.yaml")
		args = append(args, "-f", facadeFile)
	}

	//args = append(args, "--output-dir", filepath.Join(e.path, "rendered"))

	log.Printf("running: helm %s", strings.Join(args, " "))
	op, err := exec.Command("helm", args...).CombinedOutput()
	if err != nil {
		result.Error.Message = fmt.Sprintf("failed evaluating helm chart:\n %s", string(op))
		result.Status = pb.Status_EVALUATE_FAILED
		log.Print(result.Error.Message)
		return result, nil
	}

	manifests := op
	/*
		// If using output-dir we need to collect the generated files
		//args = append(args, "--output-dir", filepath.Join(e.path, "rendered"))
			manifests := []byte{}
			// parse output and read files
			lines := strings.Split(string(op), "\n")
			for _, line := range lines {
				log.Print(line)
				if strings.HasPrefix(line, "wrote ") {
					file := strings.TrimPrefix(line, "wrote ")
					// read file
					fileContent, err := os.ReadFile(file)
					if err != nil {
						newerr := fmt.Errorf("failed reading expanded file: %s, %w", file, err)
						log.Print(newerr)
						return nil, newerr
					}
					manifests = append(manifests, fileContent...)
				}
			}
	*/
	result.Manifests = manifests

	//log.Printf("Result:\n  type: %s\n  error: %s\n  status: %s\n  manifests: %s\n  values: %s\n", result.Type, result.Errors, result.Status, result.Manifests, result.Values)
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

	dir, err := os.MkdirTemp("", "validate")
	if err != nil {
		newerr := fmt.Errorf("unexpected tmp file creation failure. %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	err = os.Mkdir(filepath.Join(dir, "templates"), 0700)
	if err != nil {
		newerr := fmt.Errorf("%s/templates dir creation failed: %w", dir, err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	err = os.Mkdir(filepath.Join(dir, "crds"), 0700)
	if err != nil {
		newerr := fmt.Errorf("%s/crds dir creation failed: %w", dir, err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	// cleanup
	defer os.RemoveAll(dir)

	e.path = dir

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

	dir, err := os.MkdirTemp("", "eval")
	if err != nil {
		newerr := fmt.Errorf("tmp eval dir creation failure. %w", err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	err = os.Mkdir(filepath.Join(dir, "templates"), 0700)
	if err != nil {
		newerr := fmt.Errorf("%s/templates dir creation failed: %w", dir, err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	err = os.Mkdir(filepath.Join(dir, "crds"), 0700)
	if err != nil {
		newerr := fmt.Errorf("%s/crds dir creation failed: %w", dir, err)
		log.Print(newerr.Error())
		return nil, newerr
	}
	// cleanup
	defer os.RemoveAll(dir)

	e.path = dir

	if err = e.WriteInputsToFileSystem(); err != nil {
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
