# CloudSQL 

## [Platform Admin] Create a Team using AppTeam

First go through the AppTeam recipe and setup the `clearing` Team.

## [Platform Admin] Create the composition

```
kubectl create -f composition/hasql.yaml
```

## [clearing Team Admin] Create a CloudSQL `collateral`

```
kubectl create -f facades/cloudsql-collateral.yaml
```

Verify the relevant resources are created succesfully

```
./get_cloudsql.sh clearing-service
```

## [margin Team Admin] Create a CloudSQL `risk`
```
kubectl create -f facades/cloudsql-risk.yaml
```

Verify the relevant resources are created succesfully

```
./get_appteam.sh margin-service
```

## [Platform Admin] Cleaning up

When done with testing, cleanup the resources by deleting the `CloudSQL` CRs.

```
kubectl delete cloudsql clearing
kubectl delete cloudsql risk
```