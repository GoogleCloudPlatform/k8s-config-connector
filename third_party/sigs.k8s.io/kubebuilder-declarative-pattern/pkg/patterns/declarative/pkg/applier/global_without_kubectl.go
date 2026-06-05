//go:build without_direct_applier
// +build without_direct_applier

package applier

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DefaultApplier = NewApplySetApplier(metav1.PatchOptions{}, metav1.DeleteOptions{}, ApplysetOptions{})
