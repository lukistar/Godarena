package Godarena

import "time"

// Although its name, that's just an object moving at one direction.
type Unit struct {
	Coords
	Elements
	Stats

	// The return value should be true if the object couldn't move because of collision.
	whenCollide func(self, other Collider) bool

	// When did the object move for last time
	lastMove time.Time

	// When did the object use energy for last time
	lastEnergy time.Time

	// Comes from the last move of the object.
	directionX, directionY int

	// What kind of element attack it is.
	attackId int
}

func (u *Unit) X() int {
	return u.x
}

func (u *Unit) Y() int {
	return u.y
}

// Adds byX to the object's x.
// Adds byY to the object's y.
func (u *Unit) Move(byX, byY int) {
	if (int64)(time.Since(u.lastMove)) > (int64)(u.stats[CDMOVE]) {
		if byX > 0 {
			u.directionX = 1
		}
		if byX < 0 {
			u.directionX = -1
		}
		if byY > 0 {
			u.directionY = 1
		}
		if byY < 0 {
			u.directionY = -1
		}

		u.x += byX
		u.y += byY
		i, collide := u.Collide()
		if collide {
			goBack := u.OnCollide(COLLIDERS[i])
			COLLIDERS[i].OnCollide(u)
			if goBack {
				u.Move(-byX, -byY)
			}
		}
		u.lastMove = time.Now()
	}
}

// Finds if the object colides or not with something.
// If it does collide with B, return the position of B into COLLIDERS[].
func (u *Unit) Collide() (int, bool) {
	for i := 0; i < len(COLLIDERS); i++ {
		if COLLIDERS[i].X() == u.x && COLLIDERS[i].Y() == u.y {
			return i, true
		}
	}
	return 0, false
}

func (u *Unit) OnCollide(other Collider) bool {
	boo := u.whenCollide(u, other)
	return boo
}

// attackType could be Range(Mele), Radius
func (u *Unit) Attack(AttackType int) {
	if (int64)(time.Since(u.lastEnergy)) > (int64)(u.stats[CDENERGY]) {
		if AttackType == RANGE {
			// Create RangeAttack. Then append it to COLLIDERS.
		}
		if AttackType == RADIUS {
			// Coming "soon".
		}
		u.lastEnergy = time.Now()
	}
}

func (u *Unit) Destroy() {
	if u.sum == 0 {
		// Remove u from COLLIDERS.
	}
}

func (u *Unit) Update() {
	u.Destroy()
}

func (u *Unit) GetDamage(elementId int, amount int) {
	for i := 0; i < 7; i++ {
		u.elements[i] -= (DamageTable[elementId][i] * amount) / 100
	}
}
