# Applier

The expanded manifests need to be applied to a namespace in a cluster running KCC/ACK/.. .

Some options for implementing the applier are:

1. The composition Controller
2. A Separate Job (transient one) with appropriate permissions for the namespace.
3. A Persistent pod to which the manifests are sent via GRPC and it applies to the namespace.
4. Install a RepoSync object with the manifests pushed to a git repo or OCI repo.

## Apply from within the composition controller

- This requires the composition controller with write permissions across multiple namespaces.
+ Since it is a controller it is possible to setup watches and correct any drift on the expanded resources.

## Apply from a separate Job

Permission for applying to a namespace provided via a temporary RBAC.
The pod can be as simple as a kubectl cli or a kpt-cli bundled in, or it can be an applier implemented using an existing library.
Since it is not a persistent controller, we cannot implement drift correction.

- Separate pod with kubectl or kpt-cli would take a dependency on an external project and involves security patches for a such image.
- Don't think we need a separate pod for security reasons yet. Separate reconciler per namespace should cover most security concerns.

## A persistent pod comms via grpc

- The pod needs to be given broader permission to apply in multiple namespaces. Alternatively we can have a pod per namespace which would not scale well.
+ Since it is a persistent controller, we can implement drift correction.

## CS + Reposync

- CS implements a reconciler per namespace RepoSync. This may not scale well in some usecases.
- For CS to work we need to push the 4expanded resources to an OCI repo or a git repo, requiring further setup and integration.
- Taking explicit dependency on ConfigSync would be a problem for wider adoption.
+ CS implements drift correction.

### Libraries for implementing Applier

1. ConfigSync applier: https://github.com/GoogleContainerTools/kpt-config-sync/blob/aaa7ecefa483b8d46dcca95e3e5265097bc43455/pkg/syncer/reconcile/apply.go#L76 (not a lib ?)
2. KPT applier: https://github.com/kptdev/kpt/blob/9d58c99bac13ded5b6e279c4cb78624550fdcd07/commands/live/apply/cmdapply.go#L242
3. kdp applylib/applyset: https://acp.git.corp.google.com/configsync-operator/+/refs/heads/master/pkg/applier/applier.go#63

## POC

For POC we are implementing the Applier in the Composition controller using the kdp/applyset library.
The applyset library brings along more desirable behaviours and we could control pruning.