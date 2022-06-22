package api

import (
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/model"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/pb/protogo"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/shim"
)

// 保存文件
func SaveFile(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	fileName := string(params["file_name"])
	fileHash := string(params["file_hash"])
	time := string(params["time"])

	// 构建结构体
	file := &model.File{
		Name: fileName,
		Hash: fileHash,
		Time: time,
	}

	// 序列化
	fileByte, err := json.Marshal(file)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送事件
	stub.EmitEvent("topic_vx", []string{file.Name, file.Hash})

	// 存储数据
	err = stub.PutStateByte(model.FileKey, file.Hash, fileByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	// 返回结果
	return shim.Success(fileByte)
}

// 查询文件
func QueryFile(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()
	fileHash := string(params["file_hash"])

	if fileHash == "" {
		return shim.Error("文件哈希不能为空")
	}

	result, err := stub.GetStateByte(model.FileKey, fileHash)
	if err != nil {
		return shim.Error(fmt.Sprintf("查询文件出错: %s", err))
	}

	// var file model.File
	// err = json.Unmarshal(result, &file)
	// if err != nil {
	// 	return shim.Error(fmt.Sprintf("反序列化出错: %s", err))
	// }

	return shim.Success(result)
}
