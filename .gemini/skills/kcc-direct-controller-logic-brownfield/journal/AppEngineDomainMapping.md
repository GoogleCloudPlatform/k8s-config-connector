# AppEngineDomainMapping Direct Controller Migration Journal

## Observation

1. **Client Library**:
   - `AppEngineDomainMapping` uses GAPIC REST client (`api.NewDomainMappingsRESTClient`) from `cloud.google.com/go/appengine/apiv1`.
   
2. **Creation Request Parameters**:
   - `CreateDomainMapping` requires `Parent` set to `apps/<project>`, `DomainMapping.Id` set to the domain name (e.g. `example.com`), and handles optional `OverrideStrategy` (`STRICT` / `OVERRIDE`).
   - The creation operation is an LRO (`*api.CreateDomainMappingOperation`).

3. **Reconciler Registration**:
   - Added `ReconcilerTypeDirect` to `SupportedControllers` for `AppEngineDomainMapping` in `pkg/controller/resourceconfig/static_config.go`.
   - Kept `DefaultController` as `ReconcilerTypeTerraform` for brownfield backward compatibility.
