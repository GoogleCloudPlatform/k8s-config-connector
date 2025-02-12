// Copyright 2025 Google LLC
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

package mockgcptests

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"k8s.io/klog/v2"
)

type Proxy struct {
	httpClient *http.Client
}

func NewProxy(httpClient *http.Client) *Proxy {
	return &Proxy{httpClient: httpClient}
}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	resp, err := p.runRequest(req)
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		klog.Errorf("error proxying request: %v", err)
	}

	defer resp.Body.Close()

	// Copy headers & write response
	for k, values := range resp.Header {
		if strings.ToLower(k) == "accept-encoding" {
			// Avoid having to handle gzip from real GCP
			continue
		}
		for _, v := range values {
			rw.Header().Add(k, v)
		}
	}
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}

func (p *Proxy) runRequest(req *http.Request) (*http.Response, error) {
	klog.Infof("proxying %v %v", req.Method, req.URL)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("reading request body: %v", err)
	}

	u := req.URL
	u.Scheme = "https"

	proxyReq, err := http.NewRequest(req.Method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("creating upstream request: %v", err)
	}

	proxyReq.Header = make(http.Header)
	for k, values := range req.Header {
		switch strings.ToLower(k) {
		case "accept-encoding":
			// Avoid having to handle gzip from real GCP
			values = nil
		}

		for _, v := range values {
			proxyReq.Header.Add(k, v)
		}
	}

	resp, err := p.httpClient.Do(proxyReq)
	if err != nil {
		return nil, fmt.Errorf("doing upstream request: %v", err)
	}
	return resp, nil
}

type GcloudConfig struct {
	EnvVars []string
}

func (c *GcloudConfig) AddConfig(key string, value string) {
	envVarKey := "CLOUDSDK_" + strings.ToUpper(key)
	envVarKey = strings.ReplaceAll(envVarKey, "/", "_")
	c.EnvVars = append(c.EnvVars, fmt.Sprintf("%v=%v", envVarKey, value))
}

func (p *Proxy) BuildGcloudConfig(proxyEndpoint *net.TCPAddr, mockgcp mockgcp.Interface) GcloudConfig {
	var config GcloudConfig

	config.AddConfig("proxy/type", "http")
	config.AddConfig("proxy/address", proxyEndpoint.IP.String())
	config.AddConfig("proxy/port", strconv.Itoa(proxyEndpoint.Port))

	// We need to register services to use http, to stop gcloud trying to use TUNNEL with our proxy

	// We need a hard-coded list, because we don't always mockgcp available

	// This list should be kept in sync with the output from `gcloud config  list api_endpoint_overrides/ --all --format json | jq -r '.api_endpoint_overrides | keys[]'`
	apiEndpointOverrides := []string{
		"accessapproval",
		"accesscontextmanager",
		"ai",
		"aiplatform",
		"anthosevents",
		"anthospolicycontrollerstatus_pa",
		"apigateway",
		"apigee",
		"appengine",
		"apphub",
		"artifactregistry",
		"assuredworkloads",
		"auditmanager",
		"baremetalsolution",
		"bigtableadmin",
		"certificatemanager",
		"cloudasset",
		"cloudbilling",
		"cloudbuild",
		"cloudcommerceconsumerprocurement",
		"clouddebugger",
		"clouddeploy",
		"clouderrorreporting",
		"cloudfunctions",
		"cloudidentity",
		"cloudkms",
		"cloudresourcemanager",
		"cloudscheduler",
		"cloudtasks",
		"cloudtrace",
		"composer",
		"compute",
		"config",
		"container",
		"datacatalog",
		"dataflow",
		"datafusion",
		"datamigration",
		"datapipelines",
		"dataplex",
		"dataproc",
		"datastore",
		"datastream",
		"deploymentmanager",
		"developerconnect",
		"dns",
		"domains",
		"edgecontainer",
		"eventarc",
		"eventarcpublishing",
		"faultinjectiontesting",
		"file",
		"firebasedataconnect",
		"firestore",
		"genomics",
		"gkemulticloud",
		"healthcare",
		"iam",
		"iamcredentials",
		"iap",
		"ids",
		"krmapihosting",
		"language",
		"lifesciences",
		"logging",
		"looker",
		"managedidentities",
		"marketplacesolutions",
		"mediaasset",
		"memcache",
		"metastore",
		"monitoring",
		"netapp",
		"networkconnectivity",
		"networkmanagement",
		"networksecurity",
		"networkservices",
		"notebooks",
		"orgpolicy",
		"pam",
		"policyanalyzer",
		"privateca",
		"publicca",
		"pubsub",
		"recaptchaenterprise",
		"recommender",
		"redis",
		"resourcesettings",
		"run",
		"runtimeconfig",
		"sddc",
		"secretmanager",
		"securitycenter",
		"servicedirectory",
		"servicemanagement",
		"sourcerepo",
		"spanner",
		"speech",
		"sql",
		"storage",
		"testing",
		"transfer",
		"transferappliance",
		"vision",
		"vmmigration",
		"vmwareengine",
		"workflowexecutions",
		"workflows",
		"workstations",
	}

	for _, apiEndpointOverride := range apiEndpointOverrides {
		config.AddConfig(fmt.Sprintf("api_endpoint_overrides/%v", apiEndpointOverride), fmt.Sprintf("http://%s.googleapis.com/", apiEndpointOverride))
	}

	return config
}
