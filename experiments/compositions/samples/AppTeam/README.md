# AppTeam 

## [Platform Admin] Create a Context object

The first step is to create a context object in the namespace where compositions are installed.

```
kubectl apply -f - <<EOF
apiVersion: composition.google.com/v1alpha1
kind: Context
metadata:
  name: admin
  namespace: default
spec:
  project: <TODO: Replace with your ConfigController project ID>
EOF
```

## [Platform Admin] Create the composition

```
kubectl create -f composition/appteam.yaml
```

## [Platform Admin] Create a new team `clearing`
```
kubectl create -f facades/appteam-clearing.yaml
```

Verify the relevant resources are created succesfully

```
./get_appteam.sh clearing clearing-service
```

## [Platform Admin] Create a second team `margin`
```
kubectl create -f facades/appteam-margin.yaml
```

Verify the relevant resources are created succesfully

```
./get_appteam.sh margin margin-service
```

## [Platform Admin] Cleaning up

When done with testing, cleanup the resources by deleting the `AppTeam` CRs.

```
kubectl delete appteam clearing
kubectl delete appteam margin
```