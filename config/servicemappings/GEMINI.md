# Directory: config/servicemappings

This directory contains the service mappings for Config Connector.

The `servicemapping.yaml` file defines the mapping between GCP service hostnames (e.g., `storage.googleapis.com`) and the KCC controllers that manage them.

This mapping is used by KCC to determine which controller should be responsible for a given resource.

This file is typically not modified by users, but it is an important part of the KCC internals.

See also the root `GEMINI.md` and `config/GEMINI.md`.
