package Logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Logger       ==> 日志规范化(退出开关,日志等级,模块名称,方法名称,执行过程)
type Logger struct {
	exit     bool
	Level    string
	Module   string
	Function string
	Process  string
	field    []string
}

const LevelAll, LevelDebug, LevelInfo, LevelWarn, LevelError, LevelNone = "ALL", "DEBUG", "INFO", "WARNING", "ERROR", "NONE"

var levels map[string]int = map[string]int{
	LevelAll:   -1,
	LevelDebug: 0,
	LevelInfo:  1,
	LevelWarn:  2,
	LevelError: 3,
	LevelNone:  4,
}

// Printf       ==> 日志规范化打印(内容格式,内容参数)
func (l *Logger) Printf(level, format string, value ...interface{}) {

	if levels[strings.ToUpper(level)] < levels[strings.ToUpper(l.Level)] {
		return
	}

	var fileds string

	for _, v := range l.field {
		fileds += fmt.Sprintf("<%s>", v)
	}

	format = fmt.Sprintf("%s%s - %s\n", "%s [%s][%s][%s][%s]", fileds, format)
	var e []interface{}

	e = append(e, time.Now().Format("2006-01-02 15:04:05"), level, l.Module, l.Function, l.Process)
	for i := 0; i < len(value); i++ {
		e = append(e, value[i])
	}

	_, _ = fmt.Fprintf(os.Stdout, format, e...)
	if l.exit {
		os.Exit(1)
	}
}

func (l *Logger) Debug(format string, value ...interface{}) { l.Printf(LevelDebug, format, value...) }
func (l *Logger) Info(format string, value ...interface{})  { l.Printf(LevelInfo, format, value...) }
func (l *Logger) Warn(format string, value ...interface{})  { l.Printf(LevelWarn, format, value...) }
func (l *Logger) Error(format string, value ...interface{}) { l.Printf(LevelError, format, value...) }

// WithFiled       ==> 添加提示项
func (l *Logger) WithField(item string) *Logger {
	l.field = append(l.field, item)
	return l
}

// WithExit       ==> 提示并退出
func (l *Logger) WithExit() *Logger {
	l.exit = true
	return l
}

// Copy       ==> 衍生子日志输出
func (l *Logger) Copy() (sub Logger) {
	return *l
}
