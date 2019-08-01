/*
 * Copyright (c) 2019.
 */

package config

import (
	"sync"
)

var ConfigFile = AppPath + "/conf/app_conf.json"
var locker sync.RWMutex
var appconfig AppConfig

func defaultConfig() AppConfig {
	return AppConfig{
		Inited:      false,
		Debug:       false,
		ConfigDir:   AppPath + "/conf",
		TemplateDir: AppPath + "/temp",
		WebDir:      AppPath + "/web",
		BinDir:      AppPath + "/bin",
		DatabaseConfig: DatabaseConfig{
			Driver:   "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Database: "user",
			UserName: "root",
			Password: "",
		},
		LoggerConfig: LoggerConfig{
			Driver:      "os.stdout",
			Destination: "",
		},
		UserConfig: map[string]interface{}{
			"WeiXin": map[string]string{
				"AppId":  "",
				"Secret": "",
				"Token":  "",
			},
		},
	}
}
