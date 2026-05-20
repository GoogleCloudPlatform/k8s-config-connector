sed -i '/<<<<<<< HEAD/d' pkg/controller/direct/compute/mapper.generated.go
sed -i '/=======/d' pkg/controller/direct/compute/mapper.generated.go
sed -i '/>>>>>>> 078173d901 (Fix ComputeHealthCheck reference resolution by moving SelfLink and Type back to status)/d' pkg/controller/direct/compute/mapper.generated.go
