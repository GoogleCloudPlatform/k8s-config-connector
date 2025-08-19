# Directory: pkg/test

This directory contains the testing framework and test fixtures for Config Connector.

## Key Components

*   `resourcefixture/`: Contains the framework for defining test fixtures, which consist of `create.yaml`, `update.yaml`, and `dependencies.yaml` files. The test data is in `resourcefixture/testdata/basic`.
*   `framework/`: Contains the main testing framework, which provides helper functions for creating test environments, running tests, and making assertions.

This package is essential for writing e2e tests for KCC.

See also the root `GEMINI.md`, `tests/GEMINI.md`, and `mockgcp/GEMINI.md`.
