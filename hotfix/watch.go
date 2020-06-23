package hotfix

import (
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/agiledragon/gomonkey"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher
var patchMap sync.Map

// watch pluginPath, autoload plugins to replace origin function
// dir: watch directory of plugin files
// if dir is empty, use default value: os.Getwd() + "/tmp"
func Watch(dir string) {
	if dir != "" {
		SetPluginPath(dir)
	}
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if !strings.HasSuffix(event.Name, ".so") {
					continue
				}

				switch event.Op {
				case fsnotify.Write, fsnotify.Create:
					log.Println(event.String())
					c, err := Load(getPluginName(event.Name))
					if err != nil {
						log.Println("plugin load error", err)
						continue
					}
					err = Init(c)
					if err != nil {
						log.Println("plugin init error", err)
						continue
					}
				case fsnotify.Remove:
					log.Println(event.String())
					pluginName := getPluginName(event.Name)
					value, ok := patchMap.Load(pluginName)
					if !ok {
						log.Println("plugin remove patchMap load empty")
						continue
					}
					value.(*gomonkey.Patches).Reset()
					patchMap.Delete(pluginName)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watch error:", err)
			}
		}
	}()

	if err = watcher.Add(pluginPath); err != nil {
		log.Fatal(err)
	}
	log.Printf("start watch %s\n", pluginPath)
}

func StopWatch() {
	if watcher != nil {
		watcher.Close()
	}
}

func RegisterPatch(name string, patch *gomonkey.Patches) {
	patchMap.Store(name, patch)
}

func getPluginName(eventName string) string {
	patchName := filepath.Base(eventName)
	return patchName[:len(patchName)-3]
}
