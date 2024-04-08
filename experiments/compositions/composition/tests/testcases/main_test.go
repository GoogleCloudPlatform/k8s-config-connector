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
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"google.com/composition/tests/kind"
	"k8s.io/apimachinery/pkg/types"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	CRDManifests       = "../../release/kind/crds.yaml"
	FacadeCRDManifests = "../../release/kind/alice_crds.yaml"
	OperatorManifests  = "../../release/kind/operator.yaml"
)

var (
	images *string = flag.String("images", "", "images")
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
	var wg sync.WaitGroup
	wg.Add(clusterCount)
	// Start with 1 e2e cluster
	for i := 0; i < clusterCount; i++ {
		go func() {
			name := fmt.Sprintf("composition-e2e-%d", i)
			// kind cluster
			kc := kind.NewKindCluster(name,
				// that adds these images
				strings.Split(*images, ","),
				// and installs these manifests
				[]string{CRDManifests, OperatorManifests},
				// and waits for these deployments to be ready
				[]types.NamespacedName{
					{Namespace: "composition-system", Name: "composition-controller-manager"},
				},
			)

			// Bringup the cluster and install the operator
			err = kc.ClusterUp()
			if err != nil {
				log.Fatalf("Error creating kind cluster: %s, %v", name, err)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	exitCode := m.Run()
	//kc.Delete()
	os.Exit(exitCode)
}
