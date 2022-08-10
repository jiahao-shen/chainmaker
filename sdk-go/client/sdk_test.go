package main

import (
	"fmt"
	"reflect"
	"testing"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
)

func TestInstallContract(t *testing.T) {
	client, err := sdk.NewChainClient(sdk.WithConfPath("sdk_config.yaml"))
	if err != nil {
		fmt.Println(err)
	}

	payload, err := client.CreateContractCreatePayload("test", "1.0.0", "test.7z", common.RuntimeType_DOCKER_GO, []*common.KeyValuePair{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reflect.TypeOf(payload))

	entry, err := sdkutils.MakeEndorserWithPath("crypto-config/wx-org.chainmaker.org/user/admin1/admin1.tls.key", "crypto-config/wx-org.chainmaker.org/user/admin1/admin1.tls.crt", payload)
	if err != nil {
		fmt.Println(err)
	}

	var endorsers []*common.EndorsementEntry
	endorsers = append(endorsers, entry)

	resp, err := client.SendContractManageRequest(payload, endorsers, 5, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.TxId)

}

func TestQueryContract(t *testing.T) {
	client, err := sdk.NewChainClient(sdk.WithConfPath("sdk_config.yaml"))
	if err != nil {
		fmt.Println(err)
	}

	params := []*common.KeyValuePair{
		{
			Key:   "method",
			Value: []byte("hello"),
		},
	}

	resp, err := client.QueryContract("test", "invoke_contract", params, -1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func TestInvokeContract(t *testing.T) {
	client, err := sdk.NewChainClient(sdk.WithConfPath("sdk_config.yaml"))
	if err != nil {
		fmt.Println(err)
	}

	params := []*common.KeyValuePair{
		{
			Key:   "method",
			Value: []byte("saveFile"),
		},
		{
			Key:   "file_name",
			Value: []byte("test"),
		},
		{
			Key:   "file_hash",
			Value: []byte("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"),
		},
		{
			Key:   "time",
			Value: []byte("2022-08-10 17:15:55"),
		},
	}

	resp, err := client.InvokeContract("test", "invoke_contract", "", params, -1, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

}
