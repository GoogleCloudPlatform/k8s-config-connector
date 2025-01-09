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

	// We need to register these over http, to stop gcloud trying to use TUNNEL with our proxy
	services := mockgcp.ListServices()
	for _, service := range services {
		for _, host := range service.Hosts {
			if strings.HasSuffix(host, ".googleapis.com") {
				key := strings.TrimSuffix(host, ".googleapis.com")
				config.AddConfig(fmt.Sprintf("api_endpoint_overrides/%v", key), fmt.Sprintf("http://%s.googleapis.com/", key))
			} else {
				// Probably not actually fatal, but unexpected (today)
				klog.Fatalf("unhandled host %q", host)
			}
		}
	}

	return config
}
