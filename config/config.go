package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"sync"


)
// global vars

var Config config

var onceConfig sync.Once

func init() {
	onceConfig.Do(func() {


		if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
			fmt.Println(err)
		}


	})
}

type config struct {
	App          app
	Database     database
	SuperAdmin   SuperAdmin
	JWT          JWT
	Mail         Mail
	LambdaConfig LambdaConfig
}

type database struct {
	Connection   string
	Server   string
	Port     string
	Database string
	User     string
	Password string
}

type app struct {
	Name string
	Port string
	Migrate string
	Seed string
}

type JWT struct {
	Secret string
	Ttl int
}


type SuperAdmin struct {
	Login string
	Email string
	Password string
}

type Mail struct {
	Driver string
	Host string
	Port int
	Username string
	Password string
	Encryption string
}

type LambdaConfig struct {
	ConfigFile string
}
