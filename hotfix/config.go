package hotfix

import (
	"os"
	"path/filepath"
)

// plugin go files and so files location
// default: os.Getwd() + "/tmp"
var pluginPath string

func init() {
	temp, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pluginPath = filepath.Join(temp, "tmp")
}

func SetPluginPath(path string) {
	pluginPath = path
}

func GetPluginPath() string {
	return pluginPath
}
