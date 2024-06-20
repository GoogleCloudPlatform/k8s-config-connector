# Team Page 

## Create the composition

```
kubectl create -f composition/teampage.yaml
```

## Create a Namespace and a Context object

The first step is to create a context object in the namespace where AppTeam will be created.

```
export NAMESPACE=risk-team

kubectl create namespace $NAMESPACE

kubectl apply -f - <<EOF
apiVersion: composition.google.com/v1alpha1
kind: Context
metadata:
  name: context
  namespace: ${NAMESPACE}
spec:
  project: doesntmatter
EOF
```

## Create a team page for `risk`

Create a new `Team` CR

```
kubectl apply -f - <<EOF
apiVersion: facade.compositions.google.com/v1
kind: Team
metadata:
  name: teampage
  namespace: ${NAMESPACE}
spec:
  apps:
  - name: audit
    description: Corporate Audits
  - name: global
    description: Global Risk
  - name: market
    description: Market Risks
EOF
```

## Cleaning up

When done with testing, cleanup the resources by deleting the `Team` CRs.

```
kubectl delete team -n ${NAMESPACE} teampage
```