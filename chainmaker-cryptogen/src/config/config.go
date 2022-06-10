/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"log"

	"github.com/spf13/viper"
)

var CMViper *viper.Viper
var cryptoGenConfig *CryptoGenConfig

const (
	defaultCryptoConfigPath = "../config/crypto_config_template.yml"
)

func LoadCryptoGenConfig(path string) {
	CMViper = viper.New()
	cryptoGenConfig = &CryptoGenConfig{}

	if err := cryptoGenConfig.loadConfig(path); err != nil {
		log.Fatalf("load crypto config [%s] failed, %s",
			path, err)
	}

	cryptoGenConfig.printLog()
}

func GetCryptoGenConfig() *CryptoGenConfig {
	return cryptoGenConfig
}

func (c *CryptoGenConfig) loadConfig(path string) error {
	if path == "" {
		path = defaultCryptoConfigPath
	}

	CMViper.SetConfigFile(path)
	if err := CMViper.ReadInConfig(); err != nil {
		return err
	}

	if err := CMViper.Unmarshal(&cryptoGenConfig); err != nil {
		return err
	}

	return nil
}

func (c *CryptoGenConfig) printLog() {
	// fmt.Printf("Load crypto config success!\n")
	// fmt.Printf("%+v\n", cryptoGenConfig)
}

func PKCS11Enabled() bool {
	if cryptoGenConfig == nil || len(cryptoGenConfig.Item) == 0 {
		return false
	}
	return cryptoGenConfig.Item[0].P11Config.Enabled
}
