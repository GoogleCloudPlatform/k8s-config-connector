# Deletion of Expanded Resources

The POC code does not delete the expanded resources when the Input API (say CloudSQL) is deleted.
Abandon being the only choice is not a sufficient UX. To fix this we have a few proposals.

## Set Plan CRD as the owner

Plan CRD is already set as the parent object in the `applylib` and used as the inventory object to help with pruning. We could make it the `OwnerRef` as well. That way when the `Plan` is deleted all expanded resources are garbage collected. 

Today, the `Plan` already claims the Input API as its owner and is garbage collected when the Input API CR is deleted.

This ensures that when the Input API is deleted, the expanded resources are deleted via cascaded deletion.

## Set Input API as the owner

Input API is set as the `OwnerRef`. This allows for the usecase where the `Plan` doesnt need to exist in the same namespace as the Input API and the expanded resources. We are not yet sure of this usecase and pattern and hence would like to deletgate this option to a later date. 

## Auto-inject OwnerRef in applylib

`applylib` used in Compositions allows us to use an object as a parent. The parent object is used for bookkeeping what resources were applied and help with pruning. 

We could potentially extend `applylib` to allow injecting the parent object as `OwnerRef` also. If not we can inject it explicitly in the expander code Compositions has.

reference:
- https://github.com/kubernetes-sigs/kubebuilder-declarative-pattern/blob/83bd9c05741b108dc854a6b42bf893997b76d090/pkg/patterns/declarative/reconciler.go#L687
- https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/experiments/compositions/composition/internal/controller/expander_applier.go#L137

## Abandon chioce

We should allow the choice to abandon the expanded resources when the Input API is deleted. This could be done via an annotation on the Input API object. 

## POC decisions

We will explicitly inject `OwnerRef` in the expander code to set `Plan` as the owner. This will allow automated garbage collection when the `Plan` object is deleted or garbage collected. 

## Post POC work

- Support `abandon` as an option and prevent garbage collection on `Plan` deletion.
- Enhance `applylib` to auto-inject `OwnerRef` if deemed appropriate. 