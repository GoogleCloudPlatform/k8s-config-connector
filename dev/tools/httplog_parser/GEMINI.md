# httplog_parser

`httplog_parser` is a tool designed to make `_http.log` diffs human-readable.
It parses the custom `_http.log` format used in Config Connector e2e tests, which captures HTTP traffic (Request/Response pairs).

## Purpose

When `_http.log` golden files change, looking at a standard text diff is often confusing because:
1.  JSON formatting changes can create large noise.
2.  It's hard to correlate a change in a response body back to the request that triggered it.
3.  Timestamps and IDs often change and need normalization.

This tool parses two `_http.log` files (or a file and its git committed version), aligns the requests, and produces a summary of specific JSON fields that have changed.

## Usage

```bash
go run dev/tools/httplog_parser/main.go [file1] [file2]
```

If no arguments are provided, it defaults to comparing the `_http.log` in the current directory against `HEAD` (the last committed version).

If one argument is provided, it assumes it is a file path and compares it against `HEAD`.

If two arguments are provided, it compares `file1` (old) vs `file2` (new).

## Output Format

The output focuses on semantic differences in the JSON bodies of the responses.

```
Request: GET https://example.com/api/v1/resource
  .metadata.generation: 2 => 3
  .status.conditions[0].lastTransitionTime: "2024-01-01T00:00:00Z" => "2024-01-02T00:00:00Z"
```

## Implementation Details

The tool:
1.  Parses the `_http.log` format (headers, body, separator `---`).
2.  Aligns entries between the two log files.
3.  For each entry, compares the Response Body as JSON.
4.  Walks the JSON tree to find differences and reports them with their "dot path" (e.g., `.items[0].metadata.name`).
