---
name: kcc-greenfield-triage
description: Triage a KCC greenfield issue - check whether a new GCP resource kind can be implemented now as a direct resource (API access and allowlisting, quota, parent and dependencies in KCC, lifecycle, state on master, the factory-graph start block) and decide ready, blocked or unsure, with what clears each blocker.
---

# KCC Greenfield Triage

Check whether the greenfield kind of an issue can be implemented now -
before (or while) its issue graph works on it - and decide **ready**,
**blocked** or **unsure**. Greenfield work adds a GCP resource kind to KCC as
a direct resource, step by step (prepopulate, types-judgement, controller,
mockgcp, mockgcp-alignment, release).

Triage changes nothing: no files (but a result file you are asked to write),
no git writes, nothing posted to GitHub, no cloud resources created, changed
or deleted. Commands below are read-only.

Placeholders: `<issue_number>` is the issue's number, `<gcp_project>` the
test project (skip the project checks when there is none).

## Why

Greenfield graphs got stuck, or built the wrong thing, for reasons a
check up front finds: the API needs quota (#15037) or project
allowlisting (#14955), the parent or a dependent resource is not in KCC
yet (#13259, #15221), the resource does not follow standard CRUD
(#15221), or real GCP rejects the create although the API is GA - a
feature turned off for every model, found by an earlier attempt the
triage did not read (#12690, PR #15203).

## Checks

1. Find the kind (the issue's `factory-graph` block names it as
   `params.kind`, else its title) and its service and GCP resource:
   read its comments (`gh issue view <issue_number> --comments`)
   and its sub-issues, and find its proto message, service and
   RPCs under `.build/third_party/googleapis` (or the googleapis
   repository online).
2. Check each of these, and note what you found and how:
   - **API access**: is the API generally available, or does it
     need project allowlisting, early access or a private preview
     (its documentation, release notes, the discovery document)?
     Is it enabled, or can it be enabled, in the test project
     (`gcloud services list --enabled --project <gcp_project>`
     when a project is set; read-only)?
     **Usable, not only documented:** GA in the docs does not mean a
     create succeeds today. Read the release notes and deprecation
     notices for the resource and for anything a create request
     needs from outside KCC (a model, a machine type, a region, a
     feature): a feature turned off, deprecated or limited to some
     models or regions blocks it the same way. AIPlatformCachedContent
     (#12690) was triaged ready while real GCP rejected every create
     ("does not support explicit context caching") for every model.
   - **Earlier attempts**: what happened when this kind was tried
     before. Read, and quote what you find:
     - its sub-issues and every PR for the kind, open or closed
       (`gh pr list --state all --search "<Kind> in:title"`) - their
       descriptions ("Real GCP testing" notes) and comments;
     - `.gemini/journals/<service>.md` for the kind;
     - the test project's Cloud Audit Logs for the kind's create
       method over the last 30 days (`gcloud logging read
       'logName:"cloudaudit.googleapis.com" AND
       protoPayload.methodName:"Create<Resource>"' --freshness=30d
       --project <gcp_project>`; read-only): failed creates and their
       error messages.
     A real-GCP failure there (an API error on create, a feature not
     supported, permission or quota) means blocked - or unsure when
     it may have been fixed since - never ready, until someone shows
     it now works. A PR "underway" is not evidence that it works:
     read what it found. Siblings recorded against real GCP say
     nothing about this kind.
   - **Quota**: does creating one - or creating its parent and
     dependencies, which its E2E fixtures create too - need quota
     a test project usually lacks (dedicated hardware, nodes,
     clusters, private clouds - VMware Engine, say)? Check it the
     same way every time, so kinds that share a parent get the
     same answer:
     - Read every quota that applies to what the fixtures create,
       regional and global, for the region the service's fixtures
       use (`gcloud alpha services quota list --service=<api>
       --consumer=projects/<gcp_project>`). The
       strictest one decides: an effective limit of 0 on any quota
       the creation needs blocks it, whatever another quota allows.
     - Look at the other greenfield issues of the same service or
       parent kind (`gh issue list --label workflow/greenfield
       --search "<service> in:title" --state all`) and their
       comments: if one is stopped for quota on the same parent,
       this one is blocked the same way unless you find why not -
       and then say why.
     - Say whether the parent's own E2E fixtures were ever recorded
       against real GCP (an `_http.log` beside them, under
       `pkg/test/resourcefixture/testdata/`): recorded, its
       creation worked in a test project once.
     Name the quotas you read, their limits and the region.
   - **Parent and dependencies in KCC**: the parent resource and
     every resource the kind's references point to must already be
     a KCC kind with a controller (direct or legacy), so that E2E
     fixtures can create them in `dependencies.yaml`. Search
     `apis/`, `config/crds/` and `pkg/controller/` for each. Name
     any that is missing, or that has types but no controller.
   - **Lifecycle**: does the resource follow standard
     Create/Get/Update/Delete (and List) RPCs? A singleton, a
     "SetUp"/"Provision" RPC instead of Create, no Delete, or
     deletion through another resource needs a maintainer's design.
   - **State on master**: are its types, identity and judgement
     queue there already (and resolved)? Is it implemented already,
     or blocked in `dev/tasks/build-release-bundle`'s BLOCKED_KINDS?
   - **Start block**: the issue's description must hold a
     `factory-graph` block with `graph: kcc-greenfield` and
     `params.kind` set to this kind: labelling it overseer/graph
     starts nothing without one. Missing or wrong, the outcome is
     unsure, and "To check:" gives the block to add, written out.
3. Decide, from evidence, not from documentation alone:
   - **ready**: nothing above blocks it, and nothing found in earlier
     attempts says real GCP rejects it;
   - **blocked**: at least one thing blocks it - which, and what a
     maintainer must do (request quota, ask for allowlisting,
     implement the parent kind first - name it - or decide the
     modelling);
   - **unsure**: you could not tell; say what a person must check.
     Unsure stops it too: nothing starts on a kind nobody has
     cleared.

## Report

The report starts with one line - "**Ready to implement.**",
"**Blocked: recommend overseer/stop.**" or "**Needs a person to
check.**" - then, for each check (API access, earlier attempts,
quota, parent and dependencies, lifecycle, state on master, start
block), a bullet with
what you found and how. When blocked, end with "To unblock:" and what a
maintainer must do, each item its own bullet; name the missing KCC
kinds. When unsure, end with "To check:" and what a person must
check. When ready, end with "To start: add `overseer/graph`."

Labels: none when ready - triage does not start the graph, a maintainer
does, by adding `overseer/graph`; `overseer/stop` when blocked or unsure.
No other labels.
