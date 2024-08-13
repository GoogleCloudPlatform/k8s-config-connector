# Storage Bucket with CORS and Retention

## Create the composition

```
kubectl create -f composition/cors-bucket.yaml
```

## Create a CRBucket instance

Create a new `CRBucket` CR

```
kubectl apply -f - <<EOF
apiVersion: idp.mycompany.com/v1
kind: CRBucket
metadata:
  name: example-bucket
  namespace: config-control
spec:
  corsURL: "something.foobar.com"
  retentionDays: 10 # days
EOF
```

## Cleaning up

When done with testing, cleanup the resources by deleting the `CRBucket` CRs.

```
kubectl delete crbucket -n config-control example-bucket
```