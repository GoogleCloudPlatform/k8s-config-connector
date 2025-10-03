Step1: Create a global target TCP Proxy resource and its dependencies, operation succeeds

Step2: Apply the resource, add spec.resourceID to trigger reconcile, operation succeeds. We'll see NO `UPDATE` request in _http04.
