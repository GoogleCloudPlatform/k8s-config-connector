---
name: reviewgen-code-health
description: Comprehensive code health, architecture, and style review criteria for Go code in KCC, focusing on single responsibility, encapsulation, Stepdown ordering, collaborator ownership, and clean design.
---

# Code Health & Architecture Review Guide for KCC Go Code

This skill outlines universal code quality, style, and architectural review criteria for Go code across Config Connector (controllers, types, identities, mappers, fuzzers, and tools).

Reviewers (and authors conducting self-review) should conduct a **two-pass review**:
1. **Pass 1 (Correctness & Safety)**: First verify input/output correctness, race conditions, edge cases, error handling, pointer safety, and resource cleanup.
2. **Pass 2 (Design, Style & Architecture)**: Once functional correctness is proven, verify that the implementation adheres to the following structural and design principles.

---

## 1. Code Quality, Style & Design Invariants

### 1. 1 File = 1 Class / Struct (Single Responsibility per File)
* Follow a 1-file, 1-class (1 primary struct/concept + its methods) design with a single responsibility per file.
* Avoid dumping multiple primary structs, unrelated helper types, or bag-of-fields structs into a single file.
* If standalone unit tests are written, each `<file>.go` should have a matching `<file>_test.go` with tests structured into 3 visually separated blocks (**Arrange**, **Act**, **Assert**) separated by blank lines.

### 2. Treat Files as Private by Default (No Cross-File Package Globals)
* Even though Go shares package-level symbols across files, treat each `.go` file as a strict encapsulation boundary.
* Across files in the same package, only interact via a struct's constructor (`newX(...)`) and its exported/package methods (`x.doWork(...)`).
* Do NOT reach into another struct's unexported fields from a different file.
* Do NOT call free helper functions or reference const declarations defined in another file unless that file's explicit sole purpose is a shared domain model.

### 3. No Feature Envy on Value/Data Structs
* Value structs (e.g., parsed identifiers, reference types, options) must only contain data and validation for that data.
* Do not attach downstream domain logic (e.g., formatting IAM emails, constructing API client payloads, or computing GCS regions) onto a shared data struct; keep those as private methods/helpers in the file that actually performs that domain work.

### 4. Own Collaborators via Constructor, Not Method Threading
* If struct A uses collaborator B, wire B inside `newA(...)` (and store B as a field on A) rather than having the caller construct B and pass it into multiple methods of A (`a.foo(b)`, `a.bar(b)`).
* Initialize optional collaborators (e.g., defaulting a nil `io.Writer` to `io.Discard`) in the constructor (`newFoo`) rather than repeating `if x.field != nil` guards across methods.

### 5. Depth-First Call-Tree Ordering Within a File (Stepdown Rule)
* A reader must be able to read any file from top to bottom without jumping to other files—or bouncing to the bottom of the same file—to hunt down helpers or constants.
* Do NOT group all receiver methods at the top of a file and dump private helper functions at the bottom.
* Place every private helper function or sub-method immediately below the function that calls it, in the exact order they are called: struct -> constructor -> constructor's helpers -> primary method -> its helpers in call order -> trivial 1-line getters/delegators at the bottom.

### 6. Colocate Constants with Their First Caller
* Do not dump all `const` declarations at the top of the file if they are only used by a specific sub-tree of functions further down.
* Declare `const` blocks immediately above the first function or method that uses them.

### 7. Orchestrator vs. Leaf Method Granularity
* *Reference*: [Function Length (Martin Fowler)](https://martinfowler.com/bliki/FunctionLength.html).
* **Orchestrator methods** (methods that coordinate multiple logical steps, e.g., reconcile, provision) must stay at a single high level of abstraction and should NOT contain inline `for` loops or anonymous struct tables for sub-steps. Extract each loop or sub-step into a well-named helper placed immediately below in call order.
* **Leaf methods** (methods that perform one concrete task) should treat the 5–15 line length as a guideline, not a hard quota. Do NOT shred a cohesive 20-line leaf method into 3–4 single-use 5-line wrappers.

### 8. No `(nil, nil)` Return Contracts
* Never return `(nil, nil)` from a fetch, lookup, or parser function.
* Return either a valid non-nil value or a descriptive error (or `(T, bool)` for pure in-memory lookups) so callers never need defensive `if val == nil` checks after `if err != nil`.

---

## 2. PR Description & Metadata

### Release Note Block
Verify that the PR description includes a release note block set to `NONE` (or an appropriate user-facing note if intentional):
```release-note
NONE
```
Fail or flag the review if the release note block is missing.

---

## 3. Review Comment Checklist Snippet

When conducting code health reviews, include this section in the review summary:

```markdown
* **Code Health & Architecture**: [Pass/Fail]
  - Single Responsibility (1 File = 1 Concept/Struct): [Pass/Fail]
  - File Encapsulation (No cross-file private field access): [Pass/Fail]
  - No Feature Envy on Data Structs: [Pass/Fail]
  - Collaborator Ownership via Constructor: [Pass/Fail]
  - Depth-First Stepdown Call-Tree Ordering: [Pass/Fail]
  - Constant Colocation: [Pass/Fail]
  - Orchestrator vs Leaf Method Granularity: [Pass/Fail]
  - No (nil, nil) Return Contracts: [Pass/Fail]
* **Release Note Block Present**: [Pass/Fail]
```
