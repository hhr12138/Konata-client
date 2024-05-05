package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	defaultAddr = "localhost:59864"
)

var (
	GroupId          string
	ClientId         string
	ConfigCenterAddr string
	// init时记得判长度
	DefaultAddrs = []string{defaultAddr}
)

type config struct {
	// 配置中心地址
	ConfigCenterAddr string   `json:"config_center_addr"`
	GroupId          string   `json:"group_id"`
	ClientId         string   `json:"client_id"`
	Addrs            []string `json:"addrs"`
}

// 解析yaml文件，并判断
func init() {
	yamlFile := "/yaml/konata.yaml"
	file, err := os.Open(yamlFile)
	if err != nil {
		panic(fmt.Errorf("open config file err: %v", err))
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	var cfg config
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(fmt.Errorf("decode config file err: %v", err))
	}
	if len(cfg.Addrs) == 0 {
		panic(fmt.Errorf("len(cfg.Addrs) == 0"))
	}
	DefaultAddrs = cfg.Addrs
	GroupId = cfg.GroupId
	ClientId = cfg.ClientId
	ConfigCenterAddr = cfg.ConfigCenterAddr
}
