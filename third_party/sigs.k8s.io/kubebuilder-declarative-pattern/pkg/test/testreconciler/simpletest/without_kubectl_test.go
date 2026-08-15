package simpletest

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/ktest/testharness"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/status"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
)

func TestSSASimpleReconciler(t *testing.T) {
	Key := "ssa"
	a := applier.NewApplySetApplier(
		metav1.PatchOptions{FieldManager: "kdp-test"}, metav1.DeleteOptions{}, applier.ApplysetOptions{})
	t.Run(Key, func(t *testing.T) {
		testharness.RunGoldenTests(t, "testdata/reconcile/"+Key+"/", func(h *testharness.Harness, testdir string) {
			testSimpleReconciler(h, testdir, a, status.NewKstatusCheck(nil, nil))
		})
	})
}
