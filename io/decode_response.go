package io


import (
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
)

func DecodeMessageFromResponse(r *io.ReadCloser, m proto.Message) error {
	b, err := ioutil.ReadAll(*r)
	if err != nil {
		return err
	}
	return proto.Unmarshal(b, m)
}
