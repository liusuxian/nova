/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-19 19:32:52
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 18:13:35
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nlog/nlog.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nlog

import (
	"context"
	"fmt"
	"github.com/liusuxian/nova/nconf"
	"github.com/liusuxian/nova/utils/file"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LogConfig 日志配置
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
const defaultPath = "logs"

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

// 只能输出结构化日志，但是性能要高于SugaredLogger
var logger *zap.Logger

func init() {
	// 读取配置
	var err error
	confSlice := []LogConfig{}
	if err = nconf.StructKey("logger", &confSlice); err != nil {
		log.Fatalf("Get Logger Config Error: %+v\n", err)
	}
	// 初始化日志
	if err = initLogger(confSlice); err != nil {
		log.Fatalf("Init Logger Error: %+v\n", err)
	}
	log.Println("Init Logger Succeed !!!")
}

// 初始化日志
func initLogger(confSlice []LogConfig) (err error) {
	confSliceLen := len(confSlice)
	if confSliceLen == 0 {
		err = errors.Errorf("Logger Config Empty: %+v", confSlice)
		return
	}
	coreSlice := make([]zapcore.Core, 0, confSliceLen)
	for _, conf := range confSlice {
		// 获取日志输出方式
		var writeSyncer zapcore.WriteSyncer
		if writeSyncer, err = getWriter(conf); err != nil {
			return
		}
		// 获取编码器
		encoder := getEncoder(conf)
		// 日志打印级别
		var level zapcore.Level
		var ok bool
		if level, ok = logLevel[conf.Level]; !ok {
			level = logLevel["info"]
		}
		// 新建Core
		core := zapcore.NewCore(encoder, writeSyncer, level)
		coreSlice = append(coreSlice, core)
	}
	// 新建Logger
	coreTee := zapcore.NewTee(coreSlice...)
	logger = zap.New(coreTee, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)) // zap.Addcaller()输出日志打印文件和行数
	return
}

// 获取编码器(如何写入日志)
func getEncoder(conf LogConfig) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig() // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "file"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stack"
	encoderConfig.FunctionKey = "func"
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
	if !file.PathExists(conf.Path) {
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
	filenameList := strings.Split(conf.Filename, ".")
	filename := fmt.Sprintf("%s-%s.%s", filenameList[0], time.Now().Format("2006-01-02"), filenameList[1])
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.Path, filename),
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

func withCtxLogger(ctx context.Context, fields ...zap.Field) *zap.Logger {
	if ctx == nil {
		return logger.With(fields...)
	}
	var customCtxKey string
	if customCtxKey = nconf.GetString("customCtxKey"); customCtxKey == "" {
		return logger.With(fields...)
	}
	return logger.With(zap.Reflect(customCtxKey, ctx.Value(customCtxKey))).With(fields...)
}
