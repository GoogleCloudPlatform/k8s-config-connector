/*
Copyright 2020 The Kubernetes Authors.

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

package declarative

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/yaml"

	"sigs.k8s.io/kubebuilder-declarative-pattern/applylib/applyset"
	"sigs.k8s.io/kubebuilder-declarative-pattern/commonclient"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func fakeParent() runtime.Object {
	parent := &unstructured.Unstructured{}
	parent.SetKind("ConfigMap")
	parent.SetAPIVersion("v1")
	parent.SetName("test")
	parent.SetNamespace("default")
	return parent
}

// This test checks gvkString function
func TestGVKString(t *testing.T) {
	testCases := []struct {
		subtest string
		gvk     schema.GroupVersionKind
		want    string
	}{
		{
			subtest: "v1/Pod",
			gvk:     core.SchemeGroupVersion.WithKind("Pod"),
			want:    "v1/Pod",
		},
		{
			subtest: "apps/v1/Deployment",
			gvk:     apps.SchemeGroupVersion.WithKind("Deployment"),
			want:    "apps/v1/Deployment",
		},
	}

	for _, st := range testCases {
		t.Run(st.subtest, func(t *testing.T) {
			if got := gvkString(st.gvk); st.want != got {
				t.Errorf("want:\n%v\ngot:\n%v\n", st.want, got)
			}
		})
	}
}

// This test checks reconcileMetricsFor function & reconcieMetrics.reconcileWith method
func TestReconcileWith(t *testing.T) {
	testCases := []struct {
		subtest    string
		gvks       []schema.GroupVersionKind
		namespaces []string
		names      []string
		want       []string
	}{
		{
			subtest:    "core",
			gvks:       []schema.GroupVersionKind{core.SchemeGroupVersion.WithKind("Pod")},
			namespaces: []string{"ns1"},
			names:      []string{"n1"},
			want: []string{`
			# HELP declarative_reconciler_reconcile_count How many times reconciliation of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_count counter
			declarative_reconciler_reconcile_count {group_version_kind = "v1/Pod", name = "n1", namespace = "ns1"} 2
			`,
			},
		},
		{
			subtest:    "core&app",
			gvks:       []schema.GroupVersionKind{core.SchemeGroupVersion.WithKind("Pod"), apps.SchemeGroupVersion.WithKind("Deployment")},
			namespaces: []string{"ns1", ""},
			names:      []string{"n1", "n2"},
			want: []string{`
			# HELP declarative_reconciler_reconcile_count How many times reconciliation of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_count counter
			declarative_reconciler_reconcile_count {group_version_kind = "v1/Pod", name = "n1", namespace = "ns1"} 2
			`,
				`
			# HELP declarative_reconciler_reconcile_count How many times reconciliation of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_count counter
			declarative_reconciler_reconcile_count {group_version_kind = "apps/v1/Deployment", name = "n2", namespace = ""} 2
			`,
			},
		},
		{
			subtest:    "node - cluster scoped only",
			gvks:       []schema.GroupVersionKind{core.SchemeGroupVersion.WithKind("Node")},
			namespaces: []string{""},
			names:      []string{"n1"},
			want: []string{`
			# HELP declarative_reconciler_reconcile_count How many times reconciliation of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_count counter
			declarative_reconciler_reconcile_count {group_version_kind = "v1/Node", name = "n1", namespace = ""} 2
			`,
			},
		},
	}

	for _, st := range testCases {
		t.Run(st.subtest, func(t *testing.T) {
			for i, gvk := range st.gvks {
				rm := reconcileMetricsFor(gvk)

				rm.reconcileWith(reconcile.Request{NamespacedName: types.NamespacedName{Namespace: st.namespaces[i], Name: st.names[i]}})
				rm.reconcileWith(reconcile.Request{NamespacedName: types.NamespacedName{Namespace: st.namespaces[i], Name: st.names[i]}})

				if err := testutil.CollectAndCompare(rm.reconcileCounterVec.WithLabelValues(gvkString(gvk),
					st.namespaces[i], st.names[i]), strings.NewReader(st.want[i])); err != nil {

					t.Error(err)
				}
			}
		})

		reconcileCount.Reset()
	}
}

// This test checks reconcileMetricsFor function & reconcileMetrics.reconcileFailedWith method
func TestReconcileFailedWith(t *testing.T) {
	testCases := []struct {
		subtest    string
		gvks       []schema.GroupVersionKind
		errs       []error
		namespaces []string
		names      []string
		want       []string
	}{
		{
			subtest:    "core",
			gvks:       []schema.GroupVersionKind{core.SchemeGroupVersion.WithKind("Pod")},
			errs:       []error{errors.New("test")},
			namespaces: []string{"ns1"},
			names:      []string{"n1"},
			want: []string{`
			# HELP declarative_reconciler_reconcile_failure_count How many times reconciliation failure of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_failure_count counter
			declarative_reconciler_reconcile_failure_count {group_version_kind = "v1/Pod", name = "n1", namespace = "ns1"} 2
			`,
			},
		},
		{
			subtest:    "core&app",
			gvks:       []schema.GroupVersionKind{core.SchemeGroupVersion.WithKind("Pod"), apps.SchemeGroupVersion.WithKind("Deployment")},
			errs:       []error{errors.New("test"), errors.New("test")},
			namespaces: []string{"ns1", ""},
			names:      []string{"n1", "n2"},
			want: []string{`
			# HELP declarative_reconciler_reconcile_failure_count How many times reconciliation failure of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_failure_count counter
			declarative_reconciler_reconcile_failure_count {group_version_kind = "v1/Pod", name = "n1", namespace = "ns1"} 2
			`,
				`
			# HELP declarative_reconciler_reconcile_failure_count How many times reconciliation failure of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_failure_count counter
			declarative_reconciler_reconcile_failure_count {group_version_kind = "apps/v1/Deployment", name = "n2", namespace = ""} 2
			`,
			},
		},
		{
			subtest:    "node - cluster scoped only",
			gvks:       []schema.GroupVersionKind{core.SchemeGroupVersion.WithKind("Node")},
			errs:       []error{errors.New("test")},
			namespaces: []string{""},
			names:      []string{"n1"},
			want: []string{`
			# HELP declarative_reconciler_reconcile_failure_count How many times reconciliation failure of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_failure_count counter
			declarative_reconciler_reconcile_failure_count {group_version_kind = "v1/Node", name = "n1", namespace = ""} 2
			`,
			},
		},
		{
			subtest:    "no error",
			gvks:       []schema.GroupVersionKind{core.SchemeGroupVersion.WithKind("Node")},
			errs:       []error{nil},
			namespaces: []string{""},
			names:      []string{"n1"},
			want: []string{`
			# HELP declarative_reconciler_reconcile_failure_count How many times reconciliation failure of K8s objects managed by declarative reconciler occurs
			# TYPE declarative_reconciler_reconcile_failure_count counter
			declarative_reconciler_reconcile_failure_count {group_version_kind = "v1/Node", name = "n1", namespace = ""} 0
			`,
			},
		},
	}

	for _, st := range testCases {
		t.Run(st.subtest, func(t *testing.T) {
			for i, gvk := range st.gvks {
				rm := reconcileMetricsFor(gvk)

				rm.reconcileFailedWith(reconcile.Request{NamespacedName: types.NamespacedName{Namespace: st.namespaces[i], Name: st.names[i]}},
					reconcile.Result{}, st.errs[i])
				rm.reconcileFailedWith(reconcile.Request{NamespacedName: types.NamespacedName{Namespace: st.namespaces[i], Name: st.names[i]}},
					reconcile.Result{}, st.errs[i])

				if err := testutil.CollectAndCompare(rm.reconcileFailureCounterVec.WithLabelValues(gvkString(gvk),
					st.namespaces[i], st.names[i]), strings.NewReader(st.want[i])); err != nil {

					t.Error(err)
				}
			}
		})

		reconcileFailure.Reset()
	}
}

// This test checks *ObjectTracker.addIfNotPresent method
func TestAddIfNotPresent(t *testing.T) {
	k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
	if err != nil {
		t.Fatalf("error building mock kube-apiserver: %v", err)
	}

	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}, "namespaces", meta.RESTScopeRoot)
	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Secret"}, "secrets", meta.RESTScopeNamespace)
	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}, "configmaps", meta.RESTScopeNamespace)

	defer func() {
		if err := k8s.Stop(); err != nil {
			t.Fatalf("error closing mock kube-apiserver: %v", err)
		}
	}()

	addr, err := k8s.StartServing()
	if err != nil {
		t.Errorf("error starting mock kube-apiserver: %v", err)
	}

	klog.Infof("mock kubeapiserver will listen on %v", addr)

	restConfig := &rest.Config{
		Host: addr.String(),
	}

	// Create manager
	mgrOpt := manager.Options{}
	err = commonclient.SetMetricsBindAddress(&mgrOpt, "0")
	if err != nil {
		t.Error(err)
	}
	mgr, err := manager.New(restConfig, mgrOpt)
	if err != nil {
		t.Error(err)
	}

	ctx := context.TODO()
	go func() {
		if err := mgr.GetCache().Start(ctx); err != nil {
			klog.Errorf("error starting cache: %v", err)
		}
	}()

	// Configure globalObjectTracker
	globalObjectTracker.mgr = mgr

	testCases := []struct {
		subtest          string
		metricsDuration  int
		actions          []string
		defaultNamespace string
		objects          [][]string
		wants            []string
	}{
		// It's better to use different kind of K8s object for each test cases
		{
			subtest:          "Create K8s object",
			metricsDuration:  0,
			actions:          []string{"Create"},
			defaultNamespace: "",
			objects: [][]string{
				{
					"kind: Namespace\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: ns1\n",
				},
			},
			wants: []string{
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns1", namespace = ""} 1
				`,
			},
		},
		{
			subtest:          "Update K8s object",
			metricsDuration:  0,
			actions:          []string{"Create", "Update"},
			defaultNamespace: "",
			objects: [][]string{
				{
					"kind: Namespace\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: ns2\n",
					"kind: ConfigMap\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: cm2\n" +
						"   namespace: ns2\n" +
						"data:\n" +
						"  foo1: bar1\n",
				},
				{
					"kind: Namespace\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: ns2\n",
					"kind: ConfigMap\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: cm2\n" +
						"   namespace: ns2\n" +
						"data:\n" +
						"  foo1: bar1\n" +
						"  foo2: bar2\n",
				},
			},
			wants: []string{
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns2", namespace = ""} 1
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/ConfigMap", name = "cm2", namespace = "ns2"} 1
				`,
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns2", namespace = ""} 1
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/ConfigMap", name = "cm2", namespace = "ns2"} 1
				`,
			},
		},
		{
			subtest:          "Delete K8s object",
			metricsDuration:  0,
			actions:          []string{"Create", "Delete"},
			defaultNamespace: "ns3",
			objects: [][]string{
				{
					"kind: Namespace\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: ns3\n",
					"kind: Secret\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: s3\n" +
						"type: Opaque\n" +
						"data:\n" +
						"   name: dGVzdA==\n",
				},
				{
					"kind: Secret\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: s3\n" +
						"type: Opaque\n" +
						"data:\n" +
						"   name: dGVzdA==\n",
				},
			},
			wants: []string{
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns3", namespace = ""} 1
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Secret", name = "s3", namespace = "ns3"} 1
				`,
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns3", namespace = ""} 1
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Secret", name = "s3", namespace = "ns3"} 0
				`,
			},
		},
		{
			subtest:          "Delete metrics after specified duration(duration=2)",
			metricsDuration:  2,
			actions:          []string{"Create", "Delete", "Create", "Create"},
			defaultNamespace: "",
			objects: [][]string{
				{
					"kind: Namespace\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: ns4\n",
					"kind: Secret\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: s4\n" +
						"   namespace: ns4\n" +
						"type: Opaque\n" +
						"data:\n" +
						"   name: dGVzdA==\n",
				},
				{
					"kind: Secret\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: s4\n" +
						"   namespace: ns4\n" +
						"type: Opaque\n" +
						"data:\n" +
						"   name: dGVzdA==\n",
				},
				{
					"kind: Namespace\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: ns4\n",
				},
				{
					"kind: Namespace\n" +
						"apiVersion: v1\n" +
						"metadata:\n" +
						"   name: ns4\n",
				},
			},
			wants: []string{
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns4", namespace = ""} 1
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Secret", name = "s4", namespace = "ns4"} 1
				`,
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns4", namespace = ""} 1
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Secret", name = "s4", namespace = "ns4"} 0
				`,
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns4", namespace = ""} 1
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Secret", name = "s4", namespace = "ns4"} 0
				`,
				`
				# HELP declarative_reconciler_managed_objects_record Track the number of objects in manifest
				# TYPE declarative_reconciler_managed_objects_record gauge
				declarative_reconciler_managed_objects_record {group_version_kind = "v1/Namespace", name = "ns4", namespace = ""} 1
				`,
			},
		},
	}

	for _, st := range testCases {
		t.Run(st.subtest, func(t *testing.T) {
			globalObjectTracker.SetMetricsDuration(st.metricsDuration)

			for i, yobjList := range st.objects {
				var yobj string
				var jobjList = [][]byte{}
				var objList = []*manifest.Object{}

				for i, yitem := range yobjList {
					if i == 0 {
						yobj = yitem
					} else {
						yobj = yobj + "---\n" + yitem
					}
				}

				// YAML to JSON
				for _, yitem := range yobjList {
					jobj, err := yaml.YAMLToJSON([]byte(yitem))
					if err != nil {
						t.Error(err)
					}
					jobjList = append(jobjList, jobj)
				}

				// JSON to manifest.Object
				for _, jobj := range jobjList {
					mobj, err := manifest.ParseJSONToObject(jobj)
					if err != nil {
						t.Error(err)
					}
					objList = append(objList, mobj)
				}

				// Run addIfNotPresent
				err = globalObjectTracker.addIfNotPresent(objList, st.defaultNamespace)
				if err != nil {
					t.Error(err)
				}

				// Set up kubectl command
				if st.actions[i] == "Create" || st.actions[i] == "Update" {
					var options applier.ApplierOptions
					options.Namespace = st.defaultNamespace
					options.Objects = objList
					options.RESTMapper = mgr.GetRESTMapper()
					options.RESTConfig = restConfig
					parent := fakeParent()
					options.Client = fake.NewClientBuilder().WithRuntimeObjects(parent).Build()
					gvk := parent.GetObjectKind().GroupVersionKind()
					restmapping, err := options.RESTMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
					if err != nil {
						t.Fatalf("failed to get restmapping for parent: %v", err)
					}
					options.ParentRef = applyset.NewParentRef(parent, "kdp-test", "default", restmapping)
					applier := applier.DefaultApplier
					if err := applier.Apply(ctx, options); err != nil {
						t.Fatalf("failed to apply objects: %v", err)
					}
				} else if st.actions[i] == "Delete" {
					objects, err := manifest.ParseObjects(ctx, yobj)
					if err != nil {
						t.Fatalf("error parsing manifest: %v", err)
					}

					if err := deleteObjects(ctx, restConfig, mgr.GetRESTMapper(), st.defaultNamespace, objects); err != nil {
						t.Fatalf("error deleting objects: %v", err)
					}
				} else {
					t.Fatalf("unknown action %q", st.actions[i])
				}

				// Wait for reflector sees K8s object change in K8s API server & adds it to DeltaFIFO
				// then controller pops it and eventhandler updates metrics
				// If we omit it, there is a chance call of testutil.CollectAndCompare is too fast & fails.
				if !mgr.GetCache().WaitForCacheSync(ctx) {
					t.Errorf("could not sync caches with WaitForCacheSync")
				}
				time.Sleep(time.Second * 2)

				// Check for metrics
				err = testutil.CollectAndCompare(managedObjectsRecord, strings.NewReader(st.wants[i]))
				if err != nil {
					t.Logf("No. of action in subtest: %v\n", i)
					t.Error(err)
				}
			}
		})
		managedObjectsRecord.Reset()
	}
}

// deleteObjects is a simple helper that deletes the specified objects
func deleteObjects(ctx context.Context, restConfig *rest.Config, restMapper meta.RESTMapper, forceNamespace string, objects *manifest.Objects) error {
	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return fmt.Errorf("error building dynamic client: %w", err)
	}

	for _, obj := range objects.Items {
		gvk := obj.GroupVersionKind()
		restMapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			return fmt.Errorf("error looking up RESTMapping for %v: %w", gvk, err)
		}

		var opt metav1.DeleteOptions
		switch restMapping.Scope {
		case meta.RESTScopeNamespace:
			ns := forceNamespace
			if ns == "" {
				ns = obj.GetNamespace()
			}
			if ns == "" {
				return fmt.Errorf("no namespace for %v", gvk)
			}
			if err := dynamicClient.Resource(restMapping.Resource).Namespace(ns).Delete(ctx, obj.GetName(), opt); err != nil {
				return fmt.Errorf("error deleting object: %w", err)
			}

		case meta.RESTScopeRoot:
			if err := dynamicClient.Resource(restMapping.Resource).Delete(ctx, obj.GetName(), opt); err != nil {
				return fmt.Errorf("error deleting object: %w", err)
			}
		default:
			return fmt.Errorf("unknown rest mapping scope %v", restMapping.Scope)
		}

	}
	return nil
}
