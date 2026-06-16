package status

import (
	"fmt"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

// GetConditions pulls out the `status.conditions` field from runtime.Object
func GetConditions(instance runtime.Object) ([]metav1.Condition, error) {
	statusVal := reflect.ValueOf(instance).Elem().FieldByName("Status")
	if !statusVal.IsValid() {
		return nil, fmt.Errorf("status field not found")
	}
	conditionsVal := statusVal.FieldByName("Conditions")
	if !conditionsVal.IsValid() {
		klog.Errorf("unable to find `status.condition` in %T", instance)
		return nil, nil
	}

	v := conditionsVal.Interface()
	conditions, ok := v.([]metav1.Condition)
	if !ok {
		return nil, fmt.Errorf("unexpecetd type for status.conditions; got %T, want []metav1.Condition", v)
	}
	return conditions, nil
}

// SetConditions sets the newConditions to runtime.Object `status.conditions` field.
func SetConditions(instance runtime.Object, newConditions []metav1.Condition) error {
	statusVal := reflect.ValueOf(instance).Elem().FieldByName("Status")
	if !statusVal.IsValid() {
		// Status not ready.
		return fmt.Errorf("status field not found")
	}
	conditionsVal := statusVal.FieldByName("Conditions")
	if !conditionsVal.IsValid() {
		klog.Errorf("unable to find `status.condition` in %T", instance)
		return nil
	}

	newConditionsVal := reflect.ValueOf(newConditions)
	if !conditionsVal.CanSet() {
		return fmt.Errorf("cannot set status.conditions field")
	}
	if !newConditionsVal.CanConvert(conditionsVal.Type()) {
		return fmt.Errorf("cannot set type %v to status.conditions type %v", newConditionsVal.Type(), conditionsVal.Type())
	}
	conditionsVal.Set(newConditionsVal)

	return nil
}
