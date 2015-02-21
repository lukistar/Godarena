package draw

import "github.com/veandco/go-sdl2/sdl"

type GameState interface {
	HandleEvents(*sdl.Event, *bool)
	Draw(*sdl.Renderer, *sdl.Rect, *sdl.Rect, int, int)
}
