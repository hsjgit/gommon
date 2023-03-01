package pbjson

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func PBJsonToStr(m proto.Message) string {
	var buff = &bytes.Buffer{}
	var enc = &jsonpb.Marshaler{EmitDefaults: true}
	enc.Marshal(buff, m)
	return buff.String()
}

func PBJsonToBytes(m proto.Message) []byte {
	var buff = &bytes.Buffer{}
	var enc = &jsonpb.Marshaler{EmitDefaults: true}
	enc.Marshal(buff, m)
	return buff.Bytes()
}
