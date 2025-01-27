I need to do this task: Implement the wait method for compute operations in the mock http server

compute operations are implemented in the directory `mockcompute/`

The proto definition is this:

```
// The GlobalOperations API.
service GlobalOperations {
  option (google.api.default_host) =
    "compute.googleapis.com";

  option (google.api.oauth_scopes) =
    "https://www.googleapis.com/auth/compute,"
    "https://www.googleapis.com/auth/cloud-platform";

  // Retrieves an aggregated list of all operations. To prevent failure, Google recommends that you set the `returnPartialSuccess` parameter to `true`.
  rpc AggregatedList(AggregatedListGlobalOperationsRequest) returns (OperationAggregatedList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/aggregated/operations"
    };
    option (google.api.method_signature) = "project";
  }

  // Deletes the specified Operations resource.
  rpc Delete(DeleteGlobalOperationRequest) returns (DeleteGlobalOperationResponse) {
    option (google.api.http) = {
      delete: "/compute/v1/projects/{project}/global/operations/{operation}"
    };
    option (google.api.method_signature) = "project,operation";
  }

  // Retrieves the specified Operations resource.
  rpc Get(GetGlobalOperationRequest) returns (Operation) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/operations/{operation}"
    };
    option (google.api.method_signature) = "project,operation";
    option (mockgcp.cloud.operation_polling_method) = true;
  }

  // Retrieves a list of Operation resources contained within the specified project.
  rpc List(ListGlobalOperationsRequest) returns (OperationList) {
    option (google.api.http) = {
      get: "/compute/v1/projects/{project}/global/operations"
    };
    option (google.api.method_signature) = "project";
  }

  // Waits for the specified Operation resource to return as `DONE` or for the request to approach the 2 minute deadline, and retrieves the specified Operation resource. This method differs from the `GET` method in that it waits for no more than the default deadline (2 minutes) and then returns the current state of the operation, which might be `DONE` or still in progress. This method is called on a best-effort basis. Specifically: - In uncommon cases, when the server is overloaded, the request might return before the default deadline is reached, or might return after zero seconds. - If the default deadline is reached, there is no guarantee that the operation is actually done when the method returns. Be prepared to retry if the operation is not `DONE`. 
  rpc Wait(WaitGlobalOperationRequest) returns (Operation) {
    option (google.api.http) = {
      post: "/compute/v1/projects/{project}/global/operations/{operation}/wait"
    };
    option (google.api.method_signature) = "project,operation";
  }

}
```

* Don't write a unit test, you can use the golden test framework instead.

Implement the wait method for compute operations in the mock http server