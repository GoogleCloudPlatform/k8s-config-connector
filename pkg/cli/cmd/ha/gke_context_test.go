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

package ha

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	"k8s.io/client-go/rest"
)

func TestParseGKEContext(t *testing.T) {
	raw := "gke_my-project_us-central1_cluster-a"

	got, err := parseGKEContext(raw)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := gkeClusterInfo{
		contextName: "gke_my-project_us-central1_cluster-a",
		project:     "my-project",
		location:    "us-central1",
		clusterName: "cluster-a",
	}
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestParseGKEContextErrors(t *testing.T) {
	tests := []struct {
		name              string
		contextName       string
		wantErrSubstrings []string
	}{
		{
			name:              "non-GKE context name without gke prefix",
			contextName:       "cluster-a",
			wantErrSubstrings: []string{"does not match expected GKE context format"},
		},
		{
			name:              "missing cluster segment",
			contextName:       "gke_my-project_us-central1_",
			wantErrSubstrings: []string{"does not match expected GKE context format"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctxName := tc.contextName

			_, err := parseGKEContext(ctxName)

			if err == nil {
				t.Fatalf("expected error containing %v, got nil", tc.wantErrSubstrings)
			}
			for _, sub := range tc.wantErrSubstrings {
				if !strings.Contains(err.Error(), sub) {
					t.Fatalf("expected error %q to contain %q", err.Error(), sub)
				}
			}
		})
	}
}

func TestResolveGKECluster_FastPath(t *testing.T) {
	t.Run("standard gke context name (zero API calls)", func(t *testing.T) {
		runner := &recordingCommandRunner{}
		handle := &clusterHandle{
			contextName: "gke_my-proj_us-central1_cluster-a",
			client:      &kubecli.Client{},
		}

		got, err := resolveGKECluster(context.Background(), handle, runner)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := gkeClusterInfo{
			contextName: "gke_my-proj_us-central1_cluster-a",
			project:     "my-proj",
			location:    "us-central1",
			clusterName: "cluster-a",
		}
		if got != want {
			t.Fatalf("got %+v, want %+v", got, want)
		}
		if len(runner.commands) != 0 {
			t.Fatalf("expected 0 commands, got %v", runner.commands)
		}
	})

	t.Run("renamed context with underlying gke cluster (zero API calls)", func(t *testing.T) {
		runner := &recordingCommandRunner{}
		handle := &clusterHandle{
			contextName: "alias-cluster-a",
			clusterName: "gke_my-proj_us-central1_cluster-a",
			client:      &kubecli.Client{},
		}

		got, err := resolveGKECluster(context.Background(), handle, runner)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := gkeClusterInfo{
			contextName: "alias-cluster-a",
			project:     "my-proj",
			location:    "us-central1",
			clusterName: "cluster-a",
		}
		if got != want {
			t.Fatalf("got %+v, want %+v", got, want)
		}
		if len(runner.commands) != 0 {
			t.Fatalf("expected 0 commands, got %v", runner.commands)
		}
	})
}

func TestResolveGKECluster_ConnectGateway(t *testing.T) {
	t.Run("global connect gateway URL", func(t *testing.T) {
		runner := &recordingCommandRunner{
			outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
				wantCmd := "gcloud container fleet memberships describe mem-a --location=global --project=fleet-proj --format=json"
				gotCmd := name + " " + strings.Join(args, " ")
				if gotCmd != wantCmd {
					return nil, fmt.Errorf("unexpected command %q, want %q", gotCmd, wantCmd)
				}
				return []byte(`{
					"endpoint": {
						"gkeCluster": {
							"resourceLink": "//container.googleapis.com/projects/workload-proj/locations/us-central1/clusters/cluster-a"
						}
					}
				}`), nil
			},
		}

		handle := &clusterHandle{
			contextName: "connectgateway_fleet-proj_global_mem-a",
			client: &kubecli.Client{
				RestConfig: &rest.Config{
					Host: "https://connectgateway.googleapis.com/v1/projects/fleet-proj/locations/global/gkeMemberships/mem-a",
				},
			},
		}

		got, err := resolveGKECluster(context.Background(), handle, runner)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := gkeClusterInfo{
			contextName: "connectgateway_fleet-proj_global_mem-a",
			project:     "workload-proj",
			location:    "us-central1",
			clusterName: "cluster-a",
		}
		if got != want {
			t.Fatalf("got %+v, want %+v", got, want)
		}
	})

	t.Run("regional connect gateway URL with numeric project number", func(t *testing.T) {
		runner := &recordingCommandRunner{
			outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
				gotCmd := name + " " + strings.Join(args, " ")
				switch gotCmd {
				case "gcloud container fleet memberships describe mem-b --location=us-east1 --project=111222333 --format=json":
					return []byte(`{
						"endpoint": {
							"gkeCluster": {
								"resourceLink": "//container.googleapis.com/projects/444555666/locations/us-east1/clusters/cluster-b"
							}
						}
					}`), nil
				case "gcloud projects describe 444555666 --format=value(projectId)":
					return []byte("resolved-gke-proj\n"), nil
				default:
					return nil, fmt.Errorf("unexpected command %q", gotCmd)
				}
			},
		}

		handle := &clusterHandle{
			contextName: "connectgateway_111222333_us-east1_mem-b",
			client: &kubecli.Client{
				RestConfig: &rest.Config{
					Host: "https://us-east1-connectgateway.googleapis.com/v1/projects/111222333/locations/us-east1/gkeMemberships/mem-b",
				},
			},
		}

		got, err := resolveGKECluster(context.Background(), handle, runner)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := gkeClusterInfo{
			contextName: "connectgateway_111222333_us-east1_mem-b",
			project:     "resolved-gke-proj",
			location:    "us-east1",
			clusterName: "cluster-b",
		}
		if got != want {
			t.Fatalf("got %+v, want %+v", got, want)
		}
	})

	t.Run("cross-project fleet host and membership name != cluster name", func(t *testing.T) {
		runner := &recordingCommandRunner{
			outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
				return []byte(`{
					"endpoint": {
						"gkeCluster": {
							"resourceLink": "//container.googleapis.com/projects/actual-gke-proj/locations/europe-west1/clusters/actual-cluster"
						}
					}
				}`), nil
			},
		}

		handle := &clusterHandle{
			contextName: "custom-gateway-context",
			client: &kubecli.Client{
				RestConfig: &rest.Config{
					Host: "https://connectgateway.googleapis.com/v1beta1/projects/fleet-host-proj/locations/global/gkeMemberships/custom-mem",
				},
			},
		}

		got, err := resolveGKECluster(context.Background(), handle, runner)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := gkeClusterInfo{
			contextName: "custom-gateway-context",
			project:     "actual-gke-proj",
			location:    "europe-west1",
			clusterName: "actual-cluster",
		}
		if got != want {
			t.Fatalf("got %+v, want %+v", got, want)
		}
	})
}

func TestResolveGKECluster_NonGKEFleetMembership(t *testing.T) {
	runner := &recordingCommandRunner{
		outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return []byte(`{
				"endpoint": {
					"attachedCluster": {
						"resourceLink": "//gkeonprem.googleapis.com/projects/123/locations/global/attachedClusters/test"
					}
				}
			}`), nil
		},
	}

	handle := &clusterHandle{
		contextName: "connectgateway_fleet-proj_global_attached-mem",
		client: &kubecli.Client{
			RestConfig: &rest.Config{
				Host: "https://connectgateway.googleapis.com/v1/projects/fleet-proj/locations/global/gkeMemberships/attached-mem",
			},
		},
	}

	_, err := resolveGKECluster(context.Background(), handle, runner)
	if err == nil {
		t.Fatalf("expected error for non-GKE cluster, got nil")
	}
	if !strings.Contains(err.Error(), "multi-cluster HA requires GKE clusters") {
		t.Fatalf("expected error to contain %q, got %q", "multi-cluster HA requires GKE clusters", err.Error())
	}
}

func TestResolveGKECluster_Errors(t *testing.T) {
	t.Run("nil handle", func(t *testing.T) {
		_, err := resolveGKECluster(context.Background(), nil, nil)
		if err == nil || !strings.Contains(err.Error(), "cluster handle is nil") {
			t.Fatalf("expected nil handle error, got %v", err)
		}
	})

	t.Run("unrecognized context format", func(t *testing.T) {
		handle := &clusterHandle{
			contextName: "invalid-context",
			client:      &kubecli.Client{},
		}
		_, err := resolveGKECluster(context.Background(), handle, nil)
		if err == nil || !strings.Contains(err.Error(), "does not match expected GKE context format") {
			t.Fatalf("expected invalid format error, got %v", err)
		}
	})

	t.Run("gcloud fleet membership describe failure", func(t *testing.T) {
		runner := &recordingCommandRunner{
			outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
				return nil, fmt.Errorf("permission denied")
			},
		}
		handle := &clusterHandle{
			contextName: "connectgateway_fleet_global_mem",
			client: &kubecli.Client{
				RestConfig: &rest.Config{
					Host: "https://connectgateway.googleapis.com/v1/projects/fleet/locations/global/gkeMemberships/mem",
				},
			},
		}
		_, err := resolveGKECluster(context.Background(), handle, runner)
		if err == nil || !strings.Contains(err.Error(), "permission denied") {
			t.Fatalf("expected permission denied error, got %v", err)
		}
	})

	t.Run("gcloud projects describe failure", func(t *testing.T) {
		runner := &recordingCommandRunner{
			outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
				if strings.Contains(strings.Join(args, " "), "memberships describe") {
					return []byte(`{
						"endpoint": {
							"gkeCluster": {
								"resourceLink": "//container.googleapis.com/projects/12345/locations/global/clusters/test"
							}
						}
					}`), nil
				}
				return nil, fmt.Errorf("project not found")
			},
		}
		handle := &clusterHandle{
			contextName: "connectgateway_fleet_global_mem",
			client: &kubecli.Client{
				RestConfig: &rest.Config{
					Host: "https://connectgateway.googleapis.com/v1/projects/fleet/locations/global/gkeMemberships/mem",
				},
			},
		}
		_, err := resolveGKECluster(context.Background(), handle, runner)
		if err == nil || !strings.Contains(err.Error(), "project not found") {
			t.Fatalf("expected project not found error, got %v", err)
		}
	})
}
