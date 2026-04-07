/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Smoke test for validating kubebuilder-declarative-pattern changes.
//
// USAGE:
//   go run smoketest.go
//
// PREREQUISITES:
//   - A running Kubernetes cluster

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

const (
	// Set default namespace for your operator.
	defaultSystemNamespace = "guestbook-operator-system"
)

type (
	verifyFn   func(kubernetes.Interface, string, ...string) error
	verifyStep struct {
		fn       verifyFn
		prefixes []string
	}
)

func executeCommand(cmd string, args ...string) (string, error) {
	cmdStr := cmd
	if len(args) != 0 {
		cmdStr += " " + strings.Join(args, " ")
	}

	log.Printf("exec: %s", cmdStr)

	c := exec.Command(cmd, args...)
	var out bytes.Buffer
	c.Stdout = io.MultiWriter(os.Stdout, &out)
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		return "", fmt.Errorf("error execing command %q: %v", cmdStr, err)
	}

	return strings.TrimSuffix(out.String(), "\n"), nil
}

func mustExec(cmd string, args ...string) string {
	out, err := executeCommand(cmd, args...)
	if err != nil {
		panic(err)
	}
	return out
}

const (
	verifyTimeout   = 2 * time.Minute
	verifyFrequency = 5 * time.Second
)

var (
	imageTag            = flag.String("image-tag", "latest", "override the image tag for operator deployments")
	imageRepo           = flag.String("image-repo", "gcr.io/jrjohnson-gke", "rewrite images from gcr.io/jrjohnson-gke to a mirror")
	ignoreTests         = flag.String("ignore-tests", "", "comma separated list of test names to ignore")
	skipCustomScenarios = flag.Bool("skip-custom-scenarios", false, "Skip test scenarios custom for each operator")
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	h, err := NewRealTestHarness()
	if err != nil {
		log.Fatalf("error building test harness: %v", err)
	}

	var c CommonAddonTest
	c.Harness = h

	ignore := map[string]struct{}{}
	if *ignoreTests != "" {
		ignores := strings.Split(*ignoreTests, ",")
		for _, t := range ignores {
			ignore[t] = struct{}{}
		}
	}

	var operators []AddonTest
	for _, test := range []AddonTest{
		NewGuestbookTest(c),
	} {
		if _, ignored := ignore[test.Name()]; ignored {
			log.Printf("ignoring test: %s", test.Name())
		} else {
			operators = append(operators, test)
		}
	}

	klog.Info("Run: Deploying CRDs")
	for _, op := range operators {
		op.InstallCRDs()
	}
	klog.Info("Run: Deploying Operators")
	for _, op := range operators {
		op.InstallOperators()
	}
	klog.Info("Run: Deploying Addons (1/2)")
	for _, op := range operators {
		op.InstallResources()
	}

	klog.Info("Verify: Addons started")
	err = verifyAllUpOrTimeout(operators, verifyTimeout, "initial creation")
	if err != nil {
		klog.Errorf("verifying all up: %v", err)
	}

	klog.Info("Verify: Disrupted addons recover")
	for _, op := range operators {
		op.Disrupt()
	}
	err = verifyAllUpOrTimeout(operators, verifyTimeout, "disruption recovery")
	if err != nil {
		klog.Errorf("verifying all up: %v", err)
	}

	klog.Info("Verify: Addons delete")
	for _, op := range operators {
		op.DeleteResources()
	}
	for _, op := range operators {
		op.VerifyDown()
	}
	err = verifyAllDownOrTimeout(operators, verifyTimeout, "tear down")
	if err != nil {
		klog.Errorf("verifying all down: %v", err)
	}

	klog.Info("Run: Deploying Addons (2/2)")
	for _, op := range operators {
		op.InstallResources()
	}

	err = verifyAllUpOrTimeout(operators, verifyTimeout, "recreation")
	if err != nil {
		log.Fatal(err)
	}

	if !*skipCustomScenarios {
		klog.Info("Run/Verify: Addon specific scenarios")
		err = verifyCustomScenarios(operators, verifyTimeout, "custom testing scenarios")
		if err != nil {
			klog.Errorf("verifying custom scenarios: %v", err)
		}
	}

	klog.Infof("Clean up: Delete Addons")
	for _, op := range operators {
		op.DeleteResources()
	}
	klog.Infof("Clean up: Delete Operators")
	for _, op := range operators {
		op.DeleteOperators()
	}
	err = verifyAllDownOrTimeout(operators, verifyTimeout, "clean up")

	// TODO:
	// * Change CRD
	// * Verify that the addon has been upgraded

	// TODO:
	// * Change operator version
	// * Verify that the addon has been upgraded

	// TODO (currently not supported)
	// * Create multiple resources
	// * Verify multiple addon deployments

	if err != nil {
		log.Fatal(err)
	}
}

func verifyAllUpOrTimeout(operators []AddonTest, duration time.Duration, desc string) error {
	return verifyOrTimeout(operators, duration, desc, func(op AddonTest) error {
		return op.VerifyUp()
	})
}

func verifyCustomScenarios(operators []AddonTest, duration time.Duration, desc string) error {
	passing := len(operators)
	for _, op := range operators {
		err := op.CustomScenarios()
		if err != nil {
			log.Printf("verify error: %v", err)
			passing--
		}
	}
	log.Printf("[%d/%d] passing", passing, len(operators))
	if passing < len(operators) {
		return fmt.Errorf("error: failed to verify custom scenarios for operators")
	}
	return nil
}

func verifyAllDownOrTimeout(operators []AddonTest, duration time.Duration, desc string) error {
	return verifyOrTimeout(operators, duration, desc, func(op AddonTest) error {
		return op.VerifyDown()
	})
}

func verifyOrTimeout(operators []AddonTest, duration time.Duration, desc string, verify func(AddonTest) error) error {
	// Verify Workloads
	log.Printf("running verify tasks, retry period: %s, timeout: %s", verifyFrequency.String(), verifyTimeout.String())

	timeout := time.After(duration)
	for {
		var errs []error

		tests := 0
		for _, op := range operators {
			tests++
			if err := verify(op); err != nil {
				errs = append(errs, fmt.Errorf("%T: %v", op, err))
			}
		}

		if len(errs) == 0 {
			log.Print("all tests pass")
			return nil
		}
		for _, err := range errs {
			log.Printf("verify error: %v", err)
		}

		select {
		case <-timeout:
			return fmt.Errorf("error: failed to verify cluster (%s) after %s", desc, verifyTimeout.String())
		case <-time.After(verifyFrequency):
			log.Printf("[%d/%d] passing (%s)", tests-len(errs), tests, desc)
		}
	}
}

func verifyExistRole(clientset kubernetes.Interface, namespace string, role string) error {
	if _, err := clientset.RbacV1().Roles(namespace).Get(context.TODO(), role, metav1.GetOptions{}); err != nil {
		return err
	}
	return nil
}

func verifyExistClusterRole(clientset kubernetes.Interface, name string) error {
	if _, err := clientset.RbacV1().ClusterRoles().Get(context.TODO(), name, metav1.GetOptions{}); err != nil {
		return err
	}

	return nil
}

func verifyExistClusterRoleBinding(clientset kubernetes.Interface, name string) error {
	if _, err := clientset.RbacV1().ClusterRoleBindings().Get(context.TODO(), name, metav1.GetOptions{}); err != nil {
		return err
	}

	return nil
}

func verifySteps(clientset kubernetes.Interface, steps []verifyStep) error {
	for _, step := range steps {
		if err := step.fn(clientset, defaultSystemNamespace, step.prefixes...); err != nil {
			return err
		}
	}
	return nil
}

type PodSet []corev1.Pod

func Pods(h TestHarness, namespace string) PodSet {
	pods, err := h.Clientset().CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		h.Fatalf("error listing pods: %v", err)
	}

	return pods.Items
}

func (s *PodSet) WhereReady(ready bool) PodSet {
	return s.Where(func(pod *corev1.Pod) bool {
		isReady := false
		for _, cond := range pod.Status.Conditions {
			if cond.Type == corev1.PodReady && cond.Status == corev1.ConditionTrue {
				isReady = true
			}
		}

		return ready == isReady
	})
}

func (s *PodSet) Where(predicate func(*corev1.Pod) bool) PodSet {
	var matches []corev1.Pod

	for _, o := range *s {
		if predicate(&o) {
			matches = append(matches, o)
		}
	}

	return matches
}

func (s *PodSet) Count() int {
	return len(*s)
}

func verifyReadyPods(h TestHarness, namespace string, prefixes ...string) error {
	for _, prefix := range prefixes {
		pods := Pods(h, namespace)
		pods = pods.Where(func(pod *corev1.Pod) bool { return strings.HasPrefix(pod.Name, prefix) })

		if pods.Count() == 0 {
			return fmt.Errorf("no %s pod found in %s", prefix, namespace)
		}

		ready := pods.WhereReady(true)

		if len(ready) != len(pods) {
			return fmt.Errorf("%d pods ready with prefix %q, expected %d", len(ready), prefix, len(pods))
		}
	}

	return nil
}

func verifyNoWorkloadsWithLabel(label, namespace string) error {
	out := mustExec("kubectl", "get", "all", "-l", label, "-n", defaultSystemNamespace)
	if out != "" {
		return fmt.Errorf("Unexpected resources for label %q. Kubectl get returned %v lines of output", label, bytes.Count([]byte(out), []byte("\n")))
	}
	return nil
}

type TestHarness interface {
	KubectlApply(p string)
	KubectlDelete(p string)

	Clientset() kubernetes.Interface

	Fatalf(msg string, args ...interface{})
}

type AddonTest interface {
	Name() string

	InstallCRDs()
	InstallOperators()
	DeleteOperators()
	// TODO: InstallResources and DeleteResources should create/delete resources by name & namespace
	InstallResources()
	DeleteResources()

	// TODO: provide common implementation
	VerifyUp() error
	CustomScenarios() error
	VerifyDown() error
	Disrupt()
}

type CommonAddonTest struct {
	Base    string
	Harness TestHarness
}

func (c *CommonAddonTest) Name() string {
	return c.Base
}

// RealTestHarness is an implementation of TestHarness that runs against a real cluster
type RealTestHarness struct {
	clientset kubernetes.Interface
}

func NewRealTestHarness() (*RealTestHarness, error) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		return nil, fmt.Errorf("error building config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating clientset: %v", err)
	}

	harness := &RealTestHarness{
		clientset: clientset,
	}

	return harness, nil
}

func (h *RealTestHarness) Clientset() kubernetes.Interface {
	return h.clientset
}

func (h *RealTestHarness) KubectlApply(p string) {
	stat, err := os.Stat(p)
	if err != nil {
		h.Fatalf("error doing stat on %s: %v", p, err)
	}

	if stat.IsDir() {
		files, err := os.ReadDir(p)
		if err != nil {
			h.Fatalf("error reading directory %s: %v", p, err)
		}

		for _, f := range files {
			name := f.Name()

			// Ignore editor files & hidden files
			if strings.HasSuffix(name, "~") {
				continue
			}
			if strings.HasPrefix(name, ".") {
				continue
			}

			fp := filepath.Join(p, f.Name())
			if f.IsDir() {
				h.KubectlApply(fp)
			} else {
				h.kubectlApplyFile(fp)
			}
		}
	} else {
		h.kubectlApplyFile(p)
	}
}

func (h *RealTestHarness) KubectlDelete(p string) {
	mustExec("kubectl", "delete", "-f", p, "--recursive")
}

func (h *RealTestHarness) kubectlApplyFile(p string) {
	b, err := os.ReadFile(p)
	if err != nil {
		h.Fatalf("error reading file %s: %v", p, err)
	}

	klog.Infof("applying file %s", p)

	repo := strings.TrimSuffix(*imageRepo, "/")
	tag := *imageTag

	// Replace hardcoded controller image names with provided imageRepo/imageTag flags
	re := regexp.MustCompile(`gcr.io/jrjohnson-gke\/(.*):latest`)
	s := re.ReplaceAllString(string(b), fmt.Sprintf("%s/$1:%s", repo, tag))

	h.kubectlApplyString(s)
}

func (h *RealTestHarness) kubectlApplyString(s string) {
	cmd := "kubectl"
	args := []string{"apply", "-f", "-"}

	c := exec.Command(cmd, args...)
	c.Stdin = bytes.NewReader([]byte(s))
	var out bytes.Buffer
	c.Stdout = io.MultiWriter(os.Stdout, &out)
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		h.Fatalf("error from kubectl apply: %s", out.String())
	}
}

func (h *RealTestHarness) Fatalf(msg string, args ...interface{}) {
	h.Fatal(fmt.Sprintf(msg, args...))
}

func (h *RealTestHarness) Fatal(msg string) {
	panic(msg)
}

var _ TestHarness = &RealTestHarness{}

var _ AddonTest = &CommonAddonTest{}

func (c *CommonAddonTest) Basedir() string {
	return c.Base
}

func (c *CommonAddonTest) InstallCRDs() {
	mustExec("make", "-C", c.Basedir(), "install")
}

func (c *CommonAddonTest) InstallOperators() {
	mustExec("make", "-C", c.Basedir(), "docker-build", "docker-push", "deploy")
}

func (c *CommonAddonTest) DeleteOperators() {
	mustExec("make", "-C", c.Basedir(), "teardown")
}

func (c *CommonAddonTest) InstallResources() {
	c.Harness.KubectlApply(filepath.Join(c.Basedir(), "config", "samples"))
}

func (c *CommonAddonTest) DeleteResources() {
	c.Harness.KubectlDelete(filepath.Join(c.Basedir(), "config", "samples"))
}

func (c *CommonAddonTest) VerifyUp() error {
	return fmt.Errorf("VerifyUp not implemented for operator: %s", c.Name())
}

func (c *CommonAddonTest) CustomScenarios() error {
	// none defined by default
	return nil
}

func (c *CommonAddonTest) Disrupt() {
	// no-op
	// Specific operators should add specific disruption scenarios
	klog.Infof("Disrupt not configured for operator: %s", c.Name())
}

func (c *CommonAddonTest) VerifyDown() error {
	return fmt.Errorf("VerifyDown not implemented for operator: %s", c.Name())
}

type GuestbookTest struct {
	CommonAddonTest
}

func NewGuestbookTest(c CommonAddonTest) *GuestbookTest {
	t := &GuestbookTest{CommonAddonTest: c}
	t.Base = "../examples/guestbook-operator"
	return t
}

func (k *GuestbookTest) VerifyUp() error {
	h := k.Harness

	err := verifyReadyPods(h, defaultSystemNamespace, "guestbook-operator")
	if err != nil {
		return err
	}

	return nil
}

func (t *GuestbookTest) Disrupt() {
	_, err := executeCommand("kubectl", "delete", "all", "-l", "example-app=guestbook", "-n", defaultSystemNamespace)
	if err != nil {
		klog.Warningf("kubectl delete finished with error: %v", err)
	}
}

func (k *GuestbookTest) VerifyDown() error {
	return verifyNoWorkloadsWithLabel("example-app=guestbook", defaultSystemNamespace)
}
