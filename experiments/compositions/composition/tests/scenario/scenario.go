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

package scenario

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"google.com/composition/tests/cluster"
	"google.com/composition/tests/testclient"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const (
	// ExistTimeout - time in which object should exist after creation;
	ExistTimeout = 1 * time.Minute
	// InstantTimeout - time used to check properties that should already apply.
	// It isn't 0 to give checks time to run.
	InstantTimeout = 2 * time.Second
	// ReadyTimeout - time after creation that object should be available;
	ReadyTimeout = 4 * time.Minute
	// Composition Reconcile - bound on time
	CompositionReconcile = time.Minute
	// DeleteTimeout time to delete
	DeleteTimeout = 1 * time.Minute
)

// Scenario - context for each scenario
type Scenario struct {
	T *testing.T
	C *testclient.Client

	cluster cluster.ClusterUser
	config  *rest.Config
	ctx     context.Context

	// Scenario
	noTestData      bool
	namespace       string
	name            string
	dataFolder      string
	inputObjects    []*unstructured.Unstructured
	outputObjects   []*unstructured.Unstructured
	manifestObjects map[string][]*unstructured.Unstructured
}

// NewBasic - return a Scenario object for Basic testcases
//   - Use ../../tests/data/Test<name> for data
//   - Expect input.yaml, output.yaml (optional)
//   - Dont need KCC enabled for tests
func NewBasic(t *testing.T) *Scenario {
	// Sub-tests will include "/" in their names, which are not allowed in
	// metadata.name.
	name := strings.ReplaceAll(t.Name(), "/", "-")
	dataFolder := "../../tests/data/" + t.Name()
	// TODO (barney-s) parameterize or make it global
	logRoot := "../../"

	ctx := context.Background()
	clusterUser := cluster.ReserveCluster(t)
	config := clusterUser.Config()

	testClient := testclient.New(ctx, t, config, clusterUser.Name(), name, logRoot)

	s := &Scenario{
		T:               t,
		C:               testClient,
		cluster:         clusterUser,
		name:            name,
		ctx:             ctx,
		config:          config,
		dataFolder:      dataFolder,
		noTestData:      false,
		manifestObjects: make(map[string][]*unstructured.Unstructured),
	}

	s.inputObjects = s.loadObjects(s.inputData(), "input")
	s.outputObjects = s.loadObjects(s.outputData(), "output")
	return s
}

// NewKCCSample - return a Scenario object for testing Samples/
func NewKCCSample(t *testing.T, sampleName string, composition string) *Scenario {
	// Sub-tests will include "/" in their names, which are not allowed in
	// metadata.name.
	name := strings.ReplaceAll(t.Name(), "/", "-")
	dataFolder := "../../../samples/" + sampleName
	logRoot := "../../"

	ctx := context.Background()
	clusterUser := cluster.ReserveCluster(t)
	if !clusterUser.KCCInstalled() {
		cluster.ReleaseCluster(t, clusterUser)
		t.Logf("KCC not installed, skipping test")
		t.SkipNow()
	}
	config := clusterUser.Config()

	testClient := testclient.New(ctx, t, config, clusterUser.Name(), name, logRoot)

	s := &Scenario{
		T:               t,
		C:               testClient,
		cluster:         clusterUser,
		name:            name,
		ctx:             ctx,
		config:          config,
		dataFolder:      dataFolder,
		noTestData:      false,
		manifestObjects: make(map[string][]*unstructured.Unstructured),
	}

	s.inputObjects = s.loadObjects(s.testData("composition", composition), "composition")
	s.outputObjects = []*unstructured.Unstructured{}
	return s
}

func (s *Scenario) Cleanup() {
	s.CleanupInput()
	s.CleanupIntermediateManifests()
	s.CleanupOutput()
	s.GatherLogs()
	cluster.ReleaseCluster(s.T, s.cluster)
}

func (s *Scenario) Setup() {
	s.ApplyInput()
	s.VerifyInput()
}

func (s *Scenario) testData(path ...string) string {
	filePath := filepath.Join(append([]string{s.dataFolder}, path...)...)
	data, err := os.ReadFile(filePath)
	if err != nil {
		s.T.Errorf("Failed reading input: %s", filePath)
		s.T.FailNow()
	}
	return string(data)
}

func (s *Scenario) inputData() string {
	return s.testData("input.yaml")
}

func (s *Scenario) outputData() string {
	return s.testData("output.yaml")
}

func (s *Scenario) loadObjects(manifests string, name string) []*unstructured.Unstructured {
	output := []*unstructured.Unstructured{}
	if manifests == "" {
		return output
	}
	objects, err := manifest.ParseObjects(s.ctx, manifests)
	if err != nil {
		s.T.Errorf("Failed parsing manifests: %s", name)
		s.T.FailNow()
	}

	// loop over objects and extract unstructured
	for _, item := range objects.Items {
		output = append(output, item.UnstructuredObject())
	}
	return output
}

func (s *Scenario) applyObjects(items []*unstructured.Unstructured, updateAllowed bool) {
	// loop over objects and apply CRDs first
	crds := false
	for _, item := range items {
		if item.GroupVersionKind().Kind == "CustomResourceDefinition" {
			s.C.MustCreate(item, updateAllowed)
			crds = true
		}
	}
	if crds {
		// hacky
		time.Sleep(time.Second * 2)
	}

	namespaces := false
	// loop over objects and extract unstructured
	for _, item := range items {
		if item.GroupVersionKind().Kind == "Namespace" {
			s.C.MustCreate(item, updateAllowed)
			namespaces = true
		}
	}
	if namespaces {
		// hacky
		time.Sleep(time.Second)
	}

	// loop over objects and extract unstructured
	for _, item := range items {
		kind := item.GroupVersionKind().Kind
		if kind != "CustomResourceDefinition" && kind != "Namespace" {
			s.C.MustCreate(item, updateAllowed)
		}
	}
}

func (s *Scenario) Apply(name string, objects []*unstructured.Unstructured) {
	s.T.Logf("Applying objects: %s", name)
	s.manifestObjects[name] = objects
	s.applyObjects(s.manifestObjects[name], true)
}

func (s *Scenario) ApplyManifests(name string, path ...string) {
	if s.noTestData {
		s.T.Errorf("Scenario configured with 'nodata'. But ApplyManifest(%s) called.", path)
		s.T.FailNow()
		return
	}

	s.T.Logf("Loading manifests from: %s", name)
	s.manifestObjects[name] = s.loadObjects(s.testData(path...), name)

	s.applyObjects(s.manifestObjects[name], true)
}

func (s *Scenario) ApplyInput() {
	if s.noTestData {
		return
	}

	s.T.Log("Applying input")
	s.applyObjects(s.inputObjects, false)
}

func (s *Scenario) VerifyInput() {
	if s.noTestData {
		return
	}
	s.T.Log("Verifying input")
	s.C.MustExist(s.inputObjects, ExistTimeout)
}

func (s *Scenario) Verify(name string, matchSpec bool, objects []*unstructured.Unstructured) {
	s.manifestObjects[name] = objects

	if matchSpec {
		s.T.Logf("Verifying objects spec matches: %s", name)
		s.C.MustMatchSpec(s.manifestObjects[name], ExistTimeout)
	} else {
		s.T.Logf("Verifying objects exist: %s", name)
		s.C.MustExist(s.manifestObjects[name], ExistTimeout)
	}
}

func (s *Scenario) VerifyManifests(name string, matchSpec bool, path ...string) {
	if s.noTestData {
		s.T.Errorf("Scenario configured with 'nodata'. But VerifyManifest(%s) called.", name)
		s.T.FailNow()
		return
	}

	s.T.Logf("Loading manifests for: %s", name)
	s.manifestObjects[name] = s.loadObjects(s.testData(path...), name)

	if matchSpec {
		s.T.Logf("Verifying manifests spec matches: %s", name)
		s.C.MustMatchSpec(s.manifestObjects[name], ExistTimeout)
	} else {
		s.T.Logf("Verifying manifests exist: %s", name)
		s.C.MustExist(s.manifestObjects[name], ExistTimeout)
	}
}

func (s *Scenario) VerifyOutputExists() {
	if s.noTestData {
		return
	}
	s.T.Log("Verifying output")
	s.C.MustExist(s.outputObjects, ExistTimeout)
}

func (s *Scenario) VerifyOutputSpecMatches() {
	if s.noTestData {
		return
	}
	s.T.Log("Verifying output spec matches")
	s.C.MustMatchSpec(s.outputObjects, ExistTimeout)
}

func (s *Scenario) CleanupInput() {
	if s.noTestData {
		return
	}
	s.T.Log("Cleaning up input")
	for _, item := range s.inputObjects {
		s.C.MustDelete(item)
	}
	s.C.MustNotExist(s.inputObjects, DeleteTimeout)
}

func (s *Scenario) CleanupOutput() {
	if s.noTestData {
		return
	}
	s.T.Log("Cleaning up output")
	for _, item := range s.outputObjects {
		s.C.MustDelete(item)
	}
	s.C.MustNotExist(s.outputObjects, DeleteTimeout)
}

func (s *Scenario) CleanupIntermediateManifests() {
	if s.noTestData {
		return
	}
	for key := range s.manifestObjects {
		s.T.Logf("Cleaning up objects: %s", key)
		for _, item := range s.manifestObjects[key] {
			s.C.MustDelete(item)
		}
		s.C.MustNotExist(s.manifestObjects[key], DeleteTimeout)
	}
}

// GetName - return name
func (s *Scenario) GetName() string {
	return s.name
}

// GatherLogs - grab the logs for the test case
func (s *Scenario) GatherLogs() {
	s.C.ClearOldLogs()
	s.C.WriteNamespacePodLogs("default")
	s.C.WriteNamespacePodLogs("composition-system")
}
