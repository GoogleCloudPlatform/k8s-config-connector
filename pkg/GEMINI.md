# Directory: pkg

This directory contains shared Go packages used by the controllers and other components of Config Connector. This is where most of the core logic resides.

## Key Subdirectories

*   `clients/`: Contains the clients for interacting with the Kubernetes API server and GCP APIs.
*   `controllers/`: Contains the resource controllers for managing GCP resources. This is the heart of KCC.
*   `dcl/`: Contains the controllers that are based on the Declarative Client Library (DCL).
*   `gcp/`: Contains low-level clients and utilities for interacting with GCP services.
*   `krm/`: Contains utilities for working with the Kubernetes Resource Model (KRM).
*   `test/`: Contains the testing framework and test fixtures.
*   `tf/`: Contains the controllers that are based on the Terraform provider for Google.
*   `webhook/`: Contains the implementation of the admission webhooks.

When you are working on the core logic of KCC, you will spend most of your time in this directory.

See also the root `GEMINI.md` and the `GEMINI.md` files in the subdirectories for more details.
