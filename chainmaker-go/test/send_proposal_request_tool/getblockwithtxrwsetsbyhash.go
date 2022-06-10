/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	commonPb "chainmaker.org/chainmaker/pb-go/v2/common"
	"chainmaker.org/chainmaker/pb-go/v2/syscontract"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
)

func GetBlockWithRWSetsByHashCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getBlockWithRWSetsByHash",
		Short: "Get block with RW sets by hash",
		Long:  "Get block with RW sets by hash",
		RunE: func(_ *cobra.Command, _ []string) error {
			return getBlockWithRWSetsByHash()
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&hash, "hash", "s", "", "specify block hash")

	return cmd
}

func getBlockWithRWSetsByHash() error {
	// 构造Payload
	pairs := []*commonPb.KeyValuePair{
		{
			Key:   "blockHash",
			Value: []byte(hash),
		},
	}

	payloadBytes, err := constructQueryPayload(chainId, syscontract.SystemContract_CHAIN_QUERY.String(), "GET_BLOCK_WITH_TXRWSETS_BY_HASH", pairs)
	if err != nil {
		return err
	}

	resp, err = proposalRequest(sk3, client, payloadBytes)
	if err != nil {
		return err
	}

	blockInfo := &commonPb.BlockInfo{}
	if err = proto.Unmarshal(resp.ContractResult.Result, blockInfo); err != nil {
		return err
	}
	result := &Result{
		Code:                  resp.Code,
		Message:               resp.Message,
		ContractResultCode:    resp.ContractResult.Code,
		ContractResultMessage: resp.ContractResult.Message,
		BlockInfo:             blockInfo,
	}
	fmt.Println(result.ToJsonString())

	return nil
}
