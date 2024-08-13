# Using Samples

## Prerequisites

A Config Controller cluster with Compositions installed.

## Personas/Roles

Platform Admin:
- Responsible for creating compositions as well as the Facade CRDs
- Responsible for the AppTeam Facade CRs
- Assumed to work in `default` namespace when installing Compositions.

Team/App Owner/Admin:
- Responsible for creating Facade CRs.
- Assumed to work in their own namespace.

## AppTeam: Setting up New Teams

This recipe creates a GCP project for the team as well as sets up KCC in
namespace mode to manage the team project.

## SQLHA: Create CloudSQL in HA Mode

This recipe sets up a CloudSQL instance in your project in HA mode. If you'd
like to create the instance in a tenant namespace, do the AppTeam sample first
to create one.
