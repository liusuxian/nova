/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-04-12 18:19:13
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-05-25 19:00:45
 * @Description:
 *
 * Copyright (c) 2023 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package nconf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/liusuxian/nova/utils/nconv"
	"github.com/liusuxian/nova/utils/nenv"
	"github.com/liusuxian/nova/utils/nfile"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"strings"
	"time"
	"unicode"
)

// DecoderConfig 解码配置
type DecoderConfig = mapstructure.DecoderConfig

// DecoderConfigOption 解码配置选项
type DecoderConfigOption func(dc *DecoderConfig)

// Event 事件
type Event = fsnotify.Event

// 操作
const (
	Create = fsnotify.Create
	Write  = fsnotify.Write
	Remove = fsnotify.Remove
	Rename = fsnotify.Rename
	Chmod  = fsnotify.Chmod
)

// Config 配置结构
type Config struct {
	v *viper.Viper
}

// NewConfig 新建Config
func NewConfig(path string) (cfg *Config, err error) {
	v := viper.New()
	v.SetConfigFile(path)
	configType := nfile.ExtName(path)
	v.SetConfigType(configType)
	// 加载配置文件内容
	if err = v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.Wrapf(err, "no such config file, path: %s, configType: %s", path, configType)
		} else {
			err = errors.Wrapf(err, "read config error, path: %s, configType: %s", path, configType)
		}
		return
	}
	cfg = &Config{v: v}
	return
}

// NewRemoteConfig 新建远程Config
func NewRemoteConfig(provider, endpoint, path, configType string) (cfg *Config, err error) {
	v := viper.New()
	v.AddRemoteProvider(provider, endpoint, path)
	v.SetConfigType(configType)
	// 加载配置文件内容
	if err = v.ReadRemoteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.Wrapf(err, "no such config file, provider: %s, endpoint: %s, path: %s, configType: %s", provider, endpoint, path, configType)
		} else {
			err = errors.Wrapf(err, "read config error, provider: %s, endpoint: %s, path: %s, configType: %s", provider, endpoint, path, configType)
		}
		return
	}
	cfg = &Config{v: v}
	return
}

// NewSecureRemoteConfig 新建远程Config
func NewSecureRemoteConfig(provider, endpoint, path, secretkeyring, configType string) (cfg *Config, err error) {
	v := viper.New()
	v.AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
	v.SetConfigType(configType)
	// 加载配置文件内容
	if err = v.ReadRemoteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.Wrapf(err, "no such config file, provider: %s, endpoint: %s, path: %s, configType: %s", provider, endpoint, path, configType)
		} else {
			err = errors.Wrapf(err, "read config error, provider: %s, endpoint: %s, path: %s, configType: %s", provider, endpoint, path, configType)
		}
		return
	}
	cfg = &Config{v: v}
	return
}

// Get 获取 value
func (c *Config) Get(key string) (val any) {
	return c.v.Get(key)
}

// GetBool 获取 bool
func (c *Config) GetBool(key string) (val bool) {
	return nconv.ToBool(c.v.Get(key))
}

// GetDuration 获取 Duration
func (c *Config) GetDuration(key string) (val time.Duration) {
	return nconv.ToDuration(c.v.Get(key))
}

// GetFloat32 获取 float32
func (c *Config) GetFloat32(key string) (val float32) {
	return nconv.ToFloat32(c.v.Get(key))
}

// GetFloat64 获取 float64
func (c *Config) GetFloat64(key string) (val float64) {
	return nconv.ToFloat64(c.v.Get(key))
}

// GetInt 获取 int
func (c *Config) GetInt(key string) (val int) {
	return nconv.ToInt(c.v.Get(key))
}

// GetInt8 获取 int8
func (c *Config) GetInt8(key string) (val int8) {
	return nconv.ToInt8(c.v.Get(key))
}

// GetInt16 获取 int16
func (c *Config) GetInt16(key string) (val int16) {
	return nconv.ToInt16(c.v.Get(key))
}

// GetInt32 获取 int32
func (c *Config) GetInt32(key string) (val int32) {
	return nconv.ToInt32(c.v.Get(key))
}

// GetInt64 获取 int64
func (c *Config) GetInt64(key string) (val int64) {
	return nconv.ToInt64(c.v.Get(key))
}

// GetAnySlice 获取 []any
func (c *Config) GetAnySlice(key string) (vals []any) {
	return nconv.ToSlice(c.v.Get(key))
}

// GetBoolSlice 获取 []bool
func (c *Config) GetBoolSlice(key string) (vals []bool) {
	return nconv.ToBoolSlice(c.v.Get(key))
}

// GetStringSlice 获取 []string
func (c *Config) GetStringSlice(key string) (vals []string) {
	return nconv.ToStringSlice(c.v.Get(key))
}

// GetIntSlice 获取 []int
func (c *Config) GetIntSlice(key string) (vals []int) {
	return nconv.ToIntSlice(c.v.Get(key))
}

// GetDurationSlice 获取 []time.Duration
func (c *Config) GetDurationSlice(key string) (vals []time.Duration) {
	return nconv.ToDurationSlice(c.v.Get(key))
}

// GetSizeInBytes 获取某个配置项对应的值所占用的内存大小（以字节为单位）
func (c *Config) GetSizeInBytes(key string) (val uint) {
	sizeStr := nconv.ToString(c.v.Get(key))
	return parseSizeInBytes(sizeStr)
}

// GetString 获取 string
func (c *Config) GetString(key string) (val string) {
	return nconv.ToString(c.v.Get(key))
}

// GetStringMap 获取 map[string]any
func (c *Config) GetStringMap(key string) (val map[string]any) {
	return nconv.ToStringMap(c.v.Get(key))
}

// GetStringMapBool 获取 map[string]bool
func (c *Config) GetStringMapBool(key string) (val map[string]bool) {
	return nconv.ToStringMapBool(c.v.Get(key))
}

// GetStringMapInt 获取 map[string]int
func (c *Config) GetStringMapInt(key string) (val map[string]int) {
	return nconv.ToStringMapInt(c.v.Get(key))
}

// GetStringMapInt64 获取 map[string]int64
func (c *Config) GetStringMapInt64(key string) (val map[string]int64) {
	return nconv.ToStringMapInt64(c.v.Get(key))
}

// GetStringMapString 获取 map[string]string
func (c *Config) GetStringMapString(key string) (val map[string]string) {
	return nconv.ToStringMapString(c.v.Get(key))
}

// GetStringMapStringSlice 获取 map[string][]string
func (c *Config) GetStringMapStringSlice(key string) (val map[string][]string) {
	return nconv.ToStringMapStringSlice(c.v.Get(key))
}

// GetTime 获取 Time
func (c *Config) GetTime(key string) (val time.Time) {
	return nconv.ToTime(c.v.Get(key))
}

// GetUint 获取 uint
func (c *Config) GetUint(key string) (val uint) {
	return nconv.ToUint(c.v.Get(key))
}

// GetUint8 获取 uint8
func (c *Config) GetUint8(key string) (val uint8) {
	return nconv.ToUint8(c.v.Get(key))
}

// GetUint16 获取 uint16
func (c *Config) GetUint16(key string) (val uint16) {
	return nconv.ToUint16(c.v.Get(key))
}

// GetUint32 获取 uint32
func (c *Config) GetUint32(key string) (val uint32) {
	return nconv.ToUint32(c.v.Get(key))
}

// GetUint64 获取 uint64
func (c *Config) GetUint64(key string) (val uint64) {
	return nconv.ToUint64(c.v.Get(key))
}

// InConfig 检查给定的键(或别名)是否在配置文件中
func (c *Config) InConfig(key string) (val bool) {
	return c.v.InConfig(key)
}

// IsSet 检查是否在任何数据位置设置了键。键不区分大小写
func (c *Config) IsSet(key string) (val bool) {
	return c.v.IsSet(key)
}

// OnConfigChange 设置当配置文件更改时调用的事件处理程序(只能用于本地配置文件的变更监听)
func (c *Config) OnConfigChange(run func(e Event)) {
	c.v.OnConfigChange(run)
}

// SetDefault 设置配置项的默认值，对键不区分大小写，仅当通过flag, config或ENV没有提供值时使用默认值
func (c *Config) SetDefault(key string, value any) {
	c.v.SetDefault(key, value)
}

// Sub 返回一个新的Config实例，表示这个实例的子树，对键不区分大小写
func (c *Config) Sub(key string) (conf *Config) {
	return &Config{v: c.v.Sub(key)}
}

// Struct 将配置解析为结构体，确保标签正确设置该结构的字段
func (c *Config) Struct(rawVal any, opts ...DecoderConfigOption) (err error) {
	newOpts := make([]viper.DecoderConfigOption, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, viper.DecoderConfigOption(opt))
	}
	return c.v.Unmarshal(rawVal, newOpts...)
}

// StructExact 将配置解析为结构体，如果在目标结构体中字段不存在则报错
func (c *Config) StructExact(rawVal any, opts ...DecoderConfigOption) (err error) {
	newOpts := make([]viper.DecoderConfigOption, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, viper.DecoderConfigOption(opt))
	}
	return c.v.UnmarshalExact(rawVal, newOpts...)
}

// StructKey 接收一个键并将其解析到结构体中
func (c *Config) StructKey(key string, rawVal any, opts ...DecoderConfigOption) (err error) {
	newOpts := make([]viper.DecoderConfigOption, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, viper.DecoderConfigOption(opt))
	}
	return c.v.UnmarshalKey(key, rawVal, newOpts...)
}

// WatchConfig 监视配置文件的变化
func (c *Config) WatchConfig() {
	c.v.WatchConfig()
}

// WatchRemoteConfig 监视远程配置文件的变化(阻塞式)
func (c *Config) WatchRemoteConfig() (err error) {
	return c.v.WatchRemoteConfig()
}

// WatchRemoteConfigOnChannel 监视远程配置文件的变化(非阻塞式)
func (c *Config) WatchRemoteConfigOnChannel() (err error) {
	return c.v.WatchRemoteConfigOnChannel()
}

// 默认配置（包含启动配置文件）
var defaultConfig *Config

func init() {
	v := viper.New()
	v.SetConfigName("config")             // 设置配置文件名，不需要配置文件扩展名，配置文件的类型会自动根据扩展名自动匹配
	v.AddConfigPath("./")                 // 设置配置文件的搜索目录
	v.AddConfigPath("./config/")          // 设置配置文件的搜索目录
	v.AddConfigPath("./manifest/config/") // 设置配置文件的搜索目录
	if nenv.Contains("NOVA_CONFIG_FILE_PATH") {
		v.AddConfigPath(nenv.Get("NOVA_CONFIG_FILE_PATH")) // 设置配置文件的搜索目录
	}
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			panic(errors.Wrapf(err, "read default config error"))
		}
	}
	defaultConfig = &Config{
		v: v,
	}
	// 设置默认值
	// 服务器配置
	SetDefault("server.name", "Nova")        // 服务器应用名称，默认"Nova"
	SetDefault("server.heartBeat", "10s")    // 心跳发送间隔时间（一定要小于 maxHeartBeat 配置），默认 10秒
	SetDefault("server.maxHeartBeat", "15s") // 最长心跳检测间隔时间（一定要大于 heartBeat 配置），默认 15秒
	SetDefault("server.maxConn", 3)          // 允许的客户端连接最大数量，默认 3
	SetDefault("server.workerPoolSize", 10)  // 工作任务池最大工作 Goroutine 数量，默认 10
	SetDefault("server.maxPacketSize", 4096) // 数据包的最大值（单位:字节），默认 4096
	SetDefault("server.packetMethod", 1)     // 封包和拆包方式，1: 消息ID(2字节)-消息体长度(4字节)-消息内容，默认 1
	SetDefault("server.endian", 1)           // 字节存储次序，1: 小端 2: 大端，默认 1
	// 日志配置
	SetDefault("logger.path", "logs")             // 输出日志文件路径
	SetDefault("logger.details.type", 0)          // 日志类型 0:打印所有级别 1:打印 DEBUG、INFO、WARN 级别 2:打印 ERROR、DPANIC、PANIC、FATAL 级别，默认0
	SetDefault("logger.details.level", 0)         // 日志打印级别 0:DEBUG 1:INFO 2:WARN 3:ERROR 4:DPANIC、5:PANIC、6:FATAL，默认0
	SetDefault("logger.details.format", 1)        // 输出日志格式 0:logfmt 1:json，默认1
	SetDefault("logger.details.filename", "nova") // 输出日志文件名称
	SetDefault("logger.details.maxSize", 10)      // 单个日志文件最多存储量（单位:MB）
	SetDefault("logger.details.maxBackups", 10)   // 日志备份文件最多数量
	SetDefault("logger.details.maxAge", 7)        // 日志保留时间（单位:天）
	SetDefault("logger.details.compress", false)  // 是否压缩日志
	SetDefault("logger.details.stdout", true)     // 是否输出到控制台
}

// Get 获取 value
func Get(key string) (val any) {
	return defaultConfig.v.Get(key)
}

// GetBool 获取 bool
func GetBool(key string) (val bool) {
	return nconv.ToBool(defaultConfig.v.Get(key))
}

// GetDuration 获取 Duration
func GetDuration(key string) (val time.Duration) {
	return nconv.ToDuration(defaultConfig.v.Get(key))
}

// GetFloat32 获取 float32
func GetFloat32(key string) (val float32) {
	return nconv.ToFloat32(defaultConfig.v.Get(key))
}

// GetFloat64 获取 float64
func GetFloat64(key string) (val float64) {
	return nconv.ToFloat64(defaultConfig.v.Get(key))
}

// GetInt 获取 int
func GetInt(key string) (val int) {
	return nconv.ToInt(defaultConfig.v.Get(key))
}

// GetInt8 获取 int8
func GetInt8(key string) (val int8) {
	return nconv.ToInt8(defaultConfig.v.Get(key))
}

// GetInt16 获取 int16
func GetInt16(key string) (val int16) {
	return nconv.ToInt16(defaultConfig.v.Get(key))
}

// GetInt32 获取 int32
func GetInt32(key string) (val int32) {
	return nconv.ToInt32(defaultConfig.v.Get(key))
}

// GetInt64 获取 int64
func GetInt64(key string) (val int64) {
	return nconv.ToInt64(defaultConfig.v.Get(key))
}

// GetAnySlice 获取 []any
func GetAnySlice(key string) (vals []any) {
	return nconv.ToSlice(defaultConfig.v.Get(key))
}

// GetBoolSlice 获取 []bool
func GetBoolSlice(key string) (vals []bool) {
	return nconv.ToBoolSlice(defaultConfig.v.Get(key))
}

// GetStringSlice 获取 []string
func GetStringSlice(key string) (vals []string) {
	return nconv.ToStringSlice(defaultConfig.v.Get(key))
}

// GetIntSlice 获取 []int
func GetIntSlice(key string) (vals []int) {
	return nconv.ToIntSlice(defaultConfig.v.Get(key))
}

// GetDurationSlice 获取 []time.Duration
func GetDurationSlice(key string) (vals []time.Duration) {
	return nconv.ToDurationSlice(defaultConfig.v.Get(key))
}

// GetSizeInBytes 获取某个配置项对应的值所占用的内存大小（以字节为单位）
func GetSizeInBytes(key string) (val uint) {
	sizeStr := nconv.ToString(defaultConfig.v.Get(key))
	return parseSizeInBytes(sizeStr)
}

// GetString 获取 string
func GetString(key string) (val string) {
	return nconv.ToString(defaultConfig.v.Get(key))
}

// GetStringMap 获取 map[string]any
func GetStringMap(key string) (val map[string]any) {
	return nconv.ToStringMap(defaultConfig.v.Get(key))
}

// GetStringMapBool 获取 map[string]bool
func GetStringMapBool(key string) (val map[string]bool) {
	return nconv.ToStringMapBool(defaultConfig.v.Get(key))
}

// GetStringMapInt 获取 map[string]int
func GetStringMapInt(key string) (val map[string]int) {
	return nconv.ToStringMapInt(defaultConfig.v.Get(key))
}

// GetStringMapInt64 获取 map[string]int64
func GetStringMapInt64(key string) (val map[string]int64) {
	return nconv.ToStringMapInt64(defaultConfig.v.Get(key))
}

// GetStringMapString 获取 map[string]string
func GetStringMapString(key string) (val map[string]string) {
	return nconv.ToStringMapString(defaultConfig.v.Get(key))
}

// GetStringMapStringSlice 获取 map[string][]string
func GetStringMapStringSlice(key string) (val map[string][]string) {
	return nconv.ToStringMapStringSlice(defaultConfig.v.Get(key))
}

// GetTime 获取 Time
func GetTime(key string) (val time.Time) {
	return nconv.ToTime(defaultConfig.v.Get(key))
}

// GetUint 获取 uint
func GetUint(key string) (val uint) {
	return nconv.ToUint(defaultConfig.v.Get(key))
}

// GetUint8 获取 uint8
func GetUint8(key string) (val uint8) {
	return nconv.ToUint8(defaultConfig.v.Get(key))
}

// GetUint16 获取 uint16
func GetUint16(key string) (val uint16) {
	return nconv.ToUint16(defaultConfig.v.Get(key))
}

// GetUint32 获取 uint32
func GetUint32(key string) (val uint32) {
	return nconv.ToUint32(defaultConfig.v.Get(key))
}

// GetUint64 获取 uint64
func GetUint64(key string) (val uint64) {
	return nconv.ToUint64(defaultConfig.v.Get(key))
}

// InConfig 检查给定的键(或别名)是否在配置文件中
func InConfig(key string) (val bool) {
	return defaultConfig.v.InConfig(key)
}

// IsSet 检查是否在任何数据位置设置了键。键不区分大小写
func IsSet(key string) (val bool) {
	return defaultConfig.v.IsSet(key)
}

// OnConfigChange 设置当配置文件更改时调用的事件处理程序(只能用于本地配置文件的变更监听)
func OnConfigChange(run func(e Event)) {
	defaultConfig.v.OnConfigChange(run)
}

// SetDefault 设置配置项的默认值，对键不区分大小写，仅当通过flag, config或ENV没有提供值时使用默认值
func SetDefault(key string, value any) {
	defaultConfig.v.SetDefault(key, value)
}

// Sub 返回一个新的Config实例，表示这个实例的子树，对键不区分大小写
func Sub(key string) (conf *Config) {
	return &Config{v: defaultConfig.v.Sub(key)}
}

// Struct 将配置解析为结构体，确保标签正确设置该结构的字段
func Struct(rawVal any, opts ...DecoderConfigOption) (err error) {
	newOpts := make([]viper.DecoderConfigOption, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, viper.DecoderConfigOption(opt))
	}
	return defaultConfig.v.Unmarshal(rawVal, newOpts...)
}

// StructExact 将配置解析为结构体，如果在目标结构体中字段不存在则报错
func StructExact(rawVal any, opts ...DecoderConfigOption) (err error) {
	newOpts := make([]viper.DecoderConfigOption, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, viper.DecoderConfigOption(opt))
	}
	return defaultConfig.v.UnmarshalExact(rawVal, newOpts...)
}

// StructKey 接收一个键并将其解析到结构体中
func StructKey(key string, rawVal any, opts ...DecoderConfigOption) (err error) {
	newOpts := make([]viper.DecoderConfigOption, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, viper.DecoderConfigOption(opt))
	}
	return defaultConfig.v.UnmarshalKey(key, rawVal, newOpts...)
}

// WatchConfig 监视配置文件的变化
func WatchConfig() {
	defaultConfig.v.WatchConfig()
}

// parseSizeInBytes 将像1GB或12MB这样的字符串转换为无符号整数字节数
func parseSizeInBytes(sizeStr string) (s uint) {
	sizeStr = strings.TrimSpace(sizeStr)
	lastChar := len(sizeStr) - 1
	multiplier := uint(1)

	if lastChar > 0 {
		if sizeStr[lastChar] == 'b' || sizeStr[lastChar] == 'B' {
			if lastChar > 1 {
				switch unicode.ToLower(rune(sizeStr[lastChar-1])) {
				case 'k':
					multiplier = 1 << 10
					sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
				case 'm':
					multiplier = 1 << 20
					sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
				case 'g':
					multiplier = 1 << 30
					sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
				default:
					multiplier = 1
					sizeStr = strings.TrimSpace(sizeStr[:lastChar])
				}
			}
		}
	}

	size := nconv.ToInt(sizeStr)
	if size < 0 {
		size = 0
	}

	return safeMul(uint(size), multiplier)
}

func safeMul(a, b uint) (s uint) {
	c := a * b
	if a > 1 && b > 1 && c/b != a {
		return 0
	}
	return c
}
