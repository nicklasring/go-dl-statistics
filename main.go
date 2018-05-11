package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func handleEvent(event fsnotify.Event) {
	if event.Op == fsnotify.Create {
		log.Print(event.Name)
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				handleEvent(event)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
