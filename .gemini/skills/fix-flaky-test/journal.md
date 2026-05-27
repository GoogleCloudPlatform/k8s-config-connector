## 2026-05-09: mockgcp infinite reconcile loops
When debugging a flaky test in mockgcp (e.g., tests failing with timeouts or context cancelled in envtest), check the mockgcp service implementation for the resource:
1. GetInstance (or Get) should **not** modify fields with hardcoded values or use time.Now() for timestamps on every get. Doing so causes the KRM status.observedState to change constantly, leading to infinite reconcile loops that overload the API server.
2. storage.Update in mockgcp typically overwrites the entire object. Therefore, update endpoints (Update, Stop, Start, etc.) must first Get the object, apply the specific modifications, and then Update the object.
3. Default field population should happen exclusively in CreateInstance and must respect fields provided by the user in req.Instance.
