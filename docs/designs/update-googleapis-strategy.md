# Update googleapis strategy

## Background

Config Connector's `mockgcp` depends on `googleapis` for proto definitions. `googleapis` is an active repository, and we want to be able to use the latest features without being blocked or overwhelmed by massive updates.

## Goals

1. Contributors should be able to use the latest `googleapis` features without being blocked.
2. Updating the `googleapis` commit SHA should not automatically trigger a massive update of all generated code in `mockgcp`.
3. Contributors should be able to selectively update generated code for specific services.

## Strategy

### 1. Weekly Automatic Update of googleapis SHA

A GitHub Action will run every week to update the `git.versions` file in `mockgcp` with the latest commit SHA from the `googleapis` master branch.

This action will:
- Update `mockgcp/git.versions`.
- Run `make sync-repo` to ensure the latest protos are available locally for contributors.
- Create a Pull Request with these changes.

Importantly, it will **not** run the full code generation (`make all`).

### 2. Selective Code Generation

The `mockgcp/Makefile` will be updated to support a `SERVICE` variable. When `SERVICE` is set, only the relevant protos for that service will be regenerated.

Example usage:
```bash
make gen-proto SERVICE=storage
```

### 3. Workflow for Contributors

When a contributor needs a new feature from `googleapis`:
1. If the current `git.versions` is not recent enough, they can run `dev/tasks/update-mockgcp` to update it locally.
2. They run `make sync-repo` to download the latest protos.
3. They run `make gen-proto SERVICE=<service>` to regenerate the code for the service they are working on.
4. They commit the updated `git.versions` and the specific changes in `mockgcp/generated/`.

## Implementation Details

### GitHub Action

A new workflow `.github/workflows/update-googleapis.yaml` will be created. It will use a cron schedule.

### Makefile Changes

- Define `GEN_PROTO_SOURCES` and `GOOGLE_PROTO_SOURCES` variables.
- Use `grep` to filter these sources based on the `SERVICE` variable.
- Update `generate-protos-from-openapi` to also support `SERVICE` filtering.

### update-mockgcp Script

`dev/tasks/update-mockgcp` will be updated to only update `git.versions` and run `make sync-repo`, instead of running `make all`.
