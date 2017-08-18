package aria2cRPC

import (
	"testing"
)

func TestRPC_Ping(t *testing.T) {
	RPC{}.Init()
	RPC{}.Ping()
}
