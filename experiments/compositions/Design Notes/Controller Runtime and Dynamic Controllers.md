# Limitation of Controller runtime

## Supports starting Dynamic Controllers

Controller runtime allows creation of dynamic reconcilers for a CRD (GVK). 

Example implementation:

```
func (r *ExpanderReconciler) SetupWithManager(mgr ctrl.Manager, cr *unstructured.Unstructured) error {
	var err error
	r.Dynamic, err = dynamic.NewForConfig(r.Config)
	if err != nil {
		return fmt.Errorf("error building dynamic client: %w", err)
	}

    // Manager one per process (setup in main.go).
    // Dynamically add a new controller for the crd
	return ctrl.NewControllerManagedBy(mgr).
		For(cr).
		Complete(r)
}
```

This is allowed via the `Add(runnable)` method implemented by controller-runtime Manager.

Controller returns a runnable (`Start` method supported),
https://github.com/kubernetes-sigs/controller-runtime/blob/v0.17.2/pkg/controller/controller.go#L105

This `runnable` is added to `internal/manager/controllerManager.runnables.<type>.Add(runnable)`.
This starts a go-routine with a cancelable context that `Start()`'s the Controller interface.
This in turn waits for GVK and watched objects and calls the `Reconciler()` method. 

The runnable or Controller are not stored in a global datastructure attached to the Manager.
Multiple runnables share the same cancelable context. 

## No support for checking if a controller/reconciler exists

controller-runtime's Manager interface provides no way to check if a controller has been created for an given CRD(gvk).

`type Manager interface {}` : https://github.com/kubernetes-sigs/controller-runtime/blob/main/pkg/manager/manager.go#L55
`type Cluster interface {}` : https://github.com/kubernetes-sigs/controller-runtime/blob/main/pkg/cluster/cluster.go#L40

We would need to build an intrerface on top of the manager that supports checking if a controller exists in a manager.

## No support for stopping a controller

controller-runtime's Manager interface provides no way to `Stop` or `Remove` a runnable. In fact it does not even track the runnables for a given group. During `Add()` a go routine is created and thats it. A shared cancellable context is used across multiple reconcilers.

## Impact for Allotrope
- With naive controller-runtime library we end up starting multiple reconcilers for the same Input GVK everytime the Composition object changes.
- Without stop support we cannot stop the Input GVK controller when a Composition object is deleted.


# Workarounds in Allotrope

## Composition controller maintaining Input GVK controller status

Maintain a list of controllers started by Composition for the Input GVK.
Check it before starting a new controller.
