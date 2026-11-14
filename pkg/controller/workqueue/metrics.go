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

package workqueue

import (
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/util/workqueue"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

type queueMetrics[T comparable] interface {
	add(item T)
	get(item T)
	done(item T)
	updateUnfinishedWork()
	retry()
}

func newQueueMetrics[T comparable](mp workqueue.MetricsProvider, name string, now func() time.Time) queueMetrics[T] {
	if len(name) == 0 {
		return noMetrics[T]{}
	}
	if mp == nil {
		mp = getGlobalMetricsProvider()
	}
	if now == nil {
		now = time.Now
	}
	return &defaultQueueMetrics[T]{
		now:                     now,
		depth:                   mp.NewDepthMetric(name),
		adds:                    mp.NewAddsMetric(name),
		latency:                 mp.NewLatencyMetric(name),
		workDuration:            mp.NewWorkDurationMetric(name),
		unfinishedWorkSeconds:   mp.NewUnfinishedWorkSecondsMetric(name),
		longestRunningProcessor: mp.NewLongestRunningProcessorSecondsMetric(name),
		addTimes:                map[T]time.Time{},
		processingStartTimes:    map[T]time.Time{},
		retries:                 mp.NewRetriesMetric(name),
	}
}

type defaultQueueMetrics[T comparable] struct {
	now func() time.Time

	depth                   workqueue.GaugeMetric
	adds                    workqueue.CounterMetric
	latency                 workqueue.HistogramMetric
	workDuration            workqueue.HistogramMetric
	unfinishedWorkSeconds   workqueue.SettableGaugeMetric
	longestRunningProcessor workqueue.SettableGaugeMetric
	retries                 workqueue.CounterMetric

	mapLock              sync.RWMutex
	addTimes             map[T]time.Time
	processingStartTimes map[T]time.Time
}

func (m *defaultQueueMetrics[T]) add(item T) {
	if m == nil {
		return
	}
	if m.adds != nil {
		m.adds.Inc()
	}
	if m.depth != nil {
		m.depth.Inc()
	}

	m.mapLock.Lock()
	defer m.mapLock.Unlock()

	if _, exists := m.addTimes[item]; !exists {
		m.addTimes[item] = m.now()
	}
}

func (m *defaultQueueMetrics[T]) get(item T) {
	if m == nil {
		return
	}
	if m.depth != nil {
		m.depth.Dec()
	}

	m.mapLock.Lock()
	defer m.mapLock.Unlock()

	m.processingStartTimes[item] = m.now()
	if startTime, exists := m.addTimes[item]; exists {
		if m.latency != nil {
			m.latency.Observe(m.sinceInSeconds(startTime))
		}
		delete(m.addTimes, item)
	}
}

func (m *defaultQueueMetrics[T]) done(item T) {
	if m == nil {
		return
	}

	m.mapLock.Lock()
	defer m.mapLock.Unlock()

	if startTime, exists := m.processingStartTimes[item]; exists {
		if m.workDuration != nil {
			m.workDuration.Observe(m.sinceInSeconds(startTime))
		}
		delete(m.processingStartTimes, item)
	}
}

func (m *defaultQueueMetrics[T]) updateUnfinishedWork() {
	if m == nil {
		return
	}
	m.mapLock.RLock()
	defer m.mapLock.RUnlock()

	var total float64
	var oldest float64
	for _, t := range m.processingStartTimes {
		age := m.sinceInSeconds(t)
		total += age
		if age > oldest {
			oldest = age
		}
	}
	if m.unfinishedWorkSeconds != nil {
		m.unfinishedWorkSeconds.Set(total)
	}
	if m.longestRunningProcessor != nil {
		m.longestRunningProcessor.Set(oldest)
	}
}

func (m *defaultQueueMetrics[T]) sinceInSeconds(start time.Time) float64 {
	return m.now().Sub(start).Seconds()
}

func (m *defaultQueueMetrics[T]) retry() {
	if m == nil {
		return
	}
	if m.retries != nil {
		m.retries.Inc()
	}
}

type noMetrics[T any] struct{}

func (noMetrics[T]) add(item T)            {}
func (noMetrics[T]) get(item T)            {}
func (noMetrics[T]) done(item T)           {}
func (noMetrics[T]) updateUnfinishedWork() {}
func (noMetrics[T]) retry()                {}

type defaultMetricsProvider struct {
	depth                   *prometheus.GaugeVec
	adds                    *prometheus.CounterVec
	latency                 *prometheus.HistogramVec
	workDuration            *prometheus.HistogramVec
	unfinished              *prometheus.GaugeVec
	longestRunningProcessor *prometheus.GaugeVec
	retries                 *prometheus.CounterVec
}

var (
	metricsInitOnce       sync.Once
	globalMetricsProvider *defaultMetricsProvider
)

func getGlobalMetricsProvider() *defaultMetricsProvider {
	metricsInitOnce.Do(func() {
		globalMetricsProvider = newDefaultMetricsProvider()
	})
	return globalMetricsProvider
}

func newDefaultMetricsProvider() *defaultMetricsProvider {
	d := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "workqueue",
		Name:      "depth",
		Help:      "Current depth of workqueue",
	}, []string{"name", "controller"})
	if err := crmetrics.Registry.Register(d); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			d = are.ExistingCollector.(*prometheus.GaugeVec)
		}
	}

	a := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "workqueue",
		Name:      "adds_total",
		Help:      "Total number of adds handled by workqueue",
	}, []string{"name", "controller"})
	if err := crmetrics.Registry.Register(a); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			a = are.ExistingCollector.(*prometheus.CounterVec)
		}
	}

	l := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: "workqueue",
		Name:      "queue_duration_seconds",
		Help:      "How long in seconds an item stays in workqueue before being requested",
		Buckets:   prometheus.ExponentialBuckets(10e-9, 10, 12),
	}, []string{"name", "controller"})
	if err := crmetrics.Registry.Register(l); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			l = are.ExistingCollector.(*prometheus.HistogramVec)
		}
	}

	w := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: "workqueue",
		Name:      "work_duration_seconds",
		Help:      "How long in seconds processing an item from workqueue takes.",
		Buckets:   prometheus.ExponentialBuckets(10e-9, 10, 12),
	}, []string{"name", "controller"})
	if err := crmetrics.Registry.Register(w); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			w = are.ExistingCollector.(*prometheus.HistogramVec)
		}
	}

	u := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "workqueue",
		Name:      "unfinished_work_seconds",
		Help:      "How many seconds of work has been done that is in progress and hasn't been observed by work_duration.",
	}, []string{"name", "controller"})
	if err := crmetrics.Registry.Register(u); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			u = are.ExistingCollector.(*prometheus.GaugeVec)
		}
	}

	lrp := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "workqueue",
		Name:      "longest_running_processor_seconds",
		Help:      "How many seconds has the longest running processor for workqueue been running.",
	}, []string{"name", "controller"})
	if err := crmetrics.Registry.Register(lrp); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			lrp = are.ExistingCollector.(*prometheus.GaugeVec)
		}
	}

	r := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "workqueue",
		Name:      "retries_total",
		Help:      "Total number of retries handled by workqueue",
	}, []string{"name", "controller"})
	if err := crmetrics.Registry.Register(r); err != nil {
		var are prometheus.AlreadyRegisteredError
		if errors.As(err, &are) {
			r = are.ExistingCollector.(*prometheus.CounterVec)
		}
	}

	return &defaultMetricsProvider{
		depth:                   d,
		adds:                    a,
		latency:                 l,
		workDuration:            w,
		unfinished:              u,
		longestRunningProcessor: lrp,
		retries:                 r,
	}
}

func (p *defaultMetricsProvider) NewDepthMetric(name string) workqueue.GaugeMetric {
	return p.depth.WithLabelValues(name, name)
}

func (p *defaultMetricsProvider) NewAddsMetric(name string) workqueue.CounterMetric {
	return p.adds.WithLabelValues(name, name)
}

func (p *defaultMetricsProvider) NewLatencyMetric(name string) workqueue.HistogramMetric {
	return p.latency.WithLabelValues(name, name)
}

func (p *defaultMetricsProvider) NewWorkDurationMetric(name string) workqueue.HistogramMetric {
	return p.workDuration.WithLabelValues(name, name)
}

func (p *defaultMetricsProvider) NewUnfinishedWorkSecondsMetric(name string) workqueue.SettableGaugeMetric {
	return p.unfinished.WithLabelValues(name, name)
}

func (p *defaultMetricsProvider) NewLongestRunningProcessorSecondsMetric(name string) workqueue.SettableGaugeMetric {
	return p.longestRunningProcessor.WithLabelValues(name, name)
}

func (p *defaultMetricsProvider) NewRetriesMetric(name string) workqueue.CounterMetric {
	return p.retries.WithLabelValues(name, name)
}
