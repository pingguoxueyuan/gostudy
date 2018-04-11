package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

//2018/3/26 0:01.383 DEBUG logDebug.go:29 this is a debug log
//2006-01-02 15:04:05.999
type FileLogger struct {
	level         int
	logPath       string
	logName       string
	file          *os.File
	warnFile      *os.File
	LogDataChan   chan *LogData
	logSplitType  int
	logSplitSize  int64
	lastSplitHour int
}

func NewFileLogger(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path ")
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name ")
		return
	}

	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level ")
		return
	}

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}

	var logSplitType int = LogSplitTypeHour
	var logSplitSize int64
	logSplitStr, ok := config["log_split_type"]
	if !ok {
		logSplitStr = "hour"
	} else {
		if logSplitStr == "size" {
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok {
				logSplitSizeStr = "104857600"
			}

			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 104857600
			}

			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}

	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}

	level := getLogLevel(logLevel)
	log = &FileLogger{
		level:         level,
		logPath:       logPath,
		logName:       logName,
		LogDataChan:   make(chan *LogData, chanSize),
		logSplitSize:  logSplitSize,
		logSplitType:  logSplitType,
		lastSplitHour: time.Now().Hour(),
	}

	return
}

func (f *FileLogger) Init() {

	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filename, err))
	}

	f.file = file

	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed, err:%v", filename, err))
	}

	f.warnFile = file
	go f.writeLogBackground()
}

func (f *FileLogger) splitFileHour(warnFile bool) {
	now := time.Now()
	hour := now.Hour()
	if hour == f.lastSplitHour {
		return
	}

	f.lastSplitHour = hour
	var backupFilename string
	var filename string

	if warnFile {
		backupFilename = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)

		filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	} else {
		backupFilename = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	}

	file := f.file
	if warnFile {
		file = f.warnFile
	}

	file.Close()
	os.Rename(filename, backupFilename)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f *FileLogger) splitFileSize(warnFile bool) {

	file := f.file
	if warnFile {
		file = f.warnFile
	}

	statInfo, err := file.Stat()
	if err != nil {
		return
	}

	fileSize := statInfo.Size()
	if fileSize <= f.logSplitSize {
		return
	}

	var backupFilename string
	var filename string

	now := time.Now()
	if warnFile {
		backupFilename = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

		filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	} else {
		backupFilename = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	}

	file.Close()
	os.Rename(filename, backupFilename)

	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}

func (f *FileLogger) checkSplitFile(warnFile bool) {

	if f.logSplitType == LogSplitTypeHour {
		f.splitFileHour(warnFile)
		return
	}

	f.splitFileSize(warnFile)
}

func (f *FileLogger) writeLogBackground() {
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnFile
		}

		f.checkSplitFile(logData.WarnAndFatal)

		fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
			logData.LevelStr, logData.Filename, logData.FuncName, logData.LineNo, logData.Message)
	}
}

func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}

	logData := writeLog(LogLevelDebug, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writeLog(LogLevelTrace, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}

	logData := writeLog(LogLevelWarn, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}

	logData := writeLog(LogLevelError, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}

	logData := writeLog(LogLevelFatal, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
