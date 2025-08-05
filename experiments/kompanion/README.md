# Kompanion

Experimental KCC companion tool to help troubleshoot, analyze and gather data about KCC.

# Usage

## Build from source

```
# Assumes pwd is <REPO_ROOT>/experiments/kompanion
$ mkdir bin
$ GOWORK=off go build -o bin/kompanion
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

# Light Roadmap

* [ ] Debug/ audit logs for the tool itself
