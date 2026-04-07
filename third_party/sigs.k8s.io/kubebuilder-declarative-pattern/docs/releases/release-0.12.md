# Release notes for 0.12 release

Note: the 0.12 release is pending.  This serves to accumulate breaking/important changes prior to release.

* Name and Namespace are not longer exported by manifest.Object; use GetName() and GetNamespace() instead.
  This change allows setting the values.
  It also makes `declarative.Object` more similar to `unstructured.Unstructured`, which we model after.
