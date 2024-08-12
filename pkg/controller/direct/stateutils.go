package direct

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func SetObservedState(u *unstructured.Unstructured, typedObservedState any) error {
	observedState, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedObservedState)
	if err != nil {
		return fmt.Errorf("error converting observedState to unstructured: %w", err)
	}

	if err := unstructured.SetNestedMap(u.Object, observedState, "status", "observedState"); err != nil {
		return fmt.Errorf("setting status.observedState: %w", err)
	}

	return nil
}
