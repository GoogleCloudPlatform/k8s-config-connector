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

package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	compositionv1 "google.com/composition/api/v1"
	"google.com/composition/internal/controller"
	"google.com/composition/tests/kind"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	kstatus "sigs.k8s.io/cli-utils/pkg/kstatus/status"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

var (
	scheme             = runtime.NewScheme()
	CRDManifests       = "../../release/kind/crds.yaml"
	FacadeCRDManifests = "../../release/kind/alice_crds.yaml"
	OperatorManifests  = "../../release/kind/operator.yaml"
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(compositionv1.AddToScheme(scheme))
}

func installManifestsFromPath(kc kind.KindCluster, path string) error {
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

func InstallOperator(kc kind.KindCluster) error {
	err := installManifestsFromPath(kc, CRDManifests)
	if err != nil {
		return err
	}
	err = installManifestsFromPath(kc, OperatorManifests)
	if err != nil {
		return err
	}
	return nil
}

// isReady - is the object ready
func isReady(c client.Client, u *unstructured.Unstructured) (bool, error) {
	ctx := context.Background()

	key := types.NamespacedName{
		Name:      u.GetName(),
		Namespace: u.GetNamespace(),
	}
	err := c.Get(ctx, key, u)
	result := &kstatus.Result{}
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return false, err
		}
		return false, nil
	} else {
		result, err = kstatus.Compute(u)
		if err != nil {
			return false, err
		}
	}
	if result.Status != kstatus.CurrentStatus {
		return false, nil
	}
	return true, nil
}

func isDeploymentReady(c client.Client, namespace, name string) (bool, error) {
	u := unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"})
	u.SetName(name)
	u.SetNamespace(namespace)

	return isReady(c, &u)
}

func WaitForOperator(kc kind.KindCluster) error {
	c, err := client.New(kc.Config(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	start := time.Now()
	for true {
		time.Sleep(2)
		ready, err := isDeploymentReady(c, "composition-system", "composition-controller-manager")
		if err != nil {
			continue
		}
		if !ready {
			continue
		}
		if ready {
			break
		}
		if time.Since(start).Seconds() > 40 {
			return fmt.Errorf("timed out waiting for operator to be ready")
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
