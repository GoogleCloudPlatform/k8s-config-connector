# Labels and Annotations Handling

This document describes how KCC handles labels and annotations when creating and updating GCP resources.

## Annotations

If `spec.annotations` is specified, these annotations will be used without filtering. If `metadata.annotations` is also specified, the annotations will be filtered to remove keys that do not match the GCP requirements. The filtered `metadata.annotations` will then be merged with `spec.annotations`. In case of conflicting keys, the value from `spec.annotations` will take precedence.

The GCP requirements for annotation keys are:
* Must follow pattern `[a-z0-9A-Z]+([\.\-] _[a-z0-9A-Z]+)`
* Be less than 64 characters
* Must have a UTF encoding of less than 128 bytes

## Labels

Similarly, if `spec.labels` is specified, these labels will be used without filtering. If `metadata.labels` is also specified, the labels will be filtered to remove keys that do not match the GCP requirements. The filtered `metadata.labels` will then be merged with `spec.labels`. In case of conflicting keys, the value from `spec.labels` will take precedence.
