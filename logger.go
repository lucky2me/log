package log

import (
	"os"
	"fmt"
	"time"
	"runtime"
	"errors"
)

const (
	LoggerLevelDebug = iota
	LoggerLevelInfo
	LoggerLevelError
)

const defaultCallDepth int = 2

type Logger struct {
	rootPath string   	// desc:	absolute path
	file     *os.File 	// desc:	log file
	level    int      	// option: 	LoggerLevelDebug\LoggerLevelInfo\LoggerLevelError
	depth    int      	// default: 2
	nextDay   time.Time	// desc: 	下一次创建文件的时间
}

func NewLogger(rootPath string, level ...int) Logger {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	l := Logger{}
	l.depth =defaultCallDepth
	l.rootPath = rootPath

	var levelEnum = 0
	if len(level) > 0 {
		levelEnum = level[0]
	}
	if levelEnum != LoggerLevelDebug && levelEnum != LoggerLevelInfo && levelEnum != LoggerLevelError {
		panic("等级不存在")
	}
	l.level = levelEnum

	err := l.getLogFile()
	if err != nil {
		panic(err)
	}

	return l
}


func (this *Logger)SetCallDepth(depth int) {
	if depth > 0 {
		this.depth = depth
	}
}

func (this *Logger)Debug(args ...interface{}) {
	if LoggerLevelDebug < this.level {
		return
	}

	this.writeLogFormat(LoggerLevelDebug, fmt.Sprintf("%s", args))
}
func (this *Logger)Info(args ...interface{}) {
	if LoggerLevelInfo < this.level {
		return
	}

	this.writeLogFormat(LoggerLevelInfo, fmt.Sprintf("%s", args))
}
func (this *Logger)Error(args ...interface{}) {
	if LoggerLevelError < this.level {
		return
	}

	this.writeLogFormat(LoggerLevelError, fmt.Sprintf("%s", args))
}


func (this *Logger)getLogFile() (error) {
	rootPath := this.rootPath
	flag, err := IsExist(rootPath)

	if err != nil {

		panic(err)
	}

	if flag == false {
		os.MkdirAll(rootPath, os.ModeDir)
	}
	date := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	// ------------------


	//time.Now().
	nextD := time.Unix(time.Now().Unix() + (24 * 3600), 0);


	nextD = time.Date(nextD.Year(), nextD.Month(), nextD.Day(), 0, 0, 0, 0, nextD.Location())
	this.nextDay = nextD
	// ------------------
	logPath := fmt.Sprintf("%s/%s.log", rootPath, date)
	f, err := os.OpenFile(logPath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, os.ModePerm)

	if f == nil {
		return errors.New("log文件打开失败")
	}

	this.file = f

	return err
}

// 格式化的写入日志,level是一个枚举,如LoggerLevelError,log是日志字符串
func (this *Logger)writeLogFormat(level int, log string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()


	// 时间
	now := time.Now()
	if now.Unix() > this.nextDay.Unix() {	// 超过了原定的下次创建时间, 重新创建一个文件
		if err := this.getLogFile(); err != nil {
			panic(err)
		}
	}

	time := time.Unix(now.Unix(), 0).Format("2006-01-02 15:04:05")


	var flag string

	switch level {
	case LoggerLevelDebug:
		flag = "D"
	case LoggerLevelInfo:
		flag = "I"
	case LoggerLevelError:
		flag = "E"
	}

	_, file, line, ok := runtime.Caller(this.depth)
	if ok == false {
		panic(errors.New("获取行数失败"))
	}
	_, err := Write(this.file, fmt.Sprintf("%s [%s] [%s:%d] %s\n", time, flag, file, line, log))
	if err != nil {
		panic(err)
	}
}