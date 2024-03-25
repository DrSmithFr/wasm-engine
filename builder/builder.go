package builder

import (
	"fmt"
	go_console "github.com/DrSmithFr/go-console"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func WatchAndBuild(cmd *go_console.Script, signal chan bool) {
	forceFirstBuild := true

	lastBuild := time.Now()
	go Watch(cmd, signal)

	// build on chan signal
	go func() {
		// wait for the first build
		time.Sleep(300 * time.Millisecond)

		for {
			if forceFirstBuild || time.Since(lastBuild) > 100*time.Millisecond {
				buildWasm(cmd)
				lastBuild = time.Now()

				forceFirstBuild = false
			}
			<-signal
		}
	}()
}

func Watch(cmd *go_console.Script, signal chan bool) {
	cmd.PrintText("Starting file watcher...")

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				signal <- true
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add all the files to the watcher
	paths := make(chan string)
	go getProjectFilesToWatch(paths)
	for path := range paths {
		err = watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func getProjectFilesToWatch(paths chan string) {
	// Watch the individual files
	individualFiles := []string{"wasm.go", "go.mod", "go.sum", "public/index.html"}
	for _, file := range individualFiles {
		paths <- file
	}

	// Watch all directories except the ones in the blacklist
	blacklist := []string{"public", "node_modules", ".git", ".github", ".idea"}
	items, _ := ioutil.ReadDir(".")
	for _, item := range items {
		path := item.Name()

		if item.IsDir() && !inArray(path, blacklist) {
			paths <- path
		}
	}

	// Watch the public directory for changes, except the build directory
	items, _ = ioutil.ReadDir("./public")
	for _, item := range items {
		path := item.Name()

		if item.IsDir() && path != "build" {
			paths <- "public/" + path
		}
	}
}

func buildWasm(cmd *go_console.Script) {
	// set GOOS to js and GOARCH to wasm
	err := os.Setenv("GOOS", "js")

	if err != nil {
		panic(err)
	}

	err = os.Setenv("GOARCH", "wasm")

	if err != nil {
		panic(err)
	}

	start := time.Now()
	cmd.Output.Print("Building wasm...")

	_, err = exec.Command("go", "build", "-o", "public/build/engine.wasm", "wasm.go").CombinedOutput()

	if err != nil {
		cmd.Output.Println(fmt.Sprintf("<error>WASM build failed: %s</error>", err))
	} else {
		cmd.Output.Println(fmt.Sprintf("<info>WASM build took %s</info>", time.Since(start)))
	}
}

func inArray(val string, array []string) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}
