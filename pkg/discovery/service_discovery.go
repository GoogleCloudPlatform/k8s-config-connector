// Copyright 2026 Google LLC
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

package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
)

const cnrmGroupSuffix = ".cnrm.cloud.google.com"

// groupToServiceOverrides maps KCC API group prefixes to their canonical GCP Service Usage API names
// where the group prefix does not match "<prefix>.googleapis.com" 1:1.
var groupToServiceOverrides = map[string]string{
	"Alloydb":                 "alloydb.googleapis.com",
	"alloydb":                 "alloydb.googleapis.com",
	"apigee":                  "apigee.googleapis.com",
	"artifactregistry":        "artifactregistry.googleapis.com",
	"Bigquery":                "bigquery.googleapis.com",
	"bigquery":                "bigquery.googleapis.com",
	"bigquerybiglake":         "biglake.googleapis.com",
	"bigqueryconnection":      "bigqueryconnection.googleapis.com",
	"bigquerydatatransfer":    "bigquerydatatransfer.googleapis.com",
	"bigqueryreservation":     "bigqueryreservation.googleapis.com",
	"bigtable":                "bigtableadmin.googleapis.com",
	"billingbudgets":          "billingbudgets.googleapis.com",
	"binaryauthorization":     "binaryauthorization.googleapis.com",
	"certificatemanager":      "certificatemanager.googleapis.com",
	"cloudbuild":              "cloudbuild.googleapis.com",
	"clouddeploy":             "clouddeploy.googleapis.com",
	"cloudfunctions":          "cloudfunctions.googleapis.com",
	"cloudfunctions2":         "cloudfunctions.googleapis.com",
	"cloudidentity":           "cloudidentity.googleapis.com",
	"cloudids":                "ids.googleapis.com",
	"cloudquota":              "cloudquotas.googleapis.com",
	"cloudscheduler":          "cloudscheduler.googleapis.com",
	"cloudtasks":              "cloudtasks.googleapis.com",
	"composer":                "composer.googleapis.com",
	"compute":                 "compute.googleapis.com",
	"configcontroller":        "krmapihosting.googleapis.com",
	"container":               "container.googleapis.com",
	"containerattached":       "gkemulticloud.googleapis.com",
	"containeraws":            "gkemulticloud.googleapis.com",
	"containerazure":          "gkemulticloud.googleapis.com",
	"containernodepool":       "container.googleapis.com",
	"datacatalog":             "datacatalog.googleapis.com",
	"dataflow":                "dataflow.googleapis.com",
	"dataform":                "dataform.googleapis.com",
	"datafusion":              "datafusion.googleapis.com",
	"dataplex":                "dataplex.googleapis.com",
	"dataproc":                "dataproc.googleapis.com",
	"datastream":              "datastream.googleapis.com",
	" deploymentmanager":      "deploymentmanager.googleapis.com",
	"dlp":                     "dlp.googleapis.com",
	"dns":                     "dns.googleapis.com",
	"documentai":              "documentai.googleapis.com",
	"edgecontainer":           "edgecontainer.googleapis.com",
	"edgenetwork":             "edgenetwork.googleapis.com",
	"essentialcontacts":       "essentialcontacts.googleapis.com",
	"eventarc":                "eventarc.googleapis.com",
	"filestore":               "file.googleapis.com",
	"firestore":               "firestore.googleapis.com",
	"gkebackup":               "gkebackup.googleapis.com",
	"gkehub":                  "gkehub.googleapis.com",
	"healthcare":              "healthcare.googleapis.com",
	"iam":                     "iam.googleapis.com",
	"iap":                     "iap.googleapis.com",
	"identityplatform":        "identitytoolkit.googleapis.com",
	"kms":                     "cloudkms.googleapis.com",
	"logging":                 "logging.googleapis.com",
	"memcache":                "memcache.googleapis.com",
	"memorystore":             "memorystore.googleapis.com",
	"monitoring":              "monitoring.googleapis.com",
	"netapp":                  "netapp.googleapis.com",
	"networkconnectivity":     "networkconnectivity.googleapis.com",
	"networkmanagement":       "networkmanagement.googleapis.com",
	"networksecurity":         "networksecurity.googleapis.com",
	"networkservices":         "networkservices.googleapis.com",
	"notebooks":               "notebooks.googleapis.com",
	"orgpolicy":               "orgpolicy.googleapis.com",
	"osconfig":                "osconfig.googleapis.com",
	"oslogin":                 "oslogin.googleapis.com",
	"privateca":               "privateca.googleapis.com",
	"privilegedaccessmanager": "privilegedaccessmanager.googleapis.com",
	"pubsub":                  "pubsub.googleapis.com",
	"pubsublite":              "pubsublite.googleapis.com",
	"recaptchaenterprise":     "recaptchaenterprise.googleapis.com",
	"redis":                   "redis.googleapis.com",
	"resourcemanager":         "cloudresourcemanager.googleapis.com",
	"run":                     "run.googleapis.com",
	"securesourcemanager":     "securesourcemanager.googleapis.com",
	"secretmanager":           "secretmanager.googleapis.com",
	"securitycenter":          "securitycenter.googleapis.com",
	"servicedirectory":        "servicedirectory.googleapis.com",
	"servicenetworking":       "servicenetworking.googleapis.com",
	"serviceusage":            "serviceusage.googleapis.com",
	"sourcerepo":              "sourcerepo.googleapis.com",
	"spanner":                 "spanner.googleapis.com",
	"speech":                  "speech.googleapis.com",
	"sql":                     "sqladmin.googleapis.com",
	"storage":                 "storage.googleapis.com",
	"storagetransfer":         "storagetransfer.googleapis.com",
	"tags":                    "cloudresourcemanager.googleapis.com",
	"vertexai":                "aiplatform.googleapis.com",
	"vpcaccess":               "vpcaccess.googleapis.com",
	"workflows":               "workflows.googleapis.com",
	"workstations":            "workstations.googleapis.com",
}

// CRDGroupPrefix extracts the service group prefix from a CRD name or API group
// (for example, "storagebuckets.storage.cnrm.cloud.google.com" -> "storage").
func CRDGroupPrefix(crdOrGroup string) string {
	trimmed := strings.TrimSpace(strings.ToLower(crdOrGroup))
	if !strings.HasSuffix(trimmed, cnrmGroupSuffix) {
		return ""
	}
	withoutSuffix := strings.TrimSuffix(trimmed, cnrmGroupSuffix)
	parts := strings.Split(withoutSuffix, ".")
	return parts[len(parts)-1]
}

// ServiceNameForCRD returns the canonical GCP service name (e.g., "sqladmin.googleapis.com")
// for a given KCC CRD name or API group. Returns "" for core/customize internal CRDs.
func ServiceNameForCRD(crdOrGroup string) string {
	prefix := CRDGroupPrefix(crdOrGroup)
	if prefix == "" || prefix == "core" || prefix == "customize" {
		return ""
	}
	if svc, ok := groupToServiceOverrides[prefix]; ok {
		return svc
	}
	return prefix + ".googleapis.com"
}

// IsCRDServiceEnabled returns true if the CRD is a core/operator CRD or if its backing
// GCP service is present in enabledServices (supporting both ".googleapis.com" and custom universe suffixes).
// If enabledServices is empty, all CRDs are permitted (fail-open).
func IsCRDServiceEnabled(crdName string, enabledServices map[string]bool) bool {
	if len(enabledServices) == 0 {
		return true
	}
	svc := ServiceNameForCRD(crdName)
	if svc == "" {
		// Core and operator CRDs are always enabled.
		return true
	}
	if enabledServices[svc] {
		return true
	}
	prefix := strings.TrimSuffix(svc, ".googleapis.com")
	if enabledServices[prefix] {
		return true
	}
	for candidate, enabled := range enabledServices {
		if !enabled {
			continue
		}
		if strings.HasPrefix(candidate, prefix+".") {
			return true
		}
	}
	return false
}

// FilterCRDNames filters a list of KCC CRD names to those whose backing GCP service is enabled.
func FilterCRDNames(crdNames []string, enabledServices map[string]bool) []string {
	if len(enabledServices) == 0 {
		out := make([]string, len(crdNames))
		copy(out, crdNames)
		return out
	}
	var filtered []string
	for _, crd := range crdNames {
		if IsCRDServiceEnabled(crd, enabledServices) {
			filtered = append(filtered, crd)
		}
	}
	return filtered
}

type serviceUsageListResponse struct {
	Services []struct {
		Name   string `json:"name"`
		Config struct {
			Name string `json:"name"`
		} `json:"config"`
		State string `json:"state"`
	} `json:"services"`
	NextPageToken string `json:"nextPageToken"`
}

// DiscoverEnabledServices queries the Service Usage API in the target universe domain
// to return the set of enabled GCP services for the given project.
func DiscoverEnabledServices(ctx context.Context, client *http.Client, projectID, universeDomain string) (map[string]bool, error) {
	if client == nil {
		client = http.DefaultClient
	}
	host := gcp.ServiceHost("serviceusage", universeDomain)
	enabled := make(map[string]bool)
	pageToken := ""

	for {
		url := fmt.Sprintf("https://%s/v1/projects/%s/services?filter=state:ENABLED&pageSize=200", host, projectID)
		if pageToken != "" {
			url += "&pageToken=" + pageToken
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("building serviceusage request: %w", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("querying serviceusage API at %s: %w", host, err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("serviceusage API returned HTTP %d", resp.StatusCode)
		}
		var payload serviceUsageListResponse
		err = json.NewDecoder(resp.Body).Decode(&payload)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("decoding serviceusage response: %w", err)
		}
		for _, svc := range payload.Services {
			name := svc.Config.Name
			if name == "" {
				parts := strings.Split(svc.Name, "/")
				name = parts[len(parts)-1]
			}
			if name != "" {
				enabled[name] = true
			}
		}
		if payload.NextPageToken == "" {
			break
		}
		pageToken = payload.NextPageToken
	}

	return enabled, nil
}
