package text

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"regexp"
	"strings"
	"unsafe"

	"github.com/tidwall/gjson"
)

// Mgr ==> 按顺序合并多个string
func Mgr(target ...string) string {
	var buffer bytes.Buffer
	for _, str := range target {
		buffer.WriteString(str)
	}
	return buffer.String()
}

// Mgrs ==> 合并[]string
func Mgrs(src, target []string) []string {
	resSet := make(map[string]struct{}, len(src)+len(target))
	for _, s := range src {
		resSet[s] = struct{}{}
	}
	for _, s := range target {
		resSet[s] = struct{}{}
	}

	res := make([]string, 0, len(resSet))
	for s := range resSet {
		res = append(res, s)
	}

	return res
}

// Rms ==> 从[]sting删除sting
func Rms(src, target []string) []string {
	resSet := make(map[string]struct{}, len(src))
	for _, s := range src {
		resSet[s] = struct{}{}
	}
	for _, s := range target {
		delete(resSet, s)
	}

	res := make([]string, 0, len(resSet))
	for s := range resSet {
		res = append(res, s)
	}

	return res
}

// Atob ==> 将一个string转换成[]byte
func Atob(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// Abot ==> 将一个[]byte转换成string
func Abot(bit []byte) string {
	return *(*string)(unsafe.Pointer(&bit))
}

// Ai64ob ==> 将一个int64转换成[]byte
func Ai64ob(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// ToInt64 ==> 将一个[]byte转换成int64
func ToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// ToArrStr ==> 将一个JsonRaw转换成ArrayStr
func ToArrStr(jsonRaw interface{}) string {
	jsonStr := fmt.Sprintf(`%s`, jsonRaw)

	var re = regexp.MustCompile(`(?m)({|\[|,)".*?":"|\[|\]|"}|"]`)
	jsonStr = re.ReplaceAllString(jsonStr, ``)
	jsonStr = strings.ReplaceAll(jsonStr, `,`, `|`)
	jsonStr = strings.ReplaceAll(jsonStr, `"`, `,`)
	return jsonStr
}

// GjsonToStrArr ==> 将一个JsonRaw转换成[]string
func GjsonToStrArr(jsonRaw gjson.Result) []string {
	jsonStr := fmt.Sprintf(`%s`, jsonRaw)[2:]
	jsonStr = jsonStr[:len(jsonStr)-2]
	jsonStr = strings.ReplaceAll(jsonStr, `"`, ``)
	return strings.Split(jsonStr, `,`)
}
