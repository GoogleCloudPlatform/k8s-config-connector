### [2026-05-27] Missing instance.proto in pinned googleapis SHA
- **Context**: Implementing CloudRunInstance:Instance using direct approach.
- **Problem**: The proto `google.cloud.run.v2.Instance` was not found in the pinned `googleapis` SHA `731d7f2ab6e4e2ea15030c95039e2cb66174d4fb` because it was added in a later commit (`749628801`).
- **Solution**: Updated `apis/git.versions` to the latest `googleapis` HEAD commit (`7496288011d66f2b34be84377500d114dc74e006`), which contained `instance.proto`, and successfully generated the KRM types.
- **Impact**: When working with new resources, agents may need to dynamically bump `apis/git.versions` and run `dev/tasks/generate-all` if the required `.proto` files are missing in the pinned version.
