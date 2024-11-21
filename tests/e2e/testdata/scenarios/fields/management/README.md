This test folder contains scenarios for setting and unsetting a field.

Unsetting a native field (int, bool, string) in the underlying GCP resource that a Config Connector resource manages can be achieved by setting the corresponding field in the Config Connector resource to its 0 value. This means setting an int field to `0`, a bool field to `false`,  a string field to `""` (and etc).

The folder structure should follow:

```
tests
 └─e2e
    └─testdata
       └─scenarios
          ├─fields
          │  ├─management
          │  │  └─SERVICE_NAME
          │  │     └─RESOURCE_NAME
          │  │        └─scenario_A
          │  └─other_fields_related_scenarios
          └─other_top_level_scenarios
```