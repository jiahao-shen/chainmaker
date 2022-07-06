package api

import (
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/model"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/pb/protogo"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/shim"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/utils"
)

func Hello(stub shim.CMStubInterface) protogo.Response {
	return shim.Success([]byte("Hello World"))
}

// 生码
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
	stub.EmitEvent("GenerateQR", []string{serialNumber, generatorID, ctidHash, generateTime, validityPeriod, cityID, appID, sceneID})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, serialNumber, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 验码
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
	stub.EmitEvent("VerifyQR", []string{serialNumber, verifierID, cCodeContentHash, cityID, appID, verifyTime, generatorID})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, serialNumber, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// TODO:
// 授权用户
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
	stub.EmitEvent("AuthorizeUser", []string{authorizerID, authorizeeID, authorizeTime, authorizeContent, authorizeStatus})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 验证身份
func VerifyIdentity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	nameHash := string(params["name_hash"])
	identityHash := string(params["identity_hash"])
	telephoneHash := string(params["telephone_hash"])
	verifyStatus := string(params["verify_status"])
	verifyTime := string(params["verify_time"])

	// TODO: 数据逻辑处理

	log := &model.VerifyUserIdentityLog{
		NameHash:      nameHash,
		IdentityHash:  identityHash,
		TelephoneHash: telephoneHash,
		VerifyStatus:  verifyStatus,
		VerifyTime:    verifyTime,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("VerifyIdentity", []string{nameHash, identityHash, telephoneHash, verifyStatus})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, utils.GetSHA256String([]string{nameHash, identityHash, telephoneHash, verifyStatus, verifyTime}), logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 关联身份
func AssociateIdentity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	identityHash := string(params["identity_hash"])
	digitalIdentity := string(params["digital_identity"])
	pid := string(params["pid"])

	// TODO: 数据逻辑处理

	log := &model.AssociateIdentityLog{
		IdentityHash:    identityHash,
		DigitalIdentity: digitalIdentity,
		PID:             pid,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("AssociateIdentity", []string{identityHash, digitalIdentity, pid})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, digitalIdentity, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// TODO:
// 取消授权
func CancelAuthorize(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	authorizerID := string(params["authorizer_id"])
	authorizeeID := string(params["authorizee_id"])
	canceledContent := string(params["canceled_content"])

	// TODO: 数据逻辑处理

	log := &model.CancelAuthorizeLog{
		AuthorizerID:    authorizerID,
		AuthorizeeID:    authorizeeID,
		CanceledContent: canceledContent,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("CancelAuthorize", []string{authorizerID, authorizeeID, canceledContent})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, "", logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 注册城市
func RegisterCity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	cityID := string(params["city_id"])
	registerInfo := string(params["register_info"])
	registerTime := string(params["register_time"])

	// TODO: 数据逻辑处理

	log := &model.RegisterCityLog{
		CityID:       cityID,
		RegisterInfo: registerInfo,
		RegisterTime: registerTime,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("RegisterCity", []string{cityID, registerInfo, registerTime})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, cityID, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 禁用城市
func ForbiddenCity(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	cityID := string(params["city_id"])
	forbiddenTime := string(params["forbidden_time"])

	// TODO: 数据逻辑处理

	log := &model.ForbiddenCityLog{
		CityID:        cityID,
		ForbiddenTime: forbiddenTime,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("ForbiddenCity", []string{cityID, forbiddenTime})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, cityID, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 注册城市应用
func RegisterAPP(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	appID := string(params["app_id"])
	cityID := string(params["city_id"])
	registerInfo := string(params["register_info"])
	registerTime := string(params["register_time"])

	// TODO: 数据逻辑处理

	log := &model.RegisterAPPLog{
		APPID:        appID,
		CityID:       cityID,
		RegisterInfo: registerInfo,
		RegisterTime: registerTime,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("RegisterAPP", []string{appID, cityID, registerInfo, registerTime})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, appID, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 禁用城市应用
func ForbiddenAPP(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	appID := string(params["app_id"])
	forbiddenTime := string(params["forbidden_time"])

	// TODO: 数据逻辑处理

	log := &model.ForbiddenAPPLog{
		APPID:         appID,
		ForbiddenTime: forbiddenTime,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("ForbiddenAPP", []string{appID, forbiddenTime})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, appID, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 注册服务
func RegisterService(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	serviceID := string(params["service_id"])
	cityID := string(params["city_id"])
	appID := string(params["app_id"])
	registerInfo := string(params["register_info"])
	registerTime := string(params["register_time"])

	// TODO: 数据逻辑处理

	log := &model.RegisterServiceLog{
		ServiceID:    serviceID,
		CityID:       cityID,
		APPID:        appID,
		RegisterInfo: registerInfo,
		RegisterTime: registerTime,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("RegisterService", []string{serviceID, cityID, appID, registerInfo, registerTime})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, serviceID, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}

// 禁用服务
func ForbiddenService(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	serviceID := string(params["service_id"])
	forbiddenTime := string(params["forbidden_time"])

	// TODO: 数据逻辑处理

	log := &model.ForbiddenServiceLog{
		ServiceID:     serviceID,
		ForbiddenTime: forbiddenTime,
	}

	logByte, err := json.Marshal(log)

	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}

	// 发送日志
	stub.EmitEvent("ForbiddenService", []string{serviceID, forbiddenTime})

	// 写入账本
	err = stub.PutStateByte(model.FileKey, serviceID, logByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("保存文件出错: %s", err))
	}

	return shim.Success(logByte)
}
