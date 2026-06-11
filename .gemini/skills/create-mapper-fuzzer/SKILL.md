---
name: create-mapper-fuzzer
description: Standards and workflows for implementing missing or incomplete direct KRM mappers and fuzzers for KCC direct resources. Use this iterative "Compile -> Handwrite -> Fuzz -> Fix" loop to resolve generation panics, type mismatches, and data loss during round-trip translation.
---

Skill: Create Mapper and Fuzzer for Direct Controller

## Overview
This skill guides an automated agent through the process to implement the **Mapper** and **Fuzzer** for KCC direct resources. The goal is to ensure all KRM fields correctly map to GCP fields (Protobuf or OpenAPI) and that this mapping is validated via a lossless round-trip fuzz test.

## Workflow: The Mapper

The `generate-mapper` tool is greedy and attempts to map everything. You must use an iterative feedback loop to identify where it fails.

### 1. Audit and Generate
To generate the mappers without touching type definitions or rebuilding CRDs, you must run the `generate-mapper` tool directly. 

First, locate the `generate.sh` script for the resource (e.g., `apis/<service>/<version>/generate.sh`). 

Check if the `generate-mapper` command is already present in the script. 
- **If it exists:** Copy the exact arguments and run it from `dev/tools/controllerbuilder`.
- **If it does NOT exist:** You must construct the command based on the `generate-types` command found in the script, append it to `generate.sh` so it is saved for the future, and then run it.

The command you construct (and append to `generate.sh`) will look like this:
```bash
go run . generate-mapper \
  --service <copy-from-generate-types> \
  --api-version <copy-from-generate-types> \
  --include-skipped-output
```
*(Run this constructed command from inside `dev/tools/controllerbuilder`)*

**Identify Failures during Generation:**
- **Panics (`klog.Fatalf`):** The generator will crash if it encounters an unhandled Protobuf/OpenAPI type. Note the struct and field mentioned in the panic; you must hand-code the mapper for that struct in `<kind>_mapper.go`.
- **Missing Fields:** Grep the output file `mapper.generated.go` for `// MISSING:`. This indicates the tool couldn't find a matching KRM field. If the field *should* be supported, hand-code the mapping. If it is intentionally unsupported for now, it must be explicitly ignored in the fuzzer (see below).

### 2. Build and Fix
Run the compiler to find type signature mismatches:
```bash
go build ./pkg/controller/direct/...
```

**Common Compilation Errors and Fixes:**
- **Type Mismatches:** If the proto uses `int32` but KRM uses `int64`, the generated code (`out.Field = in.Field`) will fail. You must hand-code the mapping in `<kind>_mapper.go` with an explicit cast: `out.Field = int64(in.Field)`.
- **Resource References:** The generator cannot resolve KCC `*Ref` types (e.g., `ComputeNetworkRef`). These MUST be hand-coded to map to/from their string representations in the Protobuf/API.
- **Enums:** If KRM uses strings for enums but the proto expects integers, hand-code the translation.

**The Golden Rule:** When you define a mapping function (e.g., `KindSpec_FromProto`) in `<kind>_mapper.go`, the generator will detect it and *skip* generating that specific function in `mapper.generated.go`.

## Workflow: The Fuzzer

Once the code compiles, you must prove the mapping is lossless using KCC's `fuzztesting` framework.

### 1. Delegate to the `create-fuzzer` Skill
To create and configure the fuzzer file, you should adopt the `create-fuzzer` skill's workflow. 
Please read and follow the detailed instructions for fuzzer implementation from:
`.gemini/skills/create-fuzzer/skill.md`

That skill will guide you through:
- Creating the `<kind>_fuzzer.go` file (handling both Protobuf and OpenAPI variants).
- Categorizing fields correctly (`f.SpecField`, `f.StatusField`, `f.Unimplemented_Identity`, etc.).
- Registering the direct package in `pkg/controller/direct/register/register.go` (CRITICAL).

### 2. Run and Debug the Fuzzer
After following the `create-fuzzer` instructions, run the fuzz test:
```bash
go test -v -count=1 -timeout=20m ./pkg/fuzztesting/fuzztests/...
```
**If it fails with a `diff`:**
- Analyze the diff to see which field lost data.
- If the data loss is in a supported field, fix your mapping logic in `<kind>_mapper.go`.
- If the field isn't meant to be supported yet, add it to `Unimplemented_NotYetTriaged` as instructed in the fuzzer skill.

## Acceptance Criteria
1. `go build ./pkg/controller/direct/...` succeeds.
2. `go test -v -count=1 -timeout=20m ./pkg/fuzztesting/fuzztests/...` passes flawlessly.
3. Any `// MISSING:` comments remaining in `mapper.generated.go` correspond to fields that are intentionally unsupported and explicitly ignored in the fuzzer configuration.
4. The package is registered in `pkg/controller/direct/register/register.go`.