// Package env 所有的公共字段必须在 init() 初始化，并通过 flag.Parse() 声明参数.
package env

import (
	"flag"
	"os"
	"time"

	"bitbucket.org/pwq/tata/lib/net/ip"
)

// 部署版本
const (
	DeployEnvDev  = "dev"
	DeployEnvFat1 = "fat1"
	DeployEnvUat  = "uat"
	DeployEnvPre  = "pre"
	DeployEnvProd = "prod"
	DeployVersion = "ALPHA_VERSION"
)

// 环境变量默认值
const (
	// env
	_region     = "sh"
	_zone       = "sh001"
	_deployEnv  = "dev"
	_vifName    = "vif0"
	_vifAddr    = "10.10.10.10/24"
	_logVerbose = "false"
	_logFormat  = "pain"
)

// 环境变量配置
var System struct {
	// 应用部署区域
	Region string `json:"region"`
	// 应用服务区域
	Zone string `json:"zone"`
	// 服务主机名称
	Hostname string `json:"hostname"`
	// 部署环境变量
	DeployEnv string `json:"deployEnv"`
	// IP
	// IP = os.Getenv("POD_IP")
	IP string `json:"ip"`
	// AppID是全局唯一的应用程序ID，通过服务树进行注册
	// 例如 main.arc.discovery
	AppID string `json:"appID"`
	// Caster集群中不同实验组的标识的颜色。
	Color string `json:"color"`
	// Type 节点服务类型
	Type string `json:"type"`
	// 调试模式
	LogVerbose string `json:"-"`
	// 日志格式
	LogFormat string `json:"logFormat"`
	// 启动时间
	BootTime int64 `json:"bootTime"`
}

func init() {
	var err error
	if System.Hostname, err = os.Hostname(); err != nil || System.Hostname == "" {
		System.Hostname = os.Getenv("HOSTNAME")
	}

	System.IP = os.Getenv("POD_IP")
	if System.IP == "" {
		System.IP = ip.GetLocalIpV4()
	}

	System.BootTime = time.Now().Unix()

	addFlag(flag.CommandLine)
}

func addFlag(fs *flag.FlagSet) {
	// env
	fs.StringVar(&System.Region, "region", DefaultString("REGION", _region), "Flag - Region [env:REGION]")
	fs.StringVar(&System.Zone, "zone", DefaultString("ZONE", _zone), "Flag - Zone [env:ZONE]")
	fs.StringVar(&System.AppID, "appid", os.Getenv("APP_ID"), "Flag - App id [env:APP_ID]")
	fs.StringVar(&System.DeployEnv, "deploy.env", DefaultString("DEPLOY_ENV", _deployEnv), "Flag - Deploy env [env:DEPLOY_ENV]")
	fs.StringVar(&System.Color, "deploy.color", os.Getenv("DEPLOY_COLOR"), "Flag - Deploy color [env:DEPLOY_COLOR]")

	// debug
	fs.StringVar(&System.LogVerbose, "log.verbose", DefaultString("LOG_VERBOSE", _logVerbose), "App debug mode [env:LOG_VERBOSE]")
	fs.StringVar(&System.LogFormat, "log.format", DefaultString("LOG_FORMAT", _logFormat), "Console log print format [env:LOG_FORMAT]")
}

func DefaultString(env, value string) string {
	v := os.Getenv(env)
	if v == "" {
		return value
	}
	return v
}
