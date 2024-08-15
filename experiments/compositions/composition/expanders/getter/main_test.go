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

package main_test

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
	"testing"

	compositionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/api/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	addr = flag.String("addr", "[::]:8443", "the address to connect to")
)

// Test values
var (
	expanderClient pb.ExpanderClient

	facadeGVK schema.GroupVersionKind = schema.GroupVersionKind{
		Group:   "facade.composition.google.com",
		Version: "v1",
		Kind:    "AppTeam",
	}

	// Marshalled values
	sampleFacade             []byte
	emptyGetterConfiguration []byte

	// getter objects
	emptyGetterConfigurationObj *compositionv1alpha1.GetterConfiguration = &compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{},
		},
	}

	// Evaluate results
	emptyGetterConfigurationEvaluateResult = "{}"
)

// TestMain - umbrella test that runs all test cases
func TestMain(m *testing.M) {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	expanderClient = pb.NewExpanderClient(conn)

	// Set an empty facade object with just GVK
	facade := &unstructured.Unstructured{}
	facade.SetGroupVersionKind(facadeGVK)
	facade.SetName("foo")
	facade.SetNamespace("composition-system")
	sampleFacade, _ = json.Marshal(facade.Object)

	// Marshall GetterConfiguration config for test values
	emptyGetterConfiguration, err = json.Marshal(emptyGetterConfigurationObj)
	if err != nil {
		log.Fatalf("unable to marshall emptyGetterConfiguration: %v", err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestValidate(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:  []byte{},
			Context: []byte{},
			Facade:  []byte{},
			Value:   []byte{},
		})
	if err != nil {
		t.Fatalf("could not validate: %v", err)
	}
	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("expected SUCCESS, got: %s", r)
	}
}

func TestEvaluateEmptyConfig(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte{},
			Resource: "sqls",
			Context:  []byte("{\"project\": \"foo\"}"),
			Facade:   sampleFacade,
			Value:    []byte("{\"key\": \"val\"}"),
		})
	if err != nil {
		t.Fatalf("could not evaluate: %v", err)
	}
	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
}

func TestEvaluateEmptyContext(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   emptyGetterConfiguration,
			Resource: "sqls",
			Context:  []byte{},
			Facade:   sampleFacade,
			Value:    []byte("{\"key\": \"val\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Values) != emptyGetterConfigurationEvaluateResult {
		t.Fatalf("\nexpected: %s\n got: %s", emptyGetterConfigurationEvaluateResult, r.Values)
	}
}

func TestEvaluateEmptyResourceString(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:  emptyGetterConfiguration,
			Context: []byte{},
			Facade:  sampleFacade,
			Value:   []byte("{\"key\": \"val\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Values) != emptyGetterConfigurationEvaluateResult {
		t.Fatalf("\nexpected: %s\n got: %s", emptyGetterConfigurationEvaluateResult, r.Values)
	}
}

func TestEvaluateEmptyFacade(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   emptyGetterConfiguration,
			Resource: "sqls",
			Context:  []byte("{\"key\": \"val\"}"),
			Facade:   []byte{},
			Value:    []byte("{\"key\": \"val\"}"),
		})
	if err == nil {
		t.Fatalf("expected error. got nil")
	}
	if r != nil {
		t.Fatalf("\nexpected: nil result\n got: non-nil value")
	}
}

func evaluateGetterConfigurationMissingObject(t *testing.T, getter []byte, expectedErrorString string) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config: getter,
			Facade: sampleFacade,
		})
	if err != nil {
		t.Fatalf("expected no error. got %v", err)
	}

	if r.Status != pb.Status_EVALUATE_WAIT {
		t.Fatalf("\nexpected: EVALUATE_WAIT \n got: %s", r)
	}

	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %q \n got: %q", expectedErrorString, r)
	}
}

func TestEvaluateGetterConfigurationMissingObjectGroup(t *testing.T) {
	// Marshall GetterConfiguration config
	getter, err := json.Marshal(&compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{
				{
					Name: "missinggroup",
					ResourceRef: compositionv1alpha1.ResourceRef{
						Group:    "acme.something.com",
						Version:  "v1",
						Resource: "foobars",
						Kind:     "FooBar",
						Name:     "missing",
					},
					// Empty fieldref
					FieldRef: []compositionv1alpha1.FieldRef{},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to marshall getter: %v", err)
	}

	expectedErrorString := "Dependent object not found: GVR: foobars.acme.something.com(v1)/composition-system/missing"
	evaluateGetterConfigurationMissingObject(t, getter, expectedErrorString)
}

func TestEvaluateGetterConfigurationMissingObjectResource(t *testing.T) {
	// Marshall GetterConfiguration config
	getter, err := json.Marshal(&compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{
				{
					Name: "missingresource",
					ResourceRef: compositionv1alpha1.ResourceRef{
						Group:    "apps",
						Version:  "v1",
						Resource: "unknown",
						Kind:     "unknown",
						Name:     "missing",
					},
					// Empty fieldref
					FieldRef: []compositionv1alpha1.FieldRef{},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to marshall getter: %v", err)
	}

	expectedErrorString := "Dependent object not found: GVR: unknown.apps(v1)/composition-system/missing"
	evaluateGetterConfigurationMissingObject(t, getter, expectedErrorString)
}

func TestEvaluateGetterConfigurationMissingObject(t *testing.T) {
	// Marshall GetterConfiguration config
	getter, err := json.Marshal(&compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{
				{
					Name: "missing",
					ResourceRef: compositionv1alpha1.ResourceRef{
						Group:    "apps",
						Version:  "v1",
						Resource: "deployments",
						Kind:     "Deployment",
						Name:     "missing",
					},
					// Empty fieldref
					FieldRef: []compositionv1alpha1.FieldRef{},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to marshall getter: %v", err)
	}

	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config: getter,
			Facade: sampleFacade,
		})
	if err != nil {
		t.Fatalf("expected no error. got %v", err)
	}

	if r.Status != pb.Status_EVALUATE_WAIT {
		t.Fatalf("\nexpected: EVALUATE_WAIT \n got: %s", r)
	}

	expectedErrorString := "Dependent object not found: GVR: deployments.apps(v1)/composition-system/missing"
	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %s \n got: %s", expectedErrorString, r)
	}
}

func TestEvaluateGetterConfigurationValidObjectNoFieldRef(t *testing.T) {
	// Marshall GetterConfiguration config
	getter, err := json.Marshal(&compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{
				{
					Name: "missing",
					ResourceRef: compositionv1alpha1.ResourceRef{
						Group:    "apps",
						Version:  "v1",
						Resource: "deployments",
						Kind:     "Deployment",
						Name:     "composition-controller-manager",
					},
					// Empty fieldref
					FieldRef: []compositionv1alpha1.FieldRef{},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to marshall getter: %v", err)
	}

	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config: getter,
			Facade: sampleFacade,
		})
	if err != nil {
		t.Fatalf("expected no error. got %v", err)
	}

	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("\nexpected: SUCCESS \n got: %s", r)
	}

	if string(r.Values) != emptyGetterConfigurationEvaluateResult {
		t.Fatalf("\nexpected: %s\n got: %s", emptyGetterConfigurationEvaluateResult, r.Values)
	}
}

func TestEvaluateGetterConfigurationValidObjectWithMissingField(t *testing.T) {
	// Marshall GetterConfiguration config
	getter, err := json.Marshal(&compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{
				{
					Name: "missing",
					ResourceRef: compositionv1alpha1.ResourceRef{
						Group:    "apps",
						Version:  "v1",
						Resource: "deployments",
						Kind:     "Deployment",
						Name:     "composition-controller-manager",
					},
					// Empty fieldref
					FieldRef: []compositionv1alpha1.FieldRef{
						{
							Path: ".spec.template.foobar",
							As:   "foobar",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to marshall getter: %v", err)
	}

	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config: getter,
			Facade: sampleFacade,
		})
	if err != nil {
		t.Fatalf("expected no error. got %v", err)
	}

	if r.Status != pb.Status_EVALUATE_WAIT {
		t.Fatalf("\nexpected: EVALUATE_WAIT \n got: %s", r)
	}

	expectedErrorString := "Field path not present in object yet: " +
		"Deployment.apps(v1)/composition-system/composition-controller-manager" +
		"[.spec.template.foobar]"
	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %s \n got: %s", expectedErrorString, r)
	}
}

func TestEvaluateGetterConfigurationValidObjectWithValidField(t *testing.T) {
	resultValues := "{\"manager\":{\"replicas\":1}}"
	// Marshall GetterConfiguration config
	getter, err := json.Marshal(&compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{
				{
					Name: "manager",
					ResourceRef: compositionv1alpha1.ResourceRef{
						Group:    "apps",
						Version:  "v1",
						Resource: "deployments",
						Kind:     "Deployment",
						Name:     "composition-controller-manager",
					},
					// Empty fieldref
					FieldRef: []compositionv1alpha1.FieldRef{
						{
							Path: ".spec.replicas",
							As:   "replicas",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to marshall getter: %v", err)
	}

	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config: getter,
			Facade: sampleFacade,
		})
	if err != nil {
		t.Fatalf("expected no error. got %v", err)
	}

	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("\nexpected: SUCCESS \n got: %s", r)
	}

	if string(r.Values) != resultValues {
		t.Fatalf("\nexpected: %s\n got: %s", resultValues, r.Values)
	}
}

func TestEvaluateGetterConfigurationValidObjectWithMultipleFields(t *testing.T) {
	resultValues := "{\"manager\":{\"deadline\":600,\"replicas\":1}}"
	// Marshall GetterConfiguration config
	getter, err := json.Marshal(&compositionv1alpha1.GetterConfiguration{
		Spec: compositionv1alpha1.GetterConfigurationSpec{
			ValuesFrom: []compositionv1alpha1.ValuesFrom{
				{
					Name: "manager",
					ResourceRef: compositionv1alpha1.ResourceRef{
						Group:    "apps",
						Version:  "v1",
						Resource: "deployments",
						Kind:     "Deployment",
						Name:     "composition-controller-manager",
					},
					// Empty fieldref
					FieldRef: []compositionv1alpha1.FieldRef{
						{
							Path: ".spec.replicas",
							As:   "replicas",
						},
						{
							Path: ".spec.progressDeadlineSeconds",
							As:   "deadline",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to marshall getter: %v", err)
	}

	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config: getter,
			Facade: sampleFacade,
		})
	if err != nil {
		t.Fatalf("expected no error. got %v", err)
	}

	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("\nexpected: SUCCESS \n got: %s", r)
	}

	if string(r.Values) != resultValues {
		t.Fatalf("\nexpected: %s\n got: %s", resultValues, r.Values)
	}
}
