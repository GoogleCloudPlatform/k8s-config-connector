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

	compositionv1 "google.com/composition/api/v1"
	"google.com/composition/tests/testclient"

	"time"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(compositionv1.AddToScheme(scheme))
}

// Scenario - context for each scenario
type Scenario struct {
	T *testing.T
	C *testclient.Client

	config *rest.Config
	ctx    context.Context

	// Scenario
	namespace  string
	name       string
	dataFolder string
	objects    []*unstructured.Unstructured
}

// New - return a Scenario object
func New(t *testing.T, config *rest.Config, dataFolder string) *Scenario {
	c, err := client.New(config, client.Options{Scheme: scheme})
	require.NoError(t, err)

	// Sub-tests will include "/" in their names, which are not allowed in
	// metadata.name.
	name := strings.ReplaceAll(t.Name(), "/", "-")
	if dataFolder == "" {
		dataFolder = "../../tests/data/" + t.Name()
	}

	ctx := context.Background()

	return &Scenario{
		T: t,
		C: &testclient.Client{
			T:      t,
			Ctx:    ctx,
			Name:   name,
			Client: c,
		},

		name:       name,
		ctx:        ctx,
		config:     config,
		dataFolder: dataFolder,
		objects:    []*unstructured.Unstructured{},
	}
}

func (s *Scenario) testInputData() string {
	filePath := filepath.Join(s.dataFolder, "input.yaml")
	data, err := ioutil.ReadFile(filePath)
	require.NoErrorf(s.T, err, "Failed reading input: %s", filePath)
	return string(data)
}

func (s *Scenario) ApplyTestData() {
	s.T.Log("Loading test data")
	objects, err := manifest.ParseObjects(s.ctx, s.testInputData())
	require.NoErrorf(s.T, err, "Failed loading test data")

	s.T.Log("Applying test data")
	// loop over objects and extract unstructured
	for _, item := range objects.Items {
		s.objects = append(s.objects, item.UnstructuredObject())
		s.T.Logf("Applying %s", item.UnstructuredObject().GetName())
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
		s.T.Logf("Deleting %s", item.GetName())
		s.C.MustDelete(item)
	}
}

// GetName - return name
func (s *Scenario) GetName() string {
	return s.name
}
