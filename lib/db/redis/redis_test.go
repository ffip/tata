package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

var r = Redis{}

type Client struct {
	Id     string
	Token  string
	Bridge bool
}

func TestConnect(t *testing.T) {
	err := r.Connect("redis://127.0.0.1:6379/10")
	if err != nil {
		log.Fatalln(err)
		return
	}

	client := Client{Id: "test", Token: "test1", Bridge: true}

	cs, err := json.Marshal(client)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = r.Set("test", string(cs), 0)
	if err != nil {
		log.Fatalln(err)
		return
	}
	s, err := r.Get("test")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(s)
	err = r.Del("test")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
