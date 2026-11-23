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

//go:build unix

package ha

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const failoverSignalHelperEnv = "KCC_HA_FAILOVER_SIGNAL_HELPER"

func TestFailoverSignalHelperProcess(t *testing.T) {
	mode := os.Getenv(failoverSignalHelperEnv)
	if mode == "" {
		return
	}
	if mode == "nohup" {
		signal.Ignore(syscall.SIGHUP)
	}
	opts := newFailoverOptions()
	opts.pollInterval = 20 * time.Millisecond
	if mode == "short-bounce" {
		opts.signalBounceWindow = 50 * time.Millisecond
	}
	cmd := newFailoverCmd(opts)
	cmd.SetArgs([]string{"--contexts=cluster-a,cluster-b", "--target-cluster=cluster-b"})
	if err := cmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

// signalTestCluster serves the minimal Kubernetes endpoints that `ha failover` calls so a child
// process can run NewFailoverCmd end to end and receive real OS signals during handover.
type signalTestCluster struct {
	mu           sync.Mutex
	name         string
	holder       string
	renewTime    string
	managers     []*appsv1.StatefulSet
	syncers      []syncerv1alpha1.KRMSyncer
	patchDelay   time.Duration
	hangOnStart  bool
	stoppedEvent chan struct{}
	server       *httptest.Server
}

func newSignalTestCluster(t *testing.T, name, holder string, namespaces []string) *signalTestCluster {
	t.Helper()
	renew := "2026-11-20T00:00:00Z"
	c := &signalTestCluster{
		name:         name,
		holder:       holder,
		renewTime:    renew,
		stoppedEvent: make(chan struct{}),
	}
	leaderGKE := operatorv1beta1.GKECluster{Project: "test-project", Location: "us-central1", Name: "cluster-a"}
	isLeader := fleetIdentity(name) == holder
	for _, ns := range namespaces {
		stsName := "cnrm-controller-manager"
		syncerName := fleetLeaseName
		if ns != "" {
			stsName += "-" + ns
			syncerName += "-" + ns
		}
		sts := newManagerStatefulSet("cnrm-system", stsName, ns)
		sts.Generation = 1
		sts.Status = appsv1.StatefulSetStatus{ObservedGeneration: 1, Replicas: 1, ReadyReplicas: 1}
		c.managers = append(c.managers, sts)
		c.syncers = append(c.syncers, syncerv1alpha1.KRMSyncer{
			TypeMeta:   metav1.TypeMeta{APIVersion: "syncer.gkelabs.io/v1alpha1", Kind: "KRMSyncer"},
			ObjectMeta: metav1.ObjectMeta{Namespace: fleetLeaseNamespace, Name: syncerName},
			Spec:       newSyncerSpec(leaderGKE, isLeader),
		})
	}
	c.server = httptest.NewServer(http.HandlerFunc(c.serveHTTP))
	t.Cleanup(c.server.Close)
	return c
}

func (c *signalTestCluster) serveHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.URL.Path == "/api":
		writeJSON(w, metav1.APIVersions{TypeMeta: metav1.TypeMeta{Kind: "APIVersions"}, Versions: []string{"v1"}})
	case r.URL.Path == "/api/v1":
		writeJSON(w, metav1.APIResourceList{TypeMeta: metav1.TypeMeta{Kind: "APIResourceList", APIVersion: "v1"}, GroupVersion: "v1"})
	case r.URL.Path == "/apis":
		writeJSON(w, metav1.APIGroupList{
			TypeMeta: metav1.TypeMeta{Kind: "APIGroupList", APIVersion: "v1"},
			Groups: []metav1.APIGroup{
				apiGroup("apps", "v1"),
				apiGroup("core.cnrm.cloud.google.com", "v1beta1"),
				apiGroup("multicluster.core.cnrm.cloud.google.com", "v1alpha1"),
				apiGroup("syncer.gkelabs.io", "v1alpha1"),
			},
		})
	case r.URL.Path == "/apis/apps/v1":
		writeJSON(w, resourceList("apps/v1", metav1.APIResource{Name: "statefulsets", SingularName: "statefulset", Namespaced: true, Kind: "StatefulSet", Verbs: metav1.Verbs{"get", "list", "patch"}}))
	case r.URL.Path == "/apis/core.cnrm.cloud.google.com/v1beta1":
		writeJSON(w, resourceList("core.cnrm.cloud.google.com/v1beta1", metav1.APIResource{Name: "configconnectors", SingularName: "configconnector", Namespaced: false, Kind: "ConfigConnector", Verbs: metav1.Verbs{"get"}}))
	case r.URL.Path == "/apis/multicluster.core.cnrm.cloud.google.com/v1alpha1":
		writeJSON(w, resourceList("multicluster.core.cnrm.cloud.google.com/v1alpha1", metav1.APIResource{Name: "multiclusterleases", SingularName: "multiclusterlease", Namespaced: true, Kind: "MultiClusterLease", Verbs: metav1.Verbs{"get"}}))
	case r.URL.Path == "/apis/syncer.gkelabs.io/v1alpha1":
		writeJSON(w, resourceList("syncer.gkelabs.io/v1alpha1", metav1.APIResource{Name: "krmsyncers", SingularName: "krmsyncer", Namespaced: true, Kind: "KRMSyncer", Verbs: metav1.Verbs{"list"}}))
	case r.URL.Path == "/apis/core.cnrm.cloud.google.com/v1beta1/configconnectors/"+configConnectorResourceName:
		cc := newConfigConnectorWithMCL(fleetLeaseName, fleetLeaseNamespace, c.name)
		cc.TypeMeta = metav1.TypeMeta{APIVersion: "core.cnrm.cloud.google.com/v1beta1", Kind: "ConfigConnector"}
		writeJSON(w, cc)
	case r.URL.Path == "/apis/multicluster.core.cnrm.cloud.google.com/v1alpha1/namespaces/"+fleetLeaseNamespace+"/multiclusterleases/"+fleetLeaseName:
		c.mu.Lock()
		holder, renew := c.holder, c.renewTime
		c.mu.Unlock()
		writeJSON(w, mclv1alpha1.MultiClusterLease{
			TypeMeta:   metav1.TypeMeta{APIVersion: "multicluster.core.cnrm.cloud.google.com/v1alpha1", Kind: "MultiClusterLease"},
			ObjectMeta: metav1.ObjectMeta{Namespace: fleetLeaseNamespace, Name: fleetLeaseName},
			Status:     mclv1alpha1.MultiClusterLeaseStatus{GlobalHolderIdentity: &holder, GlobalRenewTime: &renew},
		})
	case strings.HasSuffix(r.URL.Path, "/krmsyncers"):
		c.mu.Lock()
		items := append([]syncerv1alpha1.KRMSyncer(nil), c.syncers...)
		c.mu.Unlock()
		writeJSON(w, syncerv1alpha1.KRMSyncerList{
			TypeMeta: metav1.TypeMeta{APIVersion: "syncer.gkelabs.io/v1alpha1", Kind: "KRMSyncerList"},
			Items:    items,
		})
	case r.Method == http.MethodGet && r.URL.Path == "/apis/apps/v1/statefulsets":
		c.mu.Lock()
		items := make([]appsv1.StatefulSet, len(c.managers))
		for i, m := range c.managers {
			items[i] = *m.DeepCopy()
		}
		c.mu.Unlock()
		writeJSON(w, appsv1.StatefulSetList{
			TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "StatefulSetList"},
			Items:    items,
		})
	case r.Method == http.MethodPatch && strings.Contains(r.URL.Path, "/statefulsets/"):
		body, _ := io.ReadAll(r.Body)
		var patch struct {
			Metadata struct {
				Annotations map[string]*string `json:"annotations"`
			} `json:"metadata"`
			Spec struct {
				Replicas *int32 `json:"replicas"`
			} `json:"spec"`
		}
		_ = json.Unmarshal(body, &patch)

		c.mu.Lock()
		delay := c.patchDelay
		hang := c.hangOnStart && patch.Spec.Replicas != nil && *patch.Spec.Replicas > 0
		c.mu.Unlock()
		if hang {
			<-r.Context().Done()
			return
		}
		if delay > 0 && patch.Spec.Replicas != nil && *patch.Spec.Replicas > 0 {
			time.Sleep(delay)
		}

		stsName := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		c.mu.Lock()
		var out appsv1.StatefulSet
		allStopped := true
		for _, m := range c.managers {
			if m.Name == stsName {
				if patch.Spec.Replicas != nil {
					r := *patch.Spec.Replicas
					m.Spec.Replicas = &r
					m.Generation++
					m.Status.ObservedGeneration = m.Generation
					m.Status.Replicas = r
					m.Status.ReadyReplicas = r
				}
				if ann, ok := patch.Metadata.Annotations[failoverReplicasAnnotation]; ok {
					if m.Annotations == nil {
						m.Annotations = map[string]string{}
					}
					if ann == nil {
						delete(m.Annotations, failoverReplicasAnnotation)
					} else {
						m.Annotations[failoverReplicasAnnotation] = *ann
					}
				}
				out = *m.DeepCopy()
			}
			if m.Spec.Replicas == nil || *m.Spec.Replicas > 0 {
				allStopped = false
			}
		}
		if allStopped {
			select {
			case <-c.stoppedEvent:
			default:
				close(c.stoppedEvent)
			}
		}
		c.mu.Unlock()
		out.TypeMeta = metav1.TypeMeta{APIVersion: "apps/v1", Kind: "StatefulSet"}
		writeJSON(w, out)
	default:
		http.NotFound(w, r)
	}
}

func (c *signalTestCluster) assertAllManagersRestarted(t *testing.T) {
	t.Helper()
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, m := range c.managers {
		if m.Spec.Replicas == nil || *m.Spec.Replicas != 1 {
			t.Errorf("StatefulSet %s has replicas %v, want 1", m.Name, m.Spec.Replicas)
		}
		if _, ok := m.Annotations[failoverReplicasAnnotation]; ok {
			t.Errorf("StatefulSet %s still carries %s", m.Name, failoverReplicasAnnotation)
		}
	}
}

func apiGroup(name, version string) metav1.APIGroup {
	gv := metav1.GroupVersionForDiscovery{GroupVersion: name + "/" + version, Version: version}
	return metav1.APIGroup{Name: name, Versions: []metav1.GroupVersionForDiscovery{gv}, PreferredVersion: gv}
}

func resourceList(gv string, res metav1.APIResource) metav1.APIResourceList {
	return metav1.APIResourceList{
		TypeMeta:     metav1.TypeMeta{Kind: "APIResourceList", APIVersion: "v1"},
		GroupVersion: gv,
		APIResources: []metav1.APIResource{res},
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	_ = json.NewEncoder(w).Encode(v)
}

func writeSignalKubeconfig(t *testing.T, a, b *signalTestCluster) string {
	t.Helper()
	content := fmt.Sprintf(`apiVersion: v1
kind: Config
clusters:
- name: cluster-a
  cluster:
    server: %s
- name: cluster-b
  cluster:
    server: %s
contexts:
- name: cluster-a
  context:
    cluster: cluster-a
- name: cluster-b
  context:
    cluster: cluster-b
`, a.server.URL, b.server.URL)
	path := filepath.Join(t.TempDir(), "kubeconfig")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatalf("writing kubeconfig: %v", err)
	}
	return path
}

func startFailoverSubprocess(t *testing.T, mode, kubeconfig string, stdout io.Writer, stderr *bytes.Buffer) *exec.Cmd {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=^TestFailoverSignalHelperProcess$")
	cmd.Env = append(os.Environ(),
		failoverSignalHelperEnv+"="+mode,
		"KUBECONFIG="+kubeconfig,
		"KUBECACHEDIR="+filepath.Join(t.TempDir(), "cache"),
	)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("starting helper process: %v", err)
	}
	return cmd
}

func waitStopped(t *testing.T, a *signalTestCluster) {
	t.Helper()
	select {
	case <-a.stoppedEvent:
	case <-time.After(10 * time.Second):
		t.Fatal("timed out waiting for helper process to stop cluster-a's managers")
	}
}

func waitSubprocess(t *testing.T, cmd *exec.Cmd) {
	t.Helper()
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		_ = cmd.Process.Kill()
		t.Fatal("timed out waiting for helper process to exit after signal")
	}
}

func TestFailoverSubprocessSignalsAndBrokenPipe(t *testing.T) {
	holder := fleetIdentity("cluster-a")

	t.Run("SIGINT prints feedback and restarts stopped managers", func(t *testing.T) {
		a := newSignalTestCluster(t, "cluster-a", holder, []string{""})
		b := newSignalTestCluster(t, "cluster-b", holder, []string{""})
		var stdout, stderr bytes.Buffer
		cmd := startFailoverSubprocess(t, "normal", writeSignalKubeconfig(t, a, b), &stdout, &stderr)
		waitStopped(t, a)

		_ = cmd.Process.Signal(os.Interrupt)
		waitSubprocess(t, cmd)

		a.assertAllManagersRestarted(t)
		if !strings.Contains(stderr.String(), "Interrupted; starting any stopped managers again before exiting") {
			t.Errorf("stderr missing interrupt notice:\n%s", stderr.String())
		}
	})

	t.Run("rapid double SIGINT within bounce window still finishes restarting managers", func(t *testing.T) {
		a := newSignalTestCluster(t, "cluster-a", holder, []string{"team-a", "team-b"})
		a.patchDelay = 120 * time.Millisecond
		b := newSignalTestCluster(t, "cluster-b", holder, []string{"team-a", "team-b"})
		var stdout, stderr bytes.Buffer
		cmd := startFailoverSubprocess(t, "normal", writeSignalKubeconfig(t, a, b), &stdout, &stderr)
		waitStopped(t, a)

		_ = cmd.Process.Signal(os.Interrupt)
		time.Sleep(50 * time.Millisecond)
		_ = cmd.Process.Signal(os.Interrupt)
		waitSubprocess(t, cmd)

		a.assertAllManagersRestarted(t)
	})

	t.Run("second SIGINT after bounce window force-exits when cleanup hangs", func(t *testing.T) {
		a := newSignalTestCluster(t, "cluster-a", holder, []string{""})
		a.hangOnStart = true
		b := newSignalTestCluster(t, "cluster-b", holder, []string{""})
		var stdout, stderr bytes.Buffer
		cmd := startFailoverSubprocess(t, "short-bounce", writeSignalKubeconfig(t, a, b), &stdout, &stderr)
		waitStopped(t, a)

		_ = cmd.Process.Signal(syscall.SIGINT)
		time.Sleep(120 * time.Millisecond)
		_ = cmd.Process.Signal(syscall.SIGINT)

		done := make(chan error, 1)
		go func() { done <- cmd.Wait() }()
		select {
		case err := <-done:
			var exitErr *exec.ExitError
			if !errors.As(err, &exitErr) {
				t.Fatalf("expected ExitError on second SIGINT, got %v", err)
			}
			ws, ok := exitErr.Sys().(syscall.WaitStatus)
			if !ok || !ws.Signaled() || ws.Signal() != syscall.SIGINT {
				t.Errorf("expected termination by SIGINT, got status %v", exitErr.ProcessState)
			}
		case <-time.After(5 * time.Second):
			_ = cmd.Process.Kill()
			t.Fatal("second SIGINT after bounce window did not force-exit promptly")
		}
	})

	t.Run("SIGHUP restarts stopped managers unless ignored by nohup", func(t *testing.T) {
		// Daemonized test runners (including CI agents) often start with SIGHUP set to SIG_IGN.
		// Installing a handler in the parent before execve causes the kernel to reset SIGHUP to
		// SIG_DFL in the child process, matching an interactive terminal session.
		hupReset := make(chan os.Signal, 1)
		signal.Notify(hupReset, syscall.SIGHUP)
		defer signal.Stop(hupReset)

		a := newSignalTestCluster(t, "cluster-a", holder, []string{""})
		b := newSignalTestCluster(t, "cluster-b", holder, []string{""})
		var stdout, stderr bytes.Buffer
		cmd := startFailoverSubprocess(t, "normal", writeSignalKubeconfig(t, a, b), &stdout, &stderr)
		waitStopped(t, a)

		_ = cmd.Process.Signal(syscall.SIGHUP)
		waitSubprocess(t, cmd)
		a.assertAllManagersRestarted(t)
		if !strings.Contains(stderr.String(), "Interrupted; starting any stopped managers again before exiting") {
			t.Errorf("stderr missing SIGHUP interrupt notice:\n%s", stderr.String())
		}

		// Under nohup (SIGHUP ignored), SIGHUP does not abort the handover.
		aNohup := newSignalTestCluster(t, "cluster-a", holder, []string{""})
		bNohup := newSignalTestCluster(t, "cluster-b", holder, []string{""})
		var outNohup, errNohup bytes.Buffer
		nohupCmd := startFailoverSubprocess(t, "nohup", writeSignalKubeconfig(t, aNohup, bNohup), &outNohup, &errNohup)
		waitStopped(t, aNohup)

		_ = nohupCmd.Process.Signal(syscall.SIGHUP)
		time.Sleep(80 * time.Millisecond)
		// cluster-a should still be stopped while waiting for takeover, then restart on SIGINT.
		aNohup.mu.Lock()
		stillStopped := *aNohup.managers[0].Spec.Replicas == 0
		aNohup.mu.Unlock()
		if !stillStopped {
			t.Errorf("SIGHUP aborted failover even though SIGHUP was ignored at startup")
		}
		_ = nohupCmd.Process.Signal(os.Interrupt)
		waitSubprocess(t, nohupCmd)
		aNohup.assertAllManagersRestarted(t)
	})

	t.Run("closed stdout pipe (SIGPIPE) does not abort restarting multiple managers", func(t *testing.T) {
		a := newSignalTestCluster(t, "cluster-a", holder, []string{"team-a", "team-b"})
		b := newSignalTestCluster(t, "cluster-b", holder, []string{"team-a", "team-b"})
		pr, pw, err := os.Pipe()
		if err != nil {
			t.Fatalf("os.Pipe: %v", err)
		}
		var stderr bytes.Buffer
		cmd := startFailoverSubprocess(t, "normal", writeSignalKubeconfig(t, a, b), pw, &stderr)
		_ = pw.Close()

		// Drain stdout until cluster-a is stopped, then close the reader as `| tee` does on Ctrl-C.
		go func() { _, _ = io.Copy(io.Discard, pr) }()
		waitStopped(t, a)
		_ = pr.Close()

		_ = cmd.Process.Signal(os.Interrupt)
		waitSubprocess(t, cmd)

		a.assertAllManagersRestarted(t)
	})
}
