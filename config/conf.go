package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	DbHost             string
	DbPort             string
	DbUser             string
	DbPassWord         string
	DbName             string
	Charset            string
	EtcdHost           string
	EtcdPort           string
	UserServiceAddress string
	TaskServiceAddress string
)

func Init() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录：", err)
		return
	}

	fmt.Println("当前目录：", currentDir)
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误, 请检查配置文件", err)
	}
	LoadMySqlData(file)
	LoadEtcd(file)
	LoadServer(file)
}

func LoadMySqlData(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadServer(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
}
