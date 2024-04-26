# AppTeam 

For now this would only work in CC from a specific project.

## [Platform Admin] Create a Context object

The first step is to create a context object in the namespace where compositions are installed.

```
kubectl apply -f - <<EOF
apiVersion: composition.google.com/v1alpha1
kind: Context
metadata:
  name: context
  namespace: default
spec:
  project: <Replace with CC project>
EOF
```

## [Platform Admin] Create the composition

```
kubectl create -f composition/appteam.yaml
```

## [Platform Admin] Create a new team `clearing`

Create a new `AppTeam` CR in the `config-control` namespace.

```
# Important to avoid collision
randomSuffix=$(tr -dc a-z </dev/urandom | head -c 6)

kubectl apply -f - <<EOF
apiVersion: facade.facade/v1alpha1
kind: AppTeam
metadata:
  name: clearing
  namespace: config-control
spec:
  project: clearing-${randomSuffix}
EOF
```

Verify the relevant resources are created succesfully

```
./get_appteam.sh clearing-${randomSuffix}
```

## [Platform Admin] Cleaning up

When done with testing, cleanup the resources by deleting the `AppTeam` CRs.
For now resources created in namespaces other than the `AppTeam` CR's namespace need to be manually cleaned up.

```
kubectl delete appteam clearing

# to clean up objects not in the same namespace as facade
./cleanup_appteam.sh clearing-${randomSuffix}
```