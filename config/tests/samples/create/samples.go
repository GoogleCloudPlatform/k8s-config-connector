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
	"context"
	"fmt"
	"io/ioutil"
	"path"
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
	"github.com/golang-collections/go-datastructures/queue"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
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

func RunCreateDeleteTest(t *Harness, unstructs []*unstructured.Unstructured, cleanupResources bool) {
	// Create and reconcile all resources & dependencies
	for _, u := range unstructs {
		if err := t.GetClient().Create(context.TODO(), u); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
	}
	waitForReady(t, unstructs)
	// Clean up resources on success or if cleanupResources flag is true
	if cleanupResources {
		cleanup(t, unstructs)
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
	err := wait.PollImmediate(15*time.Second, 35*time.Minute, func() (done bool, err error) {
		done = true
		logger.Info("Testing to see if resource is ready", "kind", u.GetKind(), "name", u.GetName())
		err = t.GetClient().Get(t.Ctx, name, u)
		if err != nil {
			logger.Info("Error getting resource", "kind", u.GetKind(), "name", u.GetName(), "error", err)
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
		cond := dynamic.GetConditions(t.T, u)
		if len(cond) == 0 {
			return false, nil
		}
		c := cond[0]
		if c.Reason == "UpToDate" || c.Status == "True" {
			logger.Info("resource is ready", "kind", u.GetKind(), "name", u.GetName())
			return true, nil
		}
		// This resource is not completely ready. Let's keep polling.
		logger.Info("resource is not ready", "kind", u.GetKind(), "name", u.GetName(),
			"status", c.Status, "reason", c.Reason, "message", c.Message)
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
	conditions := dynamic.GetConditions(t.T, u)
	if len(conditions) == 0 {
		t.Errorf("%v, no conditions on resource", baseMsg)
		return
	}
	c := conditions[0]
	t.Errorf("%v, final status.conditions[0] status '%v' and reason '%v': %v", baseMsg, c.Status, c.Reason, c.Message)
}

func cleanup(t *Harness, unstructs []*unstructured.Unstructured) {
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
	err := wait.PollImmediate(15*time.Second, 30*time.Minute, func() (bool, error) {
		if err := t.GetClient().Get(t.Ctx, k8s.GetNamespacedName(u), u); !errors.IsNotFound(err) {
			return false, nil
		}
		return true, nil
	})
	// TODO (b/197783299): think of better way to handle resources that take a longer time to cleanup
	if err != nil {
		t.Errorf("error while polling for resource cleanup on %v with name '%v': %v; last seen status: %v", u.GetKind(), u.GetName(), err, u.Object["status"])
	}
}

// LoadSamples loads all the samples
func LoadSamples(t *testing.T, project testgcp.GCPProject) []Sample {
	matchEverything := regexp.MustCompile(".*")
	return loadSamplesOntoUnstructs(t, matchEverything, project)
}

func loadSamplesOntoUnstructs(t *testing.T, regex *regexp.Regexp, project testgcp.GCPProject) []Sample {
	t.Helper()

	samples := make([]Sample, 0)
	sampleNamesToFiles := mapSampleNamesToFilePaths(t, regex)
	subVars := newSubstitutionVariables(t, project)
	for sample, files := range sampleNamesToFiles {
		resources := make([]*unstructured.Unstructured, 0)
		for _, f := range files {
			unstructs := readFileToUnstructs(t, f, subVars)
			resources = append(resources, unstructs...)
		}
		s := Sample{
			Name:      sample,
			Resources: resources,
		}
		samples = append(samples, s)
	}
	return samples
}

func mapSampleNamesToFilePaths(t *testing.T, regex *regexp.Regexp) map[string][]string {
	t.Helper()
	samples := make(map[string][]string)
	q := queue.New(1)
	q.Put(repo.GetResourcesSamplesPath())
	for !q.Empty() {
		items, err := q.Get(1)
		if err != nil {
			t.Fatalf("error retrieving an item from queue: %v", err)
		}
		dir := items[0].(string)
		fileInfos, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatalf("error reading directory '%v': %v", dir, err)
		}
		for _, fi := range fileInfos {
			if fi.IsDir() {
				q.Put(path.Join(dir, fi.Name()))
				continue
			}
			if !strings.HasSuffix(fi.Name(), ".yaml") {
				continue
			}
			sampleName := path.Base(dir)
			if !regex.MatchString(sampleName) {
				continue
			}
			filePath := path.Join(dir, fi.Name())
			samples[sampleName] = append(samples[sampleName], filePath)
		}
	}
	return samples
}

func newSubstitutionVariables(t *testing.T, project testgcp.GCPProject) map[string]string {
	subs := make(map[string]string)
	subs["${HOST_PROJECT_ID?}"] = project.ProjectID
	subs["${PROJECT_ID?}"] = project.ProjectID
	subs["${PROJECT_NUMBER?}"] = strconv.FormatInt(project.ProjectNumber, 10)
	subs["${FOLDER_ID?}"] = testgcp.GetFolderID(t)
	subs["${ORG_ID?}"] = testgcp.GetOrgID(t)
	subs["${BILLING_ACCOUNT_ID?}"] = testgcp.GetBillingAccountID(t)
	subs["${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}"] = testgcp.GetTestBillingAccountIDForBillingResources(t)
	subs["${GSA_EMAIL?}"] = getKCCServiceAccountEmail(t, project)
	subs["${DLP_TEST_BUCKET?}"] = testgcp.GetDLPTestBucket(t)
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

	b := testcontroller.ReadFileToBytes(t, fileName)
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
