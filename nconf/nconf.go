/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 16:30:45
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-22 19:53:55
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconf/nconf.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/liusuxian/nova/nutils/nfile"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"time"
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
			err = errors.Wrapf(err, "No Such Config File, Path: %s, ConfigType: %s", path, configType)
		} else {
			err = errors.Wrapf(err, "Read Config Error, Path: %s, ConfigType: %s", path, configType)
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
			err = errors.Wrapf(err, "No Such Config File, Provider: %s, Endpoint: %s, Path: %s, Configtype: %s", provider, endpoint, path, configType)
		} else {
			err = errors.Wrapf(err, "Read Config Error, Provider: %s, Endpoint: %s, Path: %s, Configtype: %s", provider, endpoint, path, configType)
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
			err = errors.Wrapf(err, "No Such Config File, Provider: %s, Endpoint: %s, Path: %s, Configtype: %s", provider, endpoint, path, configType)
		} else {
			err = errors.Wrapf(err, "Read Config Error, Provider: %s, Endpoint: %s, Path: %s, Configtype: %s", provider, endpoint, path, configType)
		}
		return
	}
	cfg = &Config{v: v}
	return
}

// Get 获取 value
func (c *Config) Get(key string) any {
	return c.v.Get(key)
}

// GetBool 获取 bool
func (c *Config) GetBool(key string) bool {
	return c.v.GetBool(key)
}

// GetDuration 获取 Duration
func (c *Config) GetDuration(key string) time.Duration {
	return c.v.GetDuration(key)
}

// GetFloat32 获取 float32
func (c *Config) GetFloat32(key string) float32 {
	return cast.ToFloat32(c.v.Get(key))
}

// GetFloat64 获取 float64
func (c *Config) GetFloat64(key string) float64 {
	return c.v.GetFloat64(key)
}

// GetInt 获取 int
func (c *Config) GetInt(key string) int {
	return c.v.GetInt(key)
}

// GetInt8 获取 int8
func (c *Config) GetInt8(key string) int8 {
	return cast.ToInt8(c.v.Get(key))
}

// GetInt16 获取 int16
func (c *Config) GetInt16(key string) int16 {
	return cast.ToInt16(c.v.Get(key))
}

// GetInt32 获取 int32
func (c *Config) GetInt32(key string) int32 {
	return c.v.GetInt32(key)
}

// GetInt64 获取 int64
func (c *Config) GetInt64(key string) int64 {
	return c.v.GetInt64(key)
}

// GetAnySlice 获取 []any
func (c *Config) GetAnySlice(key string) []any {
	return cast.ToSlice(c.v.Get(key))
}

// GetBoolSlice 获取 []bool
func (c *Config) GetBoolSlice(key string) []bool {
	return cast.ToBoolSlice(c.v.Get(key))
}

// GetStringSlice 获取 []string
func (c *Config) GetStringSlice(key string) []string {
	return c.v.GetStringSlice(key)
}

// GetIntSlice 获取 []int
func (c *Config) GetIntSlice(key string) []int {
	return c.v.GetIntSlice(key)
}

// GetDurationSlice 获取 []time.Duration
func (c *Config) GetDurationSlice(key string) []time.Duration {
	return cast.ToDurationSlice(c.v.Get(key))
}

// GetSizeInBytes 获取某个配置项对应的值所占用的内存大小（以字节为单位）
func (c *Config) GetSizeInBytes(key string) uint {
	return c.v.GetSizeInBytes(key)
}

// GetString 获取 string
func (c *Config) GetString(key string) string {
	return c.v.GetString(key)
}

// GetStringMap 获取 map[string]any
func (c *Config) GetStringMap(key string) map[string]any {
	return c.v.GetStringMap(key)
}

// GetStringMapBool 获取 map[string]bool
func (c *Config) GetStringMapBool(key string) map[string]bool {
	return cast.ToStringMapBool(c.v.Get(key))
}

// GetStringMapInt 获取 map[string]int
func (c *Config) GetStringMapInt(key string) map[string]int {
	return cast.ToStringMapInt(c.v.Get(key))
}

// GetStringMapInt64 获取 map[string]int64
func (c *Config) GetStringMapInt64(key string) map[string]int64 {
	return cast.ToStringMapInt64(c.v.Get(key))
}

// GetStringMapString 获取 map[string]string
func (c *Config) GetStringMapString(key string) map[string]string {
	return c.v.GetStringMapString(key)
}

// GetStringMapStringSlice 获取 map[string][]string
func (c *Config) GetStringMapStringSlice(key string) map[string][]string {
	return c.v.GetStringMapStringSlice(key)
}

// GetTime 获取 Time
func (c *Config) GetTime(key string) time.Time {
	return c.v.GetTime(key)
}

// GetUint 获取 uint
func (c *Config) GetUint(key string) uint {
	return c.v.GetUint(key)
}

// GetUint8 获取 uint8
func (c *Config) GetUint8(key string) uint8 {
	return cast.ToUint8(c.v.Get(key))
}

// GetUint16 获取 uint16
func (c *Config) GetUint16(key string) uint16 {
	return c.v.GetUint16(key)
}

// GetUint32 获取 uint32
func (c *Config) GetUint32(key string) uint32 {
	return c.v.GetUint32(key)
}

// GetUint64 获取 uint64
func (c *Config) GetUint64(key string) uint64 {
	return c.v.GetUint64(key)
}

// InConfig 检查给定的键(或别名)是否在配置文件中
func (c *Config) InConfig(key string) bool {
	return c.v.InConfig(key)
}

// IsSet 检查是否在任何数据位置设置了键。键不区分大小写
func (c *Config) IsSet(key string) bool {
	return c.v.IsSet(key)
}

// OnConfigChange 设置当配置文件更改时调用的事件处理程序
func (c *Config) OnConfigChange(run func(in fsnotify.Event)) {
	c.v.OnConfigChange(run)
}

// SetDefault 设置配置项的默认值，对键不区分大小写，仅当通过flag, config或ENV没有提供值时使用默认值
func (c *Config) SetDefault(key string, value any) {
	c.v.SetDefault(key, value)
}

// Sub 返回一个新的Config实例，表示这个实例的子树，对键不区分大小写
func (c *Config) Sub(key string) *Config {
	return &Config{v: c.v.Sub(key)}
}

// Struct 将配置解析为结构体，确保标签正确设置该结构的字段
func (c *Config) Struct(rawVal any, opts ...viper.DecoderConfigOption) error {
	return c.v.Unmarshal(rawVal, opts...)
}

// StructExact 将配置解析为结构体，如果在目标结构体中字段不存在则报错
func (c *Config) StructExact(rawVal any, opts ...viper.DecoderConfigOption) error {
	return c.v.UnmarshalExact(rawVal, opts...)
}

// StructKey 接收一个键并将其解析到结构体中
func (c *Config) StructKey(key string, rawVal any, opts ...viper.DecoderConfigOption) error {
	return c.v.UnmarshalKey(key, rawVal, opts...)
}

// WatchConfig 监视配置文件的变化
func (c *Config) WatchConfig() {
	c.v.WatchConfig()
}

// WatchRemoteConfig 监视远程配置文件的变化
func (c *Config) WatchRemoteConfig() error {
	return c.v.WatchRemoteConfig()
}

// WatchRemoteConfigOnChannel 取消监视远程配置文件的变化
func (c *Config) WatchRemoteConfigOnChannel() error {
	return c.v.WatchRemoteConfigOnChannel()
}

// 默认配置（包含启动配置文件）
var defaultConfig *Config

func init() {
	v := viper.New()
	v.SetConfigName("config") // 设置配置文件名，不需要配置文件扩展名，配置文件的类型会自动根据扩展名自动匹配
	v.AddConfigPath(".")      // 设置配置文件的搜索目录
	v.AddConfigPath("config") // 设置配置文件的搜索目录
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No Such Default Config File")
		} else {
			panic(errors.Wrapf(err, "Read Default Config Error"))
		}
	}
	defaultConfig = &Config{
		v: v,
	}
	// 设置默认值
	// 服务器配置
	SetDefault("server.name", "Nova")        // 服务器应用名称，默认"Nova"
	SetDefault("server.maxConn", 3)          // 允许的客户端连接最大数量，默认3（uint32）
	SetDefault("server.workerPoolSize", 10)  // 工作任务池最大工作Goroutine数量，默认10（uint32）
	SetDefault("server.maxPacketSize", 4096) // 数据包的最大值，默认4096（单位:字节 uint32）
	SetDefault("server.packetMethod", 1)     // 封包和拆包方式，默认1，1:消息ID(4字节)-消息体长度(4字节)-消息内容（单位:字节 uint8）
	SetDefault("server.endian", 1)           // 字节存储次序，默认小端，1:小端 2:大端（单位:字节 uint8）
	SetDefault("server.maxMsgChanLen", 1024) // SendBuffMsg发送消息的缓冲最大长度，默认1024（单位:字节 uint32）
	// 日志配置
	SetDefault("logger.level", "info")    // 日志打印级别 debug、info、warn、error、dpanic、panic、fatal
	SetDefault("logger.format", "json")   // 输出日志格式 logfmt、json
	SetDefault("logger.path", "logs")     // 输出日志文件路径
	SetDefault("logger.filename", "nova") // 输出日志文件名称
	SetDefault("logger.maxSize", 10)      // 单个日志文件最多存储量（单位:MB int）
	SetDefault("logger.maxBackups", 10)   // 日志备份文件最多数量（int）
	SetDefault("logger.maxAge", 7)        // 日志保留时间（单位:天 int）
	SetDefault("logger.compress", false)  // 是否压缩日志
	SetDefault("logger.stdout", true)     // 是否输出到控制台
	log.Println("Default Config Create Or Load Succeed !!!")
}

// Get 获取 value
func Get(key string) any {
	return defaultConfig.v.Get(key)
}

// GetBool 获取 bool
func GetBool(key string) bool {
	return defaultConfig.v.GetBool(key)
}

// GetDuration 获取 Duration
func GetDuration(key string) time.Duration {
	return defaultConfig.v.GetDuration(key)
}

// GetFloat32 获取 float32
func GetFloat32(key string) float32 {
	return cast.ToFloat32(defaultConfig.v.Get(key))
}

// GetFloat64 获取 float64
func GetFloat64(key string) float64 {
	return defaultConfig.v.GetFloat64(key)
}

// GetInt 获取 int
func GetInt(key string) int {
	return defaultConfig.v.GetInt(key)
}

// GetInt8 获取 int8
func GetInt8(key string) int8 {
	return cast.ToInt8(defaultConfig.v.Get(key))
}

// GetInt16 获取 int16
func GetInt16(key string) int16 {
	return cast.ToInt16(defaultConfig.v.Get(key))
}

// GetInt32 获取 int32
func GetInt32(key string) int32 {
	return defaultConfig.v.GetInt32(key)
}

// GetInt64 获取 int64
func GetInt64(key string) int64 {
	return defaultConfig.v.GetInt64(key)
}

// GetAnySlice 获取 []any
func GetAnySlice(key string) []any {
	return cast.ToSlice(defaultConfig.v.Get(key))
}

// GetBoolSlice 获取 []bool
func GetBoolSlice(key string) []bool {
	return cast.ToBoolSlice(defaultConfig.v.Get(key))
}

// GetStringSlice 获取 []string
func GetStringSlice(key string) []string {
	return defaultConfig.v.GetStringSlice(key)
}

// GetIntSlice 获取 []int
func GetIntSlice(key string) []int {
	return defaultConfig.v.GetIntSlice(key)
}

// GetDurationSlice 获取 []time.Duration
func GetDurationSlice(key string) []time.Duration {
	return cast.ToDurationSlice(defaultConfig.v.Get(key))
}

// GetSizeInBytes 获取某个配置项对应的值所占用的内存大小（以字节为单位）
func GetSizeInBytes(key string) uint {
	return defaultConfig.v.GetSizeInBytes(key)
}

// GetString 获取 string
func GetString(key string) string {
	return defaultConfig.v.GetString(key)
}

// GetStringMap 获取 map[string]any
func GetStringMap(key string) map[string]any {
	return defaultConfig.v.GetStringMap(key)
}

// GetStringMapBool 获取 map[string]bool
func GetStringMapBool(key string) map[string]bool {
	return cast.ToStringMapBool(defaultConfig.v.Get(key))
}

// GetStringMapInt 获取 map[string]int
func GetStringMapInt(key string) map[string]int {
	return cast.ToStringMapInt(defaultConfig.v.Get(key))
}

// GetStringMapInt64 获取 map[string]int64
func GetStringMapInt64(key string) map[string]int64 {
	return cast.ToStringMapInt64(defaultConfig.v.Get(key))
}

// GetStringMapString 获取 map[string]string
func GetStringMapString(key string) map[string]string {
	return defaultConfig.v.GetStringMapString(key)
}

// GetStringMapStringSlice 获取 map[string][]string
func GetStringMapStringSlice(key string) map[string][]string {
	return defaultConfig.v.GetStringMapStringSlice(key)
}

// GetTime 获取 Time
func GetTime(key string) time.Time {
	return defaultConfig.v.GetTime(key)
}

// GetUint 获取 uint
func GetUint(key string) uint {
	return defaultConfig.v.GetUint(key)
}

// GetUint8 获取 uint8
func GetUint8(key string) uint8 {
	return cast.ToUint8(defaultConfig.v.Get(key))
}

// GetUint16 获取 uint16
func GetUint16(key string) uint16 {
	return defaultConfig.v.GetUint16(key)
}

// GetUint32 获取 uint32
func GetUint32(key string) uint32 {
	return defaultConfig.v.GetUint32(key)
}

// GetUint64 获取 uint64
func GetUint64(key string) uint64 {
	return defaultConfig.v.GetUint64(key)
}

// InConfig 检查给定的键(或别名)是否在配置文件中
func InConfig(key string) bool {
	return defaultConfig.v.InConfig(key)
}

// IsSet 检查是否在任何数据位置设置了键。键不区分大小写
func IsSet(key string) bool {
	return defaultConfig.v.IsSet(key)
}

// OnConfigChange 设置当配置文件更改时调用的事件处理程序
func OnConfigChange(run func(in fsnotify.Event)) {
	defaultConfig.v.OnConfigChange(run)
}

// SetDefault 设置配置项的默认值，对键不区分大小写，仅当通过flag, config或ENV没有提供值时使用默认值
func SetDefault(key string, value any) {
	defaultConfig.v.SetDefault(key, value)
}

// Sub 返回一个新的Config实例，表示这个实例的子树，对键不区分大小写
func Sub(key string) *Config {
	return &Config{v: defaultConfig.v.Sub(key)}
}

// Struct 将配置解析为结构体，确保标签正确设置该结构的字段
func Struct(rawVal any, opts ...viper.DecoderConfigOption) error {
	return defaultConfig.v.Unmarshal(rawVal, opts...)
}

// StructExact 将配置解析为结构体，如果在目标结构体中字段不存在则报错
func StructExact(rawVal any, opts ...viper.DecoderConfigOption) error {
	return defaultConfig.v.UnmarshalExact(rawVal, opts...)
}

// StructKey 接收一个键并将其解析到结构体中
func StructKey(key string, rawVal any, opts ...viper.DecoderConfigOption) error {
	return defaultConfig.v.UnmarshalKey(key, rawVal, opts...)
}

// WatchConfig 监视配置文件的变化
func WatchConfig() {
	defaultConfig.v.WatchConfig()
}
