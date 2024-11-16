package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard within the given file.
func CountInFile(cb Chessboard, file string) int {
	// Check if the file exists in the chessboard
	if row, exists := cb[file]; exists {
		occupiedCount := 0
		for _, isOccupied := range row {
			if isOccupied {
				occupiedCount++
			}
		}
		return occupiedCount
	}
	// Return 0 if the file does not exist
	return 0
}

// CountInRank returns how many squares are occupied in the chessboard within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	// Validate the rank input (should be between 1 and 8 inclusive)
	if rank < 1 || rank > 8 {
		return 0
	}

	occupiedCount := 0
	// Loop through all files (columns) to count occupied squares in the given rank
	for _, row := range cb {
		if row[rank-1] {
			occupiedCount++
		}
	}
	return occupiedCount
}

// CountAll returns the total number of squares in the chessboard.
func CountAll(cb Chessboard) int {
	totalCount := 0
	// Loop through each file and count the squares
	for _, row := range cb {
		totalCount += len(row)
	}
	return totalCount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedCount := 0
	// Loop through all rows and files to count the occupied squares
	for _, row := range cb {
		for _, isOccupied := range row {
			if isOccupied {
				occupiedCount++
			}
		}
	}
	return occupiedCount
}
