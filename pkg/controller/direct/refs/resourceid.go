package refs

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetResourceID(u *unstructured.Unstructured) (string, error) {
	resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return "", fmt.Errorf("reading spec.resourceID from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	if resourceID == "" {
		resourceID = u.GetName()
	}
	return resourceID, nil
}
