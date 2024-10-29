# CloudSQL

## [Platform Admin] Create the composition

```
kubectl create -f composition/hasql.yaml
```

## [Team Admin] Create CloudSQL instance

Please note we are creating this in `config-control` namespace for this sample.
If KCC is setup in a tenant namespace (say using the `AppTeam` sample
composition), then we can use the tenant namespace instead.

```
NAMESPACE=config-control
#NAMESPACE=<app-team's namespace>

kubectl apply -f - <<EOF
apiVersion: facade.compositions.google.com/v1
kind: CloudSQL
metadata:
  name: myteam
  namespace: ${NAMESPACE}
spec:
  regions:
  - us-east1
  - us-central1
  name: myteam-db
EOF
```

Verify the relevant resources are created successfully by running:

```
./get_cloudsql.sh ${NAMESPACE}
```

## [Team Admin] Cleaning up

When done with testing, clean up the resources by deleting the `CloudSQL` CRs:

```
kubectl delete cloudsql myteam -n ${NAMESPACE}
```
