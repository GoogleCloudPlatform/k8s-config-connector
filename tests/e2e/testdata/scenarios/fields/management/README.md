This test folder contains scenarios for setting and unsetting a field.

Unsetting a native field (int, bool, string) refers to setting that field to the
0 value in Go.

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