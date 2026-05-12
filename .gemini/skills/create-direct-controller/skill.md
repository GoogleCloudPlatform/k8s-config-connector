# Create Direct Controller Skill

## Process for Migrating from Terraform/DCL to Direct Controller

When introducing a direct controller for a resource that already has an existing Terraform or DCL controller, you must do the following to ensure a safe migration:

### 1. Keep the Existing Controller as Default

In `pkg/controller/resourceconfig/static_config.go`, keep the existing controller (e.g., `k8s.ReconcilerTypeTerraform` or `k8s.ReconcilerTypeDCL`) as the `DefaultController` for now. Add `k8s.ReconcilerTypeDirect` to the list of `SupportedControllers`.

### 2. Enable Side-by-Side Testing

Update `tests/e2e/unified_test.go` to force running the golden tests with *both* the direct controller and the older controller by adding the resource Kind to the `forceDirect = true` cases in `testFixturesInSeries`. This ensures we compute diffs (which should be checked in) that make it easy to verify the direct controller's behavior against the old one.

```go
                                // Start gradually, only running for apikeyskey and tags* fixtures initially
                                forceDirect := false
                                switch fixture.GVK.Kind {
                                case "TagsTagKey", "TagsTagValue", "TagsTagBinding":
                                        forceDirect = true
                                case "APIKeysKey":
                                        forceDirect = true
                                case "TagsLocationTagBinding":
                                        forceDirect = false
+                               case "<YourResourceKind>":
+                                       forceDirect = true
                                default:
                                        forceDirect = false
                                }
```