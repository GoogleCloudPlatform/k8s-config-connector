### [2026-05-27] ModelArmorTemplate Scaffold and Greenfield Review Requirements
- **Context**: Implementing Step 1 (gen-types, CRD, identity, refs) for ModelArmorTemplate.
- **Problem**: Greenfield step 1 PRs must satisfy strict rules like requiring *string pointers for all primitives, including standard fields like Location, regardless of if they are required in the GCP API.
- **Solution**: Set Location *string in ModelArmorTemplateSpec, and manually verified the ModelArmorTemplateIdentity dereferences it with a nil check. Additionally, ModelArmorTemplate proto embeds TemplateMetadata and Labels which need to be explicitly added to ModelArmorTemplateSpec so the generator keeps them instead of pruning them.
- **Impact**: The next agent should remember to explicitly add proto fields that don't get natively scaffolded (like labels, config structures) into ModelArmorTemplateSpec and use *string for Location to satisfy .agents/greenfield-reviewgen-new-types.md.
