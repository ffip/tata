package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ffip/tata/api/user/cfg"
	xsql "github.com/ffip/tata/api/user/db/mysql"
	"github.com/ffip/tata/api/user/ext"
	"github.com/ffip/tata/lib/db/mysql"
	log "github.com/ffip/tata/lib/log"
)

func main() {
	flag.Parse()

	musql()

	ext.ListenAndServe(&log.Log{})

	fmt.Scan()
}

func musql() {
	// cfg.DB.Mysql = `{"addr":"test","dsn":"root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4","readDSN":["root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"],"archive":20,"idle":10,"idleTimeout":"4h","queryTimeout":"15s","execTimeout":"5s","tranTimeout":"5s"}`
	xsql.Session = xsql.New(mysql.GetConfig(cfg.DB.Mysql))
	fmt.Println(xsql.Session.Ping(context.TODO()))
}
