package aria2cRPC

var (
	Host       string = "127.0.0.1:6800"
	RpcVersion string = "2.0"
	Token      string = ""
	Header     string = "application/json"
)

type Data struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Method  string `json:"method"`
	Params  []interface{} `json:"params"`
}

type RPC struct {
	Host       string
	URI        string
	Header     string
	RpcVersion string
	Token      string
	Data       Data
}

func (this RPC) Init(args ...string) *RPC {
	this.Host = Host
	this.URI = "http://" + Host + "/jsonrpc"
	this.Token = "token:" + this.Token
	this.RpcVersion = RpcVersion
	this.Header = Header
	for index, v := range args {
		switch index {
		case 0:
			this.Token = "token:" + v
		case 1:
			this.Host = v
			this.URI = "http://" + v + "/jsonrpc"
		default:
			break
		}
	}

	this.Data.Jsonrpc = this.RpcVersion
	if len(this.Token) > 0 {
		this.Data.Params = append(this.Data.Params, this.Token)
	}
	return &this
}
