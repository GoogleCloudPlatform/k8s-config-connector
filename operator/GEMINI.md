# Directory: operator

This directory contains the code for the Config Connector operator.

## Role of the Operator

The KCC operator is responsible for managing the lifecycle of KCC itself. It is particularly important when running KCC in "namespace mode".

In namespace mode, the operator watches for `ConfigConnectorContext` objects in namespaces. When a `ConfigConnectorContext` is created, the operator deploys a new instance of the KCC controller manager in that namespace. This allows for per-namespace configuration and identity, which is more secure and scalable.

The operator is built from the `cmd/operator` package and is installed as part of the main KCC installation.

When you need to understand how KCC is managed in a multi-tenant environment, this is the directory to study.

See also the root `GEMINI.md`.
