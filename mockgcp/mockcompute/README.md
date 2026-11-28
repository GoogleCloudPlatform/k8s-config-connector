# MockCompute: behaviour beyond plain CRUD

Most Compute resources in mockcompute store what they are sent and return it. This page covers the places
where the mock does more than that. For the general approach, see the [mockgcp README](../README.md).
Behaviour here is modelled on GCE. Where it has not been checked against real GCP, this page says so.

## Public images

Public image projects are served from the `publicImages` table in [publicimages.go](publicimages.go).
`images.get` and `images.getFromFamily` build their responses from the table, and boot disks resolve their
`sourceImage` against it. A project is public if it appears in the table. Other projects use stored images.

| Project | Family | Image name | Notes |
|---|---|---|---|
| `debian-cloud` | `debian-11` | `debian-11-bullseye-v20231010` | |
| `debian-cloud` | `debian-12` | `debian-12-bookworm-v20261023` | |
| `cos-cloud` | `cos-81-lts` | `cos-81-12871-1196-0` | |
| `cos-cloud` | `cos-121-lts` | `cos-121-18867-90-42` | Placeholder. Name, id, status and size are invented. See the TODO in `publicimages.go`. |

## Lookup rules

For a public project, `images.get` and `images.getFromFamily` return `NotFound` for any name or family that is
not in the table. An empty family is never in the table, so it also returns `NotFound`. The messages are
`The resource 'projects/P/global/images/N' was not found` and
`The resource 'projects/P/global/images/family/F' was not found`.

Responses for public images omit `id`. See `sourceImageId` under boot disks.

### Source image references

`resolveSourceImage` resolves `initializeParams.sourceImage` on boot disks. A reference may be a path or a full
URL, where the `https://www.googleapis.com/compute/<version>/` or
`https://compute.googleapis.com/compute/<version>/` prefix is removed. These forms are accepted:

| Reference (after prefix removal) | Resolves to |
|---|---|
| `projects/P/global/images/family/F` | family F in project P |
| `global/images/family/F` | family F in the instance's project |
| `projects/P/global/images/N` | image N in project P |
| `global/images/N` | image N in the instance's project |
| `P/X` | image X in project P, then family X in project P (see below) |

`P/X` is ambiguous, and the mock resolves it as follows. It first looks for an image named X in project P. Only
if that returns `NotFound` does it look for family X in P. An image with the same name therefore takes
precedence over a family. If P is not a registered project, the name lookup returns `PermissionDenied` and the
family lookup is not tried.

Anything else is rejected with `InvalidArgument`, `source image "<ref>" is not valid`. This includes a bare
name with no project and paths with extra segments.

## How to add a public image

1. Get the image details with `gcloud compute images describe-from-family F --project P`.
1. Add an entry to `publicImages` in [publicimages.go](publicimages.go). Set `project`, `family`, `name`,
   `status`, `diskSizeGb` and `id` from that output. `description` is optional.
1. Keep the id `2443108620951880213` on the debian-cloud entries. It is shared on purpose: `disksv1.go` and
   existing goldens hard-code it.
1. Regenerate the affected goldens with `hack/compare-mock`, passing a fixture path without a trailing slash, for
   example `hack/compare-mock pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computeinstance/computeinstancefromtemplate`.
1. Check `_http_mock.diff` by hand. A regeneration run does not rewrite `.diff` goldens. `CompareGoldenFile`
   routes any file whose name contains `_http` through `compareHTTPLogs`. Diff output contains no HTTP events,
   so the comparison passes and the `.diff` golden is never rewritten
   ([#16225](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/16225)). Compare the diff with
   `diff _http_old_controller_mock.log _http_mock.log` (GNU diff).

## networkIP resolution

On instance insert (not on update), each `networkInterfaces[].networkIP` is handled as follows:

- An empty value, or a valid IPv4 or IPv6 literal, is stored unchanged.
- Any other value is treated as a reference, and must be one of:
  - `projects/P/regions/R/addresses/A`, where P is a project ID or number
  - `regions/R/addresses/A`, where the address is in the instance's project

  The stored value becomes the IP of that address.

On `instances.update`, references are stored as sent and are not resolved.

Errors, checked in this order:

| Condition | Code | Message |
|---|---|---|
| Value is not one of the two forms | `InvalidArgument` | `Invalid value for field 'resource.networkInterfaces[i].networkIP': '<value>'` |
| Project not found (ID or number) | `PermissionDenied` | `Project '<P>' not found or permission denied.` |
| Address not found | `NotFound` | `The resource 'projects/P/regions/R/addresses/A' was not found` |
| Address has no IP | `InvalidArgument` | Same as the first row |
| Address is in another project | `InvalidArgument` | `... '<value>'. IP address '<address>' (<ip>) is reserved by another project.` |
| Address is in another region | `InvalidArgument` | `... '<value>'. IP address '<address>' (<ip>) is in region R, but the instance is in region R2.` |

The order and wording of these checks are modelled on GCE. They are not verified against real GCP, which may
order them differently or word them differently.

Deferred:

- Allocating an IP when `networkIP` is omitted. The field is left empty.
- Mock addresses defaulting to `8.8.8.8`. `regionaladdress.go` and `globaladdress.go` set this when no address
  is given.

## Boot disks

On instance insert, each attached disk that is `boot`, has `initializeParams`, has no `source`, and is not
`SCRATCH` gets a persistent disk created from its parameters. Every boot disk and every networkIP is resolved
before anything is written.

- Disk name is `initializeParams.diskName`, or the instance name if that is empty. The disk is zonal, at
  `projects/P/zones/Z/disks/<name>`.
- If a disk with that name already exists in the zone, the insert fails with `AlreadyExists`.
- Size is `initializeParams.diskSizeGb`, then the source image's `diskSizeGb`, then 10 GB.
- Type is `initializeParams.diskType`, defaulting to `pd-standard`. Bare names resolve in the instance's zone.
- If `sourceImage` is set, the disk records `sourceImage` and `sourceImageId` from the resolved image. Reference
  errors are as described above.
- The disk's `users` contains the instance.
- The attached disk returned by the mock has `source` set to the disk's URL and no `initializeParams`. Missing
  `type`, `mode`, `interface`, `deviceName` and `diskSizeGb` are filled in.
- If a later disk or the instance itself cannot be created, the disks already created by that insert are deleted
  on a best-effort basis.

On `instances.delete`, the instance is removed first. Then, for each attached disk whose `source` is a zonal disk:

- If `autoDelete` is set, the disk is deleted.
- Otherwise, the instance is removed from the disk's `users`.

Failures in this step are logged, not returned. The mock does not default `autoDelete` for boot disks, so a disk
created from `initializeParams` without `autoDelete` is kept.

Deferred:

- `users` on disks attached by `source`. Inserting with an existing disk, or calling `attachDisk`, does not add the
  instance to `users`.
- Resolving project numbers in `sourceImage`. `projects/<number>/global/images/<name>` and `<number>/<name>`
  return `PermissionDenied`, because the name lookup uses `GetProjectByID`. The family forms do not look up a
  project and return `NotFound`.

## Issue #293 coverage

[#293](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/293) covers ComputeInstance
immutable-field errors on `networkInterface.0.networkIP` and `bootDisk.0.initializeParams.0.image`. This does not
fix the issue. Its relevant coverage is:

- **AC1** (address reference as networkIP): the mock stores the resolved IP, as GCE does. A test can check that
  an address reference in `networkIpRef` causes no immutable-field error and no perpetual diff.
- **AC2** (image family as `sourceImageRef`): image families (debian, COS) resolve in `sourceImage`, and the boot
  disk records the concrete image via `sourceImageId`. This lets a test check that a family reference causes no
  immutable boot-image error.
- **AC3** (restart loop with a hard-coded networkIP), **AC4** (GKE add-on version lag) and **AC5** (GKE to KCC
  version mapping docs) are out of scope.
