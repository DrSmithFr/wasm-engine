# 3D Game Engine - WebAssembly Port

This is an attempt to port the 3D Game Engine like DOOM to WebAssembly.
i'll try to rellay on the Go language to write the engine and the rendering code.
Leveraging the browser's html rendering capabilities for UI and the WebAssembly for the game rendering.
Also using html5 video element for cinematic and audio element for sound effects.

## Goals

- [x] WebAssembly and Goland to render the game via multiple methods, such as 2D Canvas, WebGL, WebGPU, etc.
- [x] Use Go to write the engine logic and the rendering code.
- [x] Render the game using the 2D Canvas API.
- [ ] Render the game using the WebGL API.
- [X] Implement multi-canvas rendering, using CSS to manage positioning and z-index.
- [ ] Dialog render using HTML elements.
- [ ] Allow video playback using the HTML5 video element for cinematic.
- [ ] Implement multiplayer support.

## Engine Features

- [x] Basic player movement (forward, backward, turn left and right).
- [X] Advance 3D rendering using vertex-base map.
- [ ] Advance collision detection.
- [ ] Texture mapping.
- [ ] Sprite rendering.
- [ ] Lighting.
- [ ] Particle effects.
- [ ] Light Bouncing:

## Deliverables

- [ ] Remake the first level of Duke Nukem 3D.

## How to try it out

1. Install Go 1.18 or later.
2. Clone the repository.
3. Run `go mod tidy` to download the dependencies.
4. Copy the `wasm_exec.js` file from the Go installation directory to the `public/build`. 

    run `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" build/`

5. Run `go run main.go server:start` to start the server.
6. Open `http://localhost:8080` in your browser.

Or run `go run main.go build:watch` to rebuild on changes. And start your own server on `./public`.
