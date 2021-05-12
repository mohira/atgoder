package topic_grid

func toStringBoard(H, W int, S []string) [][]string {
	var board = make([][]string, H)

	for i := 0; i < H; i++ {
		board[i] = make([]string, W)

		for j := 0; j < W; j++ {
			board[i][j] = string(S[i][j])
		}
	}

	return board
}
