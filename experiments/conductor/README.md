# Conductor LLM Code Automator

Conductor is a tool for automating the maintenance of Config Connector (KCC).
It is a command line tool which is largely driven by the metadata it is passed.
The complete metadata is curently held in branches-all.yaml.
We recommend you copy the metadata for the reosurces you want to work on into a separate file.
An example of this is the branches.yaml file.
Conductor will then work on that subset of resources.
We recommend when working on KCC that you point conductor at a separate git copy than the one in which conductor is running in.
This is especially true if you are making changes to the tool.
You control this through the --branch-repo flage.

## Usage

### Build
```
make build
```

### CLI Help
To get the basic help try the following
```
./bin/conductor runner
```
To get the list of supported commands
```
./bin/conductor runner --branch-repo=../.. --logging-dir=./logs
```

### Step descriptions
Descriptions of the original manual steps can be found [here](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/experiments/conductor/cmd/runner/scripts)

The code for the indivudal commands can be found [here](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/experiments/conductor/cmd/runner)

The command help frequently hints on the command type which can help identify the idnividual file.

## Contributing to Config Connector

Please refer to our [contribution guide](CONTRIBUTING.md) for more details.
