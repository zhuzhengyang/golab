package hotfix

import (
	"os"
	"path/filepath"
)

// plugin go files and so files location.
// default: os.Getwd() + "/tmp"
var pluginPath string

func init() {
	temp, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pluginPath = filepath.Join(temp, "tmp")
}

// Set the install path of plugins as you want
func SetPluginPath(path string) {
	pluginPath = path
}

func getPluginPath() string {
	return pluginPath
}
