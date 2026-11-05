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

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	// Register all services
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

type noisyFilter struct {
}

func (f *noisyFilter) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, "/compute/") {
		return len(p), nil
	}
	if strings.Contains(s, "storing operation") && strings.Contains(s, "/operations/operation-") {
		// Also filter out generic operation storage logs if they are too noisy
		return len(p), nil
	}
	return os.Stdout.Write(p)
}

func main() {
	// Silence klog (used by many k8s libs and mockgcp)
	klog.SetOutput(&noisyFilter{})
	// Silence standard log
	log.SetOutput(&noisyFilter{})

	// Ensure klog doesn't log to stderr by default
	flag.Set("logtostderr", "false")
	flag.Set("stderrthreshold", "fatal")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.GetConfig()
	var k8sClient client.Client
	if err != nil {
		fmt.Printf("Warning: could not get k8s config: %v. MockGCP might have limited functionality.\n", err)
	} else {
		k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
		if err != nil {
			log.Fatalf("failed to create k8s client: %v", err)
		}
	}

	t := &testing.T{}
	mockCloud := mockgcp.NewMockRoundTripperForTest(t, k8sClient, storage.NewInMemoryStorage())

	// Initialize mock-project
	initializeMockProject(mockCloud)

	// Start a proxy server that uses mockCloud
	server := &http.Server{
		Addr: ":8082",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origHost := r.Host
			// Strip port from host
			host, _, err := net.SplitHostPort(r.Host)
			if err == nil {
				r.Host = host
			}

			// If host is localhost or a Kind gateway IP, MockGCP won't recognize it.
			// We try to infer the host from the URL path.
			if strings.HasPrefix(r.URL.Path, "/v1/projects") || strings.HasPrefix(r.URL.Path, "/v1beta1/projects") {
				r.Host = "essentialcontacts.googleapis.com"
			} else if strings.HasPrefix(r.URL.Path, "/compute") {
				r.Host = "compute.googleapis.com"
			} else if strings.HasPrefix(r.URL.Path, "/v1/organizations") {
				r.Host = "cloudresourcemanager.googleapis.com"
			}

			// Filter out noisy requests from our proxy log
			isNoisy := false
			if strings.Contains(r.URL.Path, "/compute/") {
				isNoisy = true
			}

			if !isNoisy {
				fmt.Printf("MockGCP Proxy: %s %s (Host: %s, Orig: %s)\n", r.Method, r.URL.Path, r.Host, origHost)
			}

			resp, err := mockCloud.(http.RoundTripper).RoundTrip(r)
			if err != nil {
				if !isNoisy {
					fmt.Printf("RoundTrip error: %v\n", err)
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if resp == nil {
				if !isNoisy {
					fmt.Printf("RoundTrip returned nil response\n")
				}
				http.Error(w, "nil response from mockgcp", http.StatusInternalServerError)
				return
			}

			if resp.Body != nil {
				defer resp.Body.Close()
			}

			for k, vv := range resp.Header {
				for _, v := range vv {
					w.Header().Add(k, v)
				}
			}
			w.WriteHeader(resp.StatusCode)
			if resp.Body != nil {
				// Copy body
				var buf [32 * 1024]byte
				for {
					n, err := resp.Body.Read(buf[:])
					if n > 0 {
						w.Write(buf[:n])
					}
					if err != nil {
						break
					}
				}
			}
		}),
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("MockGCP Server listening on %s\n", ln.Addr())

	go func() {
		if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Stopping MockGCP Server...")
	server.Shutdown(ctx)
}

func initializeMockProject(mockCloud mockgcp.Interface) {
	fmt.Println("Initializing mock-project in MockGCP...")

	project := map[string]interface{}{
		"projectId": "mock-project",
		"name":      "mock-project",
	}
	body, _ := json.Marshal(project)

	req, _ := http.NewRequest("POST", "https://cloudresourcemanager.googleapis.com/v1/projects", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := mockCloud.(http.RoundTripper).RoundTrip(req)
	if err != nil {
		fmt.Printf("Warning: failed to initialize mock-project: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusConflict {
		fmt.Printf("Warning: mock-project initialization returned status: %d\n", resp.StatusCode)
	} else {
		fmt.Println("Successfully initialized mock-project.")
	}
}
