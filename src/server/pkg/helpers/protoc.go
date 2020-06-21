package helpers

import "google.golang.org/protobuf/proto"

func ProtoMarshal(v proto.Message) ([]byte, error) {
	return proto.Marshal(v)
}

func ProtoUnmarshal(data []byte, v proto.Message) error {
	return proto.Unmarshal(data, v)
}
