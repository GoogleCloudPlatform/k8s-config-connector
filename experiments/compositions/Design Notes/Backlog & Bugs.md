
1. [bug] When composition is updates, reconcile triggers a new cloudsql reconciler to be constructed.
2. [backlog] Add status to Composition, capture any bugs in parsing the template.
3. [backlog] Add status to InputAPI CRD. c:Ready, Stage[expander name]{fetched, expanded, applied}
4. [backlog] Add watches for applied resources and retrigger plan.
5. [backlog] Move applier part to plan-reconciler
6. [backlog] Explore grpc based expander pattern
7. [backlog] When Composition is deleted, let cloudsql reconciler error out. Dont delete cloudsql.
8. [backlog] When CloudSQL/InputAPI is deleted, delete the corresponding plan
9. [backlog] When plan cr is deleted, ensure kdp deletes applied object