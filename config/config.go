package config

import (
	"github.com/zieckey/goini"
)

type Config struct {
	ini *goini.INI

	DEBUG bool
}

func (this *Config) String(k, def string) string {
	v, ok := this.ini.Get(k)
	if ok {
		return v
	}
	return def
}

func (this *Config) Bool(k string) bool {
	v, ok := this.ini.GetBool(k)
	if ok {
		return v
	}
	return false
}

func (this *Config) Int(k string) int {
	v, ok := this.ini.GetInt(k)
	if ok {
		return v
	}
	return -1
}

func (this *Config) Float(k string) float64 {
	v, ok := this.ini.GetFloat(k)
	if ok {
		return v
	}
	return 0.0
}

func NewConfig(file string) (*Config, error) {
	iniObj := goini.New()
	inierr := iniObj.ParseFile(file)
	if inierr != nil {
		return nil, inierr
	}

	conf := &Config{ini: iniObj}

	return conf, nil
}

var (
	defaultConfig *Config
)

func init() {
	var err error
	defaultConfig, err = NewConfig("app.conf")
	if err != nil {
		panic(err)
		return
	}
	if defaultConfig.String("RunMode", "") != "release" {
		defaultConfig.DEBUG = true
	} else {
		defaultConfig.DEBUG = false
	}
}

func String(key string, def string) string {
	return defaultConfig.String(key, def)
}

func StringN(key string) string {
	return defaultConfig.String(key, "")
}

func Bool(key string) bool {
	return defaultConfig.Bool(key)
}

func Int(key string) int {
	return defaultConfig.Int(key)
}

func Float(key string) float64 {
	return defaultConfig.Float(key)
}
