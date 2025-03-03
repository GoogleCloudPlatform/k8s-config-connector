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
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptorand "crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"k8s.io/klog/v2"
)

type Proxy struct {
	httpClient *http.Client

	httpsEndpoint *net.TCPAddr
}

func NewProxy(httpClient *http.Client) *Proxy {
	return &Proxy{httpClient: httpClient}
}

func (p *Proxy) ListenAndServeHTTPS(ctx context.Context) error {
	if p.httpsEndpoint != nil {
		panic("StartHTTPS called twice")
	}

	httpsListener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return fmt.Errorf("net.Listen failed: %w", err)
	}

	p.httpsEndpoint = httpsListener.Addr().(*net.TCPAddr)

	httpsServer := &http.Server{}
	tlsCertificateTemplate := &x509.Certificate{
		IsCA: true,
	}
	// tlsCertificateTemplate.IPAddresses = append(tlsCertificateTemplate.IPAddresses, p.httpsEndpoint.IP)
	tlsCertificateTemplate.DNSNames = []string{"billingbudgets.googleapis.com", "tpu.googleapis.com", "play.googleapis.com"}
	tlsCertificate, err := CreateSelfSignedCertificate(tlsCertificateTemplate)
	if err != nil {
		return fmt.Errorf("creating self-signed certificate: %w", err)
	}
	httpsServer.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{*tlsCertificate},
	}
	tlsListener := tls.NewListener(httpsListener, httpsServer.TLSConfig)

	httpsServer.Handler = p

	errChan := make(chan error, 1)
	go func() {
		klog.Infof("https listener starting on %v", httpsListener.Addr())
		if err := httpsServer.Serve(tlsListener); err != nil {
			klog.Infof("https listener stopped")
			errChan <- err
		}
	}()

	go func() {
		<-ctx.Done()
		httpsServer.Shutdown(context.Background())
	}()

	err = <-errChan
	return err
}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "CONNECT" {
		p.handleConnect(rw, req)
		return
	}

	resp, err := p.runRequest(req)
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		klog.Errorf("error proxying request: %v", err)
		return
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

func (p *Proxy) handleConnect(connectResponseWriter http.ResponseWriter, connectRequest *http.Request) {
	klog.Infof("handling connect %v %v", connectRequest.Method, connectRequest.URL)

	// klog.Infof("dialing %v", p.httpsEndpoint.String())
	targetConnection, err := net.DialTimeout("tcp", p.httpsEndpoint.String(), 5*time.Second)
	if err != nil {
		http.Error(connectResponseWriter, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer targetConnection.Close()

	// Must send OK before hijacking connection
	connectResponseWriter.WriteHeader(http.StatusOK)

	inbound, ok := connectResponseWriter.(http.Hijacker)
	if !ok {
		http.Error(connectResponseWriter, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	sourceConnection, sourceReadWriter, err := inbound.Hijack()
	if err != nil {
		http.Error(connectResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	defer sourceConnection.Close()

	errors := make(chan error, 2)

	go func() {
		out := targetConnection
		in := sourceReadWriter
		for {
			buf := make([]byte, 32*1024)
			n, err := in.Read(buf)
			if err != nil {
				if err != io.EOF {
					klog.Infof("error while reading data: %v", err)
				}
				errors <- err
				return
			}
			_, err = out.Write(buf[:n])
			if err != nil {
				klog.Infof("error while writing data: %v", err)
				errors <- err
				return
			}
			// out.Flush() - no need, because out is unbuffered
			// klog.Infof("copied %v bytes", n)
		}
	}()
	go func() {
		out := sourceReadWriter
		in := targetConnection
		for {
			buf := make([]byte, 32*1024)
			n, err := in.Read(buf)
			if err != nil {
				if err != io.EOF {
					klog.Infof("error while reading data: %v", err)
				}
				errors <- err
				return
			}
			_, err = out.Write(buf[:n])
			if err != nil {
				klog.Infof("error while writing data: %v", err)
				errors <- err
				return
			}
			if err := out.Flush(); err != nil {
				klog.Infof("error from flush: %v", err)
			}
			// klog.Infof("copied %v bytes", n)
		}
	}()

	<-errors
	<-errors
}

func (p *Proxy) runRequest(req *http.Request) (*http.Response, error) {
	klog.Infof("proxying %v %v", req.Method, req.URL)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("reading request body: %v", err)
	}

	u := req.URL
	u.Scheme = "https"
	if u.Host == "" {
		// for CONNECT requests, the host appears here
		u.Host = req.Host
	}

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
	config.AddConfig("auth/disable_ssl_validation", "True")

	// We need to register services to use http, to stop gcloud trying to use TUNNEL with our proxy

	// We need a hard-coded list, because we don't always mockgcp available
	apiEndpointOverrides := []string{
		"accessapproval.googleapis.com",
		"accesscontextmanager.googleapis.com",
		"ai.googleapis.com",
		"aiplatform.googleapis.com",
		"anthosevents.googleapis.com",
		"anthospolicycontrollerstatus_pa.googleapis.com",
		"apigateway.googleapis.com",
		"apigee.googleapis.com",
		"appengine.googleapis.com",
		"apphub.googleapis.com",
		"artifactregistry.googleapis.com",
		"assuredworkloads.googleapis.com",
		"auditmanager.googleapis.com",
		"baremetalsolution.googleapis.com",
		"bigtableadmin.googleapis.com",
		"certificatemanager.googleapis.com",
		"cloudasset.googleapis.com",
		"cloudbilling.googleapis.com",
		"cloudbuild.googleapis.com",
		"cloudcommerceconsumerprocurement.googleapis.com",
		"clouddebugger.googleapis.com",
		"clouddeploy.googleapis.com",
		"clouderrorreporting.googleapis.com",
		"cloudfunctions.googleapis.com",
		"cloudidentity.googleapis.com",
		"cloudkms.googleapis.com",
		"cloudresourcemanager.googleapis.com",
		"cloudscheduler.googleapis.com",
		"cloudtasks.googleapis.com",
		"cloudtrace.googleapis.com",
		"composer.googleapis.com",
		"compute.googleapis.com",
		"config.googleapis.com",
		"container.googleapis.com",
		"datacatalog.googleapis.com",
		"dataflow.googleapis.com",
		"datafusion.googleapis.com",
		"datamigration.googleapis.com",
		"datapipelines.googleapis.com",
		"dataplex.googleapis.com",
		"dataproc.googleapis.com",
		"datastore.googleapis.com",
		"datastream.googleapis.com",
		"deploymentmanager.googleapis.com",
		"developerconnect.googleapis.com",
		"dns.googleapis.com",
		"domains.googleapis.com",
		"edgecontainer.googleapis.com",
		"eventarc.googleapis.com",
		"eventarcpublishing.googleapis.com",
		"faultinjectiontesting.googleapis.com",
		"file.googleapis.com",
		"firebasedataconnect.googleapis.com",
		"firestore.googleapis.com",
		"genomics.googleapis.com",
		"gkemulticloud.googleapis.com",
		"healthcare.googleapis.com",
		"iam.googleapis.com",
		"iamcredentials.googleapis.com",
		"iap.googleapis.com",
		"ids.googleapis.com",
		"krmapihosting.googleapis.com",
		"language.googleapis.com",
		"lifesciences.googleapis.com",
		"logging.googleapis.com",
		"looker.googleapis.com",
		"managedidentities.googleapis.com",
		"marketplacesolutions.googleapis.com",
		"mediaasset.googleapis.com",
		"memcache.googleapis.com",
		"metastore.googleapis.com",
		"monitoring.googleapis.com",
		"netapp.googleapis.com",
		"networkconnectivity.googleapis.com",
		"networkmanagement.googleapis.com",
		"networksecurity.googleapis.com",
		"networkservices.googleapis.com",
		"notebooks.googleapis.com",
		"orgpolicy.googleapis.com",
		"pam.googleapis.com",
		"policyanalyzer.googleapis.com",
		"privateca.googleapis.com",
		"publicca.googleapis.com",
		"pubsub.googleapis.com",
		"recaptchaenterprise.googleapis.com",
		"recommender.googleapis.com",
		"redis.googleapis.com",
		"resourcesettings.googleapis.com",
		"run.googleapis.com",
		"runtimeconfig.googleapis.com",
		"sddc.googleapis.com",
		"secretmanager.googleapis.com",
		"securitycenter.googleapis.com",
		"servicedirectory.googleapis.com",
		"servicemanagement.googleapis.com",
		"sourcerepo.googleapis.com",
		"spanner.googleapis.com",
		"speech.googleapis.com",
		"sql.googleapis.com",
		"storage.googleapis.com",
		"testing.googleapis.com",
		"transfer.googleapis.com",
		"transferappliance.googleapis.com",
		"vision.googleapis.com",
		"vmmigration.googleapis.com",
		"vmwareengine.googleapis.com",
		"workflowexecutions.googleapis.com",
		"workflows.googleapis.com",
		"workstations.googleapis.com",
	}

	for _, apiEndpointOverride := range apiEndpointOverrides {
		config.AddConfig(fmt.Sprintf("api_endpoint_overrides/%v", apiEndpointOverride), fmt.Sprintf("http://%s.googleapis.com/", apiEndpointOverride))
	}

	return config
}

func CreateSelfSignedCertificate(template *x509.Certificate) (*tls.Certificate, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), cryptorand.Reader)
	if err != nil {
		return nil, fmt.Errorf("creating private key: %w", err)
	}

	publicKey := &privateKey.PublicKey

	template.KeyUsage = x509.KeyUsageDigitalSignature
	template.BasicConstraintsValid = true
	template.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}

	now := time.Now()
	template.NotBefore = now.Add(-time.Hour * 24 * 7)
	template.NotAfter = now.Add(time.Hour * 24 * 7)

	template.SerialNumber = big.NewInt(1)

	template.Subject = pkix.Name{
		Organization: []string{"kcc-test-proxy"},
	}

	if template.IsCA {
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(cryptorand.Reader, template, template, publicKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf("creating certificate: %w", err)
	}

	certificatePEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// encode and load the cert and private key for the client
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("encoding private key: %w", err)
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "PRIVATE KEY", Bytes: privateKeyBytes,
	})
	tlsCertificate, err := tls.X509KeyPair(certificatePEM, privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("building TLS certificate: %w", err)
	}

	return &tlsCertificate, err
}
