package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ffip/tata/api/sms/cfg"
	"github.com/ffip/tata/api/sms/db/mysql"
	"github.com/ffip/tata/api/sms/ext"
	xsql "github.com/ffip/tata/lib/db/mysql"
	log "github.com/ffip/tata/lib/log"
)

func main() {
	flag.Parse()

	musql()

	ext.ListenAndServe(&log.Logger{})
	for {
	}
}

func musql() {
	cfg.DB.Mysql = `{"addr":"test","dsn":"root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4","readDSN":["root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"],"archive":20,"idle":10,"idleTimeout":"4h","queryTimeout":"15s","execTimeout":"5s","tranTimeout":"5s"}`
	mysql.Session = mysql.New(xsql.GetConfig(cfg.DB.Mysql))
	fmt.Println(mysql.Session.Ping(context.TODO()))
}
