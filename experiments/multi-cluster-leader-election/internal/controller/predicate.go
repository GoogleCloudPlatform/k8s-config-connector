package controller

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

var (
	ignoredNamespaces = map[string]bool{
		"cert-manager":                         true,
		"cnrm-system":                          true,
		"gke-managed-cim":                      true,
		"gmp-public":                           true,
		"gmp-system":                           true,
		"kube-node-lease":                      true,
		"kube-public":                          true,
		"multi-cluster-leader-election-system": true,
		"ssa-demo":                             true,
		"kube-system":                          true,
		"gke-managed-system":                   true,
		"default":                              true,
	}
)

// This predicate will react only to Create requests from namespaces that are not system managed.
type ManagedByLeaseControllerPredicate struct {
	predicate.Funcs
}

// Create returns true if the given resource has the KCC management label.
func (ManagedByLeaseControllerPredicate) Create(e event.CreateEvent) bool {
	return isManagedByLeaseController(e.Object)
}

// We don't care about the namespace update, thus always returns false
func (ManagedByLeaseControllerPredicate) Update(e event.UpdateEvent) bool {
	return false
}

// Delete always returns false, as deleting a namespace will delete all child resources
func (ManagedByLeaseControllerPredicate) Delete(e event.DeleteEvent) bool {
	return false
}

func isManagedByLeaseController(o metav1.Object) bool {
	if _, ok := ignoredNamespaces[o.GetName()]; ok {
		return false
	}
	return true
}
