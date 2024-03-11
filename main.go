package main

import (
    "fmt"
    go_console "github.com/DrSmithFr/go-console"
    "github.com/DrSmithFr/go-console/input/option"
    "go-webgl/builder"
    "log"
    "net/http"
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

                    // Build the wasm binary on file changes
                    signal := make(chan bool)
                    go builder.WatchAndBuild(cmd, signal)

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

    // Build the wasm binary on file changes
    signal := make(chan bool)
    go builder.WatchAndBuild(cmd, signal)

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
