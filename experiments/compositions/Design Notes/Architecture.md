# Architecture

## Personas:
- Application Owner / AppTeamOperator
- Platform Admin 

## Use Cases

1. Platform Admin creates a DRY CRD (CRD_T, Composition) and its corresponding Config CRD (CRD_V, platform.company.com/v1/Database) representing a common component that would be consumed by AppTeamOperator for her team.
2. AppTeamOperator creates an instance of the Config CRD (CRD_V, platform.company.com/v1/Database) to parameterize the DRY CRD (CRD_T, Composition) and deploy it.
3. AppTeamOperator creates another instance of the Config (CRD_V) to parameterize the DRY CRD (CRD_T) and deploy it in a second namespace for application 2
4. The platform must expand the DRY config (manifests) based on CRD_V parameters and record the expanded manifests and apply it to the cluster for materialization by downstream components like KCC, etc.
5. The platform shall provide support to gate the expanded manifests with useful filters like Quota, Approval etc before they are applied to the cluster.

## Requirements

1. Allow  Platform admin to encapsulate best practices with locked down deployment choices while allowing meaningful control for AppTeamOperator
2. Handle cloud resource dependencies and sequencing when deploying the config on behalf of AppTeamOperator.
3. Handle dynamic parameterization with values from a cloud resources for subsequent dependent cloud resources.

## K8s API/Data Model

1. DRY CRD (CRD_T, Composition)
2. Config CRD (CRD_V, platform.company.com/v1/Database)
3. Plan CRD (CRD_P, platform.company.com/v1/Plan)

Refer [API Design](API Design.md) for more detailed discussion of the API and data model.

## Manifest Sources

Refer [Manifest Storage](Manifest Storage.md) for choices on where the manifests can be stored. The goal is to provide a way for other systems to interface with the platform and control its behavior. We are envisioning external systems such as Pantheon, UI, Quota checkers, Linters, Approval workflows etc to be able to interface with the platform via the manifest storage  as well as the K8s API.

## POC Components

                            +-----------+
Admin ->   Composition ---> |           | ---> [EXPANDER JOB] --> Update CRD_V
AppTeam -> CRD_V       ---> |Composition| ---> create CRD_P
                            |Controller |
                            |           |
                            +-----------+
                 +-----------+
CRD_V/CRD_P ---> |           | ---> [APPLIER JOB] --> apply to cluster ---> [KCC] --> GCP API
                 |Applier    |
                 |Controller |
                 |           |
                 +-----------+

For POC we are using `Compositions` as the CRD_T (dry CRD).
We will have a Composition controller that will reconcile Composition instance. It will also watch the CRD_V (Config/Value CRDs) that correspond to each instance of Composition.

For POC we could club the Composition and Applier functionality into a single controller.