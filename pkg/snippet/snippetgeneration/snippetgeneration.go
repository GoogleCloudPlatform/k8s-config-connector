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

package snippetgeneration

import (
	"fmt"
	"path"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/mapslice"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/ghodss/yaml" //nolint:depguard
	goyaml "gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// preferredSampleForResource specifies the sample to be used for snippet
// generation for resources that have multiple samples. It is a map of
// 'resource samples directory name' -> 'sample subdirectory name'.
var preferredSampleForResource = map[string]string{
	"alloydbcluster":                     "regular-cluster",
	"alloydbinstance":                    "primary-instance",
	"alloydbuser":                        "database-user",
	"bigqueryjob":                        "query-bigquery-job",
	"bigtableappprofile":                 "multicluster-bigtable-app-profile",
	"bigtableinstance":                   "replicated-instance",
	"bigquerydatatransferconfig":         "bigquerydatatransferconfig-scheduledquery",
	"bigqueryconnectionconnection":       "cloudresource-connection",
	"billingbudgetsbudget":               "calendar-budget",
	"binaryauthorizationpolicy":          "cluster-policy",
	"certificatemanagercertificate":      "self-managed-certificate",
	"cloudbuildtrigger":                  "build-trigger-for-cloud-source-repo",
	"cloudbuildworkerpool":               "workerpool-with-peered-network",
	"cloudfunctionsfunction":             "httpstrigger",
	"cloudidentitymembership":            "membership-with-manager-role",
	"cloudschedulerjob":                  "scheduler-job-pubsub",
	"computehealthcheck":                 "global-health-check",
	"computeaddress":                     "global-compute-address",
	"computebackendbucket":               "basic-backend-bucket",
	"computebackendservice":              "external-load-balancing-backend-service",
	"computedisk":                        "zonal-compute-disk",
	"computefirewall":                    "allow-rule-firewall",
	"computefirewallpolicyassociation":   "association-with-folder-attachment-target",
	"computeforwardingrule":              "global-forwarding-rule-with-target-http-proxy",
	"computeimage":                       "image-from-url-raw",
	"computeinstance":                    "cloud-machine-instance",
	"computeinstancegroupmanager":        "regional-compute-instance-group-manager",
	"computenodetemplate":                "flexible-node-template",
	"computeregionnetworkendpointgroup":  "cloud-function-region-network-endpoint-group",
	"computereservation":                 "specialized-compute-reservation",
	"computeresourcepolicy":              "weekly-resource-policy-schedule",
	"computerouternat":                   "router-nat-for-all-subnets",
	"computesecuritypolicy":              "multirule-security-policy",
	"computesslcertificate":              "global-compute-ssl-certificate",
	"computesslpolicy":                   "modern-tls-1-1-ssl-policy",
	"computetargethttpsproxy":            "target-https-proxy-with-ssl-certificates",
	"computeurlmap":                      "global-compute-url-map",
	"configcontrollerinstance":           "autopilot-config-controller-instance",
	"containerattachedcluster":           "container-attached-cluster-basic",
	"containercluster":                   "vpc-native-container-cluster",
	"containernodepool":                  "basic-node-pool",
	"dataflowjob":                        "streaming-dataflow-job",
	"dataflowflextemplatejob":            "streaming-dataflow-flex-template-job",
	"dlpstoredinfotype":                  "big-query-field-stored-info-type",
	"dlpdeidentifytemplate":              "info-type-deidentify-template",
	"dlpinspecttemplate":                 "custom-inspect-template",
	"dlpjobtrigger":                      "big-query-job-trigger",
	"dnsrecordset":                       "dns-a-record-set",
	"edgecontainercluster":               "edgecontainercluster-remote-control-plane",
	"firestoredatabase":                  "basic-firestoredatabase",
	"folder":                             "folder-in-folder",
	"gkehubfeature":                      "multi-cluster-ingress-feature",
	"gkehubfeaturemembership":            "config-management-feature-membership",
	"iamauditconfig":                     "project-level-audit-config",
	"iamcustomrole":                      "project-role",
	"iampolicy":                          "external-project-level-policy",
	"iampartialpolicy":                   "project-level-iampartialpolicy",
	"iampolicymember":                    "external-project-level-policy-member",
	"iamworkforcepoolprovider":           "oidc-workforce-pool-provider",
	"iamworkloadidentitypoolprovider":    "oidc-workload-identity-pool-provider",
	"logginglogbucket":                   "project-log-bucket",
	"logginglogexclusion":                "project-exclusion",
	"logginglogmetric":                   "linear-log-metric",
	"logginglogsink":                     "project-sink",
	"logginglogview":                     "project-log-view",
	"monitoringalertpolicy":              "network-connectivity-alert-policy",
	"monitoringnotificationchannel":      "sms-monitoring-notification-channel",
	"monitoringservicelevelobjective":    "window-based-gtr-distribution-cut",
	"monitoringuptimecheckconfig":        "http-uptime-check-config",
	"osconfigospolicyassignment":         "fixed-os-policy-assignment",
	"privatecacertificate":               "basic-certificate",
	"privilegedaccessmanagerentitlement": "project-level-entitlement",
	"project":                            "project-in-folder",
	"pubsubsubscription":                 "basic-pubsub-subscription",
	"runjob":                             "basic-job",
	"recaptchaenterprisekey":             "challenge-based-web-recaptcha-enterprise-key",
	"resourcemanagerpolicy":              "organization-policy-for-project",
	"runservice":                         "run-service-secret",
	"secretmanagersecret":                "automatic-secret-replication",
	"securesourcemanagerinstance":        "securesourcemanagerinstance-basic",
	"sqlinstance":                        "mysql-sql-instance",
	"vpcaccessconnector":                 "cidr-connector",
	"vertexaidataset":                    "vertexai-dataset-encryptionkey",
	"vertexaiendpoint":                   "vertexai-endpoint-network",
	"workflowsworkflow":                  "basic-workflow",
	"workstationcluster":                 "basic-workstationcluster",
	"workstationconfig":                  "basic-workstationconfig",
	"workstation":                        "basic-workstation",
	"kmsautokeyconfig":                   "kmsautokeyconfig",
	"kmskeyhandle":                       "kmskeyhandle",
	"managedkafkacluster":                "managedkafkacluster-cmek",
	"iapsettings":                        "regionalbackendserviceiapsettings",
}

type Snippet struct {
	Label               string `yaml:"label"`
	MarkdownDescription string `yaml:"markdownDescription"`
	InsertText          string `yaml:"insertText"`
}

// PathToSampleFileUsedForSnippets gets the path to the sample file used to
// generate snippets for the given resource samples directory. Note: the given
// resource samples directory must be a subdirectory of the overall resources
// samples directory at config/samples/resources.
func PathToSampleFileUsedForSnippets(resourceDirName string) (string, error) {
	samplesPath := repo.GetResourcesSamplesPath()

	resourceDirPath := path.Join(samplesPath, resourceDirName)
	dirExists, err := fileutil.DirExists(resourceDirPath)
	if err != nil {
		return "", fmt.Errorf("error: failed to determine if directory with name %v exists in %v: %w", resourceDirName, samplesPath, err)
	}
	if !dirExists {
		return "", fmt.Errorf("error: no directory with name %v found in %v", resourceDirName, samplesPath)
	}

	hasSubdirs, err := fileutil.HasSubdirs(resourceDirPath)
	if err != nil {
		return "", fmt.Errorf("error determining if directory at %v has subdirectories: %w", resourceDirPath, err)
	}

	sampleDirPath := resourceDirPath
	if hasSubdirs {
		sampleDirPath, err = pathToPreferredSamplesSubdirForResource(resourceDirPath)
		if err != nil {
			return "", err
		}
	}

	fileNames, err := fileutil.FileNamesWithSuffixInDir(sampleDirPath, resourceDirName+".yaml")
	if err != nil {
		return "", fmt.Errorf("error getting files to use for generating snippets: %w", err)
	}
	if len(fileNames) != 1 {
		return "", fmt.Errorf("error getting exactly one file to use for generating snippets (dir=%q, suffix=%q); expected one, got %v", sampleDirPath, resourceDirName+".yaml", fileNames)
	}

	return path.Join(sampleDirPath, fileNames[0]), nil
}

func pathToPreferredSamplesSubdirForResource(resourceDirPath string) (string, error) {
	resourceDirName := path.Base(resourceDirPath)
	sampleSubdirName, ok := preferredSampleForResource[resourceDirName]
	if !ok {
		return "", fmt.Errorf("error: no sample subdirectory specified for resource directory '%v'", resourceDirName)
	}
	sampleSubdirPath := path.Join(resourceDirPath, sampleSubdirName)
	dirExists, err := fileutil.DirExists(sampleSubdirPath)
	if err != nil {
		return "", fmt.Errorf("error: failed to determine if directory at %v exists: %w", sampleSubdirPath, err)
	}
	if !dirExists {
		return "", fmt.Errorf("error: no directory found at %v", sampleSubdirPath)
	}
	return sampleSubdirPath, nil
}

func SnippifyResourceConfig(resourceConfig []byte) (Snippet, error) {
	kind, err := resourceKind(resourceConfig)
	if err != nil {
		return Snippet{}, fmt.Errorf("error parsing resource kind from resource config: %w", err)
	}
	config, err := snippifyResourceConfig(kind, resourceConfig)
	if err != nil {
		return Snippet{}, fmt.Errorf("error snippifying resource config: %w", err)
	}
	return Snippet{
		Label:               "Config Connector " + kind,
		MarkdownDescription: fmt.Sprintf("Creates yaml for a %v resource", kind),
		InsertText:          config,
	}, nil
}

func snippifyResourceConfig(kind string, config []byte) (string, error) {
	var mapSlice goyaml.MapSlice
	err := goyaml.Unmarshal(config, &mapSlice)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling bytes: %w", err)
	}

	newMapSlice := goyaml.MapSlice{}
	varNum := 1
	for _, item := range mapSlice {
		switch key := item.Key.(string); key {
		case "metadata":
			item.Value = snippifyMetadata(item.Value, kind, &varNum)
		case "spec":
			item.Value = snippifyAllLeavesInTree(item.Value, &varNum)
		}
		newMapSlice = append(newMapSlice, item)
	}

	out, err := goyaml.Marshal(newMapSlice)
	if err != nil {
		return "", fmt.Errorf("error marshalling bytes to YAML: %w", err)
	}
	return string(out), nil
}

func snippifyMetadata(metadataFields interface{}, kind string, varNum *int) interface{} {
	m := metadataFields.(goyaml.MapSlice)
	out := goyaml.MapSlice{}

	labels := mapslice.Value(m, "labels")
	name := mapslice.Value(m, "name")

	if labels != nil {
		labels := labels.(goyaml.MapSlice)
		newLabels := make([]goyaml.MapItem, 0)
		for _, l := range labels {
			newLabels = append(newLabels, goyaml.MapItem{
				Key:   snippifyVal(l.Key.(string), varNum),
				Value: snippifyVal(l.Value.(string), varNum),
			})
		}
		out = append(out, goyaml.MapItem{
			Key:   "labels",
			Value: newLabels,
		})
	}
	if name != nil {
		out = append(out, goyaml.MapItem{
			Key:   "name",
			Value: snippifyVal(strings.ToLower(kind)+"-name", varNum),
		})
	}
	return out
}

func snippifyAllLeavesInTree(node interface{}, varNum *int) interface{} {
	switch v := node.(type) {
	case goyaml.MapSlice:
		out := goyaml.MapSlice{}
		for _, item := range v {
			item.Value = snippifyAllLeavesInTree(item.Value, varNum)
			out = append(out, item)
		}
		return out
	case []interface{}:
		out := make([]interface{}, 0)
		for _, val := range v {
			out = append(out, snippifyAllLeavesInTree(val, varNum))
		}
		return out
	default:
		return snippifyVal(fmt.Sprintf("%v", v), varNum)
	}
}

func snippifyVal(s string, varNum *int) string {
	// Remove or replace characters that have special meaning in snippet strings.
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, "{", "[")
	s = strings.ReplaceAll(s, "}", "]")

	v := fmt.Sprintf("\\${%v:%v}", *varNum, s)
	*varNum++
	return v
}

func resourceKind(config []byte) (string, error) {
	u := &unstructured.Unstructured{}
	err := yaml.Unmarshal(config, u)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling bytes to CRD: %w", err)
	}
	return u.GetKind(), nil
}
