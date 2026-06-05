//go:build !without_exec_applier || !without_direct_applier
// +build !without_exec_applier !without_direct_applier

package simpletest

import (
	"testing"

	"sigs.k8s.io/kubebuilder-declarative-pattern/ktest/testharness"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/status"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
)

func TestDirectSimpleReconciler(t *testing.T) {
	Key := "direct"
	t.Run(Key, func(t *testing.T) {
		testharness.RunGoldenTests(t, "testdata/reconcile/"+Key+"/", func(h *testharness.Harness, testdir string) {
			testSimpleReconciler(h, testdir, applier.NewDirectApplier(), status.NewBasic(nil))
		})
	})
}
