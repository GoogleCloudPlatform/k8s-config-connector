This scenario verifies that the immutable field can be updated successfully
using the powertool.
1. Create a StorageBucket resource with default location.
2. Use the powertool to update the immutable `spec.location` to `EU`.
3. Read the updated StorageBucket.

The retrieved StorageBucket from step #3 should have `spec.location` field
successfully set to `EU`.