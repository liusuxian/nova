/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2023-02-20 16:30:45
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2023-02-21 00:37:49
 * @FilePath: /playlet-server/Users/liusuxian/Desktop/project-code/golang-project/nova/nconf/nconf.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package nconf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/liusuxian/nova/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"time"
)

// Config 配置类
type Config struct {
	v *viper.Viper
}

// New 创建Config
func New(path string) (cfg *Config, err error) {
	v := viper.New()
	v.SetConfigFile(path)
	configType := utils.ExtName(path)
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

// NewRemote 创建远程Config
func NewRemote(provider, endpoint, path, configType string) (cfg *Config, err error) {
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

// NewSecureRemote 创建远程Config
func NewSecureRemote(provider, endpoint, path, secretkeyring, configType string) (cfg *Config, err error) {
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
func (c *Config) Get(key string) interface{} {
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

// GetFloat64 获取 float64
func (c *Config) GetFloat64(key string) float64 {
	return c.v.GetFloat64(key)
}

// GetInt 获取 int
func (c *Config) GetInt(key string) int {
	return c.v.GetInt(key)
}

// GetInt32 获取 int32
func (c *Config) GetInt32(key string) int32 {
	return c.v.GetInt32(key)
}

// GetInt64 获取 int64
func (c *Config) GetInt64(key string) int64 {
	return c.v.GetInt64(key)
}

// GetIntSlice 获取 []int
func (c *Config) GetIntSlice(key string) []int {
	return c.v.GetIntSlice(key)
}

// GetSizeInBytes 获取 uint
func (c *Config) GetSizeInBytes(key string) uint {
	return c.v.GetSizeInBytes(key)
}

// GetString 获取 string
func (c *Config) GetString(key string) string {
	return c.v.GetString(key)
}

// GetStringMap 获取 map[string]interface{}
func (c *Config) GetStringMap(key string) map[string]interface{} {
	return c.v.GetStringMap(key)
}

// GetStringMapString 获取 map[string]string
func (c *Config) GetStringMapString(key string) map[string]string {
	return c.v.GetStringMapString(key)
}

// GetStringMapStringSlice 获取 map[string][]string
func (c *Config) GetStringMapStringSlice(key string) map[string][]string {
	return c.v.GetStringMapStringSlice(key)
}

// GetStringSlice 获取 []string
func (c *Config) GetStringSlice(key string) []string {
	return c.v.GetStringSlice(key)
}

// GetTime 获取 Time
func (c *Config) GetTime(key string) time.Time {
	return c.v.GetTime(key)
}

// GetUint 获取 uint
func (c *Config) GetUint(key string) uint {
	return c.v.GetUint(key)
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

// GetStruct 获取 Struct
func (c *Config) GetStruct(key string, output interface{}) error {
	return mapstructure.Decode(c.v.GetStringMap(key), output)
}

// GetStructs 获取 Structs
func (c *Config) GetStructs(key string, output interface{}) error {
	return mapstructure.Decode(cast.ToSlice(c.v.Get(key)), output)
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

// WatchConfig 监视配置文件的变化
func (c *Config) WatchConfig() {
	c.v.WatchConfig()
}

// Sub 返回一个新的Config实例，表示这个实例的子树，对键不区分大小写
func (c *Config) Sub(key string) *Config {
	return &Config{v: c.v.Sub(key)}
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
			err = errors.Wrapf(err, "No Such Default Config File")
		} else {
			err = errors.Wrapf(err, "Read Default Config Error")
		}
		panic(err)
	}
	defaultConfig = &Config{
		v: v,
	}
}

// Get 获取 value
func Get(key string) interface{} {
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

// GetFloat64 获取 float64
func GetFloat64(key string) float64 {
	return defaultConfig.v.GetFloat64(key)
}

// GetInt 获取 int
func GetInt(key string) int {
	return defaultConfig.v.GetInt(key)
}

// GetInt32 获取 int32
func GetInt32(key string) int32 {
	return defaultConfig.v.GetInt32(key)
}

// GetInt64 获取 int64
func GetInt64(key string) int64 {
	return defaultConfig.v.GetInt64(key)
}

// GetIntSlice 获取 []int
func GetIntSlice(key string) []int {
	return defaultConfig.v.GetIntSlice(key)
}

// GetSizeInBytes 获取 uint
func GetSizeInBytes(key string) uint {
	return defaultConfig.v.GetSizeInBytes(key)
}

// GetString 获取 string
func GetString(key string) string {
	return defaultConfig.v.GetString(key)
}

// GetStringMap 获取 map[string]interface{}
func GetStringMap(key string) map[string]interface{} {
	return defaultConfig.v.GetStringMap(key)
}

// GetStringMapString 获取 map[string]string
func GetStringMapString(key string) map[string]string {
	return defaultConfig.v.GetStringMapString(key)
}

// GetStringMapStringSlice 获取 map[string][]string
func GetStringMapStringSlice(key string) map[string][]string {
	return defaultConfig.v.GetStringMapStringSlice(key)
}

// GetStringSlice 获取 []string
func GetStringSlice(key string) []string {
	return defaultConfig.v.GetStringSlice(key)
}

// GetTime 获取 Time
func GetTime(key string) time.Time {
	return defaultConfig.v.GetTime(key)
}

// GetUint 获取 uint
func GetUint(key string) uint {
	return defaultConfig.v.GetUint(key)
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

// GetStruct 获取 Struct
func GetStruct(key string, output interface{}) error {
	return mapstructure.Decode(defaultConfig.v.GetStringMap(key), output)
}

// GetStructs 获取 Structs
func GetStructs(key string, output interface{}) error {
	return mapstructure.Decode(cast.ToSlice(defaultConfig.v.Get(key)), output)
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

// WatchConfig 监视配置文件的变化
func WatchConfig() {
	defaultConfig.v.WatchConfig()
}

// Sub 返回一个新的Config实例，表示这个实例的子树，对键不区分大小写
func Sub(key string) *Config {
	return &Config{v: defaultConfig.v.Sub(key)}
}
