# Field Detection Tool

This tool identifies changes in GCP API proto definitions by comparing the pinned version specified in the `git.version` file with the latest version at HEAD.

It can identify:
- New fields **added** to messages
- Fields **removed** from messages
- Fields that **changed type**

## Usage

```bash
# Basic usage - checks all proto messages used in "generate.sh"
$ go run . detect-new-fields

# Check specific messages
$ go run . detect-new-fields \
  --target-messages="google.cloud.bigquery.datatransfer.v1.TransferConfig"

# Ignore specific fields using configurations from metadata file
$ go run . detect-new-fields \
  --metadata-dir=metadata/
```

## Example Output

```
Changes detected in message: google.cloud.bigquery.datatransfer.v1.TransferConfig
  New field: schedule_options_v2
  New field: error
Changes detected in message: google.cloud.discoveryengine.v1.DataStore
  New field: billing_estimation
  New field: workspace_config
Changes detected in message: google.cloud.discoveryengine.v1.Engine
  New field: disable_analytics
Changes detected in message: google.spanner.admin.instance.v1.Instance
  New field: default_backup_schedule_type
  New field: replica_compute_capacity
```
