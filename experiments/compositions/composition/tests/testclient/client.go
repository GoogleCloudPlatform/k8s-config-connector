// Copyright 2024 Google LLC
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

// Package testclient defines an MCS test client with basic checks.
package testclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	compositionv1 "google.com/composition/api/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kustomize/kstatus/status"
)

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(compositionv1.AddToScheme(scheme))
}

// Client - test client of MCS target cluster
type Client struct {
	T   *testing.T
	Ctx context.Context
	cs  *kubernetes.Clientset
	client.Client
	name     string
	testName string
}

func New(ctx context.Context, t *testing.T, config *rest.Config, name, testName string) *Client {
	c, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		t.FailNow()
	}
	return &Client{
		T:        t,
		Ctx:      ctx,
		cs:       clientset,
		Client:   c,
		name:     name,
		testName: testName,
	}
}

func (c *Client) String() string {
	return c.name
}

func (c *Client) WriteNamespacePodLogs(namespace string) {
	// Get pods in namespace
	pods, err := c.cs.CoreV1().Pods(namespace).List(c.Ctx, metav1.ListOptions{})
	if err != nil {
		c.T.Errorf("unexpected error: %v", err)
		c.T.FailNow()
	}

	for _, pod := range pods.Items {
		// Get containers in pod
		for _, container := range append(pod.Spec.InitContainers, pod.Spec.Containers...) {
			// Write container logs
			c.WriteContainerLogs(namespace, pod.Name, container.Name)
		}
	}
}

func (c *Client) getFilePath(namespace, name, container string) string {
	path := filepath.Join("test-logs", c.testName, namespace)
	os.MkdirAll(path, os.ModePerm)
	filename := name + "." + container + ".txt"
	return filepath.Join("test-logs", c.testName, namespace, filename)
}

func (c *Client) ClearOldLogs() {
	path := filepath.Join("test-logs", c.testName)
	os.RemoveAll(path)
}

func (c *Client) WriteContainerLogs(namespace, name, container string) {
	podLogOpts := corev1.PodLogOptions{Container: container}

	req := c.cs.CoreV1().Pods(namespace).GetLogs(name, &podLogOpts)

	podLogs, err := req.Stream(c.Ctx)
	defer podLogs.Close()
	if err != nil {
		c.T.Errorf("Error getting log stream for pod: %s.%s %v", name, container, err)
		c.T.FailNow()
	}

	filename := c.getFilePath(namespace, name, container)
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		c.T.Errorf("Error creating log file: %s %v", filename, err)
		c.T.FailNow()
	}
	_, err = io.Copy(file, podLogs)
	if err != nil {
		c.T.Errorf("Failed to copy pod(%s.%s) logs to file: %s,  %v", name, container, filename, err)
		c.T.FailNow()
	}
}

// MustCreate - create object
func (c *Client) MustCreate(u *unstructured.Unstructured) {
	c.T.Helper()

	id := ExtractGVKNN(u)
	c.T.Logf("Creating object %q on cluster %q", id, c)
	err := c.Create(c.Ctx, u)
	if err != nil {
		if !apierrors.IsAlreadyExists(err) {
			c.T.Errorf("failed to create (absent) %q: %s", id, err)
			c.T.FailNow()
		}
		c.T.Logf("WARN object exists. Reusing. %s", err)
	}
}

// Read - read object with obj's gvknn
func (c *Client) Read(obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	c.T.Helper()

	// new encompassing struct that is printable - has String method
	id := ExtractGVKNN(obj)
	readObj := &unstructured.Unstructured{}
	readObj.SetGroupVersionKind(id.GroupVersionKind)
	err := c.Get(c.Ctx, types.NamespacedName{
		Namespace: id.Namespace,
		Name:      id.Name,
	}, readObj)
	if err != nil {
		err = fmt.Errorf("cannot read %q on %q: %w", id, c, err)
	}
	return readObj, err
}

// MustDelete - Delete object if exists
func (c *Client) MustDelete(u *unstructured.Unstructured) {
	c.T.Helper()

	id := ExtractGVKNN(u)
	c.T.Logf("Deleting object %q on cluster %q", id, c)
	err := c.Delete(c.Ctx, u)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			c.T.Errorf("failed to delete (present) %q: %s", id, err)
			c.T.FailNow()
		}
		c.T.Logf("WARN %q not found on cluster %q: %s", id, c, err)
	}
}

// MustExist - waits for all objs to exist
func (c *Client) MustExist(objs []*unstructured.Unstructured, timeout time.Duration) {
	c.T.Helper()
	c.T.Logf("Checking existence of objects %q on %q", ExtractGVKNNs(objs), c)

	toFind := objs
	retryFrequency := getFrequency(c.T, opDuration, timeout)
	err := retry.OnError(retryFrequency, func(err error) bool {
		return apierrors.IsNotFound(err)
	}, func() (err error) {
		exists := func(obj *unstructured.Unstructured) error {
			_, err := c.Read(obj)
			return err
		}
		toFind, err = c.recordFailed(exists, toFind)
		return err
	})
	if err != nil {
		c.T.Errorf("objects absent on %q", c)
		c.T.FailNow()
	}
}

// MustMatchSpec - verify specs of objects match objs
func (c *Client) MustMatchSpec(objs []*unstructured.Unstructured, timeout time.Duration) {
	c.T.Helper()

	toMatch := ExtractGVKNNs(objs)
	c.T.Logf("Matching specs of objects %q on %q", toMatch, c)

	unmatched := objs
	errMismatch := fmt.Errorf("unexpected spec")
	retryFrequency := getFrequency(c.T, opDuration, timeout)
	err := retry.OnError(retryFrequency, func(err error) bool {
		return apierrors.IsNotFound(err) || errors.Is(err, errMismatch)
	}, func() (err error) {
		match := func(obj *unstructured.Unstructured) error {
			spec := obj.UnstructuredContent()["spec"]
			read, err := c.Read(obj)
			if err != nil {
				return err
			}
			readSpec := read.UnstructuredContent()["spec"]
			if !reflect.DeepEqual(spec, readSpec) {
				diff := cmp.Diff(spec, readSpec)
				// Cannot use %q on spec strings.
				// The double quotes change the rendering of many characters.
				return fmt.Errorf(`%w: spec for object %q on %q is: %s
and not equal to: %s,
with diff %s`, errMismatch, ExtractGVKNN(obj), c, readSpec, spec, diff)
			}
			return nil
		}
		unmatched, err = c.recordFailed(match, unmatched)
		return err
	})
	if err != nil {
		c.T.Errorf("unexpected error: %v", err)
		c.T.FailNow()
	}
}

// MustBeReady - waits for all objs to be available
func (c *Client) MustBeReady(objs []*unstructured.Unstructured, timeout time.Duration) {
	c.T.Helper()
	c.T.Logf("Checking readiness of objects %q on %q", ExtractGVKNNs(objs), c)

	notReady := objs
	errNotReady := fmt.Errorf("object is not available")
	retryFrequency := getFrequency(c.T, opDuration, timeout)
	err := retry.OnError(retryFrequency, func(err error) bool {
		return apierrors.IsNotFound(err) || errors.Is(err, errNotReady)
	}, func() (err error) {
		checkReady := func(obj *unstructured.Unstructured) error {
			read, err := c.Read(obj)
			if err != nil {
				return err
			}
			result, err := status.Compute(read)
			if err != nil {
				return err
			}
			if result.Status != status.CurrentStatus {
				return fmt.Errorf("missing object %q: %w", ExtractGVKNN(read), errNotReady)
			}
			return nil
		}
		notReady, err = c.recordFailed(checkReady, notReady)
		return err
	})
	if err != nil {
		c.T.Errorf("objects unavailable on %q", c)
		c.T.FailNow()
	}
}

// MustJSONPatch - applies JSON patch to obj
func (c *Client) MustJSONPatch(obj *unstructured.Unstructured, patch map[string]any) {
	c.T.Helper()
	c.T.Logf("Applying patch %q to %q on %q", patch, ExtractGVKNN(obj), c)

	serialPatch, err := json.Marshal([]map[string]any{
		patch,
	})
	if err != nil {
		c.T.Errorf("unexpected error: %v", err)
		c.T.FailNow()
	}
	formalPatch := client.RawPatch(types.JSONPatchType, serialPatch)
	err = c.Patch(c.Ctx, obj, formalPatch)
	if err != nil {
		c.T.Errorf("unexpected error: %v", err)
		c.T.FailNow()
	}
}

// MustNotExist - checks none of objs exists
func (c *Client) MustNotExist(objs []*unstructured.Unstructured, timeout time.Duration) {
	c.T.Helper()
	c.T.Logf("Checking absence of objects %q on %q", ExtractGVKNNs(objs), c)

	existing := objs
	retryFrequency := getFrequency(c.T, opDuration, timeout)
	err := retry.OnError(retryFrequency, func(err error) bool {
		return apierrors.IsAlreadyExists(err)
	}, func() (err error) {
		doesNotExist := func(obj *unstructured.Unstructured) error {
			return c.checkNotExist(obj)
		}
		existing, err = c.recordFailed(doesNotExist, existing)
		return err
	})
	if err != nil {
		c.T.Errorf("objects should not exist on %q", c)
		c.T.FailNow()
	}
}

// recordFailed - returns all objects on which op failed and an error
// encompassing all corresponding errors
func (c *Client) recordFailed(op func(*unstructured.Unstructured) error, objs []*unstructured.Unstructured) ([]*unstructured.Unstructured, error) {
	c.T.Helper()

	var failed []*unstructured.Unstructured
	var errs []error
	for _, obj := range objs {
		if err := op(obj); err != nil {
			failed = append(failed, obj)
			errs = append(errs, err)
		}
	}
	return failed, errors.Join(errs...)
}

// checkNotExist - expects that obj does not exist
func (c *Client) checkNotExist(obj *unstructured.Unstructured) error {
	c.T.Helper()

	_, err := c.Read(obj)
	if err == nil {
		err = apierrors.NewAlreadyExists(schema.GroupResource{
			Group:    obj.GroupVersionKind().Group,
			Resource: obj.GroupVersionKind().Kind,
		}, obj.GetName())
	} else if apierrors.IsNotFound(err) {
		err = nil
	}
	return err
}
