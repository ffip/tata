package mysql

import (
	"context"

	"bitbucket.org/pwq/tata/lib/text"
	"github.com/tidwall/gjson"
)

const (
	// subject
	_getUser = "SELECT * FROM `users` u"
	_byID    = "`id` = ?"
	_find    = "(`user_account` = ? OR `idcard` = ? OR `phone` = ?)"
	_login   = "`user_pwd` = ?;"
	_newUser = "INSERT INTO `users` (`phone`,`user_pwd`) VALUES (?, ?);"
	_setUser = "UPDATE users SET `name` = ?, `nation` = ?, `sex` = ?, `idcard` = ?, `phone` = ?, `zone_province` = ?, `zone_province_code` = ?, `zone_city` = ?, `zone_city_code` = ?, `zone_county` = ?, `zone_county_code` = ?, `address` = ?, `studytype` = ?, `image_front` = ?, `image_back` = ?, `photo` = ? WHERE `user_id` = ?;"
)

var (
	// Session ==> mysql conn session
	Session = new(Dao)
)

// FindUser ==> get users info by idcard/phone/username.
func (dao *Dao) FindUser(ctx context.Context, user string) (out string, err error) {
	rows, err := dao.DB.Query(ctx, text.Mgr(_getUser, _where, _find, _end), user, user, user)
	if err == nil {
		out = rows.ToJSONStr()
	}
	return
}

// GetUser ==> get user info by idcard/phone/username and password.
func (dao *Dao) GetUser(ctx context.Context, user, password string) (out string, err error) {
	rows, err := dao.DB.Query(ctx, text.Mgr(_getUser, _where, _find, _and, _login), user, user, user, password)
	if err == nil {
		out = rows.ToJSONStr()
	}
	return
}

// New ==> new task.
func (dao *Dao) New(ctx context.Context, phone, password string) (id int64, err error) {
	row, err := dao.DB.Exec(ctx, _newUser, phone, password)
	if err != nil {
		id, _ = row.LastInsertId()
	}
	return id, err
}

// SetUser ==> set user profile item.
func (dao *Dao) SetUser(ctx context.Context, data []byte) (err error) {
	// {
	// "name": "测试",
	// "nation": "汉",
	// "sex": "1",
	// "idcard": "110100199001010000",
	// "phone": "19981741401",
	// "zone_province": "四川省",
	// "zone_province_code": "11",
	// "zone_city": "成都市",
	// "zone_city_code": "1101",
	// "zone_county": "金牛区",
	// "zone_county_code": "110101",
	// "address": "市辖区",
	// "studytype": "1",
	// "image_front": "https://n40api.chdriver.com/upload/user/65889/1649297160.png",
	// "image_back": "https://n40api.chdriver.com/upload/user/65889/1649297160.png",
	// "photo": "https://n40api.chdriver.com/upload/user/65889/1649298437.png"
	// }
	_, err = dao.DB.Exec(ctx, _setUser, gjson.GetBytes(data, "name").Str, gjson.GetBytes(data, "nation").Str, gjson.GetBytes(data, "sex").Str, gjson.GetBytes(data, "idcard").Str, gjson.GetBytes(data, "phone").Str, gjson.GetBytes(data, "zone_province").Str, gjson.GetBytes(data, "zone_province_code").Str, gjson.GetBytes(data, "zone_city").Str, gjson.GetBytes(data, "zone_city_code").Str, gjson.GetBytes(data, "zone_county").Str, gjson.GetBytes(data, "zone_county_code").Str, gjson.GetBytes(data, "address").Str, gjson.GetBytes(data, "studytype").Str, gjson.GetBytes(data, "image_front").Str, gjson.GetBytes(data, "image_back").Str, gjson.GetBytes(data, "photo").Str)
	return
}

// GetUserByID ==> get platf user by id.
func (dao *Dao) GetUserByID(ctx context.Context, id int) (out string, err error) {
	rows, err := dao.DB.Query(ctx, text.Mgr(_getUser, _where, _byID, _end), id)
	if err == nil {
		out = rows.ToJSONStr()
	}
	return
}
