# Follow `SAMPLE_XXX` format to write the content.Add commentMore actions

# Delete `SAMPLE_` and its content before publishing the release note.

# Delete the entire header if no updates.

# Run [mdformat](go/mdformat) before publishing this release notes.  

** This version is not yet released; this document is gathering release notes
for the future release **

*   Special shout-outs to ... for their contributions to this release. TODO:
    list contributors with `git log v1.128.0... | grep Merge | grep from | awk
    '{print $6}' | cut -d '/' -f 1 | sort | uniq`.

## New Beta Resources (Direct Reconciler):

* [`ComputeFutureReservation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computefuturereservation)
    
    * Manage [future reservations](https://cloud.google.com/compute/docs/instances/future-reservations-overview) provisioning
