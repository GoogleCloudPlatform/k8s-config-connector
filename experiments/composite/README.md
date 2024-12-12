An experiment into building a simple composite operator.

# Trying it out

Apply the CRD and sample:
```
kubectl apply --server-side -f config/crd
kubectl apply --server-side -f config/samples
```

The sample includes a composite definition for a new type "XNamespace":

```
kubectl get compositedefinitions
# NAME                      AGE
# xnamespaces.example.com   10m
```

The CompositeDefinition describes what happens when an XNamespace is created:

```
kubectl get compositedefinitions xnamespaces.example.com -oyaml
```
```
apiVersion: experiments.cnrm.cloud.google.com/v1alpha1
kind: CompositeDefinition
...  
spec:
  actsOn:
    apiVersion: example.com/v1alpha1
    kind: XNamespace
  definition: |
    kind: Namespace
    apiVersion: v1
    metadata:
      name: foo
  engine: yaml
```

The actsOn section describes what objects we will reconcile, in this case XNamespace objects:
```
  actsOn:
    apiVersion: example.com/v1alpha1
    kind: XNamespace
```

If we look at those objects, we can see that our sample has created one named test1:
```
kubectl get xnamespaces.example.com
# NAME    AGE
# test1   10m
```

The engine and definition in the CompositeDefinition describe how an xnamespace object is expanded:
```
  definition: |
    kind: Namespace
    apiVersion: v1
    metadata:
      name: foo
  engine: yaml
```

In this case, we will use the simple yaml engine to expand an XNamespace object into one other object, a Namespace named foo.

This namespace does not exist yet:
```
kubectl get namespace foo
# Error from server (NotFound): namespaces "foo" not found
```

If you now start the operator with `go run ./cmd/composite-controller`,  you should see it create a namespace within a few seconds.  When it settles down, hit Ctrl-C to exit.

And indeed the foo namespace has been created:
```
kubectl get namespace foo
# NAME   STATUS   AGE
# foo    Active   10m
```

To clean up:
```
go run ./cmd/composite-controller & kubectl delete -f config/samples
# Hit Control-C after a few seconds
kubectl delete namespace foo
```
