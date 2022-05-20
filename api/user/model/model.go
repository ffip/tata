package model

import (
	"context"
	"io/ioutil"
	"net/http"

	"bitbucket.org/pwq/tata/api/user/db/mysql"
	"bitbucket.org/pwq/tata/lib/text"
)

// New ==> new task
func New(writer http.ResponseWriter, request *http.Request) {
	// 取验证短信，如果成功则存入数据库
	phone := request.FormValue("phone")
	index, err := mysql.Session.New(context.TODO(), phone, phone[6:])
	if err != nil {
		_, _ = writer.Write(text.Atob(text.Mgr("{\"err\":\"", err.Error(), "\"}")))
		return
	}
	_, _ = writer.Write(text.Atob(text.Mgr("{\"id\":\"", text.Abot(text.Ai64ob(index)), "\"}")))
	return
}

// Info ==> get user info
func Info(writer http.ResponseWriter, request *http.Request) {
	id, err := getValueInt(writer, request, "id")
	if err != nil {
		fail(writer, err)
		return
	}
	result, err := mysql.Session.GetUserByID(context.TODO(), id)
	ref(writer, err, result)
	return
}

// Login ==> user login
func Login(writer http.ResponseWriter, request *http.Request) {
	// result, err := mysql.Session.GetUserByID(context.TODO(), request.FormValue("data"))
	// ref(writer, err, result)
	return
}

// Find ==> find user
func Find(writer http.ResponseWriter, request *http.Request) {
	result, err := mysql.Session.FindUser(context.TODO(), request.FormValue("user"))
	ref(writer, err, result)
	return
}

// Update ==> update user
func Update(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fail(writer, err)
		return
	}

	err = mysql.Session.SetUser(context.TODO(), data)
	if err != nil {
		fail(writer, err)
	}
	return
}
