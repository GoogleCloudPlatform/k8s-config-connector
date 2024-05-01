# CloudSQL 

## [Platform Admin] Create the composition

```
kubectl create -f composition/hasql.yaml
```

## [Platform Admin] Create a Context object

The first step is to create a context object in the namespace where CloudSQL will be created.

```
kubectl apply -f - <<EOF
apiVersion: composition.google.com/v1alpha1
kind: Context
metadata:
  name: context
  namespace: config-control
spec:
  project: <Replace with CC project>
EOF
```

## [Team Admin] Create a CloudSQL `collateral`

Please note we are creating this in `config-control` namespace for the sample.
If KCC is setup in a tenant namespace (say using `AppTeams` composition), then we can use the tenant namespace instead.

```
namespace=config-control
 
kubectl apply -f - <<EOF
apiVersion: facade.compositions.google.com/v1
kind: CloudSQL
metadata:
  name: collateral
  namespace: ${namespace}
spec:
  regions:
  - us-east1
  - us-central1
  name: collateral-db
EOF
```

Verify the relevant resources are created succesfully

```
./get_cloudsql.sh ${namespace}
```

## [Team Admin] Cleaning up

When done with testing, cleanup the resources by deleting the `CloudSQL` CRs.

```
kubectl delete cloudsql collateral
```