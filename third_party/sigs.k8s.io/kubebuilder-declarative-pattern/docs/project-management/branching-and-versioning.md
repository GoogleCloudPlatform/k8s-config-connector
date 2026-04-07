# Branching and versioning strategy

## Versioning

kubebuilder-declarative-pattern is a library, that works with controller-runtime and kubernetes.

We follow semantic versioning, similar to other tooling in the kubernetes ecosystem.

Because both controller-runtime and client-go introduce changes that require code changes, we follow their versioning.
Specifically we align with the controller-runtime version, which itself aligns with kubernetes versioning.

Thus:

| kdp version | controller-runtime version | client-go version |
|---|---|---|
| v0.20 | v0.20 | v0.32 |
| v0.19 | v0.19 | v0.31 |
| v0.18 | v0.18 | v0.30 |
| ... | ... | ... | 
| v0.x | v0.x | v0.x+12 |

If we need to release multiple versions of kdp for a single version of controller-runtime (for example if we want to fix a bug),
we use patch versions for that.
We want to avoid breaking changes, and we will only make breaking changes on minor version bumps (as far as possible).

## Branches

We maintain a `release-0.x` branch for the kdp major version `0.x`.  Patch versions are cherry-picked to the branch and released.

We cut the release branch with the beta release (i.e. we create `release-1.100` when we tag `1.100.0-beta.1`).

Before beta releases, the master branch is for the next minor version.  If we tag a version, it would be (for example) `1.101.0-alpha.1` (and then `1.101.0-alpha.2` etc), assuming the most recent release branch was `1.100`
