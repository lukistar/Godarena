// author: Jacky Boen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/lukistar/Godarena/draw"
)


var winTitle string = "Godarena"
var winWidth, winHeight int = 800, 600
var imageName string = "./data/sp/models/menu/menu00.png"

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer

	//var texture *sdl.Texture
	var src, dst *sdl.Rect

	var event sdl.Event
	var running bool

	window, renderer = draw.Init(winTitle, winWidth, winHeight)
	defer window.Destroy()
	defer renderer.Destroy()

	src = &sdl.Rect{0, 0, 1024, 768}
	dst = &sdl.Rect{0, 0, 800, 600}

	state := draw.GameState(draw.InitMenu())

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			state.HandleEvents(&event, &running)
		}
		state.Draw(renderer, src, dst, winWidth, winHeight)
	}
	// running = true
	// for running {
	// 	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch t := event.(type) {
	// 		case *sdl.QuitEvent:
	// 			running = false
	// 		case *sdl.KeyDownEvent:
	// 			switch t.Keysym.Sym {
	// 			case sdl.K_UP:
	// 				imageName = "./menu/menu01.png"
	// 			case sdl.K_DOWN:
	// 			 	imageName = "./menu/menu02.png"
	// 			case sdl.K_RETURN:
	// 				fmt.Printf("ENTER\n")
	// 			}
	// 		}
	// 	}
	// 	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	// 	renderer.Clear()
	// 	renderer.SetDrawColor(0, 0, 0, 0xFF)
	// 	texture = draw.Texture(renderer, imageName)

	// 	renderer.FillRect(&sdl.Rect{0, 0, int32(winWidth), int32(winHeight)})
	// 	renderer.Copy(texture, &src, &dst)
	// 	texture.Destroy()
	// 	renderer.Present()
	// }
}