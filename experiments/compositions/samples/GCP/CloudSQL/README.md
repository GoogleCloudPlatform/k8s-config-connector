# CloudSQL 

## [Platform Admin] Create the composition

```
kubectl create -f composition/hasql.yaml
```

## [clearing Team Admin] Create a CloudSQL `collateral`

Please note we are creating this in `config-control` namespace for the sample.
If KCC is setup in a tenant namespace (say using `AppTeams` composition), then we can use the tenant namespace instead.

```
namespace=config-control
# namespace=clearing-<suffix>
 
kubectl apply -f - <<EOF
apiVersion: facade.facade/v1alpha1
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

## [Platform Admin] Cleaning up

When done with testing, cleanup the resources by deleting the `CloudSQL` CRs.

```
kubectl delete cloudsql clearing
```