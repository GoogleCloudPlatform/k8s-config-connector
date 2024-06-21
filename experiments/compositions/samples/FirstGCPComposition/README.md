# Storage Bucket with CORS and Retention

## Create the composition

```
kubectl create -f composition/cors-bucket.yaml
```

## Create a CRBucket instance

Create a new `Team` CR

```
kubectl apply -f - <<EOF
apiVersion: facade.compositions.google.com/v1
kind: CRBucket
metadata:
  name: example-bucket
  namespace: config-control
spec:
  corsurl: "something.foobar.com"
  retention: 10 # days
EOF
```

## Cleaning up

When done with testing, cleanup the resources by deleting the `CRBucket` CRs.

```
kubectl delete crbucket -n config-control example-bucket
```