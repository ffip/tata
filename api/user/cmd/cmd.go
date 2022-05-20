package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"

	"bitbucket.org/pwq/tata/api/user/cfg"
	xsql "bitbucket.org/pwq/tata/api/user/db/mysql"
	"bitbucket.org/pwq/tata/api/user/ext"
	"bitbucket.org/pwq/tata/lib/conf/env"
	"bitbucket.org/pwq/tata/lib/db/mysql"
	"bitbucket.org/pwq/tata/lib/log"
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
	// cfg.DB.Mysql = `{"addr":"test","dsn":"root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4","readDSN":["root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"],"archive":20,"idle":10,"idleTimeout":"4h","queryTimeout":"15s","execTimeout":"5s","tranTimeout":"5s"}`
	xsql.Session = xsql.New(mysql.GetConfig(cfg.DB.Mysql))
	fmt.Println(xsql.Session.Ping(context.TODO()))
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
