
# Feedback from Demos

## Demo Session #1: 3/15/2024

|No |Feedback   |Followup Notes |
|-|---|---|
|1| Explore grpc based semi-persistent expanders (vs volume mount)| Explore part of POC soon |
|2| AppTeamOperator CRs in a separate cluster vs separate namespace | Explore as part of PuP |
|3 | Can AWS add support for S3 | Yes. Multi-Cloud can cover it |
|4| AppTeamOperator creates the cloudsql which in turn creates gcp resources.  How do we control what gcp resources are allowed on facade's behalf.  Auto generate RBAC that includes only those present in the template. | Document now and commit to security review. |

## Demo Session #2: 3/21/2024

FindInMap : Ability to map facade inputs to template values
Validations for facade input fields at which level (crd, expander, gatekeeper)
Default values for facade input fields  (crd, expander, gatekeeper)
Troubleshooting
  - make it easy to detect who needs to act Platform admin or AppTeamOperator
  - who breaks glass and what is allowed (can appteam mutate generated resources ? )
  - who carries the pager (admin & appteam OR appteam OR admin)
