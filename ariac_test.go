package aria2cRPC

import (
	"testing"
	"fmt"
)

var (
	rpc *RPC = RPC{}.Init("xiapian", "raspberry.lich.moe:6800")
)

func TestRPC_Ping(t *testing.T) {
	err := rpc.Ping()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestRPC_AddUri(t *testing.T) {
	urls := []string{
		"magnet:?xt=urn:btih:699c59f13bf572f9917d8738f638a9f91e08c0d9",
	}
	s, err := rpc.AddUri(urls)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	fmt.Println(s)
}