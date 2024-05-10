package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var GlobalConfig Config

// Config
//
// @Description: 配置文件
// @Author liuxingyu
// @Date 2024-02-16 17:18:38
type Config struct {
	Database map[string]DatabaseInfo `yaml:"database"`
	Email    map[string]EmailInfo    `yaml:"email"`
	Session  SessionInfo             `yaml:"session"`
}

// DatabaseInfo
//
// @Description: 数据库配置文件
// @Author liuxingyu
// @Date 2024-02-16 17:19:01
type DatabaseInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// EmailInfo
//
// @Description: 发送邮件配置文件
// @Author liuxingyu
// @Date 2024-02-16 17:19:13
type EmailInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// SessionInfo
//
// @Description: session配置文件
// @Author liuxingyu
// @Date 2024-02-16 17:21:14
type SessionInfo struct {
	Type   string `yaml:"type"`
	Secret string `yaml:"secret"`
	MaxAge int    `yaml:"max_age"`
	Cookie struct {
		Name string `yaml:"name"`
	} `yaml:"cookie"`
	Redis struct {
		Name         string `yaml:"name"`
		Address      string `yaml:"address"`
		DB           string `yaml:"db"`
		Password     string `yaml:"password"`
		PrefixLength int    `yaml:"prefix_length"`
		Network      string `yaml:"network"`
	} `yaml:"redis"`
	Memcached struct {
		Name    string `yaml:"name"`
		Address string `yaml:"address"`
	} `yaml:"memcached"`
}

// ParseConfig
//
// @Title ParseConfig
// @Description: 解析配置文件
// @Author liuxingyu
// @Date 2024-02-16 17:50:24
func ParseConfig() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf(".env文件无法打开：%v", err)
	}

	openPath := "./config/config.yaml"
	if os.Getenv("ENV") == "development" {
		openPath = "./config/config.development.yaml"
	}
	fmt.Println(openPath)

	// 打开文件
	file, err := os.Open(openPath)
	if err != nil {
		log.Fatalf("无法打开配置文件：%v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// 读取文件内容
	data, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatalf("无法读取配置文件：%v", err)
	}

	// 解析配置文件
	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		log.Fatalf("无法解析配置文件：%v", err)
	}
}

func init() {
	ParseConfig()
}
