# v1.125.0

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* Special shout-outs to ... for their
  contributions to this release.
TODO: list contributors with `git log v1.124.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`

## Resources promoted from alpha to beta:

* `RedisCluster` is now a v1beta1 resource.

## New Resources:

* Added support for `PlaceholderKind` (v1beta1) resource.

## New Fields:

* PlaceholderKind
  * Added `spec.placeholder` field.
