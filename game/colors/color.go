package colors

// Color is the color of the player coin
type Color int

const (
	// None represents no coin
	None Color = iota
	// Yellow represents a yellow coin
	Yellow
	// Red represents a red coin
	Red
)

func (c Color) String() string {
	return [...]string{"None", "Yellow", "Red"}[c]
}
