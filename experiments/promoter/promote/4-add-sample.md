Follow the instruction `5.4 Add samples` from docs/develop-resources/deep-dives/5-releases.md to add sample for <kind> under <service>.

If succeeds, run `go test -v -tags=integration ./config/tests/samples/create -test.run TestAll -run-tests <sample_dir_name>`. 
