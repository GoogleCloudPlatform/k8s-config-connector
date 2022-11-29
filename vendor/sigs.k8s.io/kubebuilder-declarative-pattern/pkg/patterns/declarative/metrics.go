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
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const (
	emptyNamespace = "_Empty_"
	clusterScoped  = "_Cluster_"
)

const (
	Declarative = "declarative_reconciler"
)

const (
	ReconcileCount   = "reconcile_count"
	ReconcileFailure = "reconcile_failure_count"

	ManagedObjectsRecord = "managed_objects_record"
)

var metricsRegisterOnce *sync.Once = &sync.Once{}

var (
	reconcileCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: Declarative,
		Name:      ReconcileCount,
		Help:      "How many times reconciliation of K8s objects managed by declarative reconciler occurs",
	}, []string{"group_version_kind", "namespace", "name"})

	reconcileFailure = prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: Declarative,
		Name:      ReconcileFailure,
		Help:      "How many times reconciliation failure of K8s objects managed by declarative reconciler occurs",
	}, []string{"group_version_kind", "namespace", "name"})

	managedObjectsRecord = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: Declarative,
		Name:      ManagedObjectsRecord,
		Help:      "Track the number of objects in manifest",
	}, []string{"group_version_kind", "namespace", "name"})
)

var metricsList = []prometheus.Collector{reconcileCount, reconcileFailure, managedObjectsRecord}

func gvkString(gvk schema.GroupVersionKind) string {
	if len(gvk.Group) == 0 && gvk.Version == "v1" {
		return gvk.Version + "/" + gvk.Kind
	}

	return gvk.Group + "/" + gvk.Version + "/" + gvk.Kind
}

func namespaced(mgr manager.Manager, gvk schema.GroupVersionKind) (bool, error) {
	mapping, err := mgr.GetRESTMapper().RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return false, err
	}

	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		return true, nil
	} else {
		return false, nil
	}
}

type reconcileMetrics struct {
	groupVersionKind           string
	reconcileCounterVec        *prometheus.CounterVec
	reconcileFailureCounterVec *prometheus.CounterVec
}

func reconcileMetricsFor(gvk schema.GroupVersionKind) reconcileMetrics {
	return reconcileMetrics{
		groupVersionKind:    gvkString(gvk),
		reconcileCounterVec: reconcileCount, reconcileFailureCounterVec: reconcileFailure,
	}
}

func (rm *reconcileMetrics) reconcileWith(req reconcile.Request) {
	rm.reconcileCounterVec.WithLabelValues(rm.groupVersionKind, req.Namespace, req.Name).Inc()
}

func (rm *reconcileMetrics) reconcileFailedWith(req reconcile.Request, _ reconcile.Result, err error) {
	if err != nil {
		rm.reconcileFailureCounterVec.WithLabelValues(rm.groupVersionKind, req.Namespace, req.Name).Inc()
	}
}

type objectRecorder struct {
	groupVersionKind string
	gaugeVec         *prometheus.GaugeVec
}

func objectRecorderFor(gvk schema.GroupVersionKind) objectRecorder {
	return objectRecorder{groupVersionKind: gvkString(gvk), gaugeVec: managedObjectsRecord}
}

func (or objectRecorder) Set(namespace, name string, num float64) {
	if namespace == emptyNamespace || namespace == clusterScoped {
		or.gaugeVec.WithLabelValues(or.groupVersionKind, "", name).Set(num)
	} else {
		or.gaugeVec.WithLabelValues(or.groupVersionKind, namespace, name).Set(num)
	}
}

func (or objectRecorder) Unset(namespace, name string) {
	if namespace == emptyNamespace || namespace == clusterScoped {
		or.gaugeVec.DeleteLabelValues(or.groupVersionKind, "", name)
	} else {
		or.gaugeVec.DeleteLabelValues(or.groupVersionKind, namespace, name)
	}
}

var globalObjectTracker *ObjectTracker = NewObjectTracker()

// GetMetricsDuration function returns current metricsDuration
// of package-scoped internal variable of type *ObjectTracker.
// It is safe to call this function from multiple go routines
// with another go routines calling SetMetricsDuration function.
func GetMetricsDuration() int {
	return globalObjectTracker.GetMetricsDuration()
}

// SetMetricsDuration function sets metricsDuration of package
// scoped internal variable of type *ObjectTracker.
// It is safe to call this function from multiple go routines
// with another go routines calling GetMetricsDuration function.
func SetMetricsDuration(metricsDuration int) {
	globalObjectTracker.SetMetricsDuration(metricsDuration)
}

// Type ObjectTracker manages metricsDuration of k8s objects
// managed by controller(s)
type ObjectTracker struct {
	mu              sync.Mutex
	mgr             manager.Manager
	metricsDuration int
	trackedGVK      map[schema.GroupVersionKind]*gvkTracker
}

// GetMetricsDuration method returns current metricsDuration.
// It is safe to call this function from multiple go routines
// with another go routines calling SetMetricsDuration method.
func (ot *ObjectTracker) GetMetricsDuration() int {
	ot.mu.Lock()
	defer ot.mu.Unlock()

	return ot.metricsDuration
}

// SetMetricsDuration method sets metricsDuration.
// It is safe to call this function from multiple go routines
// with another go routines calling GetMetricsDuration method..
func (ot *ObjectTracker) SetMetricsDuration(metricsDuration int) {
	ot.mu.Lock()
	defer ot.mu.Unlock()

	ot.metricsDuration = metricsDuration
}

func (ot *ObjectTracker) setMetricsDurationInternal(i int) {
	ot.mu.Lock()
	defer ot.mu.Unlock()

	if i > 0 && i > ot.metricsDuration {
		ot.metricsDuration = i
	}
}

// Call this before kubectl.Apply
func (ot *ObjectTracker) addIfNotPresent(objects []*manifest.Object, defaultNamespace string) errors.Aggregate {
	ot.mu.Lock()
	defer ot.mu.Unlock()
	defer ot.deleteMetricsIfNeeded()
	defer ot.inc()

	var errs []error

	for _, object := range objects {
		gvk, ns, name := object.GroupVersionKind(),
			object.UnstructuredObject().GetNamespace(),
			object.UnstructuredObject().GetName()

		// Check default namespace
		if defaultNamespace != "" {
			ns = defaultNamespace
		}

		namespaced, err := namespaced(ot.mgr, gvk)
		if err != nil {
			errs = append(errs, newNoRESTMapperErr(err, gvk))
			continue
		}

		if namespaced {
			if len(ns) == 0 {
				ns = emptyNamespace
				errs = append(errs, newEmptyNamespaceErr(gvk, object.UnstructuredObject().GetName()))
			}
		} else {
			ns = clusterScoped
		}

		if gvkt, ok := ot.trackedGVK[gvk]; ok {
			if !gvkt.resetIfItHas(ns, name) {
				gvkt.insert(ns, name)
			}
			continue
		}

		ot.trackedGVK[gvk] = newGVKTracker(ot.mgr, object.UnstructuredObject(), namespaced)
		ot.trackedGVK[gvk].insert(ns, name)

		// addIfNotPresent is called at Reconciler.reconcileExists,
		// so Controller & Manager is already running
		ctx := context.TODO()
		if err := ot.trackedGVK[gvk].start(ctx); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.NewAggregate(errs)
}

// Call after ot.mu.Lock()
func (ot *ObjectTracker) inc() {
	for _, gvkt := range ot.trackedGVK {
		gvkt.list.inc()
	}
}

// Call after ot.mu.Lock()
func (ot *ObjectTracker) deleteMetricsIfNeeded() {
	if ot.metricsDuration > 0 {
		for _, gvkt := range ot.trackedGVK {
			gvkt.deleteMetricsIfNeeded(ot.metricsDuration)
		}
	}
}

// NewObjectTracker returns new instance of *ObjectTracker
func NewObjectTracker() *ObjectTracker {
	ot := &ObjectTracker{}
	ot.mgr = nil
	ot.trackedGVK = make(map[schema.GroupVersionKind]*gvkTracker)

	return ot
}

type noRESTMapperErr struct {
	gvk        schema.GroupVersionKind
	wrappedErr error
}

func (noREST noRESTMapperErr) Error() string {
	return noREST.wrappedErr.Error()
}

func (noRESTMapperErr) Is(target error) bool {
	_, ok := target.(noRESTMapperErr)
	return ok
}

func newNoRESTMapperErr(err error, gvk schema.GroupVersionKind) noRESTMapperErr {
	return noRESTMapperErr{gvk: gvk, wrappedErr: err}
}

type emptyNamespaceErr struct {
	gvk  schema.GroupVersionKind
	name string
}

func (emptyNamespaceErr) Error() string {
	return "Scoped object, but no namespace specified"
}

func (emptyNamespaceErr) Is(target error) bool {
	_, ok := target.(emptyNamespaceErr)
	return ok
}

func newEmptyNamespaceErr(gvk schema.GroupVersionKind, name string) emptyNamespaceErr {
	return emptyNamespaceErr{gvk: gvk, name: name}
}

type gvkTracker struct {
	list         *items
	src          source.Source
	eventHandler handler.EventHandler
	predicate    predicate.Predicate
	recorder     objectRecorder
}

func (gvkt *gvkTracker) insert(namespace string, names ...string) {
	gvkt.list.insert(namespace, names...)
}

func (gvkt *gvkTracker) resetIfItHas(namespace, name string) bool {
	return gvkt.list.resetIfItHas(namespace, name)
}

func (gvkt *gvkTracker) deleteMetricsIfNeeded(metricsDuration int) {
	if deleted, pairs := gvkt.list.deleteIfNeeded(metricsDuration); deleted {
		for ns := range pairs {
			for _, n := range pairs[ns] {
				gvkt.recorder.Unset(ns, n)
			}
		}
	}
}

func (gvkt *gvkTracker) start(ctx context.Context) error {
	return gvkt.src.Start(ctx, gvkt.eventHandler, dummyQueue{}, gvkt.predicate)
}

func newGVKTracker(mgr manager.Manager, obj *unstructured.Unstructured, namespaced bool) (gvkt *gvkTracker) {
	gvkt = &gvkTracker{}
	gvkt.list = newItems()
	gvkt.recorder = objectRecorderFor(obj.GroupVersionKind())
	gvkt.src = source.NewKindWithCache(obj, mgr.GetCache())
	gvkt.predicate = predicate.Funcs{}
	gvkt.eventHandler = recordTrigger{gvkt.list, namespaced, gvkt.recorder}

	return
}

var _ workqueue.RateLimitingInterface = dummyQueue{}

type dummyQueue struct{}

func (dummyQueue) Add(item interface{})                              {}
func (dummyQueue) Len() int                                          { return 0 }
func (dummyQueue) Get() (item interface{}, shutdown bool)            { return struct{}{}, false }
func (dummyQueue) Done(item interface{})                             {}
func (dummyQueue) ShutDown()                                         {}
func (dummyQueue) ShuttingDown() bool                                { return false }
func (dummyQueue) AddAfter(item interface{}, duration time.Duration) {}
func (dummyQueue) AddRateLimited(item interface{})                   {}
func (dummyQueue) Forget(item interface{})                           {}
func (dummyQueue) NumRequeues(item interface{}) int                  { return 0 }
func (dummyQueue) ShutDownWithDrain()                                {}

// Namespace & list of Name pairs
type nsnPairs map[string][]string

// ids holds namespace/name pairs and is indexed as ids[namespace][name]
type items struct {
	mu  sync.Mutex
	ids map[string]record
}

func (i *items) insert(namespace string, names ...string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, ok := i.ids[namespace]; !ok {
		i.ids[namespace] = record{}
	}
	i.ids[namespace].Insert(names...)
}

func (i *items) has(namespace, name string) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, ok := i.ids[namespace]
	return ok && i.ids[namespace].Has(name)
}

// Only used by another *items methods
func (i *items) internalHas(namespace, name string) bool {
	_, ok := i.ids[namespace]
	return ok && i.ids[namespace].Has(name)
}

func (i *items) hasPairs(nsnp nsnPairs) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	for ns, nlist := range nsnp {
		for _, n := range nlist {
			if !i.internalHas(ns, n) {
				return false
			}
		}
	}
	return true
}

func (i *items) inc() {
	i.mu.Lock()
	defer i.mu.Unlock()

	for ns := range i.ids {
		i.ids[ns].Inc()
	}
}

func (i *items) resetIfItHas(namespace, name string) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, has := i.ids[namespace]; has {
		return i.ids[namespace].ResetIfItHas(name)
	}
	return false
}

func (i *items) markDeletedIfItHas(namespace, name string) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, has := i.ids[namespace]; has {
		return i.ids[namespace].MarkDeletedIfItHas(name)
	}
	return false
}

func (i *items) deleteIfNeeded(metricsDuration int) (b bool, p nsnPairs) {
	i.mu.Lock()
	defer i.mu.Unlock()

	p = make(map[string][]string)
	for ns := range i.ids {
		for n := range i.ids[ns] {
			if i.ids[ns].DeleteIfNeeded(n, metricsDuration) {
				b = true
				if _, ok := p[ns]; !ok {
					p[ns] = []string{}
				}
				p[ns] = append(p[ns], n)
			}
		}
		if len(i.ids[ns]) == 0 {
			delete(i.ids, ns)
		}
	}

	return b, p
}

func newItems() *items {
	it := &items{}
	it.ids = make(map[string]record)
	return it
}

type record map[string]*counter

type counter struct {
	count   int
	deleted bool
}

func (r record) Insert(names ...string) {
	for _, name := range names {
		r[name] = &counter{}
	}
}

func (r record) Has(name string) bool {
	_, has := r[name]
	return has
}

func (r record) Inc() {
	for i := range r {
		if r[i].deleted {
			r[i].count++
		}
	}
}

func (r record) ResetIfItHas(name string) bool {
	var has bool
	if _, has = r[name]; has && r[name].deleted {
		r[name] = &counter{}
	}
	return has
}

func (r record) MarkDeletedIfItHas(name string) bool {
	if _, has := r[name]; has {
		r[name].deleted = true
		return true
	}
	return false
}

func (r record) DeleteIfNeeded(name string, limit int) bool {
	if _, has := r[name]; has {
		if r[name].deleted && r[name].count >= limit {
			delete(r, name)
			return true
		}
	}
	return false
}

var _ handler.EventHandler = recordTrigger{}

type recordTrigger struct {
	list       *items
	namespaced bool
	recorder   objectRecorder
}

func (rt recordTrigger) Create(ev event.CreateEvent, _ workqueue.RateLimitingInterface) {
	ns, name := ev.Object.GetNamespace(), ev.Object.GetName()

	if rt.namespaced {
		if len(ns) == 0 {
			ns = emptyNamespace
		}
	} else {
		ns = clusterScoped
	}

	if rt.list.has(ns, name) {
		rt.recorder.Set(ns, name, float64(1))
	}
}

func (rt recordTrigger) Update(ev event.UpdateEvent, _ workqueue.RateLimitingInterface) {
	var nsnp nsnPairs = make(map[string][]string)
	ons, oname := ev.ObjectOld.GetNamespace(), ev.ObjectOld.GetName()
	nns, nname := ev.ObjectNew.GetNamespace(), ev.ObjectNew.GetName()

	if rt.namespaced {
		if len(ons) == 0 {
			ons = emptyNamespace
		}
		if len(nns) == 0 {
			nns = emptyNamespace
		}
	} else {
		ons = clusterScoped
		nns = clusterScoped
	}

	nsnp[ons] = append(nsnp[ons], oname)
	nsnp[nns] = append(nsnp[nns], nname)
	// Because objectTracker.addIfNotPresent is called before
	// kubectl.Apply in Reconciler.reconcileExists,
	// all pair of oldObj & newObj in DeltaFIFO is in rt.list
	if !rt.list.hasPairs(nsnp) {
		return
	}

	if ons == nns && oname == nname {
		rt.recorder.Set(ons, oname, float64(1))
		return
	}
	rt.recorder.Set(ons, oname, float64(0))
	rt.recorder.Set(nns, nname, float64(1))
}

func (rt recordTrigger) Delete(ev event.DeleteEvent, _ workqueue.RateLimitingInterface) {
	ns, name := ev.Object.GetNamespace(), ev.Object.GetName()

	if rt.namespaced {
		if len(ns) == 0 {
			ns = emptyNamespace
		}
	} else {
		ns = clusterScoped
	}

	if rt.list.markDeletedIfItHas(ns, name) {
		rt.recorder.Set(ns, name, float64(0))
	}
}

func (rt recordTrigger) Generic(ev event.GenericEvent, _ workqueue.RateLimitingInterface) {}
