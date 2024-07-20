# Team Page 

## Install the CEL expander

Install from released manifests:
```
kubectl apply -f experiments/compositions/expanders/cel-expander/release/manifest.yaml
```

## Create the composition

```
kubectl apply -f composition/teampage.yaml
```


## Create a Namespace and a Team page for `risk`

The first step is to create a namespace where Team will be created.

```
export NAMESPACE=risk-team

kubectl create namespace $NAMESPACE
```

Create a new `Team` CR

```
kubectl apply -f - <<EOF
apiVersion: facade.compositions.google.com/v1
kind: CTeam
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