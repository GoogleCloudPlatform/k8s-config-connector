Step1: Create a regional target TCP Proxy resource and its dependencies, operation succeeds

Step2: Update spec.proxyHeader field for the resource, operation fails, need to add `TEST: APPLY-10-SEC` to force stop the job. 
We'll see NO `UPDATE` request in _http03 and resource has UpdateFailed message in _object03.
