# KMSSecretCiphertext Identity and Reference Migration Journal

## Overview
Successfully migrated `KMSSecretCiphertext` under `apis/kms/v1alpha1/` to the modern `identity.IdentityV2` / `identity.ServerGeneratedIdentity` and `refs.Ref` patterns utilizing `gcpurls.Template`.

## Observations & Learnings

### Handling Client-Side / Non-REST Resources with Custom Parsers
`KMSSecretCiphertext` is a client-side wrapper and does not correspond to a standard server-side REST resource in the GCP KMS API. Its Terraform ID format is `{{crypto_key}}/{{ciphertext}}`. Because base64-encoded ciphertexts can contain standard slashes (`/`), we implemented a robust custom parser inside `FromExternal` that correctly splits the GCP resource path from the base64 ciphertext by leveraging the fixed segment count of the preceding `CryptoKey` resource path.

### GCP URL Exceptions & CAI Integration
Since `KMSSecretCiphertext` is a virtual/client-side resource not tracked by the GCP Cloud Asset Inventory (CAI), we registered its template with the `ignoredTemplates` exception registry in `pkg/gcpurls/registry_test.go` to prevent the `TestRegisteredTemplatesMatchCAI` validation test from failing.
