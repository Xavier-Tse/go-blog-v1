package core

import (
	"Backend/config"
	"Backend/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml文件配置
func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		panic(fmt.Errorf("config init unmarshal: %s", err))
	}

	log.Println("config yamlFile load init success")
	global.Config = c
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
