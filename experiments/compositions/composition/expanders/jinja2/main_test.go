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
	"os"
	"strings"
	"testing"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	trivialJsonObject          string = "{\"key\": \"val\"}"
	noAttributeZoneErrorString string = "has no attribute 'zone'"
)

var (
	addr = flag.String("addr", "[::]:8443", "the address to connect to")
)

var expanderClient pb.ExpanderClient

// TestMain - umbrella test that runs all test cases
func TestMain(m *testing.M) {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
			Context:  []byte("{\"project\": \"foo\"}"),
			Facade:   []byte(trivialJsonObject),
			Value:    []byte(trivialJsonObject),
		})
	if err != nil {
		t.Fatalf("could not evaluate: %v", err)
	}
	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
}

func TestEvaluateEmptyContext(t *testing.T) {
	config := trivialJsonObject
	manifests := trivialJsonObject
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			Context:  []byte{},
			Facade:   []byte(trivialJsonObject),
			Value:    []byte(trivialJsonObject),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Manifests) != manifests {
		t.Fatalf("\nexpected: %s\n got: %s", manifests, r.Manifests)
	}
}

func TestEvaluateEmptyFacade(t *testing.T) {
	manifests := trivialJsonObject
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(trivialJsonObject),
			Resource: "sqls",
			Context:  []byte(trivialJsonObject),
			Facade:   []byte{},
			Value:    []byte(trivialJsonObject),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Manifests) != manifests {
		t.Fatalf("\nexpected: %s\n got: %s", manifests, r.Manifests)
	}
}

func TestEvaluateEmptyValue(t *testing.T) {
	config := trivialJsonObject
	manifests := trivialJsonObject
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			Context:  []byte(trivialJsonObject),
			Facade:   []byte(trivialJsonObject),
			Value:    []byte{},
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Manifests) != manifests {
		t.Fatalf("\nexpected: %s\n got: %s", manifests, r.Manifests)
	}
}

func TestEvaluateTemplateUsesFacade(t *testing.T) {
	config := "region: {{ sqls.region }}"
	manifests := "region: us-west1"
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			Facade:   []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Manifests) != manifests {
		t.Fatalf("\nexpected: %s\n got: %s", manifests, r.Manifests)
	}
}

func TestEvaluateTemplateMissingFacadeField(t *testing.T) {
	// expects zone providing region
	config := "region: {{ sqls.zone.foobar }}"
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			// expects zone providing region
			Facade: []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.Status != pb.Status_EVALUATE_FAILED {
		t.Fatalf("\nexpected: EVALUATE_FAILED \n got: %s", r)
	}

	expectedErrorString := noAttributeZoneErrorString
	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %s \n got: %s", expectedErrorString, r)
	}
}

func TestEvaluateTemplateUsesContext(t *testing.T) {
	config := "project: {{ context.project }}, region: {{ sqls.region }}"
	manifests := "project: foobar, region: us-west1"
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Context:  []byte("{\"project\": \"foobar\"}"),
			Resource: "sqls",
			Facade:   []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Manifests) != manifests {
		t.Fatalf("\nexpected: %s\n got: %s", manifests, r.Manifests)
	}
}

func TestEvaluateTemplateMissingContextField(t *testing.T) {
	config := "region: {{ context.zone.foobar }}"
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			// expects zone providing project
			Context: []byte("{\"project\": \"foobar\"}"),
			Facade:  []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.Status != pb.Status_EVALUATE_FAILED {
		t.Fatalf("\nexpected: EVALUATE_FAILED \n got: %s", r)
	}

	expectedErrorString := noAttributeZoneErrorString
	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %s \n got: %s", expectedErrorString, r)
	}
}

func TestEvaluateTemplateUsesValues(t *testing.T) {
	config := "identity: {{ values.email }}, region: {{ sqls.region }}"
	manifests := "identity: foobar@acme.com, region: us-west1"
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Value:    []byte("{\"email\": \"foobar@acme.com\"}"),
			Resource: "sqls",
			Facade:   []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if string(r.Manifests) != manifests {
		t.Fatalf("\nexpected: %s\n got: %s", manifests, r.Manifests)
	}
}

func TestEvaluateTemplateMissingValuesField(t *testing.T) {
	config := "email: {{ values.zone.foobar }}"
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			// expects zone providing email
			Value:  []byte("{\"email\": \"foobar@acme.comm\"}"),
			Facade: []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.Status != pb.Status_EVALUATE_FAILED {
		t.Fatalf("\nexpected: EVALUATE_FAILED \n got: %s", r)
	}

	expectedErrorString := noAttributeZoneErrorString
	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %s \n got: %s", expectedErrorString, r)
	}
}

func TestEvaluateTemplateWrongTopLevelField(t *testing.T) {
	// no top level zone
	config := "email: {{ zone.foobar }}"
	r, err := expanderClient.Evaluate(context.Background(),
		&pb.EvaluateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			Facade:   []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.Status != pb.Status_EVALUATE_FAILED {
		t.Fatalf("\nexpected: EVALUATE_FAILED \n got: %s", r)
	}

	expectedErrorString := "'zone' is undefined"
	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %s \n got: %s", expectedErrorString, r)
	}
}

// --------------------------------------------------------
// ----------------- Validate Tests -----------------------
// --------------------------------------------------------

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

func TestValidateEmptyConfig(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte{},
			Resource: "sqls",
			Context:  []byte("{\"project\": \"foo\"}"),
			Facade:   []byte(trivialJsonObject),
			Value:    []byte(trivialJsonObject),
		})
	if err != nil {
		t.Fatalf("could not validate: %v", err)
	}
	if r.GetStatus() != pb.Status_SUCCESS {
		t.Fatalf("want SUCCESS, got: %s", r.GetStatus())
	}
}

func TestValidateEmptyContext(t *testing.T) {
	config := trivialJsonObject
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			Context:  []byte{},
			Facade:   []byte(trivialJsonObject),
			Value:    []byte(trivialJsonObject),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}
	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("expected success. got: %s", r)
	}
}

func TestValidateEmptyFacade(t *testing.T) {
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte(trivialJsonObject),
			Resource: "sqls",
			Context:  []byte(trivialJsonObject),
			Facade:   []byte{},
			Value:    []byte(trivialJsonObject),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("expected success. got: %s", r)
	}
}

func TestValidateEmptyValue(t *testing.T) {
	config := trivialJsonObject
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			Context:  []byte(trivialJsonObject),
			Facade:   []byte(trivialJsonObject),
			Value:    []byte{},
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("expected success. got: %s", r)
	}
}

func TestValidateTemplateUsesFacade(t *testing.T) {
	config := "region: {{ sqls.region }}"
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			Facade:   []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("expected success. got: %s", r)
	}
}

func TestValidateTemplateMissingFacadeField(t *testing.T) {
	// expects zone providing region
	config := "region: {{ sqls.zone.foobar }}"
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			// expects zone providing region
			Facade: []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	// semantic validation is not done. so it should pass
	if r.Status != pb.Status_SUCCESS {
		t.Fatalf("expected success. got: %s", r)
	}
}

func TestValidateJinja2ErrorTemplate(t *testing.T) {
	// expects zone providing region
	config := "region: {{% sqls.zone.foobar %}}"
	r, err := expanderClient.Validate(context.Background(),
		&pb.ValidateRequest{
			Config:   []byte(config),
			Resource: "sqls",
			// expects zone providing region
			Facade: []byte("{\"region\": \"us-west1\"}"),
		})
	if err != nil {
		t.Fatalf("expected no error. got: %v", err)
	}

	// semantic validation is not done. so it should pass
	if r.Status != pb.Status_VALIDATE_FAILED {
		t.Fatalf("expected VALIDATE_FAILED. got: %s", r)
	}

	expectedErrorString := "template\njinja2.exceptions.TemplateSyntaxError: unexpected '%'"
	if !strings.Contains(r.Error.Message, expectedErrorString) {
		t.Fatalf("expected error contains: %s \n got: %s", expectedErrorString, r)
	}
}
