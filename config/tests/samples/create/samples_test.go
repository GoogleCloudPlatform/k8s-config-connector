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
	"regexp"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"

	"golang.org/x/sync/semaphore"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	// Register direct controllers
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

func init() {
	// run-tests allows you to limit the tests that are run by specifying
	// regexes to be used to match test names. The "test name" in this case
	// corresponds to the directory name of the sample YAMLs.
	flag.StringVar(&runTestsRegex, "run-tests", "", "run only the tests whose names match the given regex")
	flag.StringVar(&skipTestsRegex, "skip-tests", "", "skip the tests whose names match the given regex, even those that match the run-tests regex")
	// cleanup-resources allows you to disable the cleanup of resources created during testing. This can be useful for debugging test failures.
	// The default value is true.
	//
	// To use this flag, you MUST use an equals sign as follows: go test -tags=integration -cleanup-resources=false
	flag.BoolVar(&cleanupResources, "cleanup-resources", true, "when enabled, "+
		"cloud resources created by tests will be cleaned up at the end of a test")
}

var (
	runTestsRegex    string
	skipTestsRegex   string
	cleanupResources bool
	mgr              manager.Manager
	// this manager is only used to get the rest.Config from the test framework
	unusedManager manager.Manager

	logger = log.Log
)

var testDisabledList = map[string]bool{
	// Long-term denylist -- current test framework does not support these tests
	//
	// The resources below have special test requirements
	"computeinterconnectattachment": true,
	"firestoreindex":                true,
	"edgenetworknetwork":            true,
	"edgenetworksubnet":             true,
	// The test external cluster to be attached requires special setup.
	// It cannot be attached to multiple projects, or be used multiple times if it's already registered.
	"container-attached-cluster-basic":         true,
	"container-attached-cluster-full":          true,
	"container-attached-cluster-ignore-errors": true,
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
	// The following tests require using a Service Account under an allowlisted project (with
	// label ACCESS_BOUNDARY_TRUSTED_TESTER).
	"iamaccessboundarypolicy": true,
	// The samples below requires a custom project with the test billing account
	// configured, which has a quota limit on the number of projects attached to
	// it. We should disable redundant sample tests for the same resource.
	//
	// Disable 4 out of 5 BinaryAuthorizationPolicy samples.
	"default-policy":          true,
	"namespace-policy":        true,
	"service-account-policy":  true,
	"service-identity-policy": true,

	// The below tests should stay denylisted until underlying issues are fixed
	//
	// Access Context Manager access policies are a singleton for an entire organization;
	// we need to keep on persistently around to test the access level and service policy
	// resources, so we shouldn't modify our existing one.
	"accesscontextmanageraccesspolicy": true,
	"accesscontextmanageraccesslevel":  true,
	// Behaviour of Service Perimeter(Resource) is similar to access level and service policy.
	// So disabling testing for Service Perimeter(Resource)
	"accesscontextmanagerserviceperimeter":         true,
	"accesscontextmanagerserviceperimeterresource": true,
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
	"config-management-feature-membership": true,
	"service-mesh-feature-membership":      true,
	"multi-cluster-ingress-feature":        true,
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
	// This sample test is failing because of parallel deletion failure from the API.
	// Disable the test for now while we are figuring out the long term fix with the
	// service team (b/260214463).
	"private-service-connection-region-network-endpoint-group": true,
	// This sample test is failing because configconnector.net GCP org is not allowlisted.
	// Disable the test until we have fixed b/267510222.
	"calendar-budget": true,
	// These sample test runs try to create AlloyDB backup while Instance has not been created,
	// as there is no way to enforce order of resource creation. Will re-enable once the API handles it
	// b/309167136
	"alloydbbackup":                true,
	"restored-from-backup-cluster": true,
	// This sample test need physical rack which is not suitable for e2e testing due to
	// limited budget.
	"edgecontainercluster-local-control-plane":  true,
	"edgecontainercluster-remote-control-plane": true,
	"edgecontainernodepool":                     true,
	"edgecontainervpnconnection":                true,
}

func TestAll(t *testing.T) {
	ctx, ctxCancel := context.WithCancel(signals.SetupSignalHandler())
	t.Cleanup(func() {
		ctxCancel()
	})

	project := testgcp.GetDefaultProject(t)

	setup(ctx)
	// When runTestsRegex is unset, we run all the samples.
	matchedSamples := LoadMatchingSamples(t, regexp.MustCompile(runTestsRegex), project)
	// When skipTestsRegex is unset, we don't skip any sample.
	var skippedSamples []SampleKey
	if skipTestsRegex != "" {
		skippedSamples = ListMatchingSamples(t, regexp.MustCompile(skipTestsRegex))
	}

	var samples []Sample
	skippedMap := make(map[string]bool)
	for _, skipped := range skippedSamples {
		skippedMap[skipped.Name] = true
	}
	for _, sample := range matchedSamples {
		if _, exists := skippedMap[sample.Name]; !exists {
			samples = append(samples, sample)
		}
	}
	if len(samples) == 0 {
		t.Fatalf("No tests to run for -run-tests=%s, -skip-tests=%s", runTestsRegex, skipTestsRegex)
	}

	// Sort the samples in descending order by number of resources. This is an attempt to start the samples that use
	// a network and have many dependencies sooner since they will likely be the longest running.
	sortSamplesInDescendingOrderByNumberOfResources(samples)
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
		s.Resources = updateProjectResourceWithExistingResourceIDs(t, s.Resources)
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			h := NewHarnessWithManager(ctx, t, mgr)
			SetupNamespacesAndApplyDefaults(h, s.Resources, project)

			networkCount := int64(networksInSampleCount(s))
			if networkCount > 0 {
				logger.Info("Acquiring network semaphore for test...", "testName", s.Name)
				if err := sem.Acquire(ctx, networkCount); err != nil {
					t.Fatalf("error acquiring semaphore: %v", err)
				}
				logger.Info("Acquired network semaphore for test", "testName", s.Name)
				defer releaseFunc(s, networkCount)
			}
			RunCreateDeleteTest(h, CreateDeleteTestOptions{Create: s.Resources, CleanupResources: cleanupResources})
		})
	}
}

func setup(ctx context.Context) {
	flag.Parse()
	var err error
	mgr, err = kccmanager.New(ctx, unusedManager.GetConfig(), kccmanager.Config{StateIntoSpecDefaultValue: stateintospec.StateIntoSpecDefaultValueV1Beta1})
	if err != nil {
		logging.Fatal(err, "error creating new manager")
	}
	// Register the deletion defender controller
	if err := registration.AddDeletionDefender(mgr, &controller.Deps{}); err != nil {
		logging.Fatal(err, "error adding registration controller for deletion defender controllers")
	}
	// start the manager, Start(...) is a blocking operation so it needs to be done asynchronously
	go func() {
		if err := mgr.Start(ctx); err != nil {
			logging.Fatal(err, "error starting manager")
		}
	}()
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &unusedManager)
}
