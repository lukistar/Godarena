package draw

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/lukistar/Godarena"
	"fmt"
)

const (
	//submenu 0
	MENU00 = "./data/sp/models/menu/menu00.png"
	MENU01 = "./data/sp/models/menu/menu01.png"
	MENU02 = "./data/sp/models/menu/menu02.png"
	MENU03 = "./data/sp/models/menu/menu03.png"
	MENU04 = "./data/sp/models/menu/menu04.png"

	//submenu 1
	MENU01_00 = "./data/sp/models/menu/menu01_00.png"
	MENU01_01 = "./data/sp/models/menu/menu01_01.png"
	MENU01_02 = "./data/sp/models/menu/menu01_02.png"
	MENU01_03 = "./data/sp/models/menu/menu01_03.png"
	MENU01_04 = "./data/sp/models/menu/menu01_04.png"

	//submenu 2
	MENU02_00 = "./data/sp/models/menu/menu02_00.png"
	MENU02_01 = "./data/sp/models/menu/menu02_01.png"
	MENU02_02 = "./data/sp/models/menu/menu02_02.png"
	MENU02_03 = "./data/sp/models/menu/menu02_03.png"
	MENU02_04 = "./data/sp/models/menu/menu02_04.png"
)

type Menu struct {
	Logic *logic.MenuOption
}
func InitMenu()	*Menu {
	menu := Menu{nil}
	menu00 := logic.NewMenuOption(true, MENU00)
	menu01 := logic.NewMenuOption(false, MENU01)
	menu02 := logic.NewMenuOption(false, MENU02)
	menu03 := logic.NewMenuOption(false, MENU03)
	menu04 := logic.NewMenuOption(true, MENU04)
	menu.Logic = menu00
	menu.Logic.Add(menu01, menu02, menu03, menu04)
	menu04.Add(menu00)
	menu02_01 := logic.NewMenuOption(false, MENU02_01)
	menu02_02 := logic.NewMenuOption(false, MENU02_02)
	menu02_03 := logic.NewMenuOption(false, MENU02_03)
	menu02_04 := logic.NewMenuOption(true, MENU02_04)
	menu02_04.Add(menu00)
	menu01_01 := logic.NewMenuOption(false, MENU01_01)
	menu01_02 := logic.NewMenuOption(false, MENU01_02)
	menu01_03 := logic.NewMenuOption(false, MENU01_03)
	menu01_04 := logic.NewMenuOption(true, MENU01_04)
	menu01_04.Add(menu00)
	menu01.Add(menu01_01, menu01_02, menu01_03, menu01_04)
	menu02.Add(menu02_01, menu02_02, menu02_03, menu02_04)
	return &menu
}

func (m Menu) HandleEvents(event *sdl.Event, running *bool) {
	switch t := (*event).(type) {
	case *sdl.QuitEvent:
		*running = false
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		case sdl.K_UP:
			m.Logic.Prev()
			fmt.Printf("Position %d\n", m.Logic.CurrentPosition)
		case sdl.K_DOWN:
			m.Logic.Next()
			fmt.Printf("Position %d\n", m.Logic.CurrentPosition)
		case sdl.K_RETURN:
			m.Logic = m.Logic.Activate()
			if m.Logic.Activated {
				*running = false
			}
		}
	}
}

func (m Menu) Draw(renderer *sdl.Renderer, src, dst *sdl.Rect, width, height int) {
	renderer.Clear()
 	renderer.SetDrawColor(0, 0, 0, 0xFF)

	texture := Texture(renderer, m.Logic.Options[m.Logic.CurrentPosition].Title)

	renderer.FillRect(&sdl.Rect{0, 0, int32(width), int32(height)})
 	renderer.Copy(texture, src, dst)
 	texture.Destroy()
 	renderer.Present()
}