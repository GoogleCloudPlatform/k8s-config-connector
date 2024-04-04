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
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	compositionv1 "google.com/composition/api/v1"
	"google.com/composition/internal/controller"
	"google.com/composition/tests/kind"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

var (
	scheme                    = runtime.NewScheme()
	CRDManifests              = "../../release/kind/crds.yaml"
	AliceCRDManifests         = "../../release/kind/alice_crds.yaml"
	OperatorManifests         = "../../release/kind/operator.yaml"
	images            *string = flag.String("images", "gcr.io/allotrope-barni/composition:v0.0.1.alpha,gcr.io/allotrope-barni/manifests-inline:v0.0.1.alpha,gcr.io/allotrope-barni/expander-jinja2:v0.0.1.alpha", "images")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(compositionv1.AddToScheme(scheme))
}

func InstallManifestsFromPath(kc kind.KindCluster, path string) error {
	c, err := client.New(kc.Config(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	manifests, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	objects, err := manifest.ParseObjects(context.Background(), string(manifests))
	for _, item := range objects.Items {
		err := c.Create(context.Background(), item.UnstructuredObject())
		if err != nil {
			exists := apierrors.IsAlreadyExists(err)
			if exists {
				continue
			}
			return err
		}
	}
	return nil
}

func StartLocalController(config *rest.Config, imageRegistry string) error {
	mgr, err := ctrl.NewManager(config, ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
	})

	if err != nil {
		return fmt.Errorf("Unable to start manager: %w", err)
	}

	if err = (&controller.CompositionReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		ImageRegistry: imageRegistry,
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create Composition controller: %w", err)
	}
	if err = (&controller.ContextReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create Context controller: %w", err)
	}
	if err = (&controller.PlanReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create Plan controller: %w", err)
	}

	//if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
	//	return fmt.Errorf("Problem running manager: %w", err)
	//}
	go mgr.Start(ctrl.SetupSignalHandler())
	time.Sleep(time.Second * 5)
	return nil
}

func RegisterImages(kc kind.KindCluster) error {
	for _, image := range strings.Split(*images, ",") {
		err := kc.LoadImage(image)
		if err != nil {
			return err
		}
	}
	return nil
}

func InstallController(kc kind.KindCluster) error {
	c, err := client.New(kc.Config(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	manifests, err := ioutil.ReadFile(OperatorManifests)
	if err != nil {
		return err
	}
	objects, err := manifest.ParseObjects(context.Background(), string(manifests))
	for _, item := range objects.Items {
		err := c.Create(context.Background(), item.UnstructuredObject())
		if err != nil {
			exists := apierrors.IsAlreadyExists(err)
			if exists {
				continue
			}
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
		err = InstallManifestsFromPath(kc, CRDManifests)
		if err != nil {
			log.Fatalf("Error installing CRDs in kind cluster: %s, %v", name, err)
		}
		//err = InstallManifestsFromPath(kc, AliceCRDManifests)
		//if err != nil {
		//	log.Fatalf("Error installing Alice CRDs in kind cluster: %s, %v", name, err)
		//
		// Start Local controller should be working but we seem to run into issues with CRDs not registering properly with schemes.
		//StartLocalController(kc.Config(), "")
		RegisterImages(kc)
		err = InstallManifestsFromPath(kc, OperatorManifests)
		if err != nil {
			log.Fatalf("Error installing Operator Manifests in kind cluster: %s, %v", name, err)
		}
		defer kc.Delete()
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}
