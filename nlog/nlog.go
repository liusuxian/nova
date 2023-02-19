/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 19:32:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-20 01:24:30
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nlog/nlog.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nlog

import (
	"fmt"
	"github.com/liusuxian/nova/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

// TODO LogConfig 日志配置
type LogConfig struct {
	Level      string // 日志打印级别 debug、info、warn、error、dpanic、panic、fatal
	Format     string // 输出日志格式 logfmt、json
	Path       string // 输出日志文件路径
	Filename   string // 输出日志文件名称
	MaxSize    int    // 单个日志文件最多存储量，单位(mb)
	MaxBackups int    // 日志备份文件最多数量
	MaxAge     int    // 日志保留时间，单位:天(day)
	Compress   bool   // 是否压缩日志
	Stdout     bool   // 是否输出到控制台
}

// 默认输出日志文件路径
const defaultPath = "log"

// 只能输出结构化日志，但是性能要高于SugaredLogger
var logger *zap.Logger

// 可以输出结构化日志、非结构化日志。性能差于zap.Logger
var sugarLogger *zap.SugaredLogger

// 日志打印级别
var logLevel = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func init() {
	initLogger()
}

// 初始化 Logger
func initLogger() {
	// TODO
	conf := LogConfig{
		Level:      "debug",
		Format:     "json",
		Filename:   "system",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     1,
		Compress:   false,
		Stdout:     true,
	}
	// 获取日志输出方式
	var writeSyncer zapcore.WriteSyncer
	var err error
	if writeSyncer, err = getWriter(conf); err != nil {
		panic(err)
	}
	// 获取编码器
	encoder := getEncoder(conf)
	// 日志打印级别
	var level zapcore.Level
	var ok bool
	if level, ok = logLevel[conf.Level]; !ok {
		level = logLevel["info"]
	}
	// 创建Logger
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)) // zap.Addcaller()输出日志打印文件和行数
	sugarLogger = logger.Sugar()
}

// 获取编码器(如何写入日志)
func getEncoder(conf LogConfig) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig() // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.MessageKey = "msg"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "file"
	encoderConfig.StacktraceKey = "stack"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	} // 指定时间格式
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendInt64(int64(d) / 1000000)
	}
	if conf.Format == "json" {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	return zapcore.NewConsoleEncoder(encoderConfig)              // 以logfmt格式写入
}

// 获取日志输出方式
func getWriter(conf LogConfig) (writeSyncer zapcore.WriteSyncer, err error) {
	// 判断日志路径是否存在，如果不存在就创建
	if !utils.PathExists(conf.Path) {
		if conf.Path == "" {
			conf.Path = defaultPath
		}
		if err = os.MkdirAll(conf.Path, os.ModePerm); err != nil {
			conf.Path = defaultPath
			if err = os.MkdirAll(conf.Path, os.ModePerm); err != nil {
				return
			}
		}
	}
	// 日志文件与日志切割配置
	conf.Filename = fmt.Sprintf("%s-%s.log", conf.Filename, time.Now().Format("2006-01-02"))
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.Path, conf.Filename),
		MaxSize:    conf.MaxSize,    // 单个日志文件最多存储量，单位(mb)，超过则切割
		MaxBackups: conf.MaxBackups, // 日志备份文件最多数量，超过就删除最老的日志文件
		MaxAge:     conf.MaxAge,     // 日志保留时间，单位:天(day)
		Compress:   conf.Compress,   // 是否压缩日志
	}
	// 日志输出方式
	if conf.Stdout {
		// 日志同时输出到控制台和日志文件中
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	} else {
		// 日志只输出到日志文件
		writeSyncer = zapcore.AddSync(lumberJackLogger)
	}
	return
}

// DebugFields
func DebugFields(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// InfoFields
func InfoFields(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// WarnFields
func WarnFields(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// ErrorFields
func ErrorFields(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// DPanicFields
func DPanicFields(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

// PanicFields
func PanicFields(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// FatalFields
func FatalFields(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// Debug
func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

// Debugf
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

// Info
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

// Infof
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

// Warn
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

// Warnf
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

// Error
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

// Errorf
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

// DPanic
func DPanic(args ...interface{}) {
	sugarLogger.DPanic(args...)
}

// DPanicf
func DPanicf(template string, args ...interface{}) {
	sugarLogger.DPanicf(template, args...)
}

// Panic
func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

// Panicf
func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

// Fatal
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

// Fatalf
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}
