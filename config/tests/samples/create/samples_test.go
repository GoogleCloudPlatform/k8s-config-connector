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

//go:build integration
// +build integration

package create

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/ghodss/yaml"
	"github.com/golang-collections/go-datastructures/queue"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func init() {
	// run-tests allows you to limit the tests that are run by specifying
	// regexes to be used to match test names. The "test name" in this case
	// corresponds to the directory name of the sample YAMLs.
	flag.StringVar(&runTestsRegex, "run-tests", "", "run only the tests whose names match the given regex")
	// cleanup-resources allows you to disable the cleanup of resources created during testing. This can be useful for debugging test failures.
	// The default value is true.
	//
	// To use this flag, you MUST use an equals sign as follows: go test -tags=integration -cleanup-resources=false
	flag.BoolVar(&cleanupResources, "cleanup-resources", true, "when enabled, "+
		"cloud resources created by tests will be cleaned up at the end of a test")
}

var (
	runTestsRegex    string
	cleanupResources bool
	mgr              manager.Manager
	// this manager is only used to get the rest.Config from the test framework
	unusedManager manager.Manager

	logger = klog.Log
)

var testDisabledList = map[string]bool{
	// Long-term denylist -- current test framework does not support these tests
	//
	// The resources below have special test requirements
	"computeinterconnectattachment": true,
	"firestoreindex":                true,
	// The below test requires more resources than the default quota allows
	"computenodegroup":      true,
	"sole-tenant-node-pool": true,
	// Unsupported IAMPolicy tests
	"workload-identity-policy": true, // needs a real cluster
	// Without the ignore-warnings directive, trying to delete the app profile without another one configured will fail.
	// However, ignore-warnings set to true for a single cluster profile is probably not a good idea.
	"single-cluster-bigtable-app-profile": true,
	// The below tests require logging into the superadmin account on the KCC test
	// organization, and assigning a domain admin role to the service
	// account used by the samples test runner.
	// This authentication can't be automated in our test set-up script.
	"cloudidentitygroup":              true,
	"membership-with-expiration-date": true,
	"membership-with-manager-role":    true,
	// The following tests require using a Service Account under an allowlisted project (with
	// label WORKFORCE_POOLS_TRUSTED_TESTER).
	"iamworkforcepool":             true,
	"oidc-workforce-pool-provider": true,
	"saml-workforce-pool-provider": true,
	// The samples below requires a custom project with the test billing account
	// configured, which has a quota limit on the number of projects attached to
	// it. We should disable redundant sample tests for the same resource.
	//
	// Disable 4 out of 5 BinaryAuthorizationPolicy samples.
	"default-policy":          true,
	"namespace-policy":        true,
	"service-account-policy":  true,
	"service-identity-policy": true,
	// Disable the sample for GameServicesRealm due to service deprecation.
	// Context can be found at b/240747818.
	"gameservicesrealm": true,

	// The below tests should stay denylisted until underlying issues are fixed
	//
	// Access Context Manager access policies are a singleton for an entire organization;
	// we need to keep on persistently around to test the access level and service policy
	// resources, so we shouldn't modify our existing one.
	"accesscontextmanageraccesspolicy": true,
	"accesscontextmanageraccesslevel":  true,
	// Behaviour of Service Perimeter is similar to access level and service policy.
	// So disabling testing for Service Perimeter
	"accesscontextmanagerserviceperimeter": true,
	// Cloud Build Triggers for GitHub repos require the user to connect their
	// GCP project to their GitHub repo first to work.
	"build-trigger-for-github-repo": true,
	// Samples test appends a 4-character UID to the end of project name, making it too long
	"computesharedvpcserviceproject": true,
	// DNS sample testing will cause samples to contend with previous attempts to create records for their domain name.
	"dns-a-record-set":              true,
	"dns-aaaa-record-set":           true,
	"dns-cname-record-set":          true,
	"dns-mx-record-set":             true,
	"dns-ns-record-set":             true,
	"dns-srv-record-set":            true,
	"dns-txt-record-set":            true,
	"dnssec-dnskey-record-set":      true,
	"dnssec-ds-record-set":          true,
	"dnssec-ipsecvpnkey-record-set": true,
	"dnssec-sshfp-record-set":       true,
	// The GSA that authenticates KCC has to be the owner of a Google Group in order to use the group email address as
	// the valid support email. In the integration test, we use the email address of the GSA to bypass the support email validation by IAP REST API.
	// In samples, we want to show the real use case with group/user email address.
	"iapbrand":                    true,
	"iapidentityawareproxyclient": true,
	// Utilizes IAPBrand as a dependency
	"oauth2clientid-backend-service": true,
	// SQL instance replication has complex dependencies on network connection which are unsupported
	"mysql-sql-instance-with-replication": true,
	// Org-level audit configs are uniquely identified by the org + service
	// combination, which means there will be collisions if we allow multiple
	// test runs executing in parallel to create IAMAuditConfig resources for
	// the same org and service combination.
	"external-organization-level-audit-config": true,
	// The sample below creates an org-level role. Since we currently don't
	// have the appropriate logic to periodically clean up leaked org level
	// roles, block this sample for now to mitigate maximizing the org's
	// IAMCustomRole quota (b/177688780).
	"organization-role": true,
	// The sample below creates an IAMPolicyMember that uses an org-level role.
	// Org-level roles cannot be recreated for 7-37 days after being deleted,
	// so sample test runs that inadvertently re-use the same role ID may fail
	// since you cannot create IAMPolicyMembers that use deleted roles.
	"org-level-iam-custom-role-policy-member": true,
	// Org-level org policy is uniquely identified by the org + constraint combination,
	// which affects overall org behavior and can't be safely tested from multiple
	// processes running simultaneously.
	"organization-policy-for-organization": true,
	// The below tests require manually enabling the Google Identity Platform in the Google Cloud Console.
	"identityplatformoauthidpconfig":       true,
	"identityplatformconfig":               true,
	"identityplatformtenant":               true,
	"identityplatformtenantoauthidpconfig": true,
	// TODO(b/193158731): Remove the sample below after understanding why Network can't be cleaned up.
	"vpc-native-container-cluster": true,
	// Deleting the ContainerCluster can fail if its Project was deleted first.
	// Disable these samples for now until ContainerCluster supports
	// hierarchical references or once we support proper deletion ordering
	// (b/179907721).
	"gkehubfeaturemembership":       true,
	"multi-cluster-ingress-feature": true,
	// There is a quota limit of only 10 FirewallPolicies globally in an org,
	// these samples tests are omitted to minimize the number of FirewallPolicies created.
	"computefirewallpolicy":                     true,
	"computefirewallpolicyrule":                 true,
	"association-with-folder-attachment-target": true,
	// In addition to wanting to limit the number of FirewallPolicies we create for
	// the above reason, we want to disable association-with-organization-attachment-target
	// because each organization can only be associated with one ComputeFirewallPolicyAssociation
	// at a given time, and we can't create and delete new organizations for test
	// purposes. Therefore, in order to avoid conflicts when we have multiple
	// presubmit tests running in parallel, this test needs to be disabled.
	"association-with-organization-attachment-target": true,
	// The below tests require the project to be added to the allowlist for resource creation.
	"android-recaptcha-enterprise-key": true,
	"ios-recaptcha-enterprise-key":     true,
	// This sample uses cloud source repo, as a result it is required:
	// 1. The deployer account needs to have access to the repo.
	// 2. The repo contains to-be-deployed source code.
	// This setup cannot be easily done in sample tests. Also due to issue 1, customers cannot directly
	// apply this sample before updating it with a proper cloud source repo URL.
	"eventtrigger-with-pubsubtopic": true,
	// When testing these samples, they always error out on deletion due to a lack
	// of ordered deletion of dependencies. I.e., privatecacapool is deleted before
	// privatecacertificateauthority (which has privatecacapool as a dependency),
	// so privatecacertificateauthority fails on deletion (b/225964552).
	"basic-certificate":     true,
	"cert-sign-certificate": true,
	"complex-certificate":   true,
	// TODO(b/240327420): The ApigeeEnvironment test seems to be timing out,
	// thereby causing the samples test to fail and blocking the  release.
	// Disable the test for now and re-enable it once the issue has been
	// resolved.
	"apigeeenvironment": true,
	// TODO(b/246606623): GKE API issue causing nodepool test failure. Revert after omg/57044 is mitigated.
	"basic-node-pool": true,
}

type Sample struct {
	Name      string
	Resources []*unstructured.Unstructured
}

func TestAll(t *testing.T) {
	setup()
	tfprovider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	samples := loadSamplesOntoUnstructs(t, regexp.MustCompile(runTestsRegex))
	// Sort the samples in descending order by number of resources. This is an attempt to start the samples that use
	// a network and have many dependencies sooner since they will likely be the longest running.
	sortSamplesInDescendingOrderByNumberOfResources(samples)
	setupNamespacesAndApplyDefaults(t, samples, tfprovider)
	// limit the number of in-progress test that have a network  to 14. 15 networks is the maximum number of networks
	// in a newly created GCP project for our billing account's "reputation". There is one default network leaving room
	// for 14 more networks. The reason to limit the number of tests is to avoid lengthening the test runtime due to
	// extended exponential backoff from failure to create a network due to quota.
	maximumNetworks := 14
	sem := semaphore.NewWeighted(int64(maximumNetworks))
	releaseFunc := func(s Sample, n int64) {
		logger.Info("Releasing network semaphore for test", "testName", s.Name)
		sem.Release(n)
	}
	for _, s := range samples {
		if _, ok := testDisabledList[s.Name]; ok {
			continue
		}
		s := s
		s.Resources = replaceResourceNamesWithUniqueIDs(t, s.Resources)
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			networkCount := int64(networksInSampleCount(s))
			if networkCount > 0 {
				logger.Info("Acquiring network semaphore for test...", "testName", s.Name)
				if err := sem.Acquire(context.TODO(), networkCount); err != nil {
					t.Fatalf("error acquiring semaphore: %v", err)
				}
				logger.Info("Acquired network semaphore for test", "testName", s.Name)
				defer releaseFunc(s, networkCount)
			}
			testCreate(t, s.Resources)
		})
	}
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

func setupNamespacesAndApplyDefaults(t *testing.T, samples []Sample, tfprovider *tfschema.Provider) {
	namespaceNames := getNamespaces(samples)
	setupNamespaces(t, namespaceNames)
}

func setupNamespaces(t *testing.T, namespaces []string) {
	projectID := testgcp.GetDefaultProjectID(t)
	for _, n := range namespaces {
		testcontroller.SetupNamespaceForProject(t, mgr.GetClient(), n, projectID)
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

func testCreate(t *testing.T, unstructs []*unstructured.Unstructured) {
	// Create and reconcile all resources & dependencies
	for _, u := range unstructs {
		if err := mgr.GetClient().Create(context.TODO(), u); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
	}
	waitForReady(t, unstructs)
	// Clean up resources on success or if cleanupResources flag is true
	if cleanupResources {
		cleanup(t, unstructs)
	}
}

func waitForReady(t *testing.T, unstructs []*unstructured.Unstructured) {
	var wg sync.WaitGroup
	for _, u := range unstructs {
		wg.Add(1)
		go waitForReadySingleResource(t, &wg, u)
	}
	wg.Wait()
}

func waitForReadySingleResource(t *testing.T, wg *sync.WaitGroup, u *unstructured.Unstructured) {
	name := k8s.GetNamespacedName(u)
	defer wg.Done()
	err := wait.PollImmediate(15*time.Second, 35*time.Minute, func() (done bool, err error) {
		done = true
		logger.Info("Testing to see if resource is ready", "kind", u.GetKind(), "name", u.GetName())
		err = mgr.GetClient().Get(context.TODO(), name, u)
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
		cond := dynamic.GetConditions(t, u)
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
	if err := mgr.GetClient().Get(context.TODO(), name, u); err != nil {
		t.Errorf("%v, error retrieving final status.conditions: %v", baseMsg, err)
		return
	}
	conditions := dynamic.GetConditions(t, u)
	if len(conditions) == 0 {
		t.Errorf("%v, no conditions on resource", baseMsg)
		return
	}
	c := conditions[0]
	t.Errorf("%v, final status.conditions[0] status '%v' and reason '%v': %v", baseMsg, c.Status, c.Reason, c.Message)
}

func cleanup(t *testing.T, unstructs []*unstructured.Unstructured) {
	for _, u := range unstructs {
		logger.Info("Deleting resource", "kind", u.GetKind(), "name", u.GetName())
		if err := mgr.GetClient().Delete(context.TODO(), u); err != nil {
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

func waitForDeleteToComplete(t *testing.T, wg *sync.WaitGroup, u *unstructured.Unstructured) {
	defer wg.Done()
	// Do a best-faith cleanup of the resources. Gives a 30 minute buffer for cleanup, though
	// resources that can be cleaned up quicker exit earlier.
	err := wait.PollImmediate(15*time.Second, 30*time.Minute, func() (bool, error) {
		if err := mgr.GetClient().Get(context.TODO(), k8s.GetNamespacedName(u), u); !errors.IsNotFound(err) {
			return false, nil
		}
		return true, nil
	})
	// TODO (b/197783299): think of better way to handle resources that take a longer time to cleanup
	if err != nil {
		t.Errorf("error while polling for resource cleanup on %v with name '%v': %v; last seen status: %v", u.GetKind(), u.GetName(), err, u.Object["status"])
	}
}

func loadSamplesOntoUnstructs(t *testing.T, regex *regexp.Regexp) []Sample {
	t.Helper()
	samples := make([]Sample, 0)
	sampleNamesToFiles := mapSampleNamesToFilePaths(t, regex)
	subVars := newSubstitutionVariables(t)
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

func newSubstitutionVariables(t *testing.T) map[string]string {
	subs := make(map[string]string)
	subs["${HOST_PROJECT_ID?}"] = testgcp.GetDefaultProjectID(t)
	subs["${PROJECT_ID?}"] = testgcp.GetDefaultProjectID(t)
	subs["${PROJECT_NUMBER?}"] = testgcp.GetDefaultProjectNumber(t)
	subs["${FOLDER_ID?}"] = testgcp.GetFolderID(t)
	subs["${ORG_ID?}"] = testgcp.GetOrgID(t)
	subs["${BILLING_ACCOUNT_ID?}"] = testgcp.GetBillingAccountID(t)
	subs["${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}"] = testgcp.GetTestBillingAccountIDForBillingResources(t)
	subs["${GSA_EMAIL?}"] = getKCCServiceAccountEmail(t)
	return subs
}

// getKCCServiceAccountEmail attempts to get the email address of the service
// account used by KCC.
func getKCCServiceAccountEmail(t *testing.T) string {
	// If there is a service account configured via "Application Default
	// Credentials", then assume this is the service account used by KCC. This
	// assumption holds true if the test is run by Prow.
	if sa := testgcp.GetDefaultServiceAccount(t); sa != "" {
		return sa
	}
	// Otherwise, assume the project has a standard, cluster-mode KCC service
	// account set up.
	return fmt.Sprintf("cnrm-system@%v.iam.gserviceaccount.com", testgcp.GetDefaultProjectID(t))
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

func setup() {
	flag.Parse()
	var err error
	mgr, err = kccmanager.New(unusedManager.GetConfig(), kccmanager.Config{})
	if err != nil {
		logging.Fatal(err, "error creating new manager")
	}
	// Register the deletion defender controller
	if err := registration.Add(mgr, nil, nil, nil, nil, registration.RegisterDeletionDefenderController); err != nil {
		logging.Fatal(err, "error adding registration controller for deletion defender controllers")
	}
	// start the manager, Start(...) is a blocking operation so it needs to be done asynchronously
	go func() {
		if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
			logging.Fatal(err, "error starting manager")
		}
	}()
}

func TestMain(m *testing.M) {
	testmain.TestMainForIntegrationTests(m, &unusedManager)
}
