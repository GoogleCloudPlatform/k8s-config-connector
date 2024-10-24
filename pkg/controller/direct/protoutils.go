package direct

import "google.golang.org/protobuf/proto"

// ProtoClone is a type-safe wrapper around proto.Clone
func ProtoClone[T proto.Message](t T) T {
	return proto.Clone(t).(T)
}
