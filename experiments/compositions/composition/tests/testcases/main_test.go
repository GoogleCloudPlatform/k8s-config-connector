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

package e2e

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"google.com/composition/tests/cluster"
	"google.com/composition/tests/cluster/kind"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	CRDManifests       = "../../release/kind/crds.yaml"
	FacadeCRDManifests = "../../release/kind/facade_crds.yaml"
	OperatorManifests  = "../../release/kind/operator.yaml"
)

var (
	images  *string = flag.String("images", "", "images")
	useKind *bool   = flag.Bool("use-kind", true, "use kind cluster")
)

// TestMain - umbrella test that runs all test cases
func TestMain(m *testing.M) {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	logf.SetLogger(zap.New(zap.UseDevMode(true)))

	// sanity check
	err := kind.VerifyKindIsInstalled()
	if err != nil {
		log.Fatalf("Kind not installed. Please install it from: https://kind.sigs.k8s.io/docs/user/quick-start/")
	}

	clusterCount := 1
	if *useKind {
		cluster.CreateKindClusters(clusterCount, *images)
	} else {
		cluster.CreateCCClusters(clusterCount, *images)
	}
	exitCode := m.Run()

	// TODO:
	//cluster.RemoveKindClusters()
	//cluster.RemoveCCClusters()
	os.Exit(exitCode)
}
