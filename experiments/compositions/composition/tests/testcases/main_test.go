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
	"testing"
	"time"

	"google.com/composition/tests/kind"
	"google.com/composition/tests/utils"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	images *string = flag.String("images", "", "images")
)

func RegisterImages(kc kind.KindCluster) error {
	for _, image := range strings.Split(*images, ",") {
		err := kc.LoadImage(image)
		if err != nil {
			return err
		}
	}
	return nil
}

// TestMain - umbrella test that runs all test cases
func TestMain(m *testing.M) {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	logf.SetLogger(zap.New(zap.UseDevMode(true)))

	// Start with 1 e2e cluster
	for i := 0; i < 1; i++ {
		name := fmt.Sprintf("composition-e2e-%d", i)
		kc, err := kind.NewKindCluster(name)
		if err != nil {
			log.Fatalf("Error creating kind cluster: %s, %v", name, err)
		}

		err = kc.Create()
		if err != nil {
			log.Fatalf("Error creating kind cluster: %s, %v", name, err)
		}

		err = RegisterImages(kc)
		if err != nil {
			log.Fatalf("Error registering images with kind cluster: %s, %v", name, err)
		}

		err = utils.InstallOperator(kc)
		if err != nil {
			log.Fatalf("Error installing operator manifests in kind cluster: %s, %v", name, err)
		}

		err = utils.WaitForOperator(kc)
		if err != nil {
			log.Fatalf("Error waiting for operpator to become ready")
		}

		// Start Local controller should be working but we seem to run into issues with CRDs not registering properly with schemes.
		//StartLocalController(kc.Config(), "")

		defer kc.Delete()
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}
