### 2026-06-01 Support snake_case `display_name` in MockGCP CustomClass update
- **Context**: Implementing Phase 2 direct controller and E2E fixtures for `SpeechCustomClass` ([Issue #8909](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8909)).
- **Problem**: The `SpeechCustomClass` direct controller sends update requests with path `display_name` in the field mask. However, the MockGCP implementation for `UpdateCustomClass` in `mockgcp/mockspeech/customclass.go` only handled `displayName` (camelCase) in its switch statement. This resulted in an error from the mock: `update_mask path "display_name" not valid for CustomClass update`.
- **Solution**: Updated the MockGCP switch case in `customclass.go` to handle both `"displayName"` and `"display_name"` as follows:
  ```go
  case "displayName", "display_name":
      obj.DisplayName = req.GetCustomClass().GetDisplayName()
      req.UpdateMask.Paths[i] = "display_name"
  ```
- **Impact**: Ensures that direct controllers sending snake_case update paths (as is standard in direct KCC controllers) are fully supported and correctly reconciled against MockGCP.
