package main

// Winner returns:
//
//  1 if Move-A Wins
//  0 if Neither Move Wins
// -1 if Move-B Wins
//
func Winner(moveA int, moveB int) int {
	switch {

	// Move-A Wins!
	case moveA == Rock && moveB == Scissors:
		return 1
	case moveA == Paper && moveB == Rock:
		return 1
	case moveA == Scissors && moveB == Paper:
		return 1

    // Move-B Wins!
	case moveA == Scissors && moveB == Rock:
		return -1
	case moveA == Rock && moveB == Paper:
		return -1
	case moveA == Paper && moveB == Scissors:
		return -1
	}

	return 0
}
