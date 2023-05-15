/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-03 00:32:05
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-15 16:19:22
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nlog

import (
	"bytes"
	"fmt"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/utils/nfile"
	"github.com/liusuxian/nova/utils/nslice"
	"github.com/liusuxian/nova/utils/nstr"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LogConfig 日志配置
type LogConfig struct {
	Path    string            // 输出日志文件路径
	Details []LogDetailConfig // 日志详细配置
}

// LogDetailConfig 日志详细配置
type LogDetailConfig struct {
	Type       int    // 日志类型 0:打印所有级别 1:打印 DEBUG、INFO、WARN 级别 2:打印 ERROR、DPANIC、PANIC、FATAL 级别，默认0
	Level      int    // 日志打印级别 0:DEBUG 1:INFO 2:WARN 3:ERROR 4:DPANIC、5:PANIC、6:FATAL，默认0
	Format     int    // 输出日志格式 0:logfmt 1:json，默认1
	Filename   string // 输出日志文件名称
	MaxSize    int    // 单个日志文件最多存储量（单位:MB）
	MaxBackups int    // 日志备份文件最多数量
	MaxAge     int    // 日志保留时间（单位:天）
	Compress   bool   // 是否压缩日志
	Stdout     bool   // 是否输出到控制台
}

// stdoutWriter 标准输出
type stdoutWriter struct {
	format int // 输出日志格式 0:logfmt 1:json，默认1
}

// fileWriter 文件输出
type fileWriter struct {
	lumberjack.Logger
	format int // 输出日志格式 0:logfmt 1:json，默认1
}

// nLogger 日志结构
type nLogger struct {
	zapLogger *zap.Logger                 // 只能输出结构化日志，但是性能要高于SugaredLogger
	logConfig LogConfig                   // 日志配置
	logWriter map[int]zapcore.WriteSyncer // 日志输出对象
}

// 默认输出日志文件路径
const (
	defaultPath string = "logs"
)

// 日志类型
const (
	LOGTYPE_ALL   int = iota // 打印所有级别
	LOGTYPE_INFO             // 打印 DEBUG、INFO、WARN 级别
	LOGTYPE_ERROR            // 打印 ERROR、DPANIC、PANIC、FATAL 级别
)

// 输出日志格式
const (
	FORMAT_LOGFMT int = iota
	FORMAT_JSON
)

// logger 实例
var logger = &nLogger{
	logWriter: make(map[int]zapcore.WriteSyncer),
}

func init() {
	// 读取配置
	var err error
	if err = nconf.StructKey("logger", &logger.logConfig); err != nil {
		panic(errors.Wrapf(err, "get logger config error"))
	}
	// 初始化日志
	if err = initLogger(logger.logConfig.Path, logger.logConfig.Details); err != nil {
		panic(errors.Wrapf(err, "init logger error"))
	}
}

// 初始化日志
func initLogger(logPath string, details []LogDetailConfig) (err error) {
	detailsLen := len(details)
	if detailsLen == 0 {
		err = errors.Errorf("logger details config empty: %+v", details)
		return
	}
	coreSlice := make([]zapcore.Core, 0, detailsLen)
	for index, conf := range details {
		// 日志打印级别
		if conf.Level < 0 || conf.Level > 6 {
			err = errors.Errorf("logger details config `level[%d]` undefined", conf.Level)
			return
		}
		level := zapcore.Level(conf.Level - 1)
		var levelEnabler zapcore.LevelEnabler
		switch conf.Type {
		case LOGTYPE_ALL:
			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl <= zapcore.FatalLevel && lvl >= level
			})
		case LOGTYPE_INFO:
			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl < zapcore.ErrorLevel && lvl >= level
			})
		case LOGTYPE_ERROR:
			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.ErrorLevel && lvl >= level
			})
		default:
			err = errors.Errorf("logger details config `type[%d]` undefined", conf.Type)
			return
		}
		// 获取日志输出方式
		var writeSyncer zapcore.WriteSyncer
		if writeSyncer, err = getWriter(logPath, conf); err != nil {
			return
		}
		logger.logWriter[index] = writeSyncer
		// 获取编码器
		var encoder zapcore.Encoder
		if encoder, err = getEncoder(conf); err != nil {
			return
		}
		// 新建Core
		core := zapcore.NewCore(encoder, writeSyncer, levelEnabler)
		coreSlice = append(coreSlice, core)
	}
	// 新建Logger
	coreTee := zapcore.NewTee(coreSlice...)
	logger.zapLogger = zap.New(
		coreTee,
		zap.ErrorOutput(zapcore.Lock(zapcore.AddSync(os.Stderr))),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	) // zap.Addcaller()输出日志打印文件和行数
	return
}

// Write
func (l *stdoutWriter) Write(p []byte) (int, error) {
	if l.format == FORMAT_LOGFMT {
		// 去掉大括号
		p = bytes.ReplaceAll(p, []byte("{"), []byte(""))
		p = bytes.ReplaceAll(p, []byte("}"), []byte(""))
		// 去掉引号
		p = bytes.ReplaceAll(p, []byte("\""), []byte(""))
	}
	return os.Stdout.Write(p)
}

// Write
func (l *fileWriter) Write(p []byte) (int, error) {
	if l.format == FORMAT_LOGFMT {
		// 去掉大括号
		p = bytes.ReplaceAll(p, []byte("{"), []byte(""))
		p = bytes.ReplaceAll(p, []byte("}"), []byte(""))
		// 去掉引号
		p = bytes.ReplaceAll(p, []byte("\""), []byte(""))
	}
	return l.Logger.Write(p)
}

// 获取编码器(如何写入日志)
func getEncoder(conf LogDetailConfig) (encoder zapcore.Encoder, err error) {
	encoderConfig := zap.NewProductionEncoderConfig() // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "file"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stack"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	switch conf.Format {
	case FORMAT_LOGFMT:
		encoderConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString("[" + l.CapitalString() + "]")
		}
		encoderConfig.ConsoleSeparator = " "
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
		return
	case FORMAT_JSON:
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.FunctionKey = "func"
		encoder = zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
		return
	default:
		err = errors.Errorf("logger details config `format[%d]` undefined", conf.Format)
		return
	}
}

// 获取日志输出方式
func getWriter(logPath string, conf LogDetailConfig) (writeSyncer zapcore.WriteSyncer, err error) {
	// 判断日志路径是否存在，如果不存在就创建
	logPath = strings.TrimSpace(logPath)
	if !nfile.PathExists(logPath) {
		if logPath == "" {
			logPath = defaultPath
		}
		if err = os.MkdirAll(logPath, os.ModePerm); err != nil {
			logPath = defaultPath
			if err = os.MkdirAll(logPath, os.ModePerm); err != nil {
				return
			}
		}
	}
	logger.logConfig.Path = logPath
	// 日志文件与日志切割配置
	filenameList := nstr.Split(conf.Filename, ".")
	filenameListLen := len(filenameList)
	filename := ""
	if filenameListLen == 1 {
		filename = fmt.Sprintf("%s-%s.log", filenameList[0], time.Now().Format("2006-01-02"))
	} else if filenameListLen >= 2 {
		filename = fmt.Sprintf("%s-%s.%s", strings.Join(filenameList[:filenameListLen-1], "-"), time.Now().Format("2006-01-02"), filenameList[filenameListLen-1])
	} else {
		filename = fmt.Sprintf("nova-%s.log", time.Now().Format("2006-01-02"))
	}
	logFileWriter := &fileWriter{
		Logger: lumberjack.Logger{
			Filename:   filepath.Join(logPath, filename),
			MaxSize:    conf.MaxSize,    // 单个日志文件最多存储量，单位(mb)，超过则切割
			MaxBackups: conf.MaxBackups, // 日志备份文件最多数量，超过就删除最老的日志文件
			MaxAge:     conf.MaxAge,     // 日志保留时间，单位:天(day)
			Compress:   conf.Compress,   // 是否压缩日志
		},
		format: conf.Format,
	}
	// 日志输出方式
	if conf.Stdout {
		// 日志同时输出到控制台和日志文件中
		writeSyncer = zapcore.NewMultiWriteSyncer(
			zapcore.Lock(zapcore.AddSync(logFileWriter)),
			zapcore.Lock(zapcore.AddSync(&stdoutWriter{format: conf.Format})),
		)
	} else {
		// 日志只输出到日志文件
		writeSyncer = zapcore.Lock(zapcore.AddSync(logFileWriter))
	}
	return
}

// Write 使用日志输出对象写入数据
func Write(p []byte, withoutLogType ...int) (err error) {
	for index, writer := range logger.logWriter {
		conf := logger.logConfig.Details[index]
		if !nslice.IsContains(withoutLogType, conf.Type) {
			if _, err = writer.Write(p); err != nil {
				return
			}
			if err = writer.Sync(); err != nil {
				return
			}
		}
	}
	return
}

// Debug
func Debug(msg string, fields ...logField) {
	logger.zapLogger.Debug(msg, fields...)
}

// Info
func Info(msg string, fields ...logField) {
	logger.zapLogger.Info(msg, fields...)
}

// Warn
func Warn(msg string, fields ...logField) {
	logger.zapLogger.Warn(msg, fields...)
}

// Error
func Error(msg string, fields ...logField) {
	logger.zapLogger.Error(msg, fields...)
}

// DPanic
func DPanic(msg string, fields ...logField) {
	logger.zapLogger.DPanic(msg, fields...)
}

// Panic
func Panic(msg string, fields ...logField) {
	logger.zapLogger.Panic(msg, fields...)
}

// Fatal
func Fatal(msg string, fields ...logField) {
	logger.zapLogger.Fatal(msg, fields...)
}

// Level
func Level() zapcore.Level {
	return logger.zapLogger.Level()
}

// LevelEnabled
func LevelEnabled(lvl zapcore.Level) bool {
	return logger.zapLogger.Level().Enabled(lvl)
}
