type GameState interface {
	Events()
    Logic()
    Render()
}