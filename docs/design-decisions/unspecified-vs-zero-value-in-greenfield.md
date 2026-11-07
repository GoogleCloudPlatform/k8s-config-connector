# Unspecified vs. Zero Values in KCC Direct Controllers

## Context

Config Connector (KCC) reconciles declarative Kubernetes Resource Model (KRM) specifications against Google Cloud Platform (GCP) APIs. 

A central challenge in writing KCC Direct Controllers is **semantic translation between Kubernetes and GCP**:
* **In Kubernetes (KRM)**: User intent is strictly defined via pointers. An omitted field in YAML is unmarshalled as `nil` (**Unset / Unmanaged**), while an empty block (`{}` or `[]`) is unmarshalled as an allocated empty struct/map/slice with `len == 0` (**Explicitly Managed as Empty**).
* **In GCP API Layers**: Depending on whether the underlying API contract is **Proto3**, **Proto2**, or **Legacy Discovery Document (HTTP/REST)**, the API runtime may not distinguish between an absent field and a default zero-value (`0`, `false`, `""`, or `nil` map).

This document serves as the architectural reference for how KCC Direct Controllers handle field comparison, zero values, unset fields, and drift detection across different GCP API formats.

---

## Behavior Matrix: Key Types That Can Cause Reconciliation Issues

### 1. GET Responses / Read State (Comparison & Drift Detection)

The table below highlights the key data types that **diverge in field presence between KRM user specifications and GCP GET responses**, creating risks of false drift if unhandled. (For the full matrix across all 7 data types, see [Appendix A.5: Comprehensive Behavior Matrix](#a5-comprehensive-behavior-matrix-all-data-types)).

| High-Risk Data Type | Modern SDK (Proto3 / REST Client) | Legacy Proto2 | Legacy Discovery SDK | KCC KRM Struct (`spec`) | Comparison Risk & Mitigation |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **1. Maps (`map[string]string`)** | **Not Distinguishable**<br>(`nil` and `{}` = 0 bytes) | **Not Distinguishable**<br>(`nil` and `{}` = 0 bytes) | **Not Distinguishable**<br>(server stores 0 entries) | **Distinguishable**<br>(`nil` vs `map{}`) | **Risk**: `reflect.DeepEqual(map{}, nil)` evaluates to `false`, causing false drift loops.<br>**Mitigated?**: No.<br>**Mitigation**: Use `maps.Equal(desiredPb, actualPb)`. |
| **2. Lists / Slices (`[]string`)** | **Not Distinguishable**<br>(`nil` and `[]` = 0 bytes) | **Not Distinguishable**<br>(`nil` and `[]` = 0 bytes) | **Not Distinguishable**<br>(0 elements = empty) | **Distinguishable**<br>(`nil` vs `[]`) | **Risk**: `reflect.DeepEqual([], nil)` evaluates to `false`, causing false drift loops.<br>**Mitigated?**: No.<br>**Mitigation**: Use `slices.Equal` for ordered lists or sets for unordered lists. |
| **3. Raw Primitives without `optional`** (`int`, `bool`, `string`) | **Not Distinguishable**<br>(`0` and unset = 0 bytes) | **Distinguishable**<br>(`*int32` pointer) | **Not Distinguishable**<br>(requires `ForceSendFields`) | **Distinguishable**<br>(`nil` vs `&0`) | **Risk**: Direct scalar comparison against unpopulated service defaults detects false drift.<br>**Mitigated?**: Yes.<br>**Mitigation**: Populate service defaults into `desiredPb` via `computedFieldPaths` before comparison. |
| **4. Enums without `optional`** (`enum`) | **Not Distinguishable**<br>(default `0` = 0 bytes) | **Distinguishable**<br>(`*Enum` pointer) | **Distinguishable**<br>(named string) | **Distinguishable**<br>(`nil` vs `&"ENUM"`) | **Risk**: Direct enum comparison against unpopulated server default enum constant detects false drift.<br>**Mitigated?**: Yes.<br>**Mitigation**: Populate server default enum into `desiredPb` via `computedFieldPaths` before comparison. |

> **Note**: Mitigations are all available in KCC patterns and utility functions, but this does not mean they have already been applied to all greenfield resources across the codebase. We may still need to improve older greenfield resources with these new improvements as gaps are identified.

---

### 2. Mutation Requests (Payload Construction & Intent Serialization)

The table below highlights whether and how each type's presence is **distinguished and preserved across CREATE (POST), FULL REPLACE (PUT), and PARTIAL UPDATE (PATCH)**:

> **Important (Legacy Discovery SDK Semantics)**: In Legacy Discovery SDKs (`google.golang.org/api/*`), omitting a field in any mutation request (which happens by default for empty/zero fields due to Go's `omitempty`) causes the GCP backend to treat the field as **unmanaged**, preserving existing server state (or applying creation defaults on `Create()`). To explicitly clear a field to an empty value (`""`, `0`, `{}`, `[]`), the field must be added to `ForceSendFields`. To explicitly clear/nullify a field on the server, the field must be added to `NullFields`.

| High-Risk Data Type | Modern SDK (Proto3 / REST Client) | Legacy Proto2 | Legacy Discovery SDK | KCC KRM Struct (`spec`) | Request Construction Risk & Mitigation |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **1. Maps (`map[string]string`)** | • **CREATE (POST)**: Not Distinguishable (`nil` and `{}` emit 0 bytes).<br>• **REPLACE (PUT)**: Not Distinguishable (`nil` and `{}` both wipe server map to empty).<br>• **UPDATE (PATCH)**: Can declare empty value, but not unset (via `updateMask`). | • **CREATE (POST)**: Not Distinguishable (`nil` and `{}` emit 0 bytes).<br>• **REPLACE (PUT)**: Not Distinguishable (`nil` and `{}` both wipe server map).<br>• **UPDATE (PATCH)**: Can declare empty value, but not unset (via `updateMask`). | • **CREATE (POST)**: Not Distinguishable (0 entries).<br>• **REPLACE (PUT)**: Not Distinguishable (`nil` and `{}` both wipe server map).<br>• **UPDATE (PATCH/PUT)**: Distinguishable (declare empty `{}` via `ForceSendFields` or unset/nullify via `NullFields`). | **Distinguishable**<br>(`nil` vs `map{}`) | **Risk**: In legacy Discovery SDK, failing to populate `NullFields` or `ForceSendFields` causes JSON serializer to omit empty maps or override unset intention.<br>**Mitigated?**: No.<br>**Mitigation**: In both `Create()` and `Update()`, check KRM allocation and populate `ForceSendFields` for `{}` or `NullFields` for unsetting. |
| **2. Lists / Slices (`[]string`)** | • **CREATE (POST)**: Not Distinguishable (`nil` and `[]` emit 0 bytes).<br>• **REPLACE (PUT)**: Not Distinguishable (`nil` and `[]` both wipe server list to empty).<br>• **UPDATE (PATCH)**: Can declare empty value, but not unset (via `updateMask`). | • **CREATE (POST)**: Not Distinguishable (`nil` and `[]` emit 0 bytes).<br>• **REPLACE (PUT)**: Not Distinguishable (`nil` and `[]` both wipe server list).<br>• **UPDATE (PATCH)**: Can declare empty value, but not unset (via `updateMask`). | • **CREATE (POST)**: Not Distinguishable (0 elements).<br>• **REPLACE (PUT)**: Not Distinguishable (`nil` and `[]` both wipe server list).<br>• **UPDATE (PATCH/PUT)**: Distinguishable (declare empty `[]` via `ForceSendFields` or unset/nullify via `NullFields`). | **Distinguishable**<br>(`nil` vs `[]`) | **Risk**: In legacy Discovery SDK, failing to populate `NullFields` or `ForceSendFields` causes JSON serializer to omit empty lists or override unset intention.<br>**Mitigated?**: No.<br>**Mitigation**: In both `Create()` and `Update()`, check KRM allocation and populate `ForceSendFields` for `[]` or `NullFields` for unsetting. |
| **3. Raw Primitives without `optional`** (`int`, `bool`, `string`) | • **CREATE (POST)**: Not Distinguishable (`0`/`false`/`""` emit 0 bytes -> server defaults).<br>• **REPLACE (PUT)**: Not Distinguishable (`0`/`false`/`""` emit 0 bytes -> server defaults).<br>• **UPDATE (PATCH)**: Can declare zero/empty value, but not unset (via `updateMask`). | • **CREATE, PUT & PATCH**: Always Distinguishable (`*int32` pointer presence). | • **CREATE, PUT & PATCH**: Distinguishable (declare `0`/`false`/`""` via `ForceSendFields` or nullify via `NullFields`). | **Distinguishable**<br>(`nil` vs `&0`) | **Risk**: Unspecified fields inside sub-messages default to `0`/`false`/`""` and overwrite server state under component masks (happens during `Update()` when `updateMask` targets a parent message that contains unchanged fields).<br>**Mitigated?**: Yes.<br>**Mitigation**: Use Leaf-Level Server Inheritance (`populateDesiredWithActualIfComputed`). |
| **4. Enums without `optional`** (`enum`) | • **CREATE (POST)**: Not Distinguishable (`0` emits 0 bytes -> server defaults, or rejects if backend requires an explicit non-unspecified enum).<br>• **REPLACE (PUT)**: Not Distinguishable (`0` emits 0 bytes -> server defaults).<br>• **UPDATE (PATCH)**: Can declare empty value, but not unset (via `updateMask`). | • **CREATE, PUT & PATCH**: Always Distinguishable (`*Enum` pointer presence). | • **CREATE, PUT & PATCH**: Distinguishable (declare value via string name, declare empty via `ForceSendFields`, or unset/nullify via `NullFields`). | **Distinguishable**<br>(`nil` vs `&"ENUM"`) | **Risk**: In `Create()`, omitting an enum defaults to `0` (`_UNSPECIFIED`), causing backend validation failure if the API requires a explicit value; in `Update()`, sending `0` under a parent mask can corrupt server state.<br>**Mitigated?**: Yes.<br>**Mitigation**: In `Create()`, enforce required enum validation or supply defaulting logic; in `Update()`, inherit server enum values via `computedFieldPaths` / `populateDesiredWithActualIfComputed` when omitted, and only target enum fields in `updateMask` when explicitly changed. |

> **Note**: Mitigations are all available in KCC patterns and utility functions, but this does not mean they have already been applied to all greenfield resources across the codebase. We may still need to improve older greenfield resources with these new improvements as gaps are identified.

---

## Architectural Strategy: Minimizing Hardcoded API Default Knowledge in KCC

### 1. The Ideal Solution vs. Practical Reality
In declarative resource management:
- **The Theoretical Ideal**: If KCC knew the exact default value for every single field across every GCP service, it could pre-populate `desired` state with those exact defaults. Then, comparing desired state against live state (`actualPb`) would be straightforward.
- **Why This Is Hard to Achieve**:
  1. **Massive Maintenance Burden**: GCP has hundreds of services with thousands of fields. Hardcoding default values into KCC creates significant code bloat and requires constant maintenance as GCP APIs evolve.
  2. **Dynamic Server Computing**: Many server defaults are dynamic (e.g. calculated based on region, machine size, cluster topology, or auto-scaling algorithms) rather than static constants.
  3. **High Risk of Drift & Desynchronization**: If GCP updates a default value on the backend and KCC still hardcodes the old default, KCC will continuously detect false drift and attempt to revert the resource.

---

### 2. How KCC Compares Successfully Without Hardcoding Defaults

Instead of trying to manually reverse-engineer and hardcode every API default, KCC uses structural patterns and generic utility functions that reduce API-specific knowledge while ensuring accurate drift detection:

| Technique / Utility | What It Handles | Why It Eliminates Hardcoded Knowledge |
| :--- | :--- | :--- |
| **`maps.Equal(desired, actual)`** | Maps / Labels / Environment Variables | Treats `(map)(nil)` (GCP unset/empty wire response) and `map{}` (KRM declared empty) as equal based on content (`len == 0`). Works universally across all Proto2 and Proto3 APIs without needing to know whether the API returns `nil` or `{}`. |
| **`slices.Equal(desired, actual)`** | Ordered Lists / Arrays | Treats `([]T)(nil)` and `[]T{}` as equal based on length and elements. No need to know if the backend omits empty lists. |
| **`proto.Equal(desiredPb, actualPb)`** | Sub-Messages & Dynamic Objects | Built into official Go Protobuf runtime. Automatically respects protobuf wire equivalence (ignoring unexported memory caches, struct pointers, and empty collection headers). |
| **Leaf-Level Server State Inheritance (`populateDesiredWithActualIfComputed`)** | Primitive fields inside sub-messages under parent update masks | For unspecified fields inside sub-messages (like CPU/memory configs), KCC dynamically copies the live value from `actualPb` into `desiredPb`. This preserves server-computed and dynamically assigned values without hardcoding backend defaults. |
| **Field-Presence Guards (`if spec.Field != nil`)** | Top-level optional fields & unmanaged resources | Bypasses drift comparison entirely for fields the user omitted, letting the server manage its own defaults and avoiding false drift on unmanaged state. |

---

## 3. Architectural Summary Cheat Sheet

| Field Kind (in order) | In KRM `spec` | In GCP Client | Drift Detection Method | Action on Drift |
| :--- | :--- | :--- | :--- | :--- |
| **1. Required / Raw Primitive** | `*T` | `T` (value) | Inherit via `populateDesiredWithActualIfComputed`, then compare | Send merged value with component mask |
| **2. Optional Primitive** | `*T` | `*T` (with presence) | `if spec.Field != nil && desiredPb.GetField() != actualPb.GetField()` | Send value with mask |
| **3. Map** | `map[string]string` | `map[string]string` | `if spec.Map != nil && !maps.Equal(desiredPb, actualPb)` | Send map (including `{}`) with mask |
| **4. List** | `[]string` | `[]string` | `if spec.List != nil && !slices.Equal(desiredPb, actualPb)` | Send list (including `[]`) with mask |
| **5. Enum** | `*string` | `Enum` (int32/string) | `if spec.Enum != nil && desiredPb.GetEnum() != actualPb.GetEnum()` | Send enum value with mask |
| **6. Structured Sub-Message** | `*SubStruct` | `*pb.SubMessage` (pointer) | `if spec.Sub != nil && !proto.Equal(desiredPb, actualPb)` | Merge leaf defaults, send merged sub-message |
| **7. Dynamic Object (`Struct`)** | `map[string]any` / raw | `*structpb.Struct` (pointer) | `if spec.Obj != nil && !proto.Equal(desiredPb, actualPb)` | Send dynamic JSON object with mask |
| **8. Wrapper Object (`*Value`)** | `*T` | `*wrapperspb.*Value` (pointer) | `if spec.Val != nil && !proto.Equal(desiredPb, actualPb)` | Send wrapper payload with mask |
| **9. Oneof Group** | `*SubStruct` | `isMessage_Field` (interface) | `if spec.Oneof != nil && !proto.Equal(desiredPb, actualPb)` | Send selected discriminant with mask |
| **10. Discovery Zero-Value** | `*T` | `T` + `ForceSendFields` | `if spec.Field != nil && desired.Field != actual.Field` | Populate `ForceSendFields` |

---

## Appendix: API Contract Formats & Transports

### A.1 Field Presence in Request (Mutation) vs. Response (Read)

An essential distinction in GCP API reconciliation is that **Field Presence behaves differently in Requests than in Responses**:

| Stage | Behavior & Serialization | Presence Outcome |
| :--- | :--- | :--- |
| **1. In a Request (Mutation / PATCH / CREATE)** | • Client actively declares intent.<br>• Even if a Proto3 scalar is raw `0`, including it in `updateMask` instructs GCP to apply `0` rather than ignore it.<br>• In Discovery Doc, `ForceSendFields` forces serialization of `0`/`false`. | **Request is ALWAYS Distinguishable**<br>Field intent is explicitly preserved via `updateMask` or `ForceSendFields`. |
| **2. In a Response (Read / GET / Actual State)** | • Server returns stored state over the wire.<br>• Default zeros (`0`, `false`, `""`) and empty collections (`map{}`, `[]`) are emitted as 0 bytes / omitted from JSON.<br>• Deserialization in Go unmarshals raw `0` and `nil` collections (`len == 0`). | **Response LOSES Empty Presence**<br>Client cannot distinguish unset from empty / zero based on the response struct alone. |

#### Architectural Consequence for KCC:
1. **Never use `reflect.DeepEqual` on responses**: Because a response deserializes empty maps/slices as `nil`, comparing against KRM's allocated `map{}` or `[]` causes false drift detection. Always use `maps.Equal` and `slices.Equal`.
2. **Use Leaf-Level Server Inheritance before sending requests**: Because an omitted KRM scalar translates to raw `0` in Go, KCC must copy existing server values (`actualPb`) onto unmanaged fields before generating the request payload with component updateMasks.

---

### A.2 Which Serializers / Unmarshallers Are Used?

Two primary serializer/unmarshaller systems operate across KCC and GCP:

1. **GCP Protobuf Serializers: gRPC Binary Wire (`google.golang.org/protobuf/proto`) & REST JSON Transcoding (`google.golang.org/protobuf/encoding/protojson`)**:
   * **Where Used**: All communication between KCC Go SDK and GCP APIs (whether over gRPC transport or REST HTTP transport via `NewRESTClient`).
   * **Unified Serialization & Presence Outcome by Type (in order)**:
     - **1. Required / Raw Primitive Scalars (`int32/64`, `uint32/64`, `float/double`, `bool`, `string`, `bytes` without `optional`)**: **NO field presence**. Zero values (`0`, `false`, `""`, `[]byte{}`) emit **0 bytes** on binary wire and are **omitted from JSON**.
     - **2. Optional Primitive Scalars (Proto2 primitives, or Proto3 with `optional`)**: **HAVE field presence**. Serialized with field tag / JSON property if non-nil (`*T`), omitted if `nil`.
     - **3. Maps (`map<Key, Value>`)**: **NO field presence**. When empty (`len == 0`), emits **0 bytes** on binary wire and is **omitted from JSON**.
     - **4. Lists / Repeated Fields (`repeated T`)**: **NO field presence**. When empty (`len == 0`), emits **0 bytes** on binary wire and is **omitted from JSON**.
     - **5. Enums (`enum`)**: **NO field presence** unless marked `optional`. Value `0` (e.g. `_UNSPECIFIED`) emits 0 bytes and is omitted from JSON.
     - **6. Structured Objects / Sub-Messages (`message`)**: **HAVE field presence** (`*SubMessage` is **always a pointer**). Serialized with length tag / JSON `{ ... }` when non-nil (`&SubMessage{}` emits empty length / `{}`); omitted if `nil`.
     - **7. Other Types with Field Presence**:
       - **Dynamic Objects (`google.protobuf.Struct`, `google.protobuf.Value`, `google.protobuf.Any`)**: **HAVE field presence** (sub-messages). Encoded as structured key-value maps or native JSON `{ "key": "value" }`.
       - **Scalar Wrapper Objects (`google.protobuf.Int32Value`, `google.protobuf.StringValue`, etc.)**: **HAVE field presence** (sub-messages). Explicit zero inside wrapper is serialized; omitted if pointer is `nil`.
       - **Oneof Groups (`oneof`)**: **HAVE field presence**. Only the selected field discriminant is serialized; omitted if unselected.
   * **Deserialization Outcome in Go**: Populates Go struct fields (`.pb.go`): omitted network payloads unmarshal into `nil` pointers (for objects, optional primitives, wrappers) or Go default zero-values (`0`, `false`, `""`, `nil` map/slice with `len == 0`).

2. **Kubernetes YAML / Go Standard JSON (`sigs.k8s.io/yaml` and `encoding/json`)**:
   * **Where Used**: Kubernetes API Server and KCC when reading/unmarshalling the user's KRM YAML manifests into Go KRM structs (`desired.Spec`).
   * **Deserialization Outcome by Type (in order)**:
     - **1. Required / Explicit Primitives (`0`, `false`, `""`)**: Non-nil pointers (`&0`, `&false`, `&""`).
     - **2. Optional / Omitted Primitives**: Unmarshalled as `nil` (**Unset / Unmanaged**).
     - **3. Maps**: Omitted -> `nil`; Explicit `{}` -> allocated empty map `map[string]string{}` (`len == 0`, non-nil).
     - **4. Lists**: Omitted -> `nil`; Explicit `[]` -> allocated empty slice `[]string{}` (`len == 0`, non-nil header).
     - **5. Enums**: Omitted -> `nil`; Explicit value -> non-nil pointer string/int.
     - **6. Structured Objects / Sub-Messages**: Omitted -> `nil`; Explicit `{}` -> allocated empty struct pointer `&SubStruct{}` (`!= nil`).
     - **7. Other Types**: Dynamic maps unmarshalled as `map[string]any`; wrappers unmarshalled as pointers.
   * **The Go Field Presence Collision**: 
     - K8s KRM produces an object with **Field Presence** (`map[string]string{}`, `[]string{}`).
     - Proto3 produces an object with **NO Field Presence** (`(map[string]string)(nil)`, `([]string)(nil)`).
     - Standard `reflect.DeepEqual` checks Go pointer/header allocation (field presence) and fails.
     - Content-based comparison functions (`maps.Equal`, `slices.Equal`, `proto.Equal`) ignore pointer presence and compare content lengths, successfully aligning K8s with Proto3.

---

### A.3 Transport Independence (HTTP vs gRPC)

In KCC Direct Controllers, whether a Google Cloud Go Client (`cloud.google.com/go/*`) uses **gRPC** or **REST/HTTP** does **not** change in-memory Go struct behaviors:
* Both transports serialize to and deserialize from the exact same Proto3 Go structs (`.pb.go`).
* For HTTP transport, the Go SDK utilizes `protojson` to strictly map JSON payloads to/from the Protobuf in-memory representation.
* Therefore, field comparison rules depend on the **API Contract Schema (Proto3 vs Discovery Doc)**, NOT on the transport layer (HTTP vs gRPC).

### A.4 The Three API Formats in KCC

Across the KCC codebase (`go.mod`), GCP dependencies fall into three categories:

| Format | Package Import Path | Characteristics & Presence Semantics | Typical Services in KCC |
| :--- | :--- | :--- | :--- |
| **1. Modern Proto3** | `cloud.google.com/go/*` | • Maps/Lists have **no presence** (0 bytes / omitted).<br>• Primitives have presence **only if marked `optional`**.<br>• Sub-messages always have pointer presence. | Composer, AlloyDB, BigQuery, Spanner, PubSub, Vertex AI |
| **2. Legacy Proto2** | *(Rare in new APIs)* | • Explicit field presence across all primitive fields.<br>• Generated as Go pointers (`*T`) with `Has*()` getters. | Internal legacy services |
| **3. Discovery Document** | `google.golang.org/api/*` | • Uses `ForceSendFields` to preserve explicit zero values.<br>• Uses `NullFields` to explicitly clear fields on server. | Compute Engine (v1), Cloud Storage (v1 JSON), Cloud SQL (sqladmin) |

---

### A.5 Comprehensive Behavior Matrix: All Data Types

The table below provides the comprehensive field presence behavior across all data types and serialization formats in KCC:

| Data Type | Modern SDK (Proto3 / REST Client) | Legacy Proto2 | Legacy Discovery SDK | KCC KRM Struct (`spec`) | Comparison & Reconciliation Rule |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **1. Required / Raw Primitives** (`int`, `bool`, `string`) | **Not Distinguishable**<br>(`0` and unset = 0 bytes) | **Distinguishable**<br>(`*int32` pointer) | **Not Distinguishable**<br>(requires `ForceSendFields`) | **Distinguishable**<br>(`nil` vs `&0`) | • Greenfield comparison: `desiredPb.GetX() != actualPb.GetX()`.<br>• Inherit server values via `computedFieldPaths` / `populateDesiredWithActualIfComputed` when field is omitted. |
| **2. Optional Primitives** (`optional int`, `*bool`) | **Distinguishable**<br>(`*int32` pointer) | **Distinguishable**<br>(`*int32` pointer) | **Distinguishable**<br>(via `ForceSendFields`) | **Distinguishable**<br>(`nil` vs `&0`) | • Compare pointer getters: `desiredPb.GetX() != actualPb.GetX()`. |
| **3. Maps** (`map<string, string>`) | **Not Distinguishable**<br>(`nil` and `{}` = 0 bytes) | **Not Distinguishable**<br>(`nil` and `{}` = 0 bytes) | **Not Distinguishable**<br>(server stores 0 entries) | **Distinguishable**<br>(`nil` vs `map{}`) | • **Never use `reflect.DeepEqual`**.<br>• Always use `maps.Equal(desiredPb, actualPb)`. |
| **4. Lists / Repeated** (`repeated T`) | **Not Distinguishable**<br>(`nil` and `[]` = 0 bytes) | **Not Distinguishable**<br>(`nil` and `[]` = 0 bytes) | **Not Distinguishable**<br>(0 elements = empty) | **Distinguishable**<br>(`nil` vs `[]`) | • **Never use `reflect.DeepEqual`**.<br>• Ordered lists: `slices.Equal(desiredPb, actualPb)`.<br>• Unordered sets: `sets.New(...).Equal(...)`. |
| **5. Enums** (`enum`) | **Not Distinguishable**<br>(default `0` = 0 bytes) | **Distinguishable**<br>(`*Enum` pointer) | **Distinguishable**<br>(named string) | **Distinguishable**<br>(`nil` vs `&"ENUM"`) | • Greenfield comparison: `desiredPb.GetEnum() != actualPb.GetEnum()`.<br>• Inherit server defaults via `computedFieldPaths` when omitted. |
| **6. Structured Objects** (`message`) | **Distinguishable**<br>(`*SubMessage` pointer) | **Distinguishable**<br>(`*SubMessage` pointer) | **Distinguishable**<br>(`*SubStruct` pointer) | **Distinguishable**<br>(`nil` vs `&SubStruct{}`) | • Inherit leaf computed values.<br>• Compare via `proto.Equal(desiredPb.GetSub(), actualPb.GetSub())`. |
| **7. Dynamic Objects & Wrappers** (`Struct`, `*Value`) | **Distinguishable**<br>(`*structpb.Struct`, `*wrapperspb.*`) | **Distinguishable**<br>(sub-message pointer) | **Distinguishable**<br>(`map[string]any`) | **Distinguishable**<br>(`nil` vs pointer/map) | • Compare via `proto.Equal`. |