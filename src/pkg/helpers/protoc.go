package helpers

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/protobuf/proto"
)

func ProtoMarshal(v proto.Message) ([]byte, error) {
	return proto.Marshal(v)
}

func ProtoUnmarshal(data []byte, v proto.Message) error {
	return proto.Unmarshal(data, v)
}

func ProtoMarshalJSON(v proto.Message) ([]byte, error) {
	jp := runtime.JSONPb{}
	return jp.Marshal(v)
}

func ProtoUnmarshalJSON(data []byte, v proto.Message) error {
	jp := runtime.JSONPb{}
	return jp.Unmarshal(data, v)
}
