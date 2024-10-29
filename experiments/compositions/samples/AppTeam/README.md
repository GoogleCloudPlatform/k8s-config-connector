# AppTeam

For now this will only work in CC from a specific project.

## [Platform Admin] Create a Context object

The first step is to create a context object in the namespace where AppTeam will
be created.

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

## [Platform Admin] Create the composition

```
kubectl create -f composition/appteam.yaml
```

## [Platform Admin] Create a new team

Create a new `AppTeam` CR in the `config-control` namespace.

```
# Important to avoid collision
TEAM_NAME=team-$(tr -dc a-z </dev/urandom | head -c 6)

kubectl apply -f - <<EOF
apiVersion: idp.mycompany.com/v1alpha1
kind: AppTeam
metadata:
  name: ${TEAM_NAME}
  namespace: config-control
spec:
  project: ${TEAM_NAME}
  # This is important for some human to be access the project in pantheon
  adminUser: administrator@mycompany.com
  # Please change this to your billing account to be associated with the project
  billingAccount: "000000-000000-000000"
  # Set this to the appropriate folder for the project to be created in
  folder: "000000000000"
EOF
```

Verify the relevant resources are created successfully

```
./get_appteam.sh ${TEAM_NAME}
```

## [Platform Admin] Cleaning up

When done with testing, cleanup the resources by deleting the `AppTeam` CRs.
For now resources created in namespaces other than the `AppTeam` CR's namespace need to be manually cleaned up.

```
kubectl delete appteam $TEAM_NAME

# to clean up objects not in the same namespace as facade
./cleanup_appteam.sh ${TEAM_NAME}
```
