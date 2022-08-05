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

func TestQueryContr(t *testing.T) {
	client, err := sdk.NewChainClient(sdk.WithConfPath("sdk_config.yaml"))
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.QueryContract("test", "hello", nil, 5)
	if err != nil {
		fmt.Println("fuck")
		fmt.Println(err)
	}

	fmt.Println(resp)
}
