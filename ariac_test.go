package aria2cRPC

import (
	"testing"
)

func TestRPC_Ping(t *testing.T) {
	rpc := RPC{}.Init("token1")
	err := rpc.Ping()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}
