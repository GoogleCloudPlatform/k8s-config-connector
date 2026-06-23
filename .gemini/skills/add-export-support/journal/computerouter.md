# ComputeRouter Exporter Implementation Journal

Implementing export support for `ComputeRouter` went smoothly. Here are the key observations and custom logic required for perfect alignment with the legacy controller:

## Key Observations

1. **Default Fields Omission**: 
   The GCP API response returns `encryptedInterconnectRouter: false` by default, which is mapped directly by `ComputeRouterSpec_v1beta1_FromProto`. However, the old legacy (Terraform-based) exporter does not output this field when it is `false` (the default). To ensure a perfect golden match and avoid an `_exported_object.diff` file, we explicitly nil out `obj.Spec.EncryptedInterconnectRouter` in our direct `Export` method when it is `false`.

2. **Standard Identity Setup**:
   The resource has no `projectRef` on its spec, so we use `export.SetProjectID(u, a.id.Project)` to set the standard project-id annotation. We set `obj.Spec.Region = a.id.Region` and `obj.Spec.ResourceID = &a.id.Router`.
