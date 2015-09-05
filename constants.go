package Godarena

// Id of each element.
const (
	CHAOS = iota
	LIGHT
	FIRE
	WATER
	EARTH
	WIND
	ORGANIC
)

// Id of each stats.
const (
	MASS     = iota // Your mass.
	ENERGY          // How much energy you have.
	RELEASE         //How much energy you can use at once.
	CDENERGY        // How much time until capable of using energy.
	CDMOVE          // How much time until capable of moving again.
	UNABSORB        // How much elements you can keep unabsorbed.
)

// Basic arguments for Move().
const (
	LeftX  = -1
	LeftY  = 0
	RightX = 1
	RightY = 0
	UpX    = 0
	UpY    = -1
	DownX  = 0
	DownY  = 1
)

// Used for calling attack of unit.
const (
	RANGE = iota // Range or Mele at the end it doesn't even matter.
	RADIUS
)
