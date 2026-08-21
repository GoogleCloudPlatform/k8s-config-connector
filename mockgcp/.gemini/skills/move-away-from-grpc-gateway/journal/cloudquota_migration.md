## 2026-05-17 - cloudquota migration

- For `mockcloudquota`, some services (like `QuotaAdjusterSettingsManager`) might only be available in the `v1beta` version of the official Google Cloud Go client library. If you find missing types in `apiv1`, check if `apiv1beta` has them.
- In `mockcloudquota`, I had to use `cloud.google.com/go/cloudquotas/apiv1beta/cloudquotaspb` to get access to both `CloudQuotasServer` and `QuotaAdjusterSettingsManagerServer`.
- Replaced `github.com/golang/protobuf/ptypes/wrappers` with `google.golang.org/protobuf/types/known/wrapperspb` and updated `&wrappers.Int64Value{Value: ...}` to `wrapperspb.Int64(...)`.
