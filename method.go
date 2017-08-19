package aria2cRPC

import (
	"net/http"
	"bytes"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"errors"
)

type Error struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

type Ret struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

func (this RPC) Ping() error {
	this.Data.Method = "aria2.tellActive"
	this.Data.Id = "1"
	data := this.Requset()
	if data.Error != nil {
		errInfo := new(Error)
		byteData, err := json.Marshal(data.Error)
		if err != nil {
			return err
		}
		err = json.Unmarshal(byteData, errInfo)
		if err != nil {
			return err
		}
		return errors.New(errInfo.Message)
	}
	this.Data.Params = []interface{}{this.Token}
	return nil
}

func (this RPC) AddUri(uri []string) (string, error) {
	this.Data.Method = "aria2.addUri"
	this.Data.Params = append(this.Data.Params, uri)
	data := this.Requset()
	if data.Error != nil {
		e, _ := json.Marshal(data.Error)
		return "", errors.New(string(e))
	}
	d, _ := json.Marshal(data.Result)
	return string(d), nil
}

func (this RPC) Requset() Ret {
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprint(os.Stderr, e)
		}
	}()
	tmp, err := json.Marshal(this.Data)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		panic(err)
	}
	resp, err := http.Post(this.URI, this.Header, bytes.NewBuffer(tmp))
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		panic(err)
	}
	data := new(Ret)
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		panic(err)
	}
	return *data
}
