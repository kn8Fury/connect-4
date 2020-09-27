package results

// Result of the game
type Result string

const (
	// Continue : current player did not win and the board is not full
	Continue Result = "continue"
	// Draw : current player did not win and the board is full
	Draw Result = "draw"
	// Invalid : current move is invalid
	Invalid Result = "invalid"
	// Win : current player wins the game
	Win Result = "win"
)
