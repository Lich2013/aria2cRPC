package aria2cRPC

import "fmt"

func (this RPC) Ping() {
	this.Data.Method = "aria2.tellActive"
	fmt.Println(this.Data)
}
