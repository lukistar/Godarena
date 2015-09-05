package Godarena

// Everything that collides should implement Collider.
type Collider interface {
	Collide() (int, bool)
	OnCollide(other Collider) (bool)
	X() (int)
	Y() (int)
	Update()
}

type Updater interface {
	Update()
}

type Mover interface {
	Move(byX, byY int)
}
