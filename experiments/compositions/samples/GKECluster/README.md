# GKECluster

## [Platform Admin] Create the composition

```
kubectl create -f composition/gke.yaml
```

## [Team Admin] Create GKECluster instance

Please note we are creating this in `config-control` namespace for this sample.
If KCC is setup in a tenant namespace, then we can use the tenant namespace instead.

```
NAMESPACE=config-control
#NAMESPACE=<app-team's namespace>

kubectl apply -f - <<EOF
apiVersion: facade.compositions.google.com/v1
kind: GKECluster
metadata:
  name: myteam-staging
  namespace: ${NAMESPACE}
spec:
  clusterName: staging
  nodepoolName: primary
  maxNodes: 3
  location: us-east1
EOF
```

Verify the relevant resources are created succesfully by running:

```
./get_gke.sh ${NAMESPACE}
```

## [Team Admin] Cleaning up

When done with testing, clean up the resources by deleting the `GKECluster` CR:

```
kubectl delete gkecluster myteam-staging -n ${NAMESPACE}
```
