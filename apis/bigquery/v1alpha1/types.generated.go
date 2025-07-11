// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

// +kcc:proto=google.cloud.bigquery.v2.Routine
type Routine struct {
	// Output only. A hash of this resource.
	Etag *string `json:"etag,omitempty"`

	// Required. Reference describing the ID of this routine.
	RoutineReference *RoutineReference `json:"routineReference,omitempty"`

	// Required. The type of routine.
	RoutineType *string `json:"routineType,omitempty"`

	// Output only. The time when this routine was created, in milliseconds since
	//  the epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The time when this routine was last modified, in milliseconds
	//  since the epoch.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Optional. Defaults to "SQL" if remote_function_options field is absent, not
	//  set otherwise.
	Language *string `json:"language,omitempty"`

	// Optional.
	Arguments []Routine_Argument `json:"arguments,omitempty"`

	// Optional if language = "SQL"; required otherwise.
	//  Cannot be set if routine_type = "TABLE_VALUED_FUNCTION".
	//
	//  If absent, the return type is inferred from definition_body at query time
	//  in each query that references this routine. If present, then the evaluated
	//  result will be cast to the specified returned type at query time.
	//
	//  For example, for the functions created with the following statements:
	//
	//  * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);`
	//
	//  * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));`
	//
	//  * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));`
	//
	//  The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and
	//  is absent for `Increment` (inferred as FLOAT64 at query time).
	//
	//  Suppose the function `Add` is replaced by
	//    `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);`
	//
	//  Then the inferred return type of `Increment` is automatically changed to
	//  INT64 at query time, while the return type of `Decrement` remains FLOAT64.
	ReturnType *StandardSqlDataType `json:"returnType,omitempty"`

	// Optional. Can be set only if routine_type = "TABLE_VALUED_FUNCTION".
	//
	//  If absent, the return table type is inferred from definition_body at query
	//  time in each query that references this routine. If present, then the
	//  columns in the evaluated table result will be cast to match the column
	//  types specified in return table type, at query time.
	ReturnTableType *StandardSqlTableType `json:"returnTableType,omitempty"`

	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	//  imported JAVASCRIPT libraries.
	ImportedLibraries []string `json:"importedLibraries,omitempty"`

	// Required. The body of the routine.
	//
	//  For functions, this is the expression in the AS clause.
	//
	//  If language=SQL, it is the substring inside (but excluding) the
	//  parentheses. For example, for the function created with the following
	//  statement:
	//
	//  `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))`
	//
	//  The definition_body is `concat(x, "\n", y)` (\n is not replaced with
	//  linebreak).
	//
	//  If language=JAVASCRIPT, it is the evaluated string in the AS clause.
	//  For example, for the function created with the following statement:
	//
	//  `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'`
	//
	//  The definition_body is
	//
	//  `return "\n";\n`
	//
	//  Note that both \n are replaced with linebreaks.
	DefinitionBody *string `json:"definitionBody,omitempty"`

	// Optional. The description of the routine, if defined.
	Description *string `json:"description,omitempty"`

	// Optional. The determinism level of the JavaScript UDF, if defined.
	DeterminismLevel *string `json:"determinismLevel,omitempty"`

	// Optional. The security mode of the routine, if defined. If not defined, the
	//  security mode is automatically determined from the routine's configuration.
	SecurityMode *string `json:"securityMode,omitempty"`

	// Optional. Use this option to catch many common errors. Error checking is
	//  not exhaustive, and successfully creating a procedure doesn't guarantee
	//  that the procedure will successfully execute at runtime. If `strictMode` is
	//  set to `TRUE`, the procedure body is further checked for errors such as
	//  non-existent tables or columns. The `CREATE PROCEDURE` statement fails if
	//  the body fails any of these checks.
	//
	//  If `strictMode` is set to `FALSE`, the procedure body is checked only for
	//  syntax. For procedures that invoke themselves recursively, specify
	//  `strictMode=FALSE` to avoid non-existent procedure errors during
	//  validation.
	//
	//  Default value is `TRUE`.
	StrictMode *BoolValue `json:"strictMode,omitempty"`

	// Optional. Remote function specific options.
	RemoteFunctionOptions *Routine_RemoteFunctionOptions `json:"remoteFunctionOptions,omitempty"`

	// Optional. Spark specific options.
	SparkOptions *SparkOptions `json:"sparkOptions,omitempty"`

	// Optional. If set to `DATA_MASKING`, the function is validated and made
	//  available as a masking function. For more information, see [Create custom
	//  masking
	//  routines](https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask).
	DataGovernanceType *string `json:"dataGovernanceType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Routine.Argument
type Routine_Argument struct {
	// Optional. The name of this argument. Can be absent for function return
	//  argument.
	Name *string `json:"name,omitempty"`

	// Optional. Defaults to FIXED_TYPE.
	ArgumentKind *string `json:"argumentKind,omitempty"`

	// Optional. Specifies whether the argument is input or output.
	//  Can be set for procedures only.
	Mode *string `json:"mode,omitempty"`

	// Required unless argument_kind = ANY_TYPE.
	DataType *StandardSqlDataType `json:"dataType,omitempty"`

	// Optional. Whether the argument is an aggregate function parameter.
	//  Must be Unset for routine types other than AGGREGATE_FUNCTION.
	//  For AGGREGATE_FUNCTION, if set to false, it is equivalent to adding "NOT
	//  AGGREGATE" clause in DDL; Otherwise, it is equivalent to omitting "NOT
	//  AGGREGATE" clause in DDL.
	IsAggregate *BoolValue `json:"isAggregate,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Routine.RemoteFunctionOptions
type Routine_RemoteFunctionOptions struct {
	// Endpoint of the user-provided remote service, e.g.
	//  ```https://us-east1-my_gcf_project.cloudfunctions.net/remote_add```
	Endpoint *string `json:"endpoint,omitempty"`

	// Fully qualified name of the user-provided connection object which holds
	//  the authentication information to send requests to the remote service.
	//  Format:
	//  ```"projects/{projectId}/locations/{locationId}/connections/{connectionId}"```
	Connection *string `json:"connection,omitempty"`

	// User-defined context as a set of key/value pairs, which will be sent as
	//  function invocation context together with batched arguments in the
	//  requests to the remote service. The total number of bytes of keys and
	//  values must be less than 8KB.
	UserDefinedContext map[string]string `json:"userDefinedContext,omitempty"`

	// Max number of rows in each batch sent to the remote service.
	//  If absent or if 0, BigQuery dynamically decides the number of rows in a
	//  batch.
	MaxBatchingRows *int64 `json:"maxBatchingRows,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RoutineReference
type RoutineReference struct {
	// Required. The ID of the project containing this routine.
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The ID of the dataset containing this routine.
	DatasetID *string `json:"datasetID,omitempty"`

	// Required. The ID of the routine. The ID must contain only
	//  letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	//  length is 256 characters.
	RoutineID *string `json:"routineID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.SparkOptions
type SparkOptions struct {
	// Fully qualified name of the user-provided Spark connection object. Format:
	//  ```"projects/{project_id}/locations/{location_id}/connections/{connection_id}"```
	Connection *string `json:"connection,omitempty"`

	// Runtime version. If not specified, the default runtime version is used.
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`

	// Custom container image for the runtime environment.
	ContainerImage *string `json:"containerImage,omitempty"`

	// Configuration properties as a set of key/value pairs, which will be passed
	//  on to the Spark application. For more information, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html) and the
	//  [procedure option
	//  list](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#procedure_option_list).
	Properties map[string]string `json:"properties,omitempty"`

	// The main file/jar URI of the Spark application. Exactly one of the
	//  definition_body field and the main_file_uri field must be set for Python.
	//  Exactly one of main_class and main_file_uri field
	//  should be set for Java/Scala language type.
	MainFileUri *string `json:"mainFileUri,omitempty"`

	// Python files to be placed on the PYTHONPATH for PySpark application.
	//  Supported file types: `.py`, `.egg`, and `.zip`. For more information
	//  about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	PyFileUris []string `json:"pyFileUris,omitempty"`

	// JARs to include on the driver and executor CLASSPATH.
	//  For more information about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	JarUris []string `json:"jarUris,omitempty"`

	// Files to be placed in the working directory of each executor.
	//  For more information about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	FileUris []string `json:"fileUris,omitempty"`

	// Archive files to be extracted into the working directory of each executor.
	//  For more information about Apache Spark, see
	//  [Apache Spark](https://spark.apache.org/docs/latest/index.html).
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// The fully qualified name of a class in jar_uris, for example,
	//  com.example.wordcount. Exactly one of main_class and main_jar_uri field
	//   should be set for Java/Scala language type.
	MainClass *string `json:"mainClass,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlDataType
type StandardSqlDataType struct {
	// Required. The top level type of this field.
	//  Can be any GoogleSQL data type (e.g., "INT64", "DATE", "ARRAY").
	TypeKind *string `json:"typeKind,omitempty"`

	// The type of the array's elements, if type_kind = "ARRAY".
	ArrayElementType *StandardSqlDataType `json:"arrayElementType,omitempty"`

	// The fields of this struct, in order, if type_kind = "STRUCT".
	StructType *StandardSqlStructType `json:"structType,omitempty"`

	// The type of the range's elements, if type_kind = "RANGE".
	RangeElementType *StandardSqlDataType `json:"rangeElementType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlField
type StandardSqlField struct {
	// Optional. The name of this field. Can be absent for struct fields.
	Name *string `json:"name,omitempty"`

	// Optional. The type of this parameter. Absent if not explicitly
	//  specified (e.g., CREATE FUNCTION statement can omit the return type;
	//  in this case the output parameter does not have this "type" field).
	Type *StandardSqlDataType `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlStructType
type StandardSqlStructType struct {
	// Fields within the struct.
	Fields []StandardSqlField `json:"fields,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.StandardSqlTableType
type StandardSqlTableType struct {
	// The columns in this table type
	Columns []StandardSqlField `json:"columns,omitempty"`
}

// +kcc:proto=google.protobuf.BoolValue
type BoolValue struct {
	// The bool value.
	Value *bool `json:"value,omitempty"`
}
