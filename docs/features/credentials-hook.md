# Per-object credentials hook

Config Connector can obtain Google Cloud credentials from an external local
process for each KRM object it reconciles. This is useful when one controller
process manages resources for multiple identities, including folder- and
organization-scoped resources whose API requests do not identify a project.

The hook is optional. Without a hook, existing authentication behavior is
unchanged.

## Enable the hook

Start `kcc-manager` with the path to the hook process's Unix domain socket:

```sh
kcc-manager --identity-hook-socket=/var/run/credentials-hook/hook.sock
```

The hook process and manager must share access to the socket. The manager does
not create or start the hook process. This flag configures the manager directly;
it does not add a field to the `ConfigConnector` or `ConfigConnectorContext` CRD.

Applications embedding the manager can instead set
`kccmanager.Config.IdentityResolver` to an implementation of `identity.Resolver`
from `pkg/gcp/identity`.

## Hook protocol

The hook serves HTTP over the Unix socket. Both endpoints accept JSON request
bodies and must return HTTP 200 with a JSON response on success. Other status
codes fail the operation; a bounded portion of the response body is included in
the error.

### Resolve an object's identity

Config Connector calls `POST /v1/resolve` once for each reconciliation that
requires Google Cloud credentials:

```json
{
  "group": "iam.cnrm.cloud.google.com",
  "kind": "IAMPolicyMember",
  "object": {
    "apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
    "kind": "IAMPolicyMember",
    "metadata": {
      "name": "folder-reader",
      "namespace": "team-a",
      "annotations": {
        "cnrm.cloud.google.com/project-id": "example-project"
      }
    },
    "spec": {
      "member": "serviceAccount:reader@example-project.iam.gserviceaccount.com",
      "role": "roles/browser",
      "resourceRef": {
        "kind": "Folder",
        "external": "folders/123456789"
      }
    }
  }
}
```

The `object` field contains the full reconciled object, not just its metadata.
The separate `group` and `kind` fields remain available when a typed object's
`apiVersion` and `kind` fields are empty.

Return an identity ID that is stable for the credentials being used:

```json
{"id":"controller@example-project.iam.gserviceaccount.com"}
```

### Obtain or refresh a token

Config Connector calls `POST /v1/token` when the identity needs a token:

```json
{"id":"controller@example-project.iam.gserviceaccount.com"}
```

Return an access token, its type, and a nonzero expiry:

```json
{
  "access_token": "ACCESS_TOKEN",
  "token_type": "Bearer",
  "expiry": "2027-01-01T12:00:00Z"
}
```

Tokens are cached and refreshed shortly before expiry. Objects resolving to the
same ID share a token source. Each request to the hook has a 30-second timeout.
Token refresh is bounded by that timeout because `oauth2.TokenSource` does not
accept a reconciliation context.

## Authentication behavior

Direct REST and gRPC clients use the credentials associated with the current
reconciliation. DCL clients also bind those credentials for library calls that
omit the reconciliation context. Terraform uses a provider copy with separate
HTTP clients and request batchers per identity so batching cannot mix identities.

When a hook is enabled, a failed resolution, missing credentials, token error, or
unavailable hook fails the operation. Config Connector does not fall back to
Application Default Credentials for those calls.

The background project IAM dependency tracker is disabled with a hook because
its polls have no reconciled object from which to resolve credentials.
`IAMPartialPolicy` still checks and corrects drift during its regular
reconciliations.

## Hook implementation responsibilities

Treat the hook as a privileged credential service. Restrict access to the Unix
socket and the directory containing it. Do not log returned tokens.

The hook receives object fields controlled by resource authors, including
annotations and resource references. It must authorize the namespace or tenant
to use the identity it returns. Selecting an identity solely from an unvalidated
project annotation is not an authorization check.

Use stable identity IDs. Token sources and Terraform provider copies are cached
by ID for the process lifetime, so generating a distinct ID for every request
would grow those caches without bound. A token source supplied by an in-process
resolver must refresh tokens for the lifetime of its identity ID.
