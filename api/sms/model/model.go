package model

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ffip/tata/api/sms/db/mysql"
	"github.com/ffip/tata/api/sms/stack"
)

// New ==> new task
func New(writer http.ResponseWriter, request *http.Request) {
	data, _ := ioutil.ReadAll(request.Body)
	r := stack.Message{}
	json.Unmarshal(data, &r)

	_, err := mysql.Session.New(context.TODO(), &r)
	if err != nil {
		fail(writer, err)
		return
	}
	ok(writer)
	return
}

// Used ==> sms used
func Used(writer http.ResponseWriter, request *http.Request) {
	data, _ := ioutil.ReadAll(request.Body)
	r := stack.Message{}
	json.Unmarshal(data, &r)

	_, err := mysql.Session.New(context.TODO(), &r)
	if err != nil {
		fail(writer, err)
		return
	}
	ok(writer)
	return
}
