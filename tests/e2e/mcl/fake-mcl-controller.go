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
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var leaseName, namespace, kubeconfig, kubecontext string
	flag.StringVar(&leaseName, "lease-name", "kcc-leader-lease", "Name of the lease")
	flag.StringVar(&namespace, "namespace", "cnrm-system", "Namespace of the lease")
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to kubeconfig")
	flag.StringVar(&kubecontext, "kubecontext", "", "Kubeconfig context to use")
	flag.Parse()

	if kubeconfig == "" {
		fmt.Println("Flag --kubeconfig must be set")
		os.Exit(1)
	}

	cfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{CurrentContext: kubecontext}).ClientConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building config for context %s: %v\n", kubecontext, err)
		os.Exit(1)
	}

	client, err := dynamic.NewForConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	gvr := schema.GroupVersionResource{
		Group:    "multicluster.core.cnrm.cloud.google.com",
		Version:  "v1alpha1",
		Resource: "multiclusterleases",
	}

	fmt.Printf("Fake MCL controller starting for context %s. Watching %s/%s\n", kubecontext, namespace, leaseName)
	for {
		ctx := context.Background()
		lease, err := client.Resource(gvr).Namespace(namespace).Get(ctx, leaseName, metav1.GetOptions{})
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		spec, ok := lease.Object["spec"].(map[string]interface{})
		if !ok {
			time.Sleep(1 * time.Second)
			continue
		}

		holderIdentity, _ := spec["holderIdentity"].(string)
		if holderIdentity != "" {
			candidateIdentity := holderIdentity
			candidateRenewTime, _ := spec["renewTime"].(string)

			status, _ := lease.Object["status"].(map[string]interface{})
			if status == nil {
				status = make(map[string]interface{})
			}

			globalHolderIdentity, _ := status["globalHolderIdentity"].(string)
			globalRenewTime, _ := status["globalRenewTime"].(string)

			updateStatus := false
			if globalHolderIdentity == "" {
				updateStatus = true
			} else if globalHolderIdentity == candidateIdentity {
				updateStatus = true
			} else {
				if globalRenewTime != "" {
					lastRenew, err := time.Parse(time.RFC3339Nano, globalRenewTime)
					if err == nil {
						if time.Since(lastRenew) > 10*time.Second {
							fmt.Printf("[%s] Global leader %s timed out, allowing %s to take over\n", kubecontext, globalHolderIdentity, candidateIdentity)
							updateStatus = true
						}
					}
				}
			}

			if updateStatus {
				var leaseDuration int64
				if val, ok := spec["leaseDurationSeconds"]; ok {
					switch v := val.(type) {
					case int64:
						leaseDuration = v
					case float64:
						leaseDuration = int64(v)
					case int:
						leaseDuration = int64(v)
					}
				}

				generation := lease.GetGeneration()

				status["globalHolderIdentity"] = candidateIdentity
				status["globalRenewTime"] = candidateRenewTime
				status["globalLeaseDurationSeconds"] = leaseDuration
				status["observedGeneration"] = generation
				status["conditions"] = []interface{}{
					map[string]interface{}{
						"type":               "BackendHealthy",
						"status":             "True",
						"lastTransitionTime": time.Now().Format(time.RFC3339Nano),
						"reason":             "FakeSuccess",
						"message":            "Fake MCL controller is running",
					},
				}

				lease.Object["status"] = status
				_, err := client.Resource(gvr).Namespace(namespace).UpdateStatus(ctx, lease, metav1.UpdateOptions{})
				if err != nil {
					fmt.Printf("[%s] Error updating status: %v\n", kubecontext, err)
				} else {
					fmt.Printf("[%s] Global leader is now: %s\n", kubecontext, candidateIdentity)
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
