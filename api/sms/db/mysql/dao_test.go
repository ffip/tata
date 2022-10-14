package mysql

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ffip/tata/api/sms/cfg"
	"github.com/ffip/tata/lib/db/mysql"
)

type Queue struct {
	ID        int    `json:"id" db:"id"`
	Data      string `json:"data" db:"data"`
	Type      string `json:"type" db:"type"`
	CreatedAt int64  `json:"createdAt" db:"createdAt"`
	PushedAt  int64  `json:"pushedAt" db:"pushedAt"`
}

func TestMain(m *testing.M) {
	cfg.DB.Mysql = `{"addr":"test","dsn":"root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4","readDSN":["root:test@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"],"archive":20,"idle":10,"idleTimeout":"4h","queryTimeout":"15s","execTimeout":"5s","tranTimeout":"5s"}`
	Session = New(mysql.GetConfig(cfg.DB.Mysql))
	fmt.Println(Session.Ping(context.TODO()))

	// ql, _ := Session.GetPushList(context.TODO())
	// qls := Queue{}
	// ql = strings.ReplaceAll(strings.ReplaceAll(ql, "[", ""), "]", "")
	// json.Unmarshal([]byte(ql), &qls)

	os.Exit(m.Run())
}
