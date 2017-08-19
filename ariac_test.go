package aria2cRPC

import (
	"testing"
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
		"magnet:?xt=urn:btih:44c68ca746f65291bc4139429e5c13d460323483",
	}
	rpc.AddUri(urls)
}