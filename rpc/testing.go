package rpc

import (
	"reflect"

	bin "github.com/gagliardetto/binary"
)

func newTestClient(url string) *Client {
	client := NewClient(url)
	client.requestIDGenerator = func() int {
		return 0
	}

	return client
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Ptr && rv.IsNil()
}

func puint64(v uint64) *bin.Uint64 {
	t := bin.Uint64(v)
	return &t
}
