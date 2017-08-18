package aria2cRPC

import "fmt"

var (
	Host       string = "http://127.0.0.1:6800/jsonrpc"
	RpcVersion string = "2.0"
	Token      string = ""
)

type Data struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Method  string `json:"method"`
	Params  []string `json:"params"`
}

type RPC struct {
	Host       string
	RpcVersion string
	Token      string
	Data       Data
}

func (this RPC) Init(args ...string) {
	this.Host = Host
	this.Token = Token
	this.RpcVersion = RpcVersion
	for index, v := range args {
		switch index {
		case 0:
			this.Host = "http://" + v + "/jsonrpc"
		case 1:
			this.Token = v
		case 2:
			this.RpcVersion = v
		default:
			break
		}
	}
	this.Data.Jsonrpc = this.RpcVersion
	this.Data.Params = []string{}
	fmt.Println(this)
	if len(this.Token) > 0 {
		this.Data.Params = append(this.Data.Params, "token:"+this.Token)
	}
	fmt.Println(this.Data)
}
