// Copyright 2022 Google LLC
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

package create

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	opv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/ghodss/yaml" //nolint:depguard
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const DefaultWaitForReadyTimeout = 35 * time.Minute

type Sample struct {
	Name      string
	Resources []*unstructured.Unstructured
}

func sortSamplesInDescendingOrderByNumberOfResources(samples []Sample) {
	sort.Slice(samples, func(i, j int) bool {
		return len(samples[i].Resources) > len(samples[j].Resources)
	})
}

func networksInSampleCount(sample Sample) int {
	count := 0
	for _, r := range sample.Resources {
		if r.GetKind() == "ComputeNetwork" {
			count++
		}
	}
	return count
}

func SetupNamespacesAndApplyDefaults(t *Harness, resources []*unstructured.Unstructured, project testgcp.GCPProject) {
	namespaceNames := getNamespacesIfConfigured(resources)
	setupNamespaces(t, namespaceNames, project)
}

func setupNamespaces(t *Harness, namespaces []string, project testgcp.GCPProject) {
	for _, n := range namespaces {
		testcontroller.SetupNamespaceForProject(t.T, t.GetClient(), n, project.ProjectID)
	}
}

func getNamespacesIfConfigured(resources []*unstructured.Unstructured) []string {
	namespaces := sets.NewString()
	for _, unstruct := range resources {
		if ns := unstruct.GetNamespace(); ns != "" {
			namespaces.Insert(unstruct.GetNamespace())
		}
	}
	return namespaces.List()
}

type CreateDeleteTestOptions struct { //nolint:revive
	// Create is the set of objects to create
	Create []*unstructured.Unstructured

	// Updates is the set of objects to update (after all objects have been created)
	Updates []*unstructured.Unstructured

	// CleanupResources is true if we should delete resources when we are done
	CleanupResources bool

	// SkipWaitForDelete true means that we don't wait to query that a resource has been deleted.
	SkipWaitForDelete bool

	// SkipWaitForReady true is mainly used for Paused resources as we don't emit an event for those yet.
	SkipWaitForReady bool

	// CreateInOrder true means that we create each object and wait for the object to be ready.
	// This requires that objects be sorted in creation order.
	CreateInOrder bool

	// DeleteInOrder true means that we delete each object and wait for deletion to complete.
	// This requires that objects be sorted in deletion order.
	DeleteInOrder bool

	// DoNotUseServerSideApplyForCreate uses a normal create for object creation
	// Note: we should use server-side apply for both create and update.
	// If we mix-and-match, we get surprising behaviours e.g. we can't clear a field
	DoNotUseServerSideApplyForCreate bool
}

func RunCreateDeleteTest(t *Harness, opt CreateDeleteTestOptions) {
	ctx := t.Ctx

	// Note: we should use server-side apply for both create and update.
	// If we mix-and-match, we get surprising behaviours e.g. we can't clear a field

	// Create and reconcile all resources & dependencies
	for _, u := range opt.Create {
		if opt.DoNotUseServerSideApplyForCreate {
			t.Log("using legacy create to create object (should ideally use server-side apply)")
			if err := t.GetClient().Create(ctx, u); err != nil {
				t.Fatalf("error creating resource: %v", err)
			}
		} else {
			if err := t.GetClient().Patch(ctx, u, client.Apply, client.FieldOwner("kcc-tests")); err != nil {
				t.Fatalf("error creating resource: %v", err)
			}
		}
		if opt.CreateInOrder && !opt.SkipWaitForReady {
			waitForReadySingleResource(t, u, DefaultWaitForReadyTimeout)
		}
	}

	if !opt.CreateInOrder && !opt.SkipWaitForReady {
		WaitForReady(t, DefaultWaitForReadyTimeout, opt.Create...)
	}

	if len(opt.Updates) != 0 {
		// treat as a patch
		for _, updateUnstruct := range opt.Updates {
			if err := t.GetClient().Patch(ctx, updateUnstruct, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
				t.Fatalf("error updating resource: %v", err)
			}
			if opt.CreateInOrder && !opt.SkipWaitForReady {
				waitForReadySingleResource(t, updateUnstruct, DefaultWaitForReadyTimeout)
			}
		}

		if !opt.CreateInOrder && !opt.SkipWaitForReady {
			WaitForReady(t, DefaultWaitForReadyTimeout, opt.Updates...)
		}
	}

	// Clean up resources on success if CleanupResources flag is true
	if opt.CleanupResources {
		DeleteResources(t, opt)
	}
}

func WaitForReady(h *Harness, timeout time.Duration, unstructs ...*unstructured.Unstructured) {
	var wg sync.WaitGroup
	for _, u := range unstructs {
		u := u
		wg.Add(1)
		go func() {
			defer wg.Done()
			waitForReadySingleResource(h, u, timeout)
		}()
	}
	wg.Wait()
}

func waitForReadySingleResource(t *Harness, u *unstructured.Unstructured, timeout time.Duration) {
	logger := log.FromContext(t.Ctx)

	switch u.GroupVersionKind().GroupKind() {
	case opv1beta1.ConfigConnectorGroupVersionKind.GroupKind():
		logger.Info("ConfigConnector object does not have status.conditions; assuming ready")
		return
	case opv1beta1.ConfigConnectorContextGroupVersionKind.GroupKind():
		logger.Info("ConfigConnectorContext object does not have status.conditions; assuming ready")
		return
	}

	if u.GetKind() == "StorageAnywhereCache" {
		timeout = 4 * time.Hour
	}

	name := k8s.GetNamespacedName(u)
	err := wait.PollImmediate(1*time.Second, timeout, func() (done bool, err error) {
		done = true
		logger.V(2).Info("Testing to see if resource is ready", "kind", u.GetKind(), "name", u.GetName())
		err = t.GetClient().Get(t.Ctx, name, u)
		if err != nil {
			logger.Info("Error getting resource", "kind", u.GetKind(), "name", u.GetName(), "error", err)
			if t.Ctx.Err() != nil {
				return false, t.Ctx.Err()
			}
			return false, nil
		}
		if u.GetKind() == "Secret" { // If unstruct is a Secret and it is found on the API server, then the Secret is ready
			return true, nil
		}
		if u.Object["status"] == nil {
			logger.Info("resource does not yet have status", "kind", u.GetKind(), "name", u.GetName(), "troubleshooting", "if you are testing a direct resource and keep seeing this error message, check if you've imported the service in pkg/controller/direct/register/register.go")
			return false, nil
		}

		if u.Object["status"].(map[string]interface{})["conditions"] == nil {
			logger.Info("resource does not yet have conditions", "kind", u.GetKind(), "name", u.GetName())
			return false, nil
		}
		objectStatus := dynamic.GetObjectStatus(t.T, u)
		if objectStatus.ObservedGeneration == nil {
			logger.Info("resource does not yet have status.observedGeneration", "kind", u.GetKind(), "name", u.GetName())
			return false, nil
		}
		if *objectStatus.ObservedGeneration < objectStatus.Generation {
			logger.Info("resource status.observedGeneration is behind current generation",
				"kind", u.GetKind(), "name", u.GetName(),
				"status.observedGeneration", *objectStatus.ObservedGeneration, "generation", objectStatus.Generation)
			return false, nil
		}
		for _, c := range objectStatus.Conditions {
			if c.Type == "Ready" && c.Status == "True" {
				logger.Info("resource is ready", "kind", u.GetKind(), "name", u.GetName())
				return true, nil
			}
		}
		// This resource is not completely ready. Let's keep polling.
		logger.Info("resource is not ready", "kind", u.GetKind(), "name", u.GetName(),
			"conditions", objectStatus.Conditions)
		return false, nil
	})
	if err == nil {
		return
	}
	if !wait.Interrupted(err) {
		t.Error(fmt.Errorf("error while polling for ready on %v with name '%v': %w", u.GetKind(), u.GetName(), err))
		return
	}
	baseMsg := fmt.Sprintf("timed out waiting for ready on %v with name '%v'", u.GetKind(), u.GetName())
	if err := t.GetClient().Get(t.Ctx, name, u); err != nil {
		t.Error(fmt.Errorf("%v, error retrieving final status.conditions: %w", baseMsg, err))
		return
	}
	objectStatus := dynamic.GetObjectStatus(t.T, u)
	t.Errorf("%v, final status: %+v", baseMsg, objectStatus)
}

func DeleteResources(t *Harness, opts CreateDeleteTestOptions) {
	logger := log.FromContext(t.Ctx)

	unstructs := opts.Create
	for i := len(unstructs) - 1; i >= 0; i-- {
		u := unstructs[i]
		logger.Info("Deleting resource", "kind", u.GetKind(), "name", u.GetName())
		if err := t.GetClient().Delete(t.Ctx, u); err != nil {
			if apierrors.IsNotFound(err) {
				continue
			}
			t.Errorf("error deleting: %v", err)
		}
		if opts.DeleteInOrder && !opts.SkipWaitForDelete {
			waitForDeleteToComplete(t, u)
		}
	}

	if opts.SkipWaitForDelete {
		logger.Info("Not waiting for resources to be deleted")
		return
	}

	if opts.DeleteInOrder {
		// Already deleted
		return
	}

	var wg sync.WaitGroup
	for _, u := range unstructs {
		u := u
		wg.Add(1)
		go func() {
			defer wg.Done()
			waitForDeleteToComplete(t, u)
		}()
	}
	wg.Wait()
}

func waitForDeleteToComplete(t *Harness, u *unstructured.Unstructured) {
	defer log.FromContext(t.Ctx).Info("Done waiting for resource to delete", "kind", u.GetKind(), "name", u.GetName())
	// Do a best-faith cleanup of the resources. Gives a 30 minute buffer for cleanup, though
	// resources that can be cleaned up quicker exit earlier.
	timeout := 30 * time.Minute
	if u.GetKind() == "StorageBucket" {
		timeout = 100 * time.Minute
	}

	err := wait.PollImmediate(1*time.Second, timeout, func() (bool, error) {
		if err := t.GetClient().Get(t.Ctx, k8s.GetNamespacedName(u), u); !apierrors.IsNotFound(err) {
			if t.Ctx.Err() != nil {
				return false, t.Ctx.Err()
			}
			return false, nil
		}
		return true, nil
	})
	// TODO (b/197783299): think of better way to handle resources that take a longer time to cleanup
	if err != nil {
		t.Errorf("error while polling for resource cleanup on %v with name '%v': %v; last seen status: %v", u.GetKind(), u.GetName(), err, u.Object["status"])
	}
}

// LoadAllSamples loads all the samples.
func LoadAllSamples(t *testing.T, project testgcp.GCPProject) []Sample {
	matchEverything := regexp.MustCompile(".*")
	return LoadMatchingSamples(t, matchEverything, project)
}

// LoadMatchingSamples loads the samples that match the regex
func LoadMatchingSamples(t *testing.T, regex *regexp.Regexp, project testgcp.GCPProject) []Sample {
	sampleKeys := ListMatchingSamples(t, regex)
	var samples []Sample
	for _, sampleKey := range sampleKeys {
		sample := loadSampleOntoUnstructs(t, sampleKey, project)
		samples = append(samples, sample)
	}
	return samples
}

// ListAllSamples gets the keys for all the samples without loading them.
func ListAllSamples(t *testing.T) []SampleKey {
	matchEverything := regexp.MustCompile(".*")
	return ListMatchingSamples(t, matchEverything)
}

// LoadSample loads one sample
func LoadSample(t *testing.T, sampleKey SampleKey, project testgcp.GCPProject) Sample {
	return loadSampleOntoUnstructs(t, sampleKey, project)
}

// SampleKey contains the metadata for a sample.
// This lets us defer variable substitution.
type SampleKey struct {
	Name      string
	SourceDir string
	files     []string
}

func loadSampleOntoUnstructs(t *testing.T, sampleKey SampleKey, project testgcp.GCPProject) Sample {
	t.Helper()

	subVars := newSubstitutionVariables(t, project)
	resources := make([]*unstructured.Unstructured, 0)
	for _, f := range sampleKey.files {
		unstructs := readFileToUnstructs(t, f, subVars)
		resources = append(resources, unstructs...)
	}
	s := Sample{
		Name:      sampleKey.Name,
		Resources: resources,
	}
	return s
}

// ListMatchingSamples gets the keys for all samples matching the regex, without loading them.
func ListMatchingSamples(t *testing.T, regex *regexp.Regexp) []SampleKey {
	t.Helper()
	samples := make(map[string]SampleKey)
	baseDir := repo.GetResourcesSamplesPath()
	if err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(d.Name(), ".yaml") {
			sampleName := filepath.Base(filepath.Dir(path))
			if regex.MatchString(sampleName) {
				sampleKey := samples[filepath.Dir(path)]
				sampleKey.Name = sampleName
				sampleKey.SourceDir = filepath.Dir(path)
				sampleKey.files = append(sampleKey.files, path)
				samples[filepath.Dir(path)] = sampleKey
			}
		}
		return nil
	}); err != nil {
		t.Fatalf("error walking samples directory %q: %v", baseDir, err)
	}

	var list []SampleKey
	for _, sampleKey := range samples {
		list = append(list, sampleKey)
	}
	return list
}

func newSubstitutionVariables(t *testing.T, project testgcp.GCPProject) map[string]string {
	subs := make(map[string]string)
	subs["${HOST_PROJECT_ID?}"] = project.ProjectID
	subs["${PROJECT_ID?}"] = project.ProjectID
	subs["${PROJECT_NUMBER?}"] = strconv.FormatInt(project.ProjectNumber, 10)
	subs["${FOLDER_ID?}"] = testgcp.TestFolderID.Get()
	subs["${ORG_ID?}"] = testgcp.TestOrgID.Get()
	subs["${BILLING_ACCOUNT_ID?}"] = testgcp.TestBillingAccountID.Get()
	subs["${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}"] = testgcp.TestBillingAccountIDForBillingResources.Get()
	subs["${GSA_EMAIL?}"] = getKCCServiceAccountEmail(t, project)
	subs["${DLP_TEST_BUCKET?}"] = testgcp.GetDLPTestBucket(t)
	subs["${ATTACHED_CLUSTER_NAME?}"] = testgcp.TestAttachedClusterName.Get()
	subs["${KCC_ATTACHED_CLUSTER_TEST_PROJECT?}"] = testgcp.TestKCCAttachedClusterProject.Get()
	subs["${ATTACHED_CLUSTER_PLATFORM_VERSION?}"] = testgcp.TestKCCAttachedClusterPlatformVersion.Get()
	subs["${KCC_VERTEX_AI_INDEX_TEST_BUCKET?}"] = testgcp.TestKCCVertexAIIndexBucket.Get()
	subs["${KCC_VERTEX_AI_INDEX_TEST_DATA_URI?}"] = testgcp.TestKCCVertexAIIndexDataURI.Get()
	// It needs to be a real group email, so decide to use a placeholder here to
	// avoid exposing internal information.
	subs["${GROUP_EMAIL?}"] = testgcp.TestGroupEmail.Get()
	return subs
}

// getKCCServiceAccountEmail attempts to get the email address of the service
// account used by KCC.
func getKCCServiceAccountEmail(t *testing.T, project testgcp.GCPProject) string {
	// If there is a service account configured via "Application Default
	// Credentials", then assume this is the service account used by KCC. This
	// assumption holds true if the test is run by Prow.
	if sa, err := testgcp.FindDefaultServiceAccount(); err != nil {
		t.Fatalf("error from FindDefaultServiceAccount: %v", err)
	} else if sa != "" {
		return sa
	}
	// Otherwise, assume the project has a standard, cluster-mode KCC service
	// account set up.
	return fmt.Sprintf("cnrm-system@%v.iam.gserviceaccount.com", project.ProjectID)
}

func readFileToUnstructs(t *testing.T, fileName string, subVars map[string]string) []*unstructured.Unstructured {
	t.Helper()
	var returnUnstructs []*unstructured.Unstructured

	b := test.MustReadFile(t, fileName)
	s := string(b)
	for k, v := range subVars {
		s = strings.ReplaceAll(s, k, v)
	}
	b = []byte(s)

	yamls := testyaml.SplitYAML(t, b)
	for _, b = range yamls {
		u := test.ToUnstructWithNamespace(t, b, subVars["${PROJECT_ID?}"])
		returnUnstructs = append(returnUnstructs, u)
	}
	return returnUnstructs
}

func replaceResourceNamesWithUniqueIDs(t *testing.T, unstructs []*unstructured.Unstructured) []*unstructured.Unstructured {
	namesToBeReplaced := make([]string, 0)
	for _, u := range unstructs {
		namesToBeReplaced = append(namesToBeReplaced, u.GetName())
	}

	// Replace names in order of descending length to avoid collisions. For
	// example, unstructs might have instances of names "resource-dep" and
	// "resource-dep2". If we do a string replacement of "resource-dep" first,
	// then the string "resource-dep2" will also be affected since it contains
	// "resource-dep".
	namesToBeReplaced = sortByDescendingLen(namesToBeReplaced)

	namesToUniqueIDs := make(map[string]string)
	idReg := regexp.MustCompile("[a-z]")
	for _, n := range namesToBeReplaced {
		namesToUniqueIDs[n] = testvariable.RandomIDGenerator(idReg, uint(len(n)))
	}

	newUnstructs := make([]*unstructured.Unstructured, 0)
	for _, u := range unstructs {
		b, err := yaml.Marshal(u)
		if err != nil {
			t.Fatalf("error marshalling unstruct to bytes: %v", err)
		}
		s := string(b)
		for _, name := range namesToBeReplaced {
			uniqueID := namesToUniqueIDs[name]
			s = strings.ReplaceAll(s, name, uniqueID)
		}
		b = []byte(s)
		newUnstruct := &unstructured.Unstructured{}
		err = yaml.Unmarshal(b, newUnstruct)
		if err != nil {
			t.Fatalf("error unmarshalling bytes to unstruct: %v", err)
		}
		// Folders also need to have unique values for spec.displayName
		if newUnstruct.GetKind() == "Folder" {
			newDisplayName, err := generateNewFolderDisplayName(u, idReg)
			if err != nil {
				t.Fatalf("error generating new spec.displayName value for Folder '%v': %v", u.GetName(), err)
			}
			if err := unstructured.SetNestedField(newUnstruct.Object, newDisplayName, "spec", "displayName"); err != nil {
				t.Fatal(err)
			}
		}
		newUnstructs = append(newUnstructs, newUnstruct)
	}
	return newUnstructs
}

func updateProjectResourceWithExistingResourceIDs(t *testing.T, unstructs []*unstructured.Unstructured) []*unstructured.Unstructured {
	// Hack: set abandon on delete annotation for dependent project as dynamically creation of billable GCP project is not supported
	for _, u := range unstructs {
		kind := u.GetKind()
		if kind == "Project" {
			b, found, err := unstructured.NestedString(u.Object, "spec", "billingAccountRef", "external")
			if err != nil {
				t.Fatalf("error getting billingAccountRef: %v", err)
			}
			// We cannot dynamically create GCP project with billingAccountRef, acquring pre-created project instead.
			if found && b != "" {
				annotations := u.GetAnnotations()
				if annotations == nil {
					annotations = make(map[string]string)
				}
				annotations["cnrm.cloud.google.com/deletion-policy"] = "abandon"
				u.SetAnnotations(annotations)

				var dp string
				if annotations["cnrm.cloud.google.com/auto-create-network"] == "false" {
					// We use a pre-created project without network
					dp = testgcp.TestDependentNoNetworkProjectID.Get()
				} else {
					_, projectInFolder, err := unstructured.NestedString(u.Object, "spec", "folderRef", "external")
					if err != nil {
						t.Fatalf("error getting folderRef: %v", err)
					}
					_, projectInOrg, err := unstructured.NestedString(u.Object, "spec", "organizationRef", "external")
					if err != nil {
						t.Fatalf("error getting organizationRef: %v", err)
					}

					if projectInFolder {
						dp = testgcp.TestDependentFolderProjectID.Get()
					} else if projectInOrg {
						dp = testgcp.TestDependentOrgProjectID.Get()
					}
				}

				if err := unstructured.SetNestedField(u.Object, dp, strings.Split(k8s.ResourceIDFieldPath, ".")...); err != nil {
					t.Fatalf("error setting resourceID for dependent project: %v", err)
				}
			}

		}
	}

	return unstructs
}

// generateNewFolderDisplayName returns a string that can be used as a new
// display name for the given Folder sample. It has the same length as the
// original display name used in the sample, and it contains enough randomly
// generated characters to avoid display name collisions.
func generateNewFolderDisplayName(folderUnstruct *unstructured.Unstructured, idReg *regexp.Regexp) (string, error) {
	newDisplayNamePrefix := "KCC "
	uniqueIDLen := 10
	minDisplayNameLen := len(newDisplayNamePrefix) + uniqueIDLen

	displayName, err := getFolderDisplayName(folderUnstruct)
	if err != nil {
		return "", err
	}

	if len(displayName) < minDisplayNameLen {
		return "", fmt.Errorf("Folder '%v' has a spec.displayName value of "+
			"'%v' which is too short; please use a spec.displayName with at "+
			"least '%v' characters", folderUnstruct.GetName(), displayName, minDisplayNameLen)
	}

	return newDisplayNamePrefix + testvariable.RandomIDGenerator(idReg, uint(len(displayName)-len(newDisplayNamePrefix))), nil
}

func getFolderDisplayName(folderUnstruct *unstructured.Unstructured) (string, error) {
	displayName, ok, err := unstructured.NestedString(folderUnstruct.Object, "spec", "displayName")
	if err != nil {
		return "", fmt.Errorf("error getting spec.displayName of Folder unstruct: %w", err)
	}
	if !ok {
		return "", fmt.Errorf("spec.displayName not found for Folder unstruct")
	}
	return displayName, nil
}

func sortByDescendingLen(strs []string) []string {
	strsCopy := append(make([]string, 0), strs...)
	sort.Slice(strsCopy, func(i, j int) bool {
		return len(strsCopy[i]) > len(strsCopy[j])
	})
	return strsCopy
}
