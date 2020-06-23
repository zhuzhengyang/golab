package conf

import (
	"os"
	"path/filepath"
)

var PluginPath string

func init () {
	temp, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	PluginPath = filepath.Join(temp, "tmp")
}