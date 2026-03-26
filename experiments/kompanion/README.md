# Kompanion

Experimental KCC companion tool to help troubleshoot, analyze and gather data about KCC.

# Usage

## Build from source

```
# Assumes pwd is <REPO_ROOT>/experiments/kompanion
$ GOWORK=off go build -o kompanion
```

## Export

Export function may take minutes to run as the tool searches all the api resources that are KCC related in every namespace. Filters can be applied with the supported flags: `kompanion export -h `.

```
	# export KCC resources across all namespaces, excludes \"kube\" namespaces by default
	kompanion export

	# exclude certain namespace prefixes
	kompanion export --exclude-namespaces=kube --exclude-namespaces=my-team

	# target only specific namespace prefixes
	kompanion export --target-namespaces=my-team

	# target only specific namespace prefixes AND specific object prefixes
	kompanion export --target-namespaces=my-team --target-objects=logging
```

The command will generate a timestamped report `tar.gz` file to use as a snapshot.

## MCP

Start a Model Context Protocol (MCP) server to interact with KCC resources in a cluster.

```
    # Start the MCP server using default kubeconfig
    kompanion mcp

    # Start the MCP server with a specific kubeconfig and context
    kompanion mcp --kubeconfig /path/to/config --context my-context
```

This enables AI IDEs (like Cursor) and assistants (like Claude Desktop) to:
* Get KCC CRD schemas.
* List KCC resources.
* Describe KCC resources and their status.
* Apply KCC resource YAML via Server-Side Apply.

# Light Roadmap

* [ ] Debug/ audit logs for the tool itself