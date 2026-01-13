package main

import "fmt"

/*
# 36. Valid Sudoku

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

  - Each row must contain the digits 1-9 without repetition.
  - Each column must contain the digits 1-9 without repetition.
  - Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:

  - A Sudoku board (partially filled) could be valid but is not necessarily solvable.
  - Only the filled cells need to be validated according to the mentioned rules.

Constraints:

  - board.length == 9
  - board[i].length == 9
  - board[i][j] is a digit 1-9 or '.'.
*/
func main() {
	testCases := []struct {
		i [][]byte
		e bool
	}{{i: [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	},
		e: true,
	}, {i: [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	},
		e: false},
	}

	for i, tc := range testCases {
		result := isValidSudoku(tc.i)
		fmt.Printf("Test %d: %t\n", i+1, result)
	}
}

// https://leetcode.com/problems/valid-sudoku/description/
func isValidSudoku(board [][]byte) bool {
	cm := [32]byte{}

	for i, r := range board {
		for j, c := range r {
			if c != '.' {
				num := int(c - '1')

				rByte := (i*9 + num) / 8
				rBit := (i*9 + num) % 8

				cByte := (9*9 + j*9 + num) / 8
				cBit := (9*9 + j*9 + num) % 8

				sq := ((i / 3 * 3) + j/3) * 9
				sqByte := (9*9*2 + sq + num) / 8
				sqBit := (9*9*2 + sq + num) % 8

				rExist := cm[rByte]&(1<<rBit) != 0
				cExist := cm[cByte]&(1<<cBit) != 0
				sqExist := cm[sqByte]&(1<<sqBit) != 0

				if rExist || cExist || sqExist {
					return false
				}

				cm[rByte] |= 1 << rBit
				cm[cByte] |= 1 << cBit
				cm[sqByte] |= 1 << sqBit
			}
		}
	}

	return true
}

/*
		0 1 2 j ++
						i
	  0 1 2   0
	  3 4 5   1
	  6 7 8   2

		0-8
		i=1
		j=2
		5блок - 1*3+2=5
*/

/*
				0		 0    0    1    1    1    2    2    2
	j		  0    1    2    3    4    5    6    7    8
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'}, 0 0
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'}, 1 0
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, 2 0
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, 3 1
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, 4 1
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'}, 5 1
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, 6 2
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'}, 7 2
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}, 8 2

																										 i

*/

/*
квадрат №1 = 0,1,2 | 9, 10, 11 | 18, 19, 20

  	i j
1 = 0 0
2 = 0 1
3 = 0 2
4 = 1 0
5 = 1 1
6 = 1 2
7 = 2 0
8 = 2 1
9 = 2 2
*/
