# When a field should be a Ref — and when it should not

**Read the provenance section first.** Some of the ref tooling described here is newly introduced
and has not been ratified as team policy, and it is easy to mistake one for the other.

---

## 1. Provenance: what is settled, and what is newly proposed

### Team-vetted — long-standing convention

| Source | What it establishes |
|---|---|
| `docs/develop-resources/api-conventions/resource-reference.md` | Refs are for "the dependency of another GCP resource". **External-only refs are explicitly permitted when the target has no KCC Kind.** |
| `docs/develop-resources/api-conventions/external-reference.md` | `external` uses the Cloud Asset Inventory resource name, service domain omitted; falls back to the REST resource path when a resource is not in CAI. |
| `TestMissingRefs` exclusions: `.zone`, `.location`, `.machineType`, `.acceleratorType` | Some string fields are deliberately not refs. The stated rationale is terse: *"We don't require refs for zones or regions, nor for instanceTypes."* |
| `vertexaischedule.GCSOutputURIRef` → `StorageBucketRef` | A bucket-only field is a ref. Its proto says *"Format: `gs://bucket-name`"*. |
| `composerenvironment`: `BucketRef` + `DagGCSPrefix *string` | The decomposed bucket-ref-plus-path shape already ships. |
| `StorageFolder`, `StorageManagedFolder`: `spec: [projectRef, resourceID, storagebucketRef]` | KCC already manages entities **inside** a bucket, as child resources with a bucket-ref parent. |
| `SecretManagerSecretVersion.spec.secretData` | KCC already manages data-plane *content*, not only infrastructure. |

### Newly proposed — introduced alongside the reference classifier

`isPatternField`, `notRepresentableReason`, `refs_not_representable.txt`, `refs_deferred.txt` and the
whole `reason=` taxonomy. These encode judgements made while building the classifier, not decisions
the team has ratified. They may be right; **they are not evidence.** Do not cite them as precedent.

## 2. The working principle

> A Ref points to **exactly one addressable, managed** GCP object.

This is a synthesis of the rows above, **not documented policy**. It is offered as an argument to
accept or reject on its merits. The four upstream exclusions are all taxonomy values — a zone is not
a managed object — which supports the *managed* half. Nothing upstream speaks to globs at all.

## 3. Decision table, by the value the API accepts

The discriminator is **the format the GCP API accepts**, not the field's name. Two fields in the same
service get opposite answers: `gcs_output_uri` ("Format: `gs://bucket-name`") is a ref;
`output_uri_prefix` ("URI to output directory") is not.

| Value shape | Treatment | Precedent |
|---|---|---|
| `gs://bucket` | `StorageBucketRef` | `vertexaischedule`, `dataprocbatch` |
| `gs://bucket/dir/` (prefix) | `bucketRef` + separate `path` string | `composerenvironment` |
| `gs://bucket/dir/file` (one object) | **string for now**, recorded as a warning | — (open question) |
| `gs://bucket/dir/*.jsonl` (glob) | not a reference at all — a pattern | ours, §1 |
| `bq://project.dataset.table` | not a GCP resource name; arity ambiguous | ours, §1 |

## 4. `gs://` is not the canonical external form

```go
StorageBucketIdentityFormat = gcpurls.Template("storage.googleapis.com", "projects/{project}/buckets/{bucket}")
StorageBucketURLFormat      = "projects/{{project}}/buckets/{{bucket}}"
```

That is canonical, and matches the CAI rule in `external-reference.md`. `gs://bucket` is only a
**lenient input path** in `StorageBucketIdentity.FromExternal`, and it sets `i.Project = ""` — the
resulting identity carries no project.

Consequence: the GCP API field holds `gs://bucket/...` while the ref's `external` should hold
`projects/{project}/buckets/{bucket}`. **The controller must translate between them.** Converting a
field to a bucket ref always implies mapper work; it is never only a type change.

## 5. Counter-examples worth knowing

**Deliberately not refs (upstream):** `.zone`, `.location`, `.machineType`, `.acceleratorType`.
Infrastructure taxonomy — you do not create a zone.

**Arguments that do NOT work**, because KCC's own CRDs refute them:

- *"It's data, not infrastructure."* `SecretManagerSecretVersion.spec.secretData` manages a secret's
  literal payload.
- *"It lives inside a bucket, so it is out of scope."* `StorageFolder` and `StorageManagedFolder` are
  bucket-parented child resources whose `resourceID` is the path within the bucket.

So KCC has already drawn its line **below** the bucket. Where that line sits between folders and
objects is unsettled.

**The argument that does hold:** a glob names a *set*. A Ref names *one* object. That is independent
of any data-plane framing.

## 6. Backward compatibility decides whether you *may* change the shape

`resource-reference.md`: for TF-based or DCL-based **beta** resources, KCC keeps the original CRD and
behaviour when migrating to Direct. Changing `uris []string` to `uriRefs []Ref` on a beta resource is
a breaking CRD change and is out of bounds.

For `v1alpha1` greenfield with no users, the shape is free to change. **The alpha/beta split, not the
field's meaning, decides whether a conversion is permitted.**

## 7. Open questions — not settled, do not encode as policy

- **Where is the line between folders and objects?** KCC manages `StorageFolder`; it does not manage
  objects. The only real argument for the boundary is cardinality and authorship — folders are few
  and declared deliberately, objects can number in the millions and are written by workloads. That is
  a judgement about typical usage, not a categorical rule.
- **Should globs be a category at all**, or should such fields simply be out of scope for ref
  analysis entirely?
- **Are single-object paths worth expressing** as `bucketRef` + `path`, or do they stay strings?
  Currently strings, recorded as a warning so the decision stays visible.

These belong on
[issue #12344](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/12344), not in
new checks that later get read as precedent.
