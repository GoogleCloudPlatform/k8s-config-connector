# Journal: LoggingLink Identity and Reference Implementation

## Observations & Patterns

### 1. Child Resource Dependency
The `LoggingLink` resource is a child of the `LoggingLogBucket` resource. Rather than reinventing how parent identities (such as billingAccounts, folders, organizations, or projects) are resolved and parsed, `LoggingLink` relies on the parent's reference field `LoggingLogBucketRef`. 
Since `LoggingLogBucketRef` already implements the `refs.Ref` interface and has a `ParseExternalToIdentity()` method returning `identity.Identity` (which casts to `*LogBucketIdentity`), we were able to elegantly extract all the parent details in `getIdentityFromLoggingLinkSpec`:
```go
bucketIdRaw, err := obj.Spec.LoggingLogBucketRef.ParseExternalToIdentity()
if err != nil {
    return nil, fmt.Errorf("parsing loggingLogBucketRef: %w", err)
}
bucketID, ok := bucketIdRaw.(*LogBucketIdentity)
```
This is a powerful pattern for any resources with deep hierarchies, reducing boilerplate and ensuring consistent parenting parsing behavior.

### 2. Standardizing File Names & Kind Alignment
The old implementation had files named `link_identity.go` and `link_reference.go` defining the `LinkIdentity` struct. To fully align with the standard patterns where file names and structs prefix the full `Kind` (e.g. `LoggingLink`), we deleted the old files and created `logginglink_identity.go` and `logginglink_reference.go` declaring `LoggingLinkIdentity`. We updated the direct controller to use the new identity struct. This keeps the codebase highly predictable.
