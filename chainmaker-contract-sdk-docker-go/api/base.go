package api

import (
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/model"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/pb/protogo"
	"chainmaker.org/chainmaker/chainmaker-contract-sdk-docker-go/shim"
)

func GenerateQR(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.GenerateQRLog{
		GeneratorID:    string(params["generator_id"]),
		CTIDHash:       string(params["ctid_hash"]),
		GenerateTime:   string(params["generate_time"]),
		ValidityPeriod: string(params["validity_period"]),
		CityID:         string(params["city_id"]),
		APPID:          string(params["app_id"]),
		SceneID:        string(params["scene_id"]),
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

func VerifyQRLog(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.VerifyQRLog{
		VerifierID:                   string(params["verifier_id"]),
		ComprehensiveCodeContentHash: string(params["comprehensive_code_content_hash"]),
		CityID:                       string(params["city_id"]),
		APPID:                        string(params["app_id"]),
		VerifyTime:                   string(params["verify_time"]),
		GeneratorID:                  string(params["generator_id"]),
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

func AuthorizeUser(stub shim.CMStubInterface) protogo.Response {
	params := stub.GetArgs()

	log := &model.AuthorizeUserLog{
		AuthorizerID:     string(params["authorizer_id"]),
		AuthorizeeID:     string(params["authorizee_id"]),
		AuthorizeTime:    string(params["authorize_time"]),
		AuthorizeContent: string(params["authorize_content"]),
		AuthorizeStatus:  string(params["authorize_status"]),
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

func RegisterCityLog(stub shim.CMStubInterface) protogo.Response {
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
