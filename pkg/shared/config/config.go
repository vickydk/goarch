package config

import (
	"fmt"

	Database "goarch/pkg/shared/database"
	"goarch/pkg/shared/rest"
	RpcClient "goarch/pkg/shared/rpc_client"

	"github.com/spf13/viper"
)

type Config struct {
	Apps            Apps                    `json:"apps"`
	GoarchGrpc      GoarchConfig            `json:"goarchGrpc"`
	Database        Database.ConfigDatabase `json:"database"`
	GoarchAPIConfig GoarchAPIConfig         `json:"goarchAPIConfig"`
}

type Apps struct {
	Name     string `json:"name"`
	HttpPort int    `json:"httpPort"`
	GRPCPort int    `json:"grpcPort"`
	Version  string `json:"version"`
}

type GoarchConfig struct {
	RpcOptions RpcClient.Options `json:"rpcOptions"`
}

type GoarchAPIConfig struct {
	RestOptions rest.Options `json:"restOptions"`
	Path        struct {
		GetUserDetail string `json:"getUserDetail"`
	} `json:"path"`
}

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.Apps.HttpPort)
}

func NewConfig(path string) *Config {
	fmt.Println("Try NewConfig ... ")

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
