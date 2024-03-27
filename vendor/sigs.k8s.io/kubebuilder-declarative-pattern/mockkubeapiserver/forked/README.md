This directory contains code that we copy from various places in the
kubernetes source code that are generally hard to import otherwise.

The intent here is that we demonstrate the need & utility of exporting
these functions in a stable location, and then if/when k/k exports
them we can (reasonably) easily switch, because this is only
test code.

As far as possible, we try to keep the same function names and not
change the code.  We want to reunify in the long-term.