package main

import (
    "fmt"
    go_console "github.com/DrSmithFr/go-console"
    "github.com/DrSmithFr/go-console/input/option"
    "github.com/fsnotify/fsnotify"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "time"
)

func main() {
    script := go_console.Command{
        Description: "Start the server",
        Scripts: []*go_console.Script{
            {
                Name:        "server:start",
                Description: "Start the web server.",
                Options: []go_console.Option{
                    {
                        Name:         "host",
                        Shortcut:     "H",
                        Description:  "The host address to use",
                        DefaultValue: "127.0.0.1",
                        Value:        option.Optional,
                    },
                    {
                        Name:         "port",
                        Shortcut:     "p",
                        Description:  "Port to listen on",
                        DefaultValue: "8080",
                        Value:        option.Optional,
                    },
                },
                Runner: startDevServer,
            },
            {
                Name:        "build:watch",
                Description: "Watch for file changes and rebuild the wasm binary.",
                Runner: func(cmd *go_console.Script) go_console.ExitCode {
                    cmd.PrintTitle("3D Engine Builder")

                    cmd.PrintText("First to ensure the wasm binary up to date...")
                    buildWasm(cmd)
                    lastBuild := time.Now()

                    signal := make(chan bool)
                    go WatchAndBuild(cmd, signal)

                    // build on chan signal
                    go func() {
                        for {
                            <-signal
                            if time.Since(lastBuild) > 100*time.Millisecond {
                                buildWasm(cmd)
                                lastBuild = time.Now()
                            }
                        }
                    }()

                    // Block main goroutine forever.
                    <-make(chan struct{})
                    return go_console.ExitSuccess
                },
            },
        },
    }

    script.Run()
}

func startDevServer(cmd *go_console.Script) go_console.ExitCode {
    cmd.PrintTitle("Starting 3D Engine...")

    buildWasm(cmd)
    lastBuild := time.Now()

    signal := make(chan bool)
    go WatchAndBuild(cmd, signal)

    // build on chan signal
    go func() {
        for {
            <-signal
            if time.Since(lastBuild) > 100*time.Millisecond {
                buildWasm(cmd)
                lastBuild = time.Now()
            }
        }
    }()

    host := cmd.Input.Option("host")
    port := cmd.Input.Option("port")

    cmd.PrintText(fmt.Sprintf("Server started on http://%s:%s", host, port))
    http.Handle("/", &noCache{Handler: http.FileServer(http.Dir("./public"))})

    if err := http.ListenAndServe(":8080", nil); err != nil {
        cmd.PrintError(err.Error())
        log.Fatalln(err)
    }

    return go_console.ExitSuccess
}

type noCache struct {
    http.Handler
}

func (h *noCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-cache")
    h.Handler.ServeHTTP(w, r)
}

func WatchAndBuild(cmd *go_console.Script, signal chan bool) {
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

    // Watch the individual files
    individualFiles := []string{"wasm.go", "go.mod", "go.sum", "public/index.html", "public/style.css"}
    for _, file := range individualFiles {
        err = watcher.Add(file)
        if err != nil {
            log.Fatal(err)
        }
    }

    // Watch all directories except the ones in the blacklist
    blacklist := []string{"public", "node_modules", ".git", ".github", ".idea"}
    items, _ := ioutil.ReadDir(".")
    for _, item := range items {
        path := item.Name()

        if item.IsDir() && !inArray(path, blacklist) {
            err = watcher.Add(path)

            if err != nil {
                log.Fatal(err)
            }
        }
    }

    // Watch the public directory for changes, except the build directory
    err = watcher.Add("public")
    items, _ = ioutil.ReadDir("./public")
    for _, item := range items {
        path := item.Name()

        if item.IsDir() && path != "build" {
            err = watcher.Add("public/" + path)

            if err != nil {
                log.Fatal(err)
            }
        }
    }

    // Block main goroutine forever.
    <-make(chan struct{})
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
        cmd.Output.Println(fmt.Sprintf("<warning>WASM build failed: %s</warning>\n", err))
        panic(err)
    }

    cmd.Output.Println(fmt.Sprintf("<info>WASM build took %s</info>\n", time.Since(start)))
}

func inArray(val string, array []string) bool {
    for _, v := range array {
        if v == val {
            return true
        }
    }
    return false
}
