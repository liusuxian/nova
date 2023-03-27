/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-03-08 19:20:35
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-03-27 18:17:05
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nlog/log.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nlog

import (
	"context"
	"fmt"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/utils/nfile"
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
	CtxKeys []string          // 自定义 Context 上下文变量名称，自动打印 Context 的变量到日志中。默认为空
	Path    string            // 输出日志文件路径
	Details []LogDetailConfig // 日志详细配置
}

// LogDetailConfig 日志详细配置
type LogDetailConfig struct {
	Type       int    // 日志类型 0:打印所有级别 1:打印 DEBUG、INFO、WARN 级别 2:打印 ERROR、DPANIC、PANIC、FATAL 级别，默认0
	Level      int    // 日志打印级别 0:DEBUG 1:INFO 2:WARN 3:ERROR 4:DPANIC、5:PANIC、6:FATAL，默认0
	Format     int    // 输出日志格式 0:logfmt 1:json，默认0
	Filename   string // 输出日志文件名称
	MaxSize    int    // 单个日志文件最多存储量（单位:MB）
	MaxBackups int    // 日志备份文件最多数量
	MaxAge     int    // 日志保留时间（单位:天）
	Compress   bool   // 是否压缩日志
	Stdout     bool   // 是否输出到控制台
}

const (
	// 默认输出日志文件路径
	defaultPath = "logs"
	// 日志类型
	LOGTYPE_ALL   = 0 // 打印所有级别
	LOGTYPE_INFO  = 1 // 打印 DEBUG、INFO、WARN 级别
	LOGTYPE_ERROR = 2 // 打印 ERROR、DPANIC、PANIC、FATAL 级别
	// 输出日志格式
	FORMAT_LOGFMT = 0
	FORMAT_JSON   = 1
)

var (
	// 只能输出结构化日志，但是性能要高于SugaredLogger
	logger *zap.Logger
	// 输出日志文件路径
	loggerPath string
	// 日志配置
	logConfig LogConfig
)

func init() {
	// 读取配置
	var err error
	if err = nconf.StructKey("logger", &logConfig); err != nil {
		panic(errors.Wrapf(err, "Get Logger Config Error"))
	}
	// 初始化日志
	if err = initLogger(logConfig.Path, logConfig.Details); err != nil {
		panic(errors.Wrapf(err, "Init Logger Error"))
	}
}

// 初始化日志
func initLogger(logPath string, details []LogDetailConfig) (err error) {
	detailsLen := len(details)
	if detailsLen == 0 {
		err = errors.Errorf("Logger Details Config Empty: %+v", details)
		return
	}
	coreSlice := make([]zapcore.Core, 0, detailsLen)
	for _, conf := range details {
		// 日志打印级别
		if conf.Level < 0 || conf.Level > 6 {
			err = errors.Errorf("Logger Details Config `Level[%d]` Undefined", conf.Level)
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
			err = errors.Errorf("Logger Details Config `Type[%d]` Undefined", conf.Type)
			return
		}
		// 获取日志输出方式
		var writeSyncer zapcore.WriteSyncer
		if writeSyncer, err = getWriter(logPath, conf); err != nil {
			return
		}
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
	logger = zap.New(coreTee, zap.ErrorOutput(zapcore.AddSync(zapcore.Lock(os.Stderr))), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)) // zap.Addcaller()输出日志打印文件和行数
	return
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
	switch conf.Format {
	case FORMAT_LOGFMT:
		encoderConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString("[" + l.CapitalString() + "]")
		}
		encoderConfig.EncodeCaller = func(ec zapcore.EntryCaller, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(ec.TrimmedPath() + ":")
		}
		encoderConfig.ConsoleSeparator = " "
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
		return
	case FORMAT_JSON:
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.FunctionKey = "func"
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
		return
	default:
		err = errors.Errorf("Logger Details Config `Format[%d]` Undefined", conf.Format)
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
	loggerPath = logPath
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
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logPath, filename),
		MaxSize:    conf.MaxSize,    // 单个日志文件最多存储量，单位(mb)，超过则切割
		MaxBackups: conf.MaxBackups, // 日志备份文件最多数量，超过就删除最老的日志文件
		MaxAge:     conf.MaxAge,     // 日志保留时间，单位:天(day)
		Compress:   conf.Compress,   // 是否压缩日志
	}
	// 日志输出方式
	if conf.Stdout {
		// 日志同时输出到控制台和日志文件中
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(zapcore.Lock(os.Stdout)))
	} else {
		// 日志只输出到日志文件
		writeSyncer = zapcore.AddSync(lumberJackLogger)
	}
	return
}

// Debug
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	withCtxLogger(ctx, fields...).Debug(msg)
}

// Info
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	withCtxLogger(ctx, fields...).Info(msg)
}

// Warn
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	withCtxLogger(ctx, fields...).Warn(msg)
}

// Error
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	withCtxLogger(ctx, fields...).Error(msg)
}

// DPanic
func DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	withCtxLogger(ctx, fields...).DPanic(msg)
}

// Panic
func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	withCtxLogger(ctx, fields...).Panic(msg)
}

// Fatal
func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	withCtxLogger(ctx, fields...).Fatal(msg)
}

// Level
func Level() zapcore.Level {
	return logger.Level()
}

// LevelEnabled
func LevelEnabled(lvl zapcore.Level) bool {
	return logger.Level().Enabled(lvl)
}

// 获取输出日志文件路径
func GetLoggerPath() string {
	return loggerPath
}

func withCtxLogger(ctx context.Context, fields ...zap.Field) *zap.Logger {
	if ctx == nil {
		return logger.With(fields...)
	}
	ctxKeys := logConfig.CtxKeys
	ctxKeysLen := len(ctxKeys)
	if ctxKeysLen == 0 {
		return logger.With(fields...)
	}
	fieldList := make([]zapcore.Field, 0, ctxKeysLen)
	for _, key := range ctxKeys {
		fieldList = append(fieldList, zap.Any(key, ctx.Value(key)))
	}
	return logger.With(fieldList...).With(fields...)
}
