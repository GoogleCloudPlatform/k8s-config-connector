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
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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

	cluster kind.KindClusterReader
	config  *rest.Config
	ctx     context.Context

	// Scenario
	namespace  string
	name       string
	dataFolder string
	objects    []*unstructured.Unstructured
}

// New - return a Scenario object
func New(t *testing.T, dataFolder string) *Scenario {
	// Sub-tests will include "/" in their names, which are not allowed in
	// metadata.name.
	name := strings.ReplaceAll(t.Name(), "/", "-")
	if dataFolder == "" {
		dataFolder = "../../tests/data/" + t.Name()
	}

	ctx := context.Background()
	cluster := kind.ReserveCluster(t)
	config := cluster.Config()

	testClient := testclient.New(ctx, t, config, cluster.Name(), name)

	return &Scenario{
		T:          t,
		C:          testClient,
		cluster:    cluster,
		name:       name,
		ctx:        ctx,
		config:     config,
		dataFolder: dataFolder,
		objects:    []*unstructured.Unstructured{},
	}
}

func (s *Scenario) ReleaseResources() {
	kind.ReleaseCluster(s.T, s.cluster)
}

func (s *Scenario) testInputData() string {
	filePath := filepath.Join(s.dataFolder, "input.yaml")
	data, err := ioutil.ReadFile(filePath)
	require.NoErrorf(s.T, err, "Failed reading input: %s", filePath)
	return string(data)
}

func (s *Scenario) ApplyTestData() {
	objects, err := manifest.ParseObjects(s.ctx, s.testInputData())
	require.NoErrorf(s.T, err, "Failed loading test data")

	s.T.Log("Applying test data")
	// loop over objects and extract unstructured
	for _, item := range objects.Items {
		s.objects = append(s.objects, item.UnstructuredObject())
		s.C.MustCreate(item.UnstructuredObject())
	}
}

func (s *Scenario) VerifyTestData() {
	s.T.Log("Verifying test data")
	s.C.MustExist(s.objects, ExistTimeout)
}

func (s *Scenario) CleanupTestData() {
	s.T.Log("Cleaning up test data")
	for _, item := range s.objects {
		s.C.MustDelete(item)
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
