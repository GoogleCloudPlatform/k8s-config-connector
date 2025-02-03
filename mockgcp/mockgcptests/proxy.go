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
	services := []string{
		"compute.googleapis.com",
		"pubsub.googleapis.com",
		"storage.googleapis.com",
		"cloudresourcemanager.googleapis.com",
		"serviceusage.googleapis.com",
	}
	for _, service := range services {
		if strings.HasSuffix(service, ".googleapis.com") {
			key := strings.TrimSuffix(service, ".googleapis.com")
			config.AddConfig(fmt.Sprintf("api_endpoint_overrides/%v", key), fmt.Sprintf("http://%s.googleapis.com/", key))
		} else {
			// Probably not actually fatal, but unexpected (today)
			klog.Fatalf("unhandled host %q", service)
		}
	}

	return config
}
