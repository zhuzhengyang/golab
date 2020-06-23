package watch

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err = watcher.Add("./plugin"); err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op == fsnotify.Write || event.Op == fsnotify.Create {
				}
			}
		}
	}()
	<-done
}
