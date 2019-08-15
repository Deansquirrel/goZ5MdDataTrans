package object

import (
	"encoding/json"
	"fmt"
	"github.com/Deansquirrel/goToolSecret"
	"strings"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	defaultSSAddress = "MCo7KcSDw9hFF6q8al2mH9l0mscROSA8PBorCjU1OTIROyErLTlkVWopLjQ0NwspMSkiOjoLOnd5cDMxTSwnZWpxZkJ9"
)

//系统配置（Server|Client）
type SystemConfig struct {
	Total        systemConfigTotal        `toml:"total"`
	SSConfig     systemConfigSSConfig     `toml:"ssConfig"`
	RunMode      systemConfigRunMode      `toml:"runMode"`
	OnlineConfig systemConfigOnlineConfig `toml:"onlineConfig"`
	SvrConfig    systemConfigSvrConfig    `toml:"svrConfig"`
	LocalDb      systemConfigLocalDB      `toml:"localDb"`
	Task         systemConfigTask         `toml:"task"`
	Service      systemConfigService      `toml:"service"`
}

func (sc *SystemConfig) FormatConfig() {
	sc.Total.FormatConfig()
	sc.RunMode.FormatConfig()
	sc.OnlineConfig.FormatConfig()
	sc.SvrConfig.FormatConfig()
	sc.LocalDb.FormatConfig()
	sc.Task.FormatConfig()
	sc.Service.FormatConfig()
	sc.SSConfig.FormatConfig()
}

func (sc *SystemConfig) ToString() string {
	d, err := json.Marshal(sc)
	if err != nil {
		log.Warn(fmt.Sprintf("SystemConfig转换为字符串时遇到错误：%s", err.Error()))
		return ""
	}
	return string(d)
}

//通用配置
type systemConfigTotal struct {
	StdOut   bool   `toml:"stdOut"`
	LogLevel string `toml:"logLevel"`
}

func (t *systemConfigTotal) FormatConfig() {
	//去除首尾空格
	t.LogLevel = strings.Trim(t.LogLevel, " ")
	//设置默认日志级别
	if t.LogLevel == "" {
		t.LogLevel = "warn"
	}
	//设置字符串转换为小写
	t.LogLevel = strings.ToLower(t.LogLevel)
	t.LogLevel = t.checkLogLevel(t.LogLevel)
}

//校验SysConfig中iris日志级别设置
func (t *systemConfigTotal) checkLogLevel(level string) string {
	switch level {
	case "debug", "info", "warn", "error":
		return level
	default:
		return "warn"
	}
}

type systemConfigSSConfig struct {
	Address string `toml:"address"`
}

//格式化
func (sc *systemConfigSSConfig) FormatConfig() {
	sc.Address = strings.Trim(sc.Address, " ")
	if strings.Trim(sc.Address, " ") == "" {
		sc.Address = defaultSSAddress
	}
	address, err := goToolSecret.DecryptFromBase64Format(sc.Address, "ServiceSupport")
	if err != nil {
		sc.Address = ""
		log.Error(fmt.Sprintf("convert address err: %s", err.Error()))
		sc.Address, _ = goToolSecret.DecryptFromBase64Format(defaultSSAddress, "ServiceSupport")
	} else {
		sc.Address = address
	}
}

type systemConfigRunMode struct {
	Mode string `toml:"mode"`
}

func (sc *systemConfigRunMode) FormatConfig() {
	sc.Mode = strings.Trim(sc.Mode, " ")
	if sc.Mode == "" || (sc.Mode != string(RunModeMdCollect) && sc.Mode != string(RunModeBbRestore)) {
		sc.Mode = string(RunModeMdCollect)
	}
}

type systemConfigOnlineConfig struct {
	Address string `toml:"address"`
}

func (sc *systemConfigOnlineConfig) FormatConfig() {
	sc.Address = strings.Trim(sc.Address, " ")
}

type systemConfigSvrConfig struct {
	Address string `toml:"address"`
}

func (sc *systemConfigSvrConfig) FormatConfig() {
	sc.Address = strings.Trim(sc.Address, " ")
}

//本地库配置库
type systemConfigLocalDB struct {
	Server string `toml:"server"`
	Port   int    `toml:"port"`
	DbName string `toml:"dbName"`
	User   string `toml:"user"`
	Pwd    string `toml:"pwd"`
}

func (c *systemConfigLocalDB) FormatConfig() {
	c.Server = strings.Trim(c.Server, " ")
	if c.Port == 0 {
		c.Port = 1433
	}
	c.DbName = strings.Trim(c.DbName, " ")
	c.User = strings.Trim(c.User, " ")
	c.Pwd = strings.Trim(c.Pwd, " ")
}

//Task
type systemConfigTask struct {
	Cron string `toml:"cron"`
}

func (sc *systemConfigTask) FormatConfig() {
	sc.Cron = strings.Trim(sc.Cron, " ")
	if sc.Cron == "" {
		sc.Cron = "0 0/5 * * * ?"
	}
}

//服务配置
type systemConfigService struct {
	Name        string `toml:"name"`
	DisplayName string `toml:"displayName"`
	Description string `toml:"description"`
}

//格式化
func (sc *systemConfigService) FormatConfig() {
	sc.Name = strings.Trim(sc.Name, " ")
	sc.DisplayName = strings.Trim(sc.DisplayName, " ")
	sc.Description = strings.Trim(sc.Description, " ")
	if sc.Name == "" {
		sc.Name = "Z5MdDataTrans"
	}
	if sc.DisplayName == "" {
		sc.DisplayName = "Z5MdDataTrans"
	}
	if sc.Description == "" {
		sc.Description = sc.Name
	}
}
