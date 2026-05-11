# Journal for Create Direct Controller

## 2026-05-11 - FirestoreIndex
- When replacing existing resources or adding GVKs, ensure that `var <Kind>GVK = GroupVersion.WithKind("<Kind>")` is not inserted inside or immediately before the struct documentation block that contains generator annotations like `// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object`. Doing so breaks the `make generate` process, preventing the `DeepCopyObject` method from being created and causing compilation errors (`*Kind does not implement client.Object`).
- `proto.CloneOf(a.desired)` returns a `proto.Message` interface in Go instead of the typed message. Ensure you use type assertion to cast it back to the concrete protobuf struct if you need it (e.g. `proto.CloneOf(a.desired).(*pb.Index)`).
- When a GCP resource field is completely immutable (like `FirestoreIndex`), the `Update` method in the direct controller adapter should simply update and return the status, or return an error if it isn't expected to be modified. KCC will detect immutable fields if they are correctly annotated and handle re-creation if needed, so the `Update` method doesn't need to try to modify them on the backend.
