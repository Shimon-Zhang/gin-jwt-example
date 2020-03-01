package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var Configs *ini.Section

func InitCnf(env string) {
	cnf, err := ini.Load("app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	Configs = cnf.Section(env)
}
