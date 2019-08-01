/*
 * Copyright (c) 2019.
 */

package clog

import (
	"os"
	"testing"
)

func TestClog(t *testing.T) {
	lg := Clog{
		w:      os.Stdout,
		level:  Lvl_Info | Lvl_Debug | Lvl_Error | Lvl_Warning,
		format: `{"fn":"$fn","line":$ln,"data":${data}}`,
	}
	lg.SetDataFormat(FMT_Json)
	lg.Error(map[string]string{"aaa": "ccc"})
	lg.Debug("aaa")
}

func TestCC(t *testing.T) {
	Warning("www", "tttt")
}

func TestCommonClog(t *testing.T) {
	lg := NewClog()
	lg.SetShowLevel(Lvl_Info | Lvl_Warning | Lvl_Error | Lvl_Debug)
	lg.Debug("debug info")
	lg.Error("error info")
	lg.Warning("warning info")
	lg.Info("info ")
	lg.SetFormat(`FN:$F t:$t`)
	lg.Info("fffff")
}

func TestAddFunc(t *testing.T) {
	lg := NewClog()
	lg.SetFormat(`[$l] $T file:$f line:$ln func:$fn custom:$my ${data}`)
	lg.AddCustomFormatFunc("my", func(level LogLevel, skip int) string {
		return "my func"
	})

	lg.Info("ssssssssssssssss")
}
