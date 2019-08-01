/*
 * Copyright (c) 2019.
 */

package config

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//go:generate   echo `uname`

var AppPath = "/cooker"

func hotReload() {
	var sig = make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR1)
	go func() {
		for {
			<-sig
			conf, err := LoadConfigFromFile(ConfigFile)
			if err != nil {
				fmt.Println("重新加载配置出错", err)
				continue
			}
			locker.Lock()
			appconfig = conf
			locker.Unlock()
		}
	}()
}
