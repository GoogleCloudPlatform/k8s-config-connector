package jitter

import (
	"fmt"
	"time"

	leaderelectionv1 "github.com/600lyy/multi-cluster-leader-election/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	JitterFactor = 2.0
)

func GenerateJitterReenqueuePeriod(obj metav1.Object) (time.Duration, error) {
	if obj, ok := obj.(*leaderelectionv1.Lease); ok {
		if obj.Spec.LeaseDurationSeconds < 0 {
			return 0, fmt.Errorf("reconcileInvervalInAnnotation can't be negative")
		}
		reconcileInseconds := time.Duration(obj.Spec.LeaseDurationSeconds) * time.Second
		return wait.Jitter(reconcileInseconds, JitterFactor), nil
	}
	return 0, fmt.Errorf("cannot hanlde other objects than lease")
}
