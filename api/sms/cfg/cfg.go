// Package cfg 所有的公共字段必须在 init() 初始化，并通过 flag.Parse() 声明参数.
package cfg

import (
	"flag"
	"os"

	"bitbucket.org/pwq/tata/lib/conf/env"
)

// app default value.
const (
	_httpPort   = "127.0.0.1:8000"
	_logVerbose = "false"
	_logFormat  = "pain"
)

// web cfg.
var Web struct {
	HttpAddr string `json:"httpAddr"`
	HttpsCrt string `json:"httpsCrt"`
	HttpsKey string `json:"httpsKey"`
}

// db dsn
var DB struct {
	Mysql string `json:mysql`
	Redis string `json:redis`
	Mongo string `json:mongo`
	MSSQL string `json:mssql`
	Psql  string `json:psql`
}

func init() {
	addFlag(flag.CommandLine)
}

func addFlag(fs *flag.FlagSet) {
	// db
	fs.StringVar(&DB.Mysql, "db.mysql", os.Getenv("DSN_MYSQL"), "Mysql DSN [env:DSN_MYSQL]")
	fs.StringVar(&DB.Redis, "db.redis", os.Getenv("DSN_REDIS"), "Redis DSN [env:DSN_REDIS]")
	fs.StringVar(&DB.Redis, "db.mongo", os.Getenv("DSN_MONGO"), "Redis DSN [env:DSN_MONGO]")
	fs.StringVar(&DB.Redis, "db.mssql", os.Getenv("DSN_MSSQL"), "Redis DSN [env:DSN_MSSQL]")
	fs.StringVar(&DB.Redis, "db.pssql", os.Getenv("DSN_PSSQL"), "Redis DSN [env:DSN_PSSQL]")
	// web
	fs.StringVar(&Web.HttpAddr, "http.addr", env.DefaultString("HTTP_PORT", _httpPort), "Http services listen addr [env:HTTP_PORT]")
	fs.StringVar(&Web.HttpsCrt, "https.cert", os.Getenv("HTTPS_CRT"), "Https crt [env:HTTPS_CRT]")
	fs.StringVar(&Web.HttpsKey, "https.key", os.Getenv("HTTPS_KEY"), "Https key [env:HTTPS_KEY]")
}
