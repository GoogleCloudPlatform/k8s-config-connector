Fuzz is required for Alpha, we assume fuzz pass.  
<!-- 
(TODO) The current fuzz check runs all Fuzz tests, and takes too long (~20 min). We want to run a specific fuzz (the original approach). 
Run fuzz test by service
Run `./dev/tasks/find-missing-fields` and try to fix the error, you should look at the `*_fuzz.go` file under `pkg/controller/direct/<SERVICE>`. If you cannot find the file path, try `ls pkg/controller/direct/<SERVICE>/*_fuzzer.go` -->