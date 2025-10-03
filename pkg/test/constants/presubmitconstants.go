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

// testsconstants contains constants used in tests.
package testconstants

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"
)

var (
	// RepresentativeCRUDTestsForAllServices specify which test case should be used for each service
	// in the Presubmit-lite job. The format of this service-to-testname mapping is:
	// "<service>" : {"<kind/testname>", "optional testname2"}
	//
	// When adding a new service/updating this list, prioritize test cases with more
	// dependencies for more resource coverage in our presubmit-lite test.
	RepresentativeCRUDTestsForAllServices = map[string][]string{
		"accesscontextmanager":    {"accesscontextmanagerserviceperimeter"},
		"alloydb":                 {"fullalloydbcluster"},
		"apigee":                  {"apigeeenvironment"},
		"apikeys":                 {"apikeyskeybasic"},
		"artifactregistry":        {"artifactregistryrepository"},
		"bigquery":                {"bigqueryjob"},
		"bigtable":                {"bigtablegcpolicy"},
		"billingbudgets":          {"custombudget"},
		"binaryauthorization":     {"namespacepolicytoclusterpolicy"},
		"certificatemanager":      {"certificatemanagercertificatemapentry"},
		"cloudbuild":              {"cloudbuildtrigger"},
		"cloudfunctions":          {"httpsfunction"},
		"cloudidentity":           {"addexpirydatecloudidentitymembership"},
		"cloudids":                {"cloudidsendpoint"},
		"cloudscheduler":          {"cloudschedulerjob"},
		"compute":                 {"globalcomputeforwardingrule"},
		"configcontroller":        {"configcontrollerinstance"},
		"container":               {"containernodepool"},
		"containeranalysis":       {"containeranalysisnote"},
		"dataflow":                {"streamingdataflowjobupdatetemplate"},
		"dataform":                {"dataformrepository"},
		"datafusion":              {"datafusioninstance"},
		"dataproc":                {"dataproccluster"},
		"dlp":                     {"cloudstoragepathstoredinfotype"},
		"dns":                     {"dnsrecordset"},
		"edgecontainer":           {"edgecontainercluster"},
		"eventarc":                {"eventarctrigger"},
		"filestore":               {"filestorebackup"},
		"gkehub":                  {"gkehubfeaturemembership"},
		"firestore":               {"firestoreindex"},
		"iam":                     {"oidcworkloadidentitypoolprovider"},
		"iap":                     {"iapidentityawareproxyclient"},
		"identityplatform":        {"identityplatformoauthidpconfig"},
		"kms":                     {"kmscryptokey"},
		"logging":                 {"logginglogview"},
		"memcache":                {"memcacheinstance"},
		"monitoring":              {"monitoringalertpolicy"},
		"networkconnectivity":     {"networkconnectivityhub"},
		"networksecurity":         {"networksecurityclienttlspolicy"},
		"networkservices":         {"networkservicesgrpcroute"},
		"osconfig":                {"osconfigguestpolicy"},
		"privateca":               {"basiccertificate"},
		"privilegedaccessmanager": {"privilegedaccessmanagerentitlement"},
		"pubsub":                  {"basicpubsubsubscription"},
		"pubsublite":              {"pubsublitereservation"},
		"recaptchaenterprise":     {"androidrecaptchaenterprisekey"},
		"redis":                   {"redisinstance"},
		"resourcemanager":         {"resourcemanagerlien"},
		"run":                     {"runservicebasic"},
		"secretmanager":           {"secretmanagersecretversion"},
		"securesourcemanager":     {"securesourcemanagerinstancebasic"},
		"servicedirectory":        {"servicedirectorynamespace"},
		"servicenetworking":       {"servicenetworkingconnection"},
		"serviceusage":            {"service"},
		"sourcerepo":              {"sourcereporepository"},
		"spanner":                 {"spannerdatabase"},
		"sql":                     {"sqluser"},
		"storage":                 {"storagenotification"},
		"storagetransfer":         {"storagetransferjob"},
		"workstations":            {"workstationcluster-minimal"},
		"vpcaccess":               {"subnetconnector"},
		"vertexai":                {"vertexaidatasetbasic"},
	}
	longRunningCRUDTests = []string{
		"cidrconnector",
		"configcontrollerinstance",
		"containercluster",
		"containernodepool",
		"datafusioninstance",
		"filestorebackup",
		"filestoreinstance",
		"gkehubmembership",
		"gkehubfeaturemembership",
		"memcacheinstance",
		"redisinstance",
		"removedefaultnodepool",
	}
	periodicCRUDTests = []string{
		"cloudidentitygroup",
		"addexpirydatecloudidentitymembership",
		"addrolecloudidentitymembership",
		"removerolecloudidentitymembership",
		"computeinterconnectattachment",
		"computefirewallpolicy",
		"computefirewallpolicyassociation",
		"computefirewallpolicyrule",
	}
	// Services with special testing requirements that should be skipped in presubmit
	skipCRUDTests = map[string]bool{
		"containerattached":    true,
		"edgecontainercluster": true,
		"edgenetwork":          true,
	}
	DynamicTestPackagePath = "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic/..."
)

func GetPresubmitLiteRegexStringArray() []string {
	excludedTests := append(longRunningCRUDTests, periodicCRUDTests...)
	var testList []string
	for _, testCases := range RepresentativeCRUDTestsForAllServices {
		for _, tc := range testCases {
			if slice.StringSliceContains(excludedTests, tc) {
				continue
			}
			testList = append(testList, tc)
		}
	}
	return testList
}

// JoinTestNamesWithRegexFormat ensures that only tests that match the "-foobar"
// test name will run. For example, ServiceUsage's test name 'service' is a common
// substring that can be used by many tests, but specifying the '-' prefix and
// using '$' to indicate where the string should end will ensure that only the
// specified test will run.
// The expected output will look something like this: "-pubsubtopic$|-service$|-sqluser$"
func JoinTestNamesWithRegexFormat(testNames []string) string {
	return "-" + strings.Join(testNames, "$|-") + "$"
}
