# Journal: DataflowJob Direct Controller Implementation

## Observations & Learnings

1. **Classic Templates vs. Flex Templates:**
   While `DataflowFlexTemplateJob` uses the Flex Templates service via its own custom launcher, `DataflowJob` represents traditional Dataflow Jobs created from classic templates. This required using the `TemplatesService` (`TemplatesClient` from `"cloud.google.com/go/dataflow/apiv1beta3"`) and its `CreateJobFromTemplate` method instead of the custom REST template launcher.

2. **Legacy Terraform Provider and SdkPipelineOptions:**
   Traditional Dataflow Jobs are largely immutable batch jobs. When running in a fallback or dual-controller test scenario, KCC's unified testing framework runs both the direct and legacy Terraform controllers.
   The legacy Terraform provider expects any GET/List response to contain a populated `sdkPipelineOptions` structure inside the environment:
   Specifically, `temp_gcs_location`, `network`, `subnetwork`, `machine_type`, `max_workers`, `service_account_email`, and `ip_configuration` are read from the `sdkPipelineOptions.options` sub-map in the JSON payload returned by mockgcp.
   If these fields are missing or not returned in `options`, the Terraform provider detects drift / difference against the desired KRM spec, attempts to perform an Update, and subsequently crashes/fails with `Batch jobs cannot be updated.`.

3. **Protobuf structpb.Struct gotcha:**
   When storing and serializing nested `map[string]any` values inside `job.Environment.SdkPipelineOptions` (which is a `*structpb.Struct`), standard `structpb.NewStruct` can fail silently on custom types, enums, or non-`[]any` slices (e.g. `[]string`).
   To resolve this robustly, we implemented a `buildStruct` helper that round-trips the maps through `json.Marshal` and `json.Unmarshal` to clean up slice/scalar types before building the `structpb.Struct`. This ensures that all mock responses deserialize cleanly into a valid Go/Terraform representation.
