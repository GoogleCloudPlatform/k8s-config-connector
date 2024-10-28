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
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sigs.k8s.io/yaml"
)

var (
	addr = flag.String("addr", "[::]:8443", "the address to connect to")

	// Empty Chart with just chart.yaml
	emptyCEL = `apiVersion: composition.google.com/v1alpha1
kind: CELConfiguration
metadata:
  name: foobar
  namespace: config-control
spec:
`
	emptyCELRenderedManifests = ""

	// Simple chart with a static template file (no templating)
	staticSimpleChart = `apiVersion: composition.google.com/v1alpha1
kind: CELConfiguration
metadata:
  name: foobar
  namespace: config-control
spec:
  resources:
  - name: configmap
    definition:
      apiVersion: v1
      kind: ConfigMap
      metadata:
        name: demo
      data:
        foo: "3"
        bar: "interface"
`
	staticSimpleChartRenderedManifests = `
---
apiVersion: v1
data:
  bar: interface
  foo: "3"
kind: ConfigMap
metadata:
  name: demo
`

	simpleCELConfig = `apiVersion: composition.google.com/v1alpha1
kind: CELConfiguration
metadata:
  name: foobar
  namespace: config-control
spec:
  resources:
  - name: configmap
    definition:
      apiVersion: v1
      kind: ConfigMap
      metadata:
        name: demo
      data:
        foo: ${sqls.spec.foo}
        car: ${sqls.spec.car}
`
	simpleCELConfigRenderedManifests = `
---
apiVersion: v1
data:
  car: sedan
  foo: bar
kind: ConfigMap
metadata:
  name: demo
`

	simpleConfigUsingFetched = `apiVersion: composition.google.com/v1alpha1
kind: CELConfiguration
metadata:
  name: foobar
  namespace: config-control
spec:
  resources:
  - name: configmap
    definition:
      apiVersion: v1
      kind: ConfigMap
      metadata:
        name: demo
      data:
        foo: ${sqls.spec.foo}
        car: ${fetched.car}
`
)

var expanderClient pb.ExpanderClient

func dummyValues(t *testing.T) []byte {
	y := `foo: bar
car: sedan
`
	j, err := yaml.YAMLToJSON([]byte(y))
	if err != nil {
		t.Fatalf("error marshalling to json: %v\n %s", err, y)
	}
	return j
}

func configFrom(t *testing.T, config string) []byte {
	j, err := yaml.YAMLToJSON([]byte(config))
	if err != nil {
		t.Fatalf("error marshalling to json: %v\n %s", err, config)
	}
	return j
}

func testFacade(t *testing.T, facade string) []byte {
	if facade == "" {
		facade = `apiVersion: facade.foobar.com/v1alpha1
kind: Foo
metadata:
  name: appteam-sample
  namespace: default
spec:
  foo: bar
  car: sedan
`
	}
	j, err := yaml.YAMLToJSON([]byte(facade))
	if err != nil {
		t.Fatalf("error marshalling to json: %v\n %s", err, facade)
	}
	return j
}

func testContext(t *testing.T) []byte {
	y := `apiVersion: composition.google.com/v1alpha1
kind: Context
metadata:
  name: context
  namespace: config-control
spec:
  project: test-project
`
	j, err := yaml.YAMLToJSON([]byte(y))
	if err != nil {
		t.Fatalf("error marshalling to json: %v\n %s", err, y)
	}
	return j
}

// TestMain - umbrella test that runs all test cases
func TestMain(m *testing.M) {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	expanderClient = pb.NewExpanderClient(conn)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestEvaluateEmptyConfig(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte{},
			Resource: "sqls",
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	if err != nil {
		t.Fatalf("could not evaluate: %v", err)
	}
	if r.GetStatus() != pb.Status_EVALUATE_FAILED {
		t.Fatalf("want FAILURE, got: %s", r.GetStatus())
	}
	t.Logf("status returned: %s, %s", r.GetStatus(), r.GetError())
}

func TestEvaluateBadConfig(t *testing.T) {
	config := `dummy config`
	_, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, config),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	if err == nil {
		t.Fatalf("expected error. got none")
	}

	errMessage := "cannot unmarshal string into Go value of type v1alpha1.CELConfiguration"
	if !strings.Contains(err.Error(), errMessage) {
		t.Fatalf("Did not find expected string in err: %s, got: %s", errMessage, err.Error())
	}
}

func TestEvaluateNoResourcesConfig(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, emptyCEL),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
	if string(r.Manifests) != emptyCELRenderedManifests {
		t.Fatalf("\nexpected: %s\n got: %s", emptyCELRenderedManifests, r.Manifests)
	}
}

func TestEvaluateEmptyContext(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, emptyCEL),
			Context:  []byte{},
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
	t.Logf("status returned: %s, %s", r.GetStatus(), r.GetError())
	if string(r.Manifests) != emptyCELRenderedManifests {
		t.Fatalf("\nexpected: %s\n got: %s", emptyCELRenderedManifests, r.Manifests)
	}
}

func TestEvaluateEmptyFacade(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, emptyCEL),
			Context:  testContext(t),
			Facade:   []byte{},
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err == nil {
		t.Fatalf("expected error. got none")
	}

	errMessage := "Empty Facade for an Evaluate call"
	if !strings.Contains(err.Error(), errMessage) {
		t.Fatalf("Did not find expected string in err: %s, got: %s", errMessage, err.Error())
	}
}

func TestEvaluateEmptyValue(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, emptyCEL),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    []byte{},
		})
	t.Logf("status returned: %s", r)
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
	t.Logf("status returned: %s, %s", r.GetStatus(), r.GetError())
	if string(r.Manifests) != emptyCELRenderedManifests {
		t.Fatalf("\nexpected: %s\n got: %s", emptyCELRenderedManifests, r.Manifests)
	}
}

func TestEvaluateStaticConfig(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, staticSimpleChart),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
	if string(r.Manifests) != staticSimpleChartRenderedManifests {
		t.Fatalf("\nexpected: %s\n got: %s", staticSimpleChartRenderedManifests, r.Manifests)
	}
}

func TestEvaluateUsingFacade(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, simpleCELConfig),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
	if string(r.Manifests) != simpleCELConfigRenderedManifests {
		t.Fatalf("\nexpected: %s\n got: %s", simpleCELConfigRenderedManifests, r.Manifests)
	}
}

func TestEvaluateMissingFacadeField(t *testing.T) {
	facade := `apiVersion: facade.foobar.com/v1alpha1
kind: Foo
metadata:
  name: appteam-sample
  namespace: default
spec:
  # Missing .Values.spec.car
   foo: bar
  #car: sedan
`
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, simpleCELConfig),
			Context:  testContext(t),
			Facade:   testFacade(t, facade),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err == nil {
		t.Fatalf("expected error. Got none")
	}

	errMessage := "error Evaluating expression: ${sqls.spec.car}, no such key: car"
	if !strings.Contains(err.Error(), errMessage) {
		t.Fatalf("Did not find expected string in err: %s, got: %s", errMessage, err.Error())
	}
}

func TestEvaluateTemplateUsesValues(t *testing.T) {
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Resource: "sqls",
			Config:   configFrom(t, simpleConfigUsingFetched),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS , got: %s", r.GetStatus())
	}

	if string(r.Manifests) != simpleCELConfigRenderedManifests {
		t.Fatalf("\nexpected: %s\ngot: %s", simpleCELConfigRenderedManifests, r.Manifests)
	}
}

// --------------------------------------------------------
// ----------------- Validate Tests -----------------------
// --------------------------------------------------------

func TestValidateEmptyConfig(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte{},
			Resource: "sqls",
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	if err != nil {
		t.Fatalf("could not validate: %v", err)
	}
	if r.GetStatus() != pb.Status_VALIDATE_FAILED {
		t.Fatalf("want FAILURE, got: %s", r.GetStatus())
	}
	t.Logf("status returned: %s, %s", r.GetStatus(), r.GetError())
}

func TestValidateEmptyContext(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   configFrom(t, emptyCEL),
			Resource: "sqls",
			Context:  []byte{},
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s, %s", r.GetStatus(), r.GetError())
	if err != nil {
		t.Fatalf("could not validate: %v", err)
	}
	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
}

func TestValidateEmptyFacade(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   configFrom(t, emptyCEL),
			Resource: "sqls",
			Context:  testContext(t),
			Facade:   []byte{},
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s, %s", r.GetStatus(), r.GetError())
	if err != nil {
		t.Fatalf("could not validate: %v", err)
	}
	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
}

func TestValidateEmptyValue(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   configFrom(t, emptyCEL),
			Resource: "sqls",
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    []byte{},
		})
	t.Logf("status returned: %s, %s", r.GetStatus(), r.GetError())
	if err != nil {
		t.Fatalf("could not validate: %v", err)
	}
	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
}

func TestValidateBadConfig(t *testing.T) {
	config := `dummy config`
	_, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Resource: "sqls",
			Config:   configFrom(t, config),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	if err == nil {
		t.Fatalf("expected error. got none")
	}

	errMessage := "cannot unmarshal string into Go value of type v1alpha1.CELConfiguration"
	if !strings.Contains(err.Error(), errMessage) {
		t.Fatalf("Did not find expected string in err: %s, got: %s", errMessage, err.Error())
	}
}

func TestValidateNoTemplateConfig(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Resource: "sqls",
			Config:   configFrom(t, emptyCEL),
			Context:  testContext(t),
			Facade:   testFacade(t, ""),
			Value:    dummyValues(t),
		})
	t.Logf("status returned: %s", r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
}
