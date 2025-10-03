This scenario verifies that the `spec` fields not specified by the users are
cleaned up properly using the powertool.
1. Create two identical StorageBucket resources but with different
   `state-into-spec` annotation values: `merge` and `absent`.
    1. Create the StorageBucket with `state-into-spec: merge`.
    2. Create the StorageBucket with `state-into-spec: absent`.
2. Use the powertool to update the `state-into-spec` annotation from `merge` to
   `absent` for the first StorageBucket.
3. Read the updated StorageBucket.

The retrieved StorageBucket from step #3 should be the same as the StorageBucket
created by step #1.ii except for the user-specified list field.