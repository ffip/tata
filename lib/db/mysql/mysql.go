package mysql

import (
	"encoding/json"
	"fmt"
	xtime "time"

	"github.com/ffip/tata/lib/net/netutil/breaker"
	"github.com/ffip/tata/lib/time"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// Config mysql config.
type Config struct {
	Addr         string          // for trace
	DSN          string          // write data source name.
	ReadDSN      []string        // read data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  time.Duration   // connect max life time.
	QueryTimeout time.Duration   // query sql timeout
	ExecTimeout  time.Duration   // execute sql timeout
	TranTimeout  time.Duration   // transaction sql timeout
	Breaker      *breaker.Config // breaker
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *DB) {
	if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err := Open(c)
	if err != nil {
		fmt.Printf("open mysql error(%v) \n", err)
		panic(err)
	}
	return
}

// GetConfig ==> get mysql uri json.
func GetConfig(dsn string) (cfg Config) {
	err := json.Unmarshal([]byte(dsn), &cfg)
	fmt.Println(err)
	cfg.Breaker = &breaker.Config{
		Window:  time.Duration(10 * xtime.Second),
		Sleep:   time.Duration(10 * xtime.Second),
		Bucket:  10,
		Ratio:   0.5,
		Request: 100,
	}
	return cfg
}
