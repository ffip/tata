package Logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Logger       ==> 日志规范化(退出开关,日志等级,模块名称,方法名称,执行过程)
type Log struct {
	Exit     bool
	Level    string
	Model    string
	Function string
	Process  string
}

const LevelAll, LevelDebug, LevelInfo, LevelWarning, LevelError, LevelNone = "ALL", "DEBUG", "INFO", "WARNING", "ERROR", "None"

var levels map[string]int = map[string]int{
	LevelAll:     -1,
	LevelDebug:   0,
	LevelInfo:    1,
	LevelWarning: 2,
	LevelError:   3,
	LevelNone:    4,
}

// Printf       ==> 日志规范化打印(内容格式,内容参数)
func (l *Log) Printf(level, format string, value ...interface{}) {

	if levels[strings.ToLower(level)] < levels[strings.ToLower(l.Level)] {
		return
	}

	format = fmt.Sprintf("%s%s", "%s [%s][%s][%s][%s] - ", format)
	var e []interface{}

	e = append(e, time.Now().Format("2006-01-02 15:04:05"), level, l.Model, l.Function, l.Process)
	for i := 0; i < len(value); i++ {
		e = append(e, value[i])
	}

	_, _ = fmt.Fprintf(os.Stdout, format, e...)
	if l.Exit {
		os.Exit(1)
	}
}

func (l *Log) Debug(format string, value ...interface{})   { l.Printf(LevelDebug, format, value...) }
func (l *Log) Info(format string, value ...interface{})    { l.Printf(LevelInfo, format, value...) }
func (l *Log) Warning(format string, value ...interface{}) { l.Printf(LevelWarning, format, value...) }
func (l *Log) Error(format string, value ...interface{})   { l.Printf(LevelError, format, value...) }

// Copy       ==> 衍生子日志输出
func (l *Log) Copy() (sub Log) {
	return *l
}
