The `preview` command is a CLI tool that enables you to visualize potential changes to your Google Cloud resources managed by Config Connector. It executes the current version of KCC against your cluster, intercepting and recording all intended modifications without actually applying any changes to the GCP resources.

## Usage

For running the preview command directly with your current cluster context, you can use the following command:

```sh
REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/experiments/kompanion

go run main.go preview
```

For more information about the options, you can run:

```sh
go run main.go preview --help
```

## Example result

```text
Finish reconciled 7 out of 7 resouces.
Detect 6 good and 1 bad objects.
-----------------------------------------------------------------
Namespace: config-control
GROUP                            KIND              GOOD   BAD
alloydb.cnrm.cloud.google.com    AlloyDBInstance   1      0
spanner.cnrm.cloud.google.com    SpannerInstance   1      0
bigquery.cnrm.cloud.google.com   BigQueryTable     1      0
bigquery.cnrm.cloud.google.com   BigQueryDataset   1      0
storage.cnrm.cloud.google.com    StorageBucket     1      1
-----------------------------------------------------------------
Namespace: test-autopilot
GROUP                       KIND          GOOD   BAD
sql.cnrm.cloud.google.com   SQLInstance   1      0
```

The preview tool generates a summary report in the current directory, similar to the example above. This result indicates that 6 out of 7 resources are valid ("good"), while 1 resource is invalid ("bad"). Specifically, the invalid resource is a `StorageBucket` (group: `storage.cnrm.cloud.google.com`) in the `config-control` namespace. This discrepancy arises from changes made by the KCC controller which were intercepted by the preview tool.

If any invalid resources are detected, the preview tool generates an additional file with a `-detail` suffix in the current directory containing detailed information about the issues found.