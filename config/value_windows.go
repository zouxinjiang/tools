/*
 * Copyright (c) 2019.
 */

package config

import (
	"os"
)

//go:generate   echo `uname`

var AppPath = os.TempDir() + "/cooker"

func hotReload() {

}
