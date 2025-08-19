# Directory: build

This directory contains the Dockerfiles and build-related scripts for the various Config Connector components.

## Components

The main components built from this directory are:
*   `manager`: The main KCC controller manager.
*   `deletiondefender`: A webhook that prevents accidental deletion of GCP resources.
*   `recorder`: A tool for recording GCP API traffic for testing.
*   `unmanageddetector`: A tool to detect unmanaged GCP resources.
*   `webhook`: The main admission webhook for KCC.

## Building

The components are built using the main `Makefile` in the root of the project. For example, to build the docker images, you can run `make docker-build`.

The Dockerfiles in this directory are used to package the binaries from the `cmd` directory into container images.

See also the root `GEMINI.md` and the `cmd/GEMINI.md` for more context.
