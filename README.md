# 3D Game Engine - WebAssembly Port

This is attempt to port the 3D Game Engine of DOOM to WebAssembly. 
The original engine is written in C++ and uses OpenGL for rendering. 

## Goals

- [x] WebAssembly and Goland to render the game via multiple methods, such as 2D Canvas, WebGL, WebGPU, etc.
- [x] Use Go to write the engine logic and the rendering code.
- [x] Render the game using the 2D Canvas API.
- [ ] Render the game using the WebGL API.
- [ ] Implement multi-canvas rendering, using CSS to manage positioning and z-index.
- [ ] Dialog render using HTML elements.
- [ ] Allow video playback using the HTML5 video element for cinematic.
- [ ] Implement multiplayer support.


## Deliverables

- [ ] Remake the first level of Duke Nukem 3D.

## How to try it out

1. Install Go 1.18 or later.
2. Clone the repository.
3. Copy the `wasm_exec.js` file from the Go installation directory to the `public/build`. 

    run `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" public/build/`

4. Run `go build -o public/build/engine.wasm engine.go` to build the WebAssembly binary.
5. Run `go run server.go` to start the server.
6. Open `http://localhost:8080` in your browser.