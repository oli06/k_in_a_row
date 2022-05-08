package k_in_a_row

import (
	"testing"
)

func Test_isGameWon_returns_False_OnInit(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{0, 0}, 1, &game)

	result, usr := isGameWonWithToken(game)

	if result != false || usr != 0 {
		t.Fatalf("Result must be false and user must be 0, %t, %d", result, usr)
	}
}

func Test_getValueAt_returns_Nil_onInit(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	usr := getValueAt(point{0, 0}, game)

	if usr != 0 {
		t.Fatalf("User must be 0 on Init, %d", usr)
	}
}

func Test_getValueAt_returns_Usr1_afterPlacingToken(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{0, 0}, 1, &game)

	usr := getValueAt(point{0, 0}, game)

	if usr != 1 {
		t.Fatalf("User must be 1, %d", usr)
	}
}

func Test_placeToken_FailsForSameUserTwice(t *testing.T) {
	defer func() {
		recover()
	}()

	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{0, 0}, 1, &game)
	placeToken(&point{0, 1}, 1, &game)

	t.Fatalf("placeToken should have paniced!")
}

func Test_placeToken_FailsForSameUserBeforeAllOtherUsersPlacedAToken(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2, 3}, k: 4, lastMove: nil}

	defer func() {
		recover()
	}()

	placeToken(&point{0, 0}, 2, &game)
	placeToken(&point{0, 1}, 1, &game)

	t.Fatalf("placeToken should have paniced since user placed before it is his turn!")
}

func Test_placeToken_FailsForZeroUser(t *testing.T) {
	defer func() {
		recover()
	}()

	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{0, 0}, 0, &game)

	t.Fatalf("placeToken should fail for user 0!")
}

func Test_placeToken_FailsForSameField(t *testing.T) {
	defer func() {
		recover()
	}()

	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{0, 0}, 1, &game)
	placeToken(&point{0, 0}, 2, &game)

	t.Fatalf("placeToken should fail for same field, different user!")
}

/*func Test_placeToken_Fails_on_not_possible_field(t *testing.T) {
	defer func() {
		recover()
	}()

	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{2, 2}, 1, &game)

	t.Fatalf("placeToken should fail for unaccessible field!")
}*/

func Test_isGameWon_returns_True_onVerticalWin(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{0, 0}, 1, &game)
	placeToken(&point{1, 1}, 2, &game)
	placeToken(&point{0, 1}, 1, &game)
	placeToken(&point{2, 2}, 2, &game)
	placeToken(&point{0, 2}, 1, &game)
	placeToken(&point{3, 3}, 2, &game)
	placeToken(&point{0, 3}, 1, &game)

	result, usr := isGameWonWithToken(game)

	if !result || usr != 1 {
		t.Fatalf("User 1 should have won the game by now, %t, %d", result, usr)
	}
}

func Test_isGameWon_returns_True_onHorizontalWin(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2}, k: 4, lastMove: nil}

	placeToken(&point{0, 0}, 1, &game)
	placeToken(&point{1, 1}, 2, &game)
	placeToken(&point{1, 0}, 1, &game)
	placeToken(&point{2, 2}, 2, &game)
	placeToken(&point{2, 0}, 1, &game)
	placeToken(&point{3, 3}, 2, &game)
	placeToken(&point{3, 0}, 1, &game)

	result, usr := isGameWonWithToken(game)

	if !result || usr != 1 {
		t.Fatalf("User 1 should have won the game by now, %t, %d", result, usr)
	}
}

func Test_isGameWon_returns_True_onDiagonalTopLeftWin(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2, 3}, k: 4, lastMove: nil}

	//setup:
	//x
	//ex
	//oox
	//eeox

	placeToken(&point{3, 0}, 1, &game)
	placeToken(&point{2, 0}, 2, &game)
	placeToken(&point{1, 0}, 3, &game)
	placeToken(&point{2, 1}, 1, &game)
	placeToken(&point{1, 1}, 2, &game)
	placeToken(&point{0, 0}, 3, &game)
	placeToken(&point{1, 2}, 1, &game)
	placeToken(&point{0, 1}, 2, &game)
	placeToken(&point{0, 2}, 3, &game)
	placeToken(&point{0, 3}, 1, &game)

	result, usr := isGameWonWithToken(game)

	if !result || usr != 1 {
		t.Fatalf("User 1 should have won the game by now, %t, %d", result, usr)
	}
}

func Test_isGameWon_returns_True_onDiagonalBottomLeftWin(t *testing.T) {
	game := kInARowGame{board: make(map[point]int), users: []int{1, 2, 3}, k: 4, lastMove: nil}

	//setup:
	//   x
	//  xe
	// xoo
	//xoee

	placeToken(&point{0, 0}, 1, &game)
	placeToken(&point{1, 0}, 2, &game)
	placeToken(&point{2, 0}, 3, &game)
	placeToken(&point{1, 1}, 1, &game)
	placeToken(&point{2, 1}, 2, &game)
	placeToken(&point{3, 0}, 3, &game)
	placeToken(&point{2, 2}, 1, &game)
	placeToken(&point{3, 1}, 2, &game)
	placeToken(&point{3, 2}, 3, &game)
	placeToken(&point{3, 3}, 1, &game)

	result, usr := isGameWonWithToken(game)

	if !result || usr != 1 {
		t.Fatalf("User 1 should have won the game by now, %t, %d", result, usr)
	}
}
