package k_in_a_row

import "fmt"

type point struct {
	x int
	y int
}

type kInARowGame struct {
	board    map[point]int
	users    []int
	k        int
	lastMove *point
	border   point
}

func isGameWonWithToken(game kInARowGame) (bool, int) {
	if game.lastMove == nil {
		return false, 0
	}

	usr := getValueAt(*game.lastMove, game)
	if usr == 0 {
		return false, usr
	}

	x, y := game.lastMove.x, game.lastMove.y
	required_ks := game.k - 1

	//check horizontal, vertical, diagonal_tleft, diagonal_tright
	//horizontal:
	var count int
	index := x - 1
	for game.board[point{index, y}] == usr {
		index -= 1
		count++
	}
	index = x + 1
	for game.board[point{index, y}] == usr {
		index += 1
		count++
	}

	if count == required_ks {
		return true, usr
	}

	//vertical
	index = y + 1
	for game.board[point{x, index}] == usr {
		index += 1
		count++
	}
	index = y - 1
	for game.board[point{x, index}] == usr {
		index -= 1
		count++
	}

	if count == required_ks {
		return true, usr
	}

	//diagonal top_left
	index_x := x - 1
	index_y := y + 1
	for game.board[point{index_x, index_y}] == usr {
		index_x -= 1
		index_y += 1
		count++
	}
	index_x = x + 1
	index_y = y - 1
	for game.board[point{index_x, index_y}] == usr {
		index_x += 1
		index_y -= 1
		count++
	}

	if count == required_ks {
		return true, usr
	}

	//diagonal top_right
	index_x = x + 1
	index_y = y + 1
	for game.board[point{index_x, index_y}] == usr {
		index_x += 1
		index_y += 1
		count++
	}
	index_x = x - 1
	index_y = y - 1
	for game.board[point{index_x, index_y}] == usr {
		index_x -= 1
		index_y -= 1
		count++
	}

	if count == required_ks {
		return true, usr
	}

	return false, 0
}

func getValueAt(coords point, game kInARowGame) int {
	return game.board[coords]
}

func placeToken(placement *point, usr int, game *kInARowGame) {
	if usr == 0 {
		panic("usr == 0 is not allowed to place tokens!")
	}

	currentUser := getCurrentUser(*game)
	if currentUser != usr && currentUser != -2 {
		panic("user is not allowed to play!")
	}

	if isOutsideBorders(placement, *game) {
		panic("token cant be placed outside the border!")
	}

	if isGravityUsed(placement, *game) {
		panic("token missplacement! Cant place token in the air")
	}

	if game.board[*placement] == 0 {
		game.board[*placement] = usr
		game.lastMove = placement
	} else {
		panic("resetting a token is not allowed!")
	}
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func getCurrentUser(game kInARowGame) int {
	if game.lastMove == nil {
		return -2
	}

	lastUsr := game.board[*game.lastMove]
	lastUsrIndex := SliceIndex(len(game.users), func(i int) bool { return game.users[i] == lastUsr })
	if lastUsrIndex == -1 {
		return -1
	}

	return game.users[(lastUsrIndex+1)%len(game.users)]
}

func isGravityUsed(placement *point, game kInARowGame) bool {
	//check if all x coordinates below placement.x are placed (i.e != 0)

	y_index := placement.y - 1
	for y_index > -1 {
		if getValueAt(point{placement.x, y_index}, game) == 0 {
			return true
		}

		y_index--
	}

	return false
}

func isOutsideBorders(placement *point, game kInARowGame) bool {
	return !(placement.x < game.border.x && placement.y < game.border.y)
}

//TODO: test
func print(game kInARowGame) {
	for i := game.border.x - 1; i >= 0; i-- {
		should_print := false
		row := ""
		for j := game.border.y - 1; j >= 0; j-- {
			val := getValueAt(point{i, j}, game)
			if val != 0 {
				should_print = true
				row += fmt.Sprint(val)
			} else {
				row += "_"
			}
		}

		if should_print {
			fmt.Println(row)
		}
	}
}
