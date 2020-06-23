package hotfix

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/agiledragon/gomonkey"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher
var patchMap sync.Map

// watch pluginPath, autoload plugins to replace origin function.
//
// dir: watch directory of plugin files.
// if dir is empty, use default value: os.Getwd() + "/tmp".
//
// if loadAll is true, will load and run all existing plugins under dir before start watching
func Watch(dir string, loadAll bool) {
	if dir != "" {
		SetPluginPath(dir)
	}
	if loadAll {
		loadALlPlugins()
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
					c, err := load(getPluginName(event.Name))
					if err != nil {
						log.Println("plugin load error", err)
						continue
					}
					err = runNewFunc(c)
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

// close *fsnotify.Watcher
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

func loadALlPlugins() {
	_ = filepath.Walk(pluginPath, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".so") {
			return nil
		}
		log.Println("watch load ", path)
		c, err := load(getPluginName(path))
		if err != nil {
			log.Println("plugin load error", err)
			return nil
		}
		err = runNewFunc(c)
		if err != nil {
			log.Println("plugin init error", err)
			return nil
		}
		return nil
	})
}
