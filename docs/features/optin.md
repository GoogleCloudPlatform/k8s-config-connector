# Opt in to Enable Experimental Direct Controllers

> `Feature State`: `alpha` as of version v1.123.1

Config Connector can be configured to enable experimental versions of direct
controllers for reconciling specific resources. This will allow users to test
the new direct controller code for bug fixes or features not available in
legacy controllers.

## Enabling via ConfigConnectorContext (Recommended)

You can enable direct controllers for all resources of a specific kind
within a namespace using the `experiments.controllerOverrides` field in the
`ConfigConnectorContext` resource. For more details and examples, see the
[Controller Implementation Overrides](./controller-configuration.md) documentation.

## Verifying

To verify the direct controller code is in-use for a particular resource,
check the logs for the phrase "Running reconcile" associated with the resource.

```
$ kubectl -n cnrm-system logs pod/cnrm-controller-manager-0 | grep -e "Running reconcile" -e "my-sqlinstance"
```

You should see output similar to the following:

> {"severity":"info","timestamp":"2024-11-04T21:56:28.704Z","msg":"Running reconcile","controller":"sqlinstance-controller","controllerGroup":"sql.cnrm.cloud.google.com","controllerKind":"SQLInstance","SQLInstance":{"name":"my-sqlinstance","namespace":"default"},"namespace":"default","name":"my-sqlinstance","reconcileID":"577fe58d-14eb-4e5d-9642-c9ec4b1a3137","resource":{"name":"my-sqlinstance","namespace":"default"}}

If the legacy controller is enabled (instead of the new direct controller), the
logs will show "starting reconcile" for the resource instead of "Running reconcile".

> {"severity":"info","timestamp":"2024-10-30T23:47:49.726Z","logger":"sqlinstance-controller","msg":"starting reconcile","resource":{"name":"my-sqlinstance","namespace":"default"}}


## Applicability

If you're unsure if a particular resource has an experimental direct controller
available and want to try out the new code wherever possible (or if you're just
curious to learn more about the internals of this feature in Config Connector),
read on.

To find out if an experimental direct controller is available for a resource,
check CRD for the resource.

If the CRD specifies either of the following legacy
controller labels, then it is potentially possible to enable an experimental
direct controller for resources of that kind.

* `cnrm.cloud.google.com/tf2crd: true`
* `cnrm.cloud.google.com/dcl2crd: true`

To check the CRD labels for a resource type, you can run the following command:
```
kubectl get crd sqlinstances.sql.cnrm.cloud.google.com -ojson | jq '.metadata.labels'
```

Alternatively, you can check in the source tree under `config/crds/resources/`
for the CRD YAML, but remember to double-check your source checkout matches
your running Config Connector version.

However, not all resources with the legacy controller annotation on their CRDs
will have an experimental direct controller available. This is because not all
resources with the legacy controller annotation have a direct controller
implemented yet; many of the direct controllers are not yet implemented. To
confirm if there is an experimental direct controller available for a
particular resource, check to make sure there is an implementation of the
controller for that type in `pkg/controller/direct`. The controllers are found
in `*_controller.go` [[example]](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/pkg/controller/direct/sql/sqlinstance_controller.go).
Also, verify that the direct controller is registered in `pkg/controller/direct/register/register.go`
[[example]](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/f485d486553f1fc939eb8e487ccabaf2f2f032ed/pkg/controller/direct/register/register.go#L43).

At this point, if there is both a legacy controller label on the CRD and a
direct controller implemented + registered, then you can be reasonably sure
there is a direct controller available. To try it out, add the override in the
`ConfigConnectorContext` for a resource of that kind, and use the steps above
to verify the direct controller is enabled.
