# Directory: pkg/webhook

This directory contains the implementation of the admission webhooks for Config Connector.

## Webhooks

KCC uses several admission webhooks to provide features such as:
*   **Validation**: Validating KCC resources before they are created or updated.
*   **Mutation**: Defaulting values in KCC resources.
*   **Deletion Defender**: Preventing accidental deletion of GCP resources.

The webhooks are implemented using the `controller-runtime` library.

When you need to understand how KCC validates and mutates resources, this is the directory to look at.

See also the root `GEMINI.md` and `pkg/GEMINI.md`.
