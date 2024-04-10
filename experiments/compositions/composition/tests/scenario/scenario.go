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

	"google.com/composition/tests/kind"
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

	cluster kind.KindClusterUser
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

// New - return a Scenario object
func New(t *testing.T, dataFolder string) *Scenario {
	// Sub-tests will include "/" in their names, which are not allowed in
	// metadata.name.
	name := strings.ReplaceAll(t.Name(), "/", "-")
	if dataFolder == "" {
		dataFolder = "../../tests/data/" + t.Name()
	}
	// TODO (barney-s) parameterize or make it global
	logRoot := "../../"
	noTestData := false
	if dataFolder == "none" {
		noTestData = true
	}

	ctx := context.Background()
	cluster := kind.ReserveCluster(t)
	config := cluster.Config()

	testClient := testclient.New(ctx, t, config, cluster.Name(), name, logRoot)

	s := &Scenario{
		T:               t,
		C:               testClient,
		cluster:         cluster,
		name:            name,
		ctx:             ctx,
		config:          config,
		dataFolder:      dataFolder,
		noTestData:      noTestData,
		manifestObjects: make(map[string][]*unstructured.Unstructured),
	}

	if noTestData {
		s.inputObjects = []*unstructured.Unstructured{}
		s.outputObjects = []*unstructured.Unstructured{}
	} else {
		s.inputObjects = s.loadObjects(s.inputData(), "input")
		s.outputObjects = s.loadObjects(s.outputData(), "output")
	}
	return s
}

func (s *Scenario) Cleanup() {
	s.CleanupInput()
	s.CleanupIntermediateManifests()
	s.CleanupOutput()
	s.GatherLogs()
	kind.ReleaseCluster(s.T, s.cluster)
}

func (s *Scenario) Setup() {
	s.ApplyInput()
	s.VerifyInput()
}

func (s *Scenario) testData(filename string) string {
	filePath := filepath.Join(s.dataFolder, filename)
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

func (s *Scenario) ApplyManifests(filename string) {
	if s.noTestData {
		s.T.Errorf("Scenario configured with 'nodata'. But ApplyManifest(%s) called.", filename)
		s.T.FailNow()
		return
	}

	s.T.Logf("Loading manifests from: %s", filename)
	s.manifestObjects[filename] = s.loadObjects(s.testData(filename), filename)

	s.applyObjects(s.manifestObjects[filename], true)
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

func (s *Scenario) VerifyManifests(filename string, matchSpec bool) {
	if s.noTestData {
		s.T.Errorf("Scenario configured with 'nodata'. But VerifyManifest(%s) called.", filename)
		s.T.FailNow()
		return
	}

	s.T.Logf("Loading manifests from: %s", filename)
	s.manifestObjects[filename] = s.loadObjects(s.testData(filename), filename)

	if matchSpec {
		s.T.Logf("Verifying manifests spec matches: %s", filename)
		s.C.MustMatchSpec(s.manifestObjects[filename], ExistTimeout)
	} else {
		s.T.Logf("Verifying manifests exist: %s", filename)
		s.C.MustExist(s.manifestObjects[filename], ExistTimeout)
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
		s.T.Logf("Cleaning up intermediate Manifests: %s", key)
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
