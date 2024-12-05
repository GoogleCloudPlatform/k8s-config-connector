
# API Design

We are proposing these KRM APIs that would enable composition for users consuming cloud infrastructure via KRM:

1. CRD_T: Template CRD, `Composition`
2. CRD_V: Value CRD, `facade.facade/cloudsqls`
3. CRD_P: Plan CRD

## CRD_T: ( Template)

This is created by Platform Administrator
For the purpose of the POC this is cluster scoped.

This API captures the template and a list of expanders.
It also contains the input API that is consumed by the expanders.

An instance of the CRD_T for POC is `Composition`

Example:
```
apiVersion: composition.kubernetes.io/v1
kind: Composition
metadata:
  name: SimpleDB
spec:
  template: ...
  expanders:
    - name:
      type:
      parameters:
      image:
  inputAPI:
     group: zplatform.easybank.com
     resource: SimpleDB
     createAPI: yes
```

## CRD_V: (Value)

This is a namespaced CRD, instance of which are Created by AppTeamOperator (App/team owner) to consume the CRD_T. This would match the `inputAPI` section of CRD_T. The CRD spec section contains the values for expanding the template.

```
apiVersion: zplatform.easybank.com/v1
resource: SimpleDB
metadata:
  name: db1
spec:
  region: us-east-1
  size: small
  backup: no
```

## CRD_P: (Plan)

A Plan CRD contains the expanded manifests and gating filters.
Contains a list of changes that are being applied to a cluster (diff)
We could use the Plan CRD to apply gates such as:
* Quota
* Approvals
* Linting
* Predictions etc.

This CRD is used by different gates and finally used by the applier.


## Open Questions.

Q1. Should CRD_T be Namespace-scoped or Cluster-scoped ?
> For POC it is cluster scoped. To be revisited if we need a namespace scoped version as well.

Q2. Choice for CRD_V: Single CRD (Say CompositionValues) for all CRD_T's or Separate CRD for each CRD_T instance ?
>  Single CRD => Too generic typing, less scope for access control
>  Separate GVR => Better Access controls based on types, Structured API implies better error checking
>  We choose to go with Separate GVR, one per each CRD_T instance.
