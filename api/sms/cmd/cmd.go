package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"

	"github.com/ffip/tata/api/sms/cfg"
	"github.com/ffip/tata/api/sms/db/mysql"
	"github.com/ffip/tata/api/sms/ext"
	"github.com/ffip/tata/lib/conf/env"
	xsql "github.com/ffip/tata/lib/db/mysql"
	"github.com/ffip/tata/lib/log"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	musql()

	log := logger()
	ext.ListenAndServe(log)
	for {
	}
}

func musql() {
	cfg.DB.Mysql = `{"addr":"test","dsn":"root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4","readDSN":["root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"],"archive":20,"idle":10,"idleTimeout":"4h","queryTimeout":"15s","execTimeout":"5s","tranTimeout":"5s"}`
	mysql.Session = mysql.New(xsql.GetConfig(cfg.DB.Mysql))
	fmt.Println(mysql.Session.Ping(context.TODO()))
}

func logger() (out *logrus.Logger) {
	logLevel := logrus.InfoLevel
	out = log.NewLogger(logLevel)

	if env.System.LogFormat == "json" {
		out.SetFormatter(&logrus.JSONFormatter{})
	} else if runtime.GOOS == "windows" {
		out.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	}
	return
}
