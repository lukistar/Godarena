package draw

import (
	"github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "fmt"
    "os"
)

func Init (title string, width, height int) (*sdl.Window, *sdl.Renderer) {
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}

	return window, renderer
}

func Texture (renderer *sdl.Renderer, imageName string) (*sdl.Texture) {
	image, err := img.Load(imageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(3)
	}
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(4)
	}

	return texture
}