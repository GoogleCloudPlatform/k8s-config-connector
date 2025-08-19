# Directory: cmd

This directory contains the `main` packages for the various Config Connector binaries. Each subdirectory corresponds to a different executable.

## Structure

*   `config-connector`: The main entrypoint for the KCC controller manager.
*   `deletiondefender`: The entrypoint for the deletion defender webhook.
*   `manager`: A helper for the main controller manager.
*   `recorder`: The entrypoint for the recorder tool.
*   `unmanageddetector`: The entrypoint for the unmanaged detector tool.
*   `webhook`: The entrypoint for the main admission webhook.

These binaries are built using the `Makefile` in the root of the project and packaged into container images using the Dockerfiles in the `build` directory.

When you need to understand how a KCC component is started and configured, this is the place to look.

See also the root `GEMINI.md` and the `build/GEMINI.md` for more context.
