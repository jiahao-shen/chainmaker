package api

import (
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/model"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/pb/protogo"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/shim"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/utils"
)

func GenerateQR(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	serialNumber := string(params["serial_number"])
	generatorID := string(params["generator_id"])
	ctidHash := string(params["ctid_hash"])
	generateTime := string(params["generate_time"])
	validityPeriod := string(params["validity_period"])
	cityID := string(params["city_id"])
	appID := string(params["app_id"])
	sceneID := string(params["scene_id"])

	// TODO: 数据逻辑处理

	log := &model.GenerateQRLog{
		SerialNumber:   serialNumber,
		GeneratorID:    generatorID,
		CTIDHash:       ctidHash,
		GenerateTime:   generateTime,
		ValidityPeriod: validityPeriod,
		CityID:         cityID,
		APPID:          appID,
		SceneID:        sceneID,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("topic_vx", []string{serialNumber, generatorID, ctidHash, generateTime, validityPeriod, cityID, appID, sceneID})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, utils.GetSHA256String([]string{serialNumber}), logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func VerifyQR(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	serialNumber := string(params["serial_number"])
	verifierID := string(params["verifier_id"])
	cCodeContentHash := string(params["c_code_content_hash"])
	cityID := string(params["city_id"])
	appID := string(params["app_id"])
	verifyTime := string(params["verify_time"])
	generatorID := string(params["generator_id"])

	// TODO: 数据逻辑处理

	log := &model.VerifyQRLog{
		SerialNumber:     serialNumber,
		VerifierID:       verifierID,
		CCodeContentHash: cCodeContentHash,
		CityID:           cityID,
		APPID:            appID,
		VerifyTime:       verifyTime,
		GeneratorID:      generatorID,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("topic_vx", []string{serialNumber, verifierID, cCodeContentHash, cityID, appID, verifyTime, generatorID})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, utils.GetSHA256String([]string{serialNumber}), logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// TODO:
func AuthorizeUser(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	authorizerID := string(params["authorizer_id"])
	authorizeeID := string(params["authorizee_id"])
	authorizeTime := string(params["authorize_time"])
	authorizeContent := string(params["authorize_content"])
	authorizeStatus := string(params["authorize_status"])

	// TODO: 数据逻辑处理

	log := &model.AuthorizeUserLog{
		AuthorizerID:     authorizerID,
		AuthorizeeID:     authorizeeID,
		AuthorizeTime:    authorizeTime,
		AuthorizeContent: authorizeContent,
		AuthorizeStatus:  authorizeStatus,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("topic_vx", []string{})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func VerifyUserIdentity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.VerifyUserIdentityLog{
		NameHash:         string(params["name_hash"]),
		IdentityCardHash: string(params["identity_card_hash"]),
		TelephoneHash:    string(params["telephone_hash"]),
		VerifyStatus:     string(params["verify_status"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func AssociateIdentity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.AssociateIdentityLog{
		IdentityCardHash: string(params["identity_card_hash"]),
		UserID:           string(params["user_id"]),
		PID:              string(params["pid"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func CancelAuthorize(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.CancelAuthorizeLog{
		AuthorizerID:    string(params["authorizer_id"]),
		AuthorizeeID:    string(params["authorizee_id"]),
		CanceledContent: string(params["canceled_content"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func RegisterCity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.RegisterCityLog{
		CityID:       string(params["city_id"]),
		RegisterInfo: string(params["register_info"]),
		RegisterTime: string(params["register_time"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func ForbiddenCity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.ForbiddenCityLog{
		CityID:        string(params["city_id"]),
		ForbiddenTime: string(params["forbidden_time"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func RegisterAPP(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.RegisterAPPLog{
		APPID:        string(params["app_id"]),
		CityID:       string(params["city_id"]),
		RegisterInfo: string(params["register_info"]),
		RegisterTime: string(params["register_time"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func ForbiddenAPP(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.ForbiddenAPPLog{
		APPID:         string(params["app_id"]),
		ForbiddenTime: string(params["forbidden_time"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func RegisterService(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.RegisterServiceLog{
		ServiceID:    string(params["service_id"]),
		CityID:       string(params["city_id"]),
		APPID:        string(params["app_id"]),
		RegisterInfo: string(params["register_info"]),
		RegisterTime: string(params["register_time"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

func ForbiddenService(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.ForbiddenServiceLog{
		ServiceID:     string(params["service_id"]),
		ForbiddenTime: string(params["forbidden_time"]),
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// TODOO
	stub.EmitEvent("topic_vx", []string{})

	// TODO
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}
