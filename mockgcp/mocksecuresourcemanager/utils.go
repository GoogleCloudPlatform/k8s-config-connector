package mocksecuresourcemanager

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
