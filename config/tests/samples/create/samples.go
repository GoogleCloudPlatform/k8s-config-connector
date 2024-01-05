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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

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
			count += 1
		}
	}
	return count
}

func SetupNamespacesAndApplyDefaults(t *Harness, samples []Sample, project testgcp.GCPProject) {
	namespaceNames := getNamespaces(samples)
	setupNamespaces(t, namespaceNames, project)
}

func setupNamespaces(t *Harness, namespaces []string, project testgcp.GCPProject) {
	for _, n := range namespaces {
		testcontroller.SetupNamespaceForProject(t.T, t.GetClient(), n, project.ProjectID)
	}
}

func getNamespaces(samples []Sample) []string {
	namespaces := make(map[string]bool)
	for _, sample := range samples {
		for _, unstruct := range sample.Resources {
			namespaces[unstruct.GetNamespace()] = true
		}
	}
	results := make([]string, 0, len(namespaces))
	for k := range namespaces {
		results = append(results, k)
	}
	return results
}

type CreateDeleteTestOptions struct {
	// Create is the set of objects to create
	Create []*unstructured.Unstructured

	// Updates is the set of objects to update (after all objects have been created)
	Updates []*unstructured.Unstructured

	// CleanupResources is true if we should delete resources when we are done
	CleanupResources bool
}

func RunCreateDeleteTest(t *Harness, opt CreateDeleteTestOptions) {
	ctx := t.Ctx

	// Create and reconcile all resources & dependencies
	for _, u := range opt.Create {
		if err := t.GetClient().Create(ctx, u); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
	}

	waitForReady(t, opt.Create)

	if len(opt.Updates) != 0 {
		// treat as a patch
		for _, updateUnstruct := range opt.Updates {
			if err := t.GetClient().Patch(ctx, updateUnstruct, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
				t.Fatalf("error updating resource: %v", err)
			}
		}
		waitForReady(t, opt.Updates)
	}

	// Clean up resources on success if CleanupResources flag is true
	if opt.CleanupResources {
		DeleteResources(t, opt.Create)
	}
}

func waitForReady(t *Harness, unstructs []*unstructured.Unstructured) {
	var wg sync.WaitGroup
	for _, u := range unstructs {
		wg.Add(1)
		go waitForReadySingleResource(t, &wg, u)
	}
	wg.Wait()
}

func waitForReadySingleResource(t *Harness, wg *sync.WaitGroup, u *unstructured.Unstructured) {
	logger := log.FromContext(t.Ctx)

	name := k8s.GetNamespacedName(u)
	defer wg.Done()
	err := wait.PollImmediate(1*time.Second, 35*time.Minute, func() (done bool, err error) {
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
		if u.Object["status"] == nil ||
			u.Object["status"].(map[string]interface{})["conditions"] == nil { // status not ready
			logger.Info("resource does not yet have status or conditions", "kind", u.GetKind(), "name", u.GetName())
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
	if err != wait.ErrWaitTimeout {
		t.Errorf("error while polling for ready on %v with name '%v': %v", u.GetKind(), u.GetName(), err)
		return
	}
	baseMsg := fmt.Sprintf("timed out waiting for ready on %v with name '%v'", u.GetKind(), u.GetName())
	if err := t.GetClient().Get(t.Ctx, name, u); err != nil {
		t.Errorf("%v, error retrieving final status.conditions: %v", baseMsg, err)
		return
	}
	objectStatus := dynamic.GetObjectStatus(t.T, u)
	t.Errorf("%v, final status: %+v", baseMsg, objectStatus)
}

func DeleteResources(t *Harness, unstructs []*unstructured.Unstructured) {
	logger := log.FromContext(t.Ctx)

	for _, u := range unstructs {
		logger.Info("Deleting resource", "kind", u.GetKind(), "name", u.GetName())
		if err := t.GetClient().Delete(t.Ctx, u); err != nil {
			t.Errorf("error deleting: %v", err)
		}
	}
	var wg sync.WaitGroup
	for _, u := range unstructs {
		wg.Add(1)
		go waitForDeleteToComplete(t, &wg, u)
	}
	wg.Wait()
}

func waitForDeleteToComplete(t *Harness, wg *sync.WaitGroup, u *unstructured.Unstructured) {
	defer wg.Done()
	// Do a best-faith cleanup of the resources. Gives a 30 minute buffer for cleanup, though
	// resources that can be cleaned up quicker exit earlier.
	err := wait.PollImmediate(1*time.Second, 30*time.Minute, func() (bool, error) {
		if err := t.GetClient().Get(t.Ctx, k8s.GetNamespacedName(u), u); !errors.IsNotFound(err) {
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
	Name  string
	files []string
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
				sampleKey := samples[sampleName]
				sampleKey.Name = sampleName
				sampleKey.files = append(sampleKey.files, path)
				samples[sampleName] = sampleKey
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
	subs["${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}"] = testgcp.GetTestBillingAccountIDForBillingResources(t)
	subs["${GSA_EMAIL?}"] = getKCCServiceAccountEmail(t, project)
	subs["${DLP_TEST_BUCKET?}"] = testgcp.GetDLPTestBucket(t)
	subs["${ATTACHED_CLUSTER_NAME?}"] = testgcp.TestAttachedClusterName.Get()
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
		namesToUniqueIDs[n] = testvariable.RandomIdGenerator(idReg, uint(len(n)))
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
			unstructured.SetNestedField(newUnstruct.Object, newDisplayName, "spec", "displayName")
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
					dp = testgcp.GetDependentNoNetworkProjectID(t)
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
						dp = testgcp.GetDependentFolderProjectID(t)
					} else if projectInOrg {
						dp = testgcp.GetDependentOrgProjectID(t)
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

	return newDisplayNamePrefix + testvariable.RandomIdGenerator(idReg, uint(len(displayName)-len(newDisplayNamePrefix))), nil
}

func getFolderDisplayName(folderUnstruct *unstructured.Unstructured) (string, error) {
	displayName, ok, err := unstructured.NestedString(folderUnstruct.Object, "spec", "displayName")
	if err != nil {
		return "", fmt.Errorf("error getting spec.displayName of Folder unstruct: %v", err)
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
