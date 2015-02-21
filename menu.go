package logic

//import "math"
/*
	Just realized that Menu would be the same as MenuOption,
	i.e every MenuOption is a Menu.
*/
type MenuOption struct {
	CurrentPosition uint8
	LastPosition uint8
	Options []*MenuOption
	HasOptions bool
	// Needed whenever HasOptions == false
	Activated bool
	//	You may want to identify the current Option.
	//	I assume you may need to keep some image path, without recreating yourself positions'logic.
	//	To be honest I need this field, i.e that's not really an assumption.
	Title string
}
func NewMenuOption(hasOptions bool, title string) *MenuOption {
	option := MenuOption{}
	option.CurrentPosition = 0
	option.LastPosition = 0
	option.HasOptions = hasOptions
	option.Options = []*MenuOption{}
	option.Activated = false
	option.Title = title
	return 	&option
}

func (m *MenuOption) Add(options ...*MenuOption) {
	m.Options = append(m.Options, options...)
	m.LastPosition += uint8(len(options))
}

func (m *MenuOption) Activate() *MenuOption{
	if m.HasOptions {
		position := m.CurrentPosition
		m.CurrentPosition = 0
		m.Activated = false //just to be sure
		return m.Options[position]
	} else {
		m.Activated = true
	}
	return m 
}

func (m *MenuOption) Deactivate() {
	m.Activated = false
}

func (m *MenuOption) Next() {
	m.CurrentPosition = (m.CurrentPosition + 1) % m.LastPosition
}

func (m *MenuOption) Prev() {
	m.CurrentPosition = (m.CurrentPosition - 1) % m.LastPosition
}