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

package configcontroller

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	compositionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/utils"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	kstatus "sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const (
	Region = "us-west1"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	//TODO (barney-s) remove compositionv1alpha1 from test package
	utilruntime.Must(compositionv1alpha1.AddToScheme(scheme))
}

type CCCluster interface {
	ClusterUp() error
	//Delete() error
	//Exists() (bool, error)

	// ClusterUser
	Config() *rest.Config
	Name() string
	RestartWorkloads() error
	WaitForWorkloads() error
	KCCInstalled() bool
	Context() map[string]string
}

type ccCluster struct {
	name             string
	masterCIDR       string
	config           *rest.Config
	manifestPaths    []string
	deployments      []types.NamespacedName
	ctx              context.Context
	gcpProject       string
	gcpProjectNumber string
	client.Client
}

func NewCluster(name string, masterCidr string, manifestPaths []string,
	deployments []types.NamespacedName) CCCluster {
	return &ccCluster{
		name:          name,
		masterCIDR:    masterCidr,
		manifestPaths: manifestPaths,
		deployments:   deployments,
		ctx:           context.Background(),
	}
}

func (c *ccCluster) KCCInstalled() bool { return true }

func (c *ccCluster) Config() *rest.Config {
	return c.config
}

func (c *ccCluster) Context() map[string]string {
	return map[string]string{
		"clusterName":      c.name,
		"clusterLocation":  Region,
		"gcpProject":       c.gcpProject,
		"gcpProjectNumber": c.gcpProjectNumber,
	}
}

func (c *ccCluster) Name() string {
	return c.name
}

func (c *ccCluster) RestartWorkloads() error {
	return nil
}

func (c *ccCluster) exec(cmd string, args ...string) string {
	commandString := fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
	op, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed running cmd: %s \noutput: %s\nerr: %v", commandString, string(op), err)
	}
	return strings.TrimSuffix(string(op), "\n")
}

func (c *ccCluster) create() error {
	// reuse existing project
	// gcloud config list --format 'value(core.project)'
	project := c.exec("gcloud", "config", "list", "--format", "value(core.project)")
	c.gcpProject = project

	//gcloud projects describe $PROJECT_ID --format="value(projectNumber)"
	projectNumber := c.exec("gcloud", "projects", "describe",
		project, "--format", "value(projectNumber)")
	c.gcpProjectNumber = projectNumber

	// create or reuse cc if it exists
	getcc, err := exec.Command("gcloud", "anthos", "config", "controller",
		"describe", c.name, "--location", Region).CombinedOutput()
	if err != nil {
		// If exiterror (failed to get CC), ignore
		if _, ok := err.(*exec.ExitError); !ok {
			return err
		}
	}
	if strings.Contains(string(getcc), "not found") {
		// create CC instance
		c.exec("gcloud", "anthos", "config", "controller",
			"create", c.name, "--master-ipv4-cidr-block", c.masterCIDR, "--location", Region)
	}

	// point kubectl to cc
	c.exec("gcloud", "anthos", "config", "controller",
		"get-credentials", c.name, "--location", Region)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// read from default kubectl context
	kubectlContext := "gke_" + project + "_" + Region + "_krmapihost-" + c.name
	c.config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: homeDir + "/.kube/config"},
		&clientcmd.ConfigOverrides{
			ClusterInfo: clientcmdapi.Cluster{
				Server: "",
			},
			CurrentContext: kubectlContext,
		}).ClientConfig()
	if err != nil {
		return err
	}

	// grant cc permissions
	sa := fmt.Sprintf("serviceAccount:service-%s@gcp-sa-yakima.iam.gserviceaccount.com", projectNumber)
	//gcloud projects add-iam-policy-binding "${PROJECT_ID}" \
	// --member "serviceAccount:${SA_EMAIL}" --role "roles/owner" \
	// --project "${PROJECT_ID}"
	c.exec(
		"gcloud", "projects", "add-iam-policy-binding", project,
		"--member", sa,
		"--role", "roles/owner",
		"--project", project,
	)

	// allow custom workloads

	//kubectl patch k8sallowedresources.constraints.gatekeeper.sh block-workloads \
	// --patch '{"spec":{"enforcementAction":"dryrun"}}' --type merge
	// Not needed anymore: https://source.corp.google.com/h/acp/acp/+/51a8ebb6fbb9de6637dd0eda72cf6c16937fd107
	/*
		c.exec("kubectl", "patch",
			"k8sallowedresources.constraints.gatekeeper.sh",
			"block-workloads", "--patch",
			"{\"spec\":{\"enforcementAction\":\"dryrun\"}}", "--type", "merge")
	*/

	//kubectl label validatingwebhookconfigurations.admissionregistration.k8s.io \
	//  gatekeeper-validating-webhook-configuration policycontroller.configmanagement.gke.io/managed-by-operator-
	c.exec("kubectl", "label",
		"validatingwebhookconfigurations.admissionregistration.k8s.io",
		"gatekeeper-validating-webhook-configuration",
		"policycontroller.configmanagement.gke.io/managed-by-operator-")

	//kubectl patch validatingwebhookconfigurations.admissionregistration.k8s.io \
	// gatekeeper-validating-webhook-configuration \
	// --type=json -p '[ {"op":"remove","path":"/webhooks"} ]'
	c.exec("kubectl", "patch",
		"validatingwebhookconfigurations.admissionregistration.k8s.io",
		"gatekeeper-validating-webhook-configuration", "--type", "json",
		"-p", "[ {\"op\":\"remove\",\"path\":\"/webhooks\"} ]")

	// this is reverted:
	//  kubectl patch K8sAllowedResources forbidden-namespaces --type=json --patch \
	//  '[ {"op":"add","path":"/spec/parameters/exemptedNamespaces/-", "value": "composition-system"} ]'
	// so do this instead:
	//  kubectl patch K8sAllowedResources forbidden-namespaces \
	//  --patch  '{"spec":{"enforcementAction":"dryrun"}}' --type merge
	c.exec("kubectl", "patch",
		"K8sAllowedResources",
		"forbidden-namespaces",
		"--patch", "{\"spec\":{\"enforcementAction\":\"dryrun\"}}",
		"--type", "merge",
	)
	/*
		// Enable 3p workloads
		c.exec("gcloud", "container", "clusters", "update", c.name,
			"--enable-autoprovisioning",
			"--min-cpu", "1",
			"--min-memory", "1",
			"--max-cpu", "10",
			"--max-memory", "64",
			"--autoprovisioning-scopes",
			"https://www.googleapis.com/auth/logging.write," +
			 https://www.googleapis.com/auth/monitoring," +
			 https://www.googleapis.com/auth/devstorage.read_only",
		)
	*/

	// allow image pulls
	//export defaultGCESA="$(gcloud iam service-accounts list --format=json | \
	// jq '.[] | select(.displayName == "Compute Engine default service account") | .email' | xargs echo)"
	defaultGCESA := fmt.Sprintf("%s-compute@developer.gserviceaccount.com", projectNumber)
	registryBucket := fmt.Sprintf("gs://artifacts.%s.appspot.com/", project)
	iamPermissions := fmt.Sprintf("serviceAccount:%s:roles/storage.objectViewer", defaultGCESA)
	c.exec("gsutil", "iam", "ch", iamPermissions, registryBucket)

	c.Client, err = client.New(c.Config(), client.Options{Scheme: scheme})
	if err != nil {
		return err
	}
	return nil
}

func (c *ccCluster) installManifests() error {
	for _, path := range c.manifestPaths {
		manifests, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		objects, err := manifest.ParseObjects(context.Background(), string(manifests))
		if err != nil {
			return err
		}
		for _, item := range objects.Items {
			err := c.Client.Create(context.Background(), item.UnstructuredObject())
			if err != nil {
				exists := apierrors.IsAlreadyExists(err)
				if exists {
					continue
				}
				return err
			}
		}
	}
	return nil
}

func (c *ccCluster) ClusterUp() error {
	err := c.create()
	if err != nil {
		return fmt.Errorf("Error Creating Cluster. err: %v", err)
	}

	err = c.installManifests()
	if err != nil {
		return fmt.Errorf("Error Installing Manifests. err: %v", err)
	}

	err = c.WaitForWorkloads()
	if err != nil {
		return fmt.Errorf("Error Waiting for Deloyments. err: %v", err)
	}

	// HACK - Ties this lib to Composition. Shall figure out callback later on.
	// May be parameterize manifests with context
	contextObj := utils.GetContextObj(c.Context())
	err = c.Client.Create(context.Background(), contextObj)
	if err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return fmt.Errorf("Error creating config-control/context. err: %v", err)
		}
	}

	return nil
}

// isReady - is the object ready
func isReady(ctx context.Context, c client.Client, u *unstructured.Unstructured) (bool, error) {
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

func isDeploymentReady(ctx context.Context, c client.Client, nn types.NamespacedName) (bool, error) {
	u := unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"})
	u.SetName(nn.Name)
	u.SetNamespace(nn.Namespace)

	return isReady(ctx, c, &u)
}

func (c *ccCluster) WaitForWorkloads() error {
	start := time.Now()
	for {
		allReady := true
		for _, workload := range c.deployments {
			ready, err := isDeploymentReady(c.ctx, c.Client, workload)
			if err != nil {
				continue
			}
			if !ready {
				allReady = false
				break
			}
		}
		if allReady {
			break
		}
		if time.Since(start).Seconds() > 40 {
			return fmt.Errorf("timed out waiting for operator to be ready")
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
