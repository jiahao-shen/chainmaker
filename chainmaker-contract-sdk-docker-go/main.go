/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/api"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/pb/protogo"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/shim"
)

type FactContract struct {
}

func (f *FactContract) InitContract(stub shim.CMStubInterface) protogo.Response {
	return shim.Success([]byte("Init Success"))
}

func (f *FactContract) InvokeContract(stub shim.CMStubInterface) protogo.Response {

	// 获取参数
	method := string(stub.GetArgs()["method"])

	switch method {
	case "saveFile":
		return api.SaveFile(stub)
	case "queryFile":
		return api.QueryFile(stub)
	default:
		return shim.Error("invalid method")
	}

}

func main() {

	err := shim.Start(new(FactContract))
	if err != nil {
		log.Fatal(err)
	}
}
