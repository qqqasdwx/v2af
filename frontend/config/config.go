package config

import (
	"encoding/json"
	"fmt"
	"github.com/v2af/file"
	"log"
	"sync"
)

type GlobalConfig struct {
	Debug    bool                    `json:"debug"`
	Salt     string                  `json:"salt"`
	ShowSql  bool                    `json:"showSql"`
	Http     *HttpConfig             `json:"http"`
	Database map[string]*MysqlConfig `json:"database"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
	Secret string `json:"secret"`
	Access string `json:"access"`
}

type MysqlConfig struct {
	Addr string `json:"addr"`
	Idle int    `json:"idle"`
	Max  int    `json:"max"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("使用 -c 指定配置文件")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("配置文件%s不存在", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("读取配置文件 %s 失败,原因:  %s", cfg, err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("解析配置文件 %s 失败,原因: %s", cfg, err.Error())
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("读取配置文件", cfg, "成功")
	return nil
}
