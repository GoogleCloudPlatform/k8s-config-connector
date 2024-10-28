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
	"strings"

	compositionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/api/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	port = flag.Int("port", 8443, "The server port")
)

type Getter struct {
	req    *pb.EvaluateRequest
	facade *unstructured.Unstructured
	getter *compositionv1alpha1.GetterConfiguration
	ctx    context.Context
	client dynamic.Interface
	values map[string]interface{}
}

func NewGetter(ctx context.Context, client dynamic.Interface, req *pb.EvaluateRequest) *Getter {
	return &Getter{
		req:    req,
		client: client,
		ctx:    ctx,
		values: make(map[string]interface{}),
	}
}

func (g *Getter) LoadInputs() error {
	g.getter = &compositionv1alpha1.GetterConfiguration{}
	if string(g.req.Config) != "" {
		err := json.Unmarshal(g.req.Config, g.getter)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Config into GetterConfiguration{}: %w", err)
		}
	} else {
		g.getter.Spec.ValuesFrom = []compositionv1alpha1.ValuesFrom{}
	}

	context := &compositionv1alpha1.Context{}
	if string(g.req.Context) != "" {
		err := json.Unmarshal(g.req.Context, &context)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Context: %w", err)
		}
	}

	g.facade = &unstructured.Unstructured{}
	err := g.facade.UnmarshalJSON(g.req.Facade)
	if err != nil {
		return fmt.Errorf("error unmarshalling req.Facade into Unstructured: %w", err)
	}

	if g.facade.GetNamespace() == "" {
		return fmt.Errorf("missing namespace in req.Facade object")
	}

	getterValues := map[string]interface{}{}
	if string(g.req.Value) != "" {
		err := json.Unmarshal(g.req.Value, &getterValues)
		if err != nil {
			return fmt.Errorf("error unmarshalling req.Values: %w", err)
		}
	}

	return nil
}

func (g *Getter) updateValues(obj *unstructured.Unstructured, vf *compositionv1alpha1.ValuesFrom) (string, bool) {
	wait := false
	for index := range vf.FieldRef {
		fr := &vf.FieldRef[index]
		path := strings.Split(strings.TrimLeft(fr.Path, "."), ".")
		identifier := ""
		gvk := obj.GroupVersionKind()
		if obj.GetNamespace() != "" {
			identifier = fmt.Sprintf("%s.%s(%s)/%s/%s[%s]",
				gvk.Kind, gvk.Group, gvk.Version,
				obj.GetNamespace(), obj.GetName(), fr.Path)
		} else {
			identifier = fmt.Sprintf("%s.%s(%s)/%s[%s]",
				gvk.Kind, gvk.Group, gvk.Version, obj.GetName(), fr.Path)
		}
		v, ok, err := unstructured.NestedFieldCopy(obj.Object, path...)
		if err != nil {
			message := fmt.Sprintf("Error traversing field path: %s", identifier)
			log.Print(message)
			return message, wait
		}
		if ok {
			if g.values[vf.Name] == nil {
				g.values[vf.Name] = map[string]interface{}{}
			}
			g.values[vf.Name].(map[string]interface{})[fr.As] = v
		} else {
			message := fmt.Sprintf("Field path not present in object yet: %s", identifier)
			log.Print(message)
			wait = true
			return message, wait
		}
	}
	return "", wait
}

func (g *Getter) getObject(vf *compositionv1alpha1.ValuesFrom,
	name string) (*unstructured.Unstructured, string, bool, error) {
	failMessage := ""
	wait := false
	gvr := schema.GroupVersionResource{
		Group:    vf.ResourceRef.Group,
		Version:  vf.ResourceRef.Version,
		Resource: vf.ResourceRef.Resource,
	}
	namespace := g.facade.GetNamespace()
	identifier := ""
	if namespace != "" {
		identifier = fmt.Sprintf("%s.%s(%s)/%s/%s", gvr.Resource, gvr.Group, gvr.Version, namespace, name)
	} else {
		identifier = fmt.Sprintf("%s.%s(%s)/%s", gvr.Resource, gvr.Group, gvr.Version, name)
	}
	log.Printf("Fetching :%s", identifier)
	obj, err := g.client.Resource(gvr).Namespace(namespace).Get(g.ctx, name, metav1.GetOptions{})
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			log.Printf("Error getting dependent object: %s", identifier)
			return nil, failMessage, wait, err
		}
		failMessage = fmt.Sprintf("Dependent object not found: GVR: %s", identifier)
		wait := true
		return nil, failMessage, wait, nil
	}
	return obj, failMessage, wait, nil
}

func (g *Getter) Fetch() (string, bool, error) {
	wait := false
	message := ""
	for index := range g.getter.Spec.ValuesFrom {
		vf := &g.getter.Spec.ValuesFrom[index]
		name := vf.ResourceRef.Name
		if name == "" {
			name = g.facade.GetName() + vf.ResourceRef.NameSuffix
		}
		obj, failMessage, wait, err := g.getObject(vf, name)
		if err != nil || failMessage != "" || wait {
			return failMessage, wait, err
		}

		message, wait = g.updateValues(obj, vf)
		if message != "" || wait {
			return message, wait, nil
		}
	}
	return "", wait, nil
}

func (g *Getter) GetValues() map[string]interface{} {
	return g.values
}

// ------------- GRPC Server implementation ----------------

// server is used to implement expander.Evaluator
type grpcServer struct {
	pb.UnimplementedExpanderServer
	dynamicClient dynamic.Interface
}

// Verify the expander config/template
func (s *grpcServer) Validate(context.Context, *pb.ValidateRequest) (*pb.ValidateResult, error) {
	return &pb.ValidateResult{Status: pb.Status_SUCCESS}, nil
}

// Evaluate the expander config in context of inputs and return manifests
func (s *grpcServer) Evaluate(ctx context.Context, req *pb.EvaluateRequest) (*pb.EvaluateResult, error) {
	result := &pb.EvaluateResult{
		Status: pb.Status_SUCCESS,
		Type:   pb.ResultType_VALUES,
		Error:  &pb.Error{},
	}

	log.Printf("Evaluate called")
	g := NewGetter(ctx, s.dynamicClient, req)
	err := g.LoadInputs()
	if err != nil {
		return nil, err
	}
	failMessage, wait, err := g.Fetch()
	if err != nil {
		return nil, err
	}
	if failMessage != "" {
		result.Error.Message = failMessage
		if wait {
			result.Status = pb.Status_EVALUATE_WAIT
		} else {
			result.Status = pb.Status_EVALUATE_FAILED
		}
		return result, nil
	}

	values, err := json.Marshal(g.GetValues())
	if err != nil {
		log.Printf("Error marshalling extracted values: %v", err)
		return nil, fmt.Errorf("Failed to marshal extracted values: %v", err)
	}
	result.Values = values
	return result, nil
}

// -------------------- MAIN -------------------------------
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := &grpcServer{}

	config := ctrl.GetConfigOrDie()
	server.dynamicClient, err = dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("failed to get dynamic client error: %v", err)
	}

	pb.RegisterExpanderServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
