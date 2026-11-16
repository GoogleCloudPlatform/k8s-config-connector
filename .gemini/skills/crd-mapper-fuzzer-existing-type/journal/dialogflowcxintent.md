# DialogflowCXIntent Journal

## Observations
- Resource: `DialogflowCXIntent`
- Group: `dialogflowcx.cnrm.cloud.google.com`
- Service Proto: `google.cloud.dialogflow.cx.v3`
- Version: `v1alpha1`

## Key Learnings
1. **Directory Location**: Created new service directory `apis/dialogflowcx/v1alpha1` for group `dialogflowcx.cnrm.cloud.google.com`.
2. **OpenAPI Integer Formatting**: Baseline CRD specified `type: integer` without format for `priority`, `repeatCount`, and `observedGeneration`. Used Go `*int` in `DialogflowCXIntentSpec`, `IntentTrainingPhrases`, and `DialogflowCXIntentStatus` to preserve `type: integer` without emitting `format: int32`, satisfying `TestCRDBackwardCompatibility`.
3. **Custom Field Mappers**: Added handcoded mappers in `pkg/controller/direct/dialogflowcx/mappers.go` to convert between KRM `*int` and proto `int32` for `Priority` and `RepeatCount`.
4. **Field Name Conventions**: Renamed struct fields `Id` -> `ID` and `ParameterId` -> `ParameterID` to align with Go conventions while preserving exact JSON tags (`json:"id"`, `json:"parameterId,omitempty"`), allowing generated mappers to match proto fields automatically.
