package main

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

func main() {
	flag.Parse()
	path := ""
	args := flag.Args()
	if len(args) == 0 {
		path = "."
	} else {
		path = args[0]
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal("Can't monitor something that doesn't exist")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Watcher is broken")
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error", err)
			}
		}
	}()
	err = watcher.Add(path)
	if err != nil {
		log.Fatal("err:", err)
	}
	<-done
}
