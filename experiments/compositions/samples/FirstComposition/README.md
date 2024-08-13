# Team Page 

## Create the composition

```
kubectl create -f composition/teampage.yaml
```


## Create a Namespace and a Team page

The first step is to create a namespace for the team:

```
export NAMESPACE=my-team
kubectl create namespace $NAMESPACE
```

Create a new `TeamPage` CR in the namespace:

```
kubectl apply -f - <<EOF
apiVersion: idp.mycompany.com/v1alpha1
kind: TeamPage
metadata:
  name: landing
  namespace: ${NAMESPACE}
spec:
  members:
  - name: Jo
    role: Eng Manager
  - name: Jane
    role: Lead
  - name: Bob
    role: Developer
EOF
```

## Cleaning up

When done with testing, cleanup the resources by deleting the `Team` CRs.

```
kubectl delete teampage -n ${NAMESPACE} landing
```