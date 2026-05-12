sed -i '/<<<<<<< HEAD/d' pkg/controller/direct/networkmanagement/connectivitytest_fuzzer.go
sed -i '/=======/d' pkg/controller/direct/networkmanagement/connectivitytest_fuzzer.go
sed -i '/>>>>>>> eaae97e075 (Fix ConnectivityTest fuzzer test failure by ignoring newly added upstream proto field)/d' pkg/controller/direct/networkmanagement/connectivitytest_fuzzer.go
