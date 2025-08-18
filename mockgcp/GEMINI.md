# Directory: mockgcp

This directory contains the mock GCP environment used for testing Config Connector.

## Purpose

The mock GCP layer allows us to run KCC tests without needing a real GCP project. This has several advantages:
*   **Hermeticity**: Tests can be run in a self-contained environment, which is great for CI.
*   **Speed**: Mocking GCP APIs can be much faster than making real API calls.
*   **Fault Injection**: We can simulate errors and edge cases that are difficult to reproduce with a real GCP project.

## Implementation

The mock GCP layer is implemented as a set of Go packages that mimic the behavior of the real GCP APIs. It uses a library called `grpc-replay` to record and replay gRPC traffic.

When a test is run against the mock GCP layer, the KCC controllers make API calls to the mock implementation instead of the real GCP APIs.

See also the root `GEMINI.md` and the `tests/GEMINI.md`.
